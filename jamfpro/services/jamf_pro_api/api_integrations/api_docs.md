Get the current API Integrations
get
https://yourServer.jamfcloud.com/api/v1/api-integrations

Get Jamf|Pro API Integrations with Search Criteria

Query Params
page
integer
Defaults to 0
0
page-size
integer
Defaults to 100
100
sort
array of strings
Defaults to id:asc
Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. Fields allowed in the query: id, displayName. Example: sort=displayName:desc


string

id:asc

ADD string
filter
string
Query in the RSQL format, allowing to filter app titles collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: id, displayName. Example: displayName=="IntegrationName"

Response

200
A list of the current Jamf API Integrations

Response body
object
totalCount
integer
required
≥ 0
results
array of objects
required
length ≥ 0
object
id
integer
required
authorizationScopes
array of strings
required
displayName
string
required
enabled
boolean
required
accessTokenLifetimeSeconds
integer
appType
string
enum
required
Type of API Client: * CLIENT_CREDENTIALS - A client ID and secret have been generated for this integration. * NATIVE_APP_OAUTH - A native app (i.e., Jamf Reset) has been linked to this integration for auth code grant type via Managed App Config. * NONE - No client is currently associated with this integration.

CLIENT_CREDENTIALS NATIVE_APP_OAUTH NONE

clientId
string
required

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v1/api-integrations?page=0&page-size=100&sort=id%3Aasc' \
     --header 'accept: application/json'

{
  "totalCount": 1,
  "results": [
    {
      "id": 1,
      "authorizationScopes": [
        "Tootsie Roal",
        "Jamf Reset"
      ],
      "displayName": "My API Integration",
      "enabled": true,
      "accessTokenLifetimeSeconds": 300,
      "appType": "CLIENT_CREDENTIALS",
      "clientId": "538878d4-9744-43ed-8732-3df99d502bd6"
    }
  ]
}

-----

Create API integration object
post
https://yourServer.jamfcloud.com/api/v1/api-integrations

Create API integration object

Body Params
api integration object to create

authorizationScopes
array of strings
required
API Role display names.


ADD string
displayName
string
required
enabled
boolean

true
accessTokenLifetimeSeconds
integer
Response

201
Api Integration object was created

Response body
object
id
integer
required
authorizationScopes
array of strings
required
displayName
string
required
enabled
boolean
required
accessTokenLifetimeSeconds
integer
appType
string
enum
required
Type of API Client: * CLIENT_CREDENTIALS - A client ID and secret have been generated for this integration. * NATIVE_APP_OAUTH - A native app (i.e., Jamf Reset) has been linked to this integration for auth code grant type via Managed App Config. * NONE - No client is currently associated with this integration.

CLIENT_CREDENTIALS NATIVE_APP_OAUTH NONE

clientId
string
required

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/api-integrations \
     --header 'accept: application/json' \
     --header 'content-type: application/json'

{
  "id": 1,
  "authorizationScopes": [
    "Tootsie Roal",
    "Jamf Reset"
  ],
  "displayName": "My API Integration",
  "enabled": true,
  "accessTokenLifetimeSeconds": 300,
  "appType": "CLIENT_CREDENTIALS",
  "clientId": "538878d4-9744-43ed-8732-3df99d502bd6"
}

------

Get specified API integration object
get
https://yourServer.jamfcloud.com/api/v1/api-integrations/{id}

Gets specified API integration object

Path Params
id
string
required
instance id of api integration object

Responses

200
Details of api integration object were found.

Response body
object
id
integer
required
authorizationScopes
array of strings
required
displayName
string
required
enabled
boolean
required
accessTokenLifetimeSeconds
integer
appType
string
enum
required
Type of API Client: * CLIENT_CREDENTIALS - A client ID and secret have been generated for this integration. * NATIVE_APP_OAUTH - A native app (i.e., Jamf Reset) has been linked to this integration for auth code grant type via Managed App Config. * NONE - No client is currently associated with this integration.

CLIENT_CREDENTIALS NATIVE_APP_OAUTH NONE

clientId
string
required

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/api-integrations/ \
     --header 'accept: application/json'

{
  "id": 1,
  "authorizationScopes": [
    "Tootsie Roal",
    "Jamf Reset"
  ],
  "displayName": "My API Integration",
  "enabled": true,
  "accessTokenLifetimeSeconds": 300,
  "appType": "CLIENT_CREDENTIALS",
  "clientId": "538878d4-9744-43ed-8732-3df99d502bd6"
}

-----

Update specified API integration object
put
https://yourServer.jamfcloud.com/api/v1/api-integrations/{id}

Update specified API integration object

Path Params
id
string
required
instance id of api integration object

Body Params
api object to update

authorizationScopes
array of strings
required
API Role display names.


ADD string
displayName
string
required
enabled
boolean

true
accessTokenLifetimeSeconds
integer
Responses

200
Api Integration updated

Response body
object
id
integer
required
authorizationScopes
array of strings
required
displayName
string
required
enabled
boolean
required
accessTokenLifetimeSeconds
integer
appType
string
enum
required
Type of API Client: * CLIENT_CREDENTIALS - A client ID and secret have been generated for this integration. * NATIVE_APP_OAUTH - A native app (i.e., Jamf Reset) has been linked to this integration for auth code grant type via Managed App Config. * NONE - No client is currently associated with this integration.

CLIENT_CREDENTIALS NATIVE_APP_OAUTH NONE

clientId
string
required

curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v1/api-integrations/ \
     --header 'accept: application/json' \
     --header 'content-type: application/json'

{
  "id": 1,
  "authorizationScopes": [
    "Tootsie Roal",
    "Jamf Reset"
  ],
  "displayName": "My API Integration",
  "enabled": true,
  "accessTokenLifetimeSeconds": 300,
  "appType": "CLIENT_CREDENTIALS",
  "clientId": "538878d4-9744-43ed-8732-3df99d502bd6"
}

-----

Remove specified API integration
delete
https://yourServer.jamfcloud.com/api/v1/api-integrations/{id}


Removes specified API integration

Path Params
id
string
required
instance id of api integration object

Responses

curl --request DELETE \
     --url https://yourserver.jamfcloud.com/api/v1/api-integrations/ \
     --header 'accept: application/json'

-----

Create client credentials for specified API integration
post
https://yourServer.jamfcloud.com/api/v1/api-integrations/{id}/client-credentials


Create client credentials for specified API integration

Path Params
id
string
required
instance id of api integration object

Responses

200
Client credentials have been created

Response body
object
clientId
string
required
clientSecret
string
required

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/api-integrations//client-credentials \
     --header 'accept: application/json'

{
  "clientId": "admin",
  "clientSecret": "12345"
}