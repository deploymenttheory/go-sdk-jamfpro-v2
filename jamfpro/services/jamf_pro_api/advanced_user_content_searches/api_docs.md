Get All Advanced User Content Search objects
get
https://yourServer.jamfcloud.com/api/v1/advanced-user-content-searches


Get All Advanced User Content Search Objects

Response

200
Successful response - Advanced User Content Searches retrieved

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

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/advanced-user-content-searches \
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
        "Content Name",
        "Price"
      ],
      "siteId": "-1"
    }
  ]
}

-----

Create Advanced User Content Search object
post
https://yourServer.jamfcloud.com/api/v1/advanced-user-content-searches


Creates Advanced User Content Search Object

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
Successful response - Advanced User Content Search created

Response body
object
id
string
length ≥ 1
href
string

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/advanced-user-content-searches \
     --header 'accept: application/json' \
     --header 'content-type: application/json'

{
  "id": "1",
  "href": "https://yourJamfProUrl.jamf/api/v1/resource/1"
}

-----

Get Specified Advanced User Content Search object
get
https://yourServer.jamfcloud.com/api/v1/advanced-user-content-searches/{id}


Gets Specified Advanced User Content Search Object

Path Params
id
string
required
id of target Advanced User Content Search

Responses

200
Successful response - Advanced User Content Search retrieved

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
     --url https://yourserver.jamfcloud.com/api/v1/advanced-user-content-searches/ \
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
    "Content Name",
    "Price"
  ],
  "siteId": "-1"
}

-----

Get Specified Advanced User Content Search object
put
https://yourServer.jamfcloud.com/api/v1/advanced-user-content-searches/{id}


Gets Specified Advanced User Content Search Object

Path Params
id
string
required
id of target Advanced User Content Search

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
Successful response - Advanced User Content Search updated

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
     --url https://yourserver.jamfcloud.com/api/v1/advanced-user-content-searches/ \
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
    "Content Name",
    "Price"
  ],
  "siteId": "-1"
}

-----

Remove specified Advanced User Content Search object
delete
https://yourServer.jamfcloud.com/api/v1/advanced-user-content-searches/{id}


Removes specified Advanced User Content Search Object

Path Params
id
string
required
instance id of Advanced User Content Search record

curl --request DELETE \
     --url https://yourserver.jamfcloud.com/api/v1/advanced-user-content-searches/ \
     --header 'accept: application/json'