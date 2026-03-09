
Retrieve an existing declaration
get
https://yourServer.jamfcloud.com/api/v1/dss-declarations/{declarationId}


Retrieves a stored declaration based on the provided declaration id

Path Params
declarationId
string
required
Declaration UUID

538F90D7-9383-4A2B-B1C8-81E845A9CFD7
Responses

200
The Declaration request was sent successfully.

Response body
object
declarations
array of objects
Defaults to
object
uuid
string
payloadJson
string | null
type
string | null
group
string | null
enum
ACTIVATION ASSET CONFIGURATION MANAGEMENT

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/dss-declarations/538F90D7-9383-4A2B-B1C8-81E845A9CFD7 \
     --header 'accept: application/json'

{
  "declarations": [
    {
      "uuid": "72676372-af55-432f-acd8-12984522e472",
      "payloadJson": {},
      "type": "com.apple.configuration.management.status-subscriptions",
      "group": "activation"
    }
  ]
}