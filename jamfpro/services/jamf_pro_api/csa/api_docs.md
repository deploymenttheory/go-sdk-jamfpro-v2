Returns the CSA tenant ID.
get
https://yourServer.jamfcloud.com/api/v1/csa/tenant-id

Returns the CSA tenant ID.

Response

200
CSA tenant ID is returned.

Response body
object
tenantId
string | null
The tenant ID

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/csa/tenant-id \
     --header 'accept: application/json'

{
  "tenantId": "123"
}
-----
Get details regarding the CSA token exchange
get
https://yourServer.jamfcloud.com/api/v1/csa/token

Get details regarding the CSA token exchange

Responses

200
Success

Response body
object
tenantId
string | null
The tenant ID

subject
string
Salesforce CRM account ID

refreshExpiration
int64
scopes
array of strings
legacyJamfSalesforceIds
array of strings

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/csa/token \
     --header 'accept: application/json'

{
  "tenantId": "12345678",
  "subject": "001C000000wkzDLIAY",
  "refreshExpiration": 1584544108,
  "scopes": [
    "read mobile_device",
    "write mobile_device"
  ],
  "legacyJamfSalesforceIds": [
    "001C000000wkzXXXXX",
    "001C000000wkzYYYYY"
  ]
}
-----
Delete the CSA token exchange - This will disable Jamf Pro's ability to authenticate with cloud-hosted services
delete
https://yourServer.jamfcloud.com/api/v1/csa/token

Delete the CSA token exchange - This will disable Jamf Pro's ability to authenticate with cloud-hosted services

Responses
204
CSA Token Exchange successfully deleted

curl --request DELETE \
     --url https://yourserver.jamfcloud.com/api/v1/csa/token \
     --header 'accept: application/json'