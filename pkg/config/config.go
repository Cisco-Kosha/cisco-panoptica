package config

import (
	"flag"
	"net/url"
	"os"
	"strings"
)

type Config struct {
	serverUrl string
	accessKey string
	secretKey string
}

func Get() *Config {
	conf := &Config{}
	flag.StringVar(&conf.accessKey, "accessKey", os.Getenv("ACCESS_KEY"), "Panoptica User Access Key")
	flag.StringVar(&conf.secretKey, "secretKey", os.Getenv("SECRET_KEY"), "Panoptica User Secret Key")
	flag.StringVar(&conf.serverUrl, "serverUrl", os.Getenv("SERVER_URL"), "Server Url")

	flag.Parse()

	return conf
}

func (c *Config) GetAccessAndSecretKey() (string, string) {
	return c.accessKey, c.secretKey
}

func (c *Config) GetServerURL() string {
	c.serverUrl = strings.TrimSuffix(c.serverUrl, "/")
	u, _ := url.Parse(c.serverUrl)
	if u.Scheme == "" {
		return "https://" + c.serverUrl
	} else {
		return c.serverUrl
	}
}

func (c *Config) GetServerHost() string {
	c.serverUrl = strings.TrimSuffix(c.serverUrl, "/")
	u, _ := url.Parse(c.serverUrl)
	if u.Scheme == "" {
		return u.Host
	} else {
		return u.Host
	}
}
