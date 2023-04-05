package httpclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/EscherAuth/escher/config"
	"github.com/EscherAuth/escher/request"
	"github.com/EscherAuth/escher/signer"
	"github.com/itchyny/timefmt-go"

	"github.com/kosha/panoptica-connector/pkg/logger"
)

func MakeEscherAuthCall(headers map[string]string, accessKey, secretKey, method, url string, body interface{}, log logger.Logger) (interface{}, int, error) {
	var req *http.Request
	var bodyString string
	if body != nil {
		jsonReq, _ := json.Marshal(body)
		bodyString = string(jsonReq)
	}

	req, _ = http.NewRequest(method, url, nil)

	for k, v := range headers {
		req.Header.Add(k, v)
	}
	t, err := timefmt.Parse(time.Now().UTC().Format("2006-01-02T15:04:05Z"), "%Y-%m-%dT%H:%M:%SZ")
	if err != nil {
		fmt.Println(err)
		return nil, 500, err
	}
	req.Header.Add("X-Escher-Date", timefmt.Format(t, "%Y%m%dT%H%M%SZ"))
	req.Header.Set("Accept-Encoding", "identity")

	escherHeaders := createEscherHeadersFromHTTPHeaders(req)

	escherReq := request.New(
		method,
		req.URL.Path,
		escherHeaders,
		bodyString,
		300,
	)
	// Escher Config
	c := config.Config{
		ApiSecret:       secretKey,
		AccessKeyId:     accessKey,
		CredentialScope: "global/services/portshift_request",
	}

	config.SetDefaults(&c)

	signedRequest, err := signer.New(c).SignRequest(escherReq, []string{})
	if err != nil {
		log.Error("Unable to sign request with EscherAuth")
		log.Error(err)
		return nil, 400, err
	}

	httpRequest, err := signedRequest.HTTPRequest(url)
	if err != nil {
		log.Error("Unable to turn signedRequest to http.Request")
		log.Error(err)
		return nil, 400, err
	}

	var response interface{}

	res, statusCode := makeEscherAuthReq(httpRequest, log)
	if string(res) == "" {
		return nil, statusCode, fmt.Errorf("nil")
	}
	// Convert response body to target struct
	err = json.Unmarshal(res, &response)
	if err != nil {
		log.Error("Unable to parse response as json")
		log.Error(err)
		return nil, 500, err
	}
	return response, statusCode, nil
}

func makeEscherAuthReq(req *http.Request, log logger.Logger) ([]byte, int) {

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		log.Error("unable to make http call to url: %s", req.URL.String())
		log.Error(err)
		return nil, 500
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
	}
	return bodyBytes, resp.StatusCode
}

func createEscherHeadersFromHTTPHeaders(r *http.Request) request.Headers {
	headers := request.Headers{}

	for key, values := range r.Header {
		for _, value := range values {
			headers = append(headers, [2]string{key, value})
		}
	}

	if r.Header.Get("host") == "" {
		// https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.23
		headers = append(headers, [2]string{"host", r.Host})
	}

	return headers
}
