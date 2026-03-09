Get the user preferences for the authenticated user and key.
get
https://yourServer.jamfcloud.com/api/v1/user/preferences/settings/{keyId}


Gets the user preferences for the authenticated user and key.

Path Params
keyId
string
required
user setting to be retrieved

1
Response

200
Successful response

Response body
object
username
string
required
key
string
required
values
array of strings
required
List of preferences for the specific key and user.

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/user/preferences/settings/1 \
     --header 'accept: application/json'

{
  "username": "jamfuser",
  "key": "mobileDeviceDisplayHeadings",
  "values": [
    "Device name"
  ]
}
-----
Get the user setting for the authenticated user and key
get
https://yourServer.jamfcloud.com/api/v1/user/preferences/{keyId}


Gets the user setting for the authenticated user and key.

Path Params
keyId
string
required
user setting to be retrieved

1
Response

200
Successful response

Response body
object
string

View Additional Properties

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/user/preferences/1 \
     --header 'accept: application/json'
-----
Persist the user setting
put
https://yourServer.jamfcloud.com/api/v1/user/preferences/{keyId}


Persists the user setting

Path Params
keyId
string
required
unique key of user setting to be persisted

Body Params
user setting value to be persisted

string
newKey
New Value
string
newKey-1
New Value

Add Field
Response

curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v1/user/preferences/ \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "newKey": "New Value",
  "newKey-1": "New Value"
}
'
-----
Remove specified setting for authenticated user
delete
https://yourServer.jamfcloud.com/api/v1/user/preferences/{keyId}


Remove specified setting for authenticated user

Path Params
keyId
string
required
unique key of user setting to be persisted

1

curl --request DELETE \
     --url https://yourserver.jamfcloud.com/api/v1/user/preferences/1