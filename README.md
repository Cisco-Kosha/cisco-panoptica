# Cisco Panoptica Connector

The Cisco Panoptica REST API has endpoints to perform all activities that can be performed on the console UI, including endpoints to create environments and clusters, create policies, add images and registries.

The REST API uses Escher authentication.

![panoptica](images/panoptica.png)

## Useful Actions

Panoptica APIs allow the users to retrieve and manage :
1. login, passwords
2. Dashboard statistics
3. Define environments
4. Telemetry information
5. Environment and Connection policies.
6. API Security policies
7. CI/CD policies
8. Registries
9. Kubernetes Clusters
10. API Security
11. Image Hashes
12. Audit Logs
13. Apps
14. Image Hashes
15. Vulnerabilities
16. Serveless Policies
17. System Settings

and many more..

## Endpoints

Refer to the Panoptica connector [API specification](openapi.json) for details.

## Example Usage

< sdk example? >

## Authentication

The REST API uses Escher authentication. This method uses a unique token for each request. The token is a hash generated from fixed Access and Secret keys (obtained from Panoptica), the request URL, and the request time.

### Access & Secret keys

Generate unique Access and Secret keys from the Panoptica console. These will be used for each request.

Navigate to the System page, and select MANAGE USERS.
Click New User and select Service User.
Enter a name for the user. Leave the status as 'Active'.
Click FINISH to receive the Token window.
Copy the values of Access Key and Secret Key, for use in the Escher authentication.