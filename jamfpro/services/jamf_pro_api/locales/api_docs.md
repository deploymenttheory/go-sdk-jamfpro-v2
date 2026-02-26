Return locales that can be used in other features
get
https://yourServer.jamfcloud.com/api/v1/locales

Returns locales that can be used in other features.

Response

200
Successful response

Response body
array of objects
object
description
string
identifier
string

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/locales \
     --header 'accept: application/json'

[
  {
    "description": "English (United States)",
    "identifier": "en_US"
  }
]
-----