Get Advanced Search objects
get
https://yourServer.jamfcloud.com/api/v1/advanced-mobile-device-searches


Gets Advanced Search Objects

Response

200
Successful response - Advanced searches retrieved

Response body
object
totalCount
integer
≥ 0
results
array of objects
length ≥ 0
object
id
string
length ≥ 1
name
string
required
criteria
array of objects
object
name
string
required
priority
integer
andOr
string
required
searchType
string
required
value
string
required
openingParen
boolean
closingParen
boolean
displayFields
array of strings
siteId
string | null
Updated 12 days ago

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/advanced-mobile-device-searches \
     --header 'accept: application/json'

{
  "totalCount": 3,
  "results": [
    {
      "id": "1",
      "name": "Andy's Search",
      "criteria": [
        {
          "name": "Account",
          "priority": 0,
          "andOr": "and",
          "searchType": "is",
          "value": "test",
          "openingParen": false,
          "closingParen": false
        }
      ],
      "displayFields": [
        "AirPlay Password",
        "App Analytics Enabled"
      ],
      "siteId": "-1"
    }
  ]
}

-----

Create Advanced Search object
post
https://yourServer.jamfcloud.com/api/v1/advanced-mobile-device-searches


Creates Advanced Search Object

Body Params
name
string
required
criteria
array of objects

ADD object
displayFields
array of strings

ADD string
siteId
string | null
Response

201
Successful response - Advanced search created

Response body
object
id
string
length ≥ 1
href
string

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/advanced-mobile-device-searches \
     --header 'accept: application/json' \
     --header 'content-type: application/json'

{
  "id": "1",
  "href": "https://yourJamfProUrl.jamf/api/v1/resource/1"
}

-----

Get Mobile Device Advanced Search criteria choices
get
https://yourServer.jamfcloud.com/api/v1/advanced-mobile-device-searches/choices


Gets Mobile Device Advanced Search criteria choices. A list of potentially valid choices can be found by navigating to the Criteria page of the Advanced Mobile Device Search creation process. A few are "App Name", "Building", and "Display Name".

Query Params
criteria
string
required
site
string
Defaults to -1
-1
contains
string
Defaults to null
null
Response

200
Successful response - Criteria choices retrieved

Response body
object
choices
array of strings

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v1/advanced-mobile-device-searches/choices?site=-1&contains=null' \
     --header 'accept: application/json'

{
  "choices": [
    "Option 1",
    "Option 2"
  ]
}

-----

Remove specified Advanced Search objects
post
https://yourServer.jamfcloud.com/api/v1/advanced-mobile-device-searches/delete-multiple


Removes specified Advanced Search Objects

Body Params
ids of the building to be deleted

ids
array of strings

string


ADD string
Responses

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/advanced-mobile-device-searches/delete-multiple \
     --header 'accept: application/json' \
     --header 'content-type: application/json'

-----

Get specified Advanced Search object
get
https://yourServer.jamfcloud.com/api/v1/advanced-mobile-device-searches/{id}


Gets Specified Advanced Search Object

Path Params
id
string
required
id of target Advanced Search

Responses

200
Successful response - Advanced Search retrieved

Response body
object
id
string
length ≥ 1
name
string
required
criteria
array of objects
object
name
string
required
priority
integer
andOr
string
required
searchType
string
required
value
string
required
openingParen
boolean
closingParen
boolean
displayFields
array of strings
siteId
string | null

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/advanced-mobile-device-searches/ \
     --header 'accept: application/json'

{
  "id": "1",
  "name": "Andy's Search",
  "criteria": [
    {
      "name": "Account",
      "priority": 0,
      "andOr": "and",
      "searchType": "is",
      "value": "test",
      "openingParen": false,
      "closingParen": false
    }
  ],
  "displayFields": [
    "AirPlay Password",
    "App Analytics Enabled"
  ],
  "siteId": "-1"
}

-----

Get specified Advanced Search object
put
https://yourServer.jamfcloud.com/api/v1/advanced-mobile-device-searches/{id}


Gets Specified Advanced Search Object

Path Params
id
string
required
id of target Advanced Search

Body Params
name
string
required
criteria
array of objects

ADD object
displayFields
array of strings

ADD string
siteId
string | null
Responses

200
Successful response - Advanced Search updated

Response body
object
id
string
length ≥ 1
name
string
required
criteria
array of objects
object
name
string
required
priority
integer
andOr
string
required
searchType
string
required
value
string
required
openingParen
boolean
closingParen
boolean
displayFields
array of strings
siteId
string | null

curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v1/advanced-mobile-device-searches/ \
     --header 'accept: application/json' \
     --header 'content-type: application/json'

{
  "id": "1",
  "name": "Andy's Search",
  "criteria": [
    {
      "name": "Account",
      "priority": 0,
      "andOr": "and",
      "searchType": "is",
      "value": "test",
      "openingParen": false,
      "closingParen": false
    }
  ],
  "displayFields": [
    "AirPlay Password",
    "App Analytics Enabled"
  ],
  "siteId": "-1"
}

-----

Remove specified Advanced Search object
delete
https://yourServer.jamfcloud.com/api/v1/advanced-mobile-device-searches/{id}


Removes specified Advanced Search Object

Path Params
id
string
required
instance id of advanced search record

curl --request DELETE \
     --url https://yourserver.jamfcloud.com/api/v1/advanced-mobile-device-searches/ \
     --header 'accept: application/json'