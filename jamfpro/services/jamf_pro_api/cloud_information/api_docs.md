Retrieve information related to cloud setup.
get
https://yourServer.jamfcloud.com/api/v1/cloud-information

Retrieve information related to cloud setup. Retrieves information related to cloud setup. Provides details about cloud instance configuration.

Response

200
Result of verification for being cloud-based instance

Response body
object
cloudInstance
boolean
rampInstance
boolean
govCloudInstance
boolean
managedServiceProviderInstance
boolean
Information whether this instance is managed by managed service provider

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/cloud-information \
     --header 'accept: application/json'

{
  "cloudInstance": true,
  "rampInstance": true,
  "govCloudInstance": true,
  "managedServiceProviderInstance": true
}