Get Client Check-In settings
get
https://yourServer.jamfcloud.com/api/v3/check-in

Gets Client Check-In object.

Response

200
Successful response

Response body
object
checkInFrequency
int32
Defaults to 15
Suggested values are 5, 15, 30, or 60. Web interface will not display correctly if not one of those. Minimum is 5, maximum is 60.

createHooks
boolean
Defaults to false
hookLog
boolean
Defaults to false
hookPolicies
boolean
Defaults to false
createStartupScript
boolean
Defaults to false
startupLog
boolean
Defaults to false
startupPolicies
boolean
Defaults to false
startupSsh
boolean
Defaults to false
enableLocalConfigurationProfiles
boolean
Defaults to false

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v3/check-in \
     --header 'accept: application/json'

{
  "checkInFrequency": 15,
  "createHooks": false,
  "hookLog": false,
  "hookPolicies": false,
  "createStartupScript": false,
  "startupLog": false,
  "startupPolicies": false,
  "startupSsh": false,
  "enableLocalConfigurationProfiles": false
}
-----

Update Client Check-In object
put
https://yourServer.jamfcloud.com/api/v3/check-in

Update Client Check-In object

Body Params
Client Check-In object to update

checkInFrequency
int32
Defaults to 15
Suggested values are 5, 15, 30, or 60. Web interface will not display correctly if not one of those. Minimum is 5, maximum is 60.

15
createHooks
boolean
Defaults to false

false
hookLog
boolean
Defaults to false

false
hookPolicies
boolean
Defaults to false

false
createStartupScript
boolean
Defaults to false

false
startupLog
boolean
Defaults to false

false
startupPolicies
boolean
Defaults to false

false
startupSsh
boolean
Defaults to false

false
enableLocalConfigurationProfiles
boolean
Defaults to false

false
Response

200
Client Check-In was updated

Response body
object
checkInFrequency
int32
Defaults to 15
Suggested values are 5, 15, 30, or 60. Web interface will not display correctly if not one of those. Minimum is 5, maximum is 60.

createHooks
boolean
Defaults to false
hookLog
boolean
Defaults to false
hookPolicies
boolean
Defaults to false
createStartupScript
boolean
Defaults to false
startupLog
boolean
Defaults to false
startupPolicies
boolean
Defaults to false
startupSsh
boolean
Defaults to false
enableLocalConfigurationProfiles
boolean
Defaults to false

curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v3/check-in \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "checkInFrequency": 15,
  "createHooks": false,
  "hookLog": false,
  "hookPolicies": false,
  "createStartupScript": false,
  "startupLog": false,
  "startupPolicies": false,
  "startupSsh": false,
  "enableLocalConfigurationProfiles": false
}
'

{
  "checkInFrequency": 15,
  "createHooks": false,
  "hookLog": false,
  "hookPolicies": false,
  "createStartupScript": false,
  "startupLog": false,
  "startupPolicies": false,
  "startupSsh": false,
  "enableLocalConfigurationProfiles": false
}

-----

Get Client Check-In history object
get
https://yourServer.jamfcloud.com/api/v3/check-in/history

Gets Client Check-In history object

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
Defaults to date:desc
Sorting criteria in the format: property:asc/desc. Default sort is name:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,username:asc


string

date:desc

ADD string
filter
string
Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: username, date, note, details. This param can be combined with paging and sorting. Example: filter=username!=admin and details==disabled and date<2019-12-15

Response

200
Details of Client Check-In history were found

Response body
object
totalCount
integer
≥ 0
results
array of objects
object
id
string
length ≥ 1
username
string
date
string
note
string
details
string | null

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v3/check-in/history?page=0&page-size=100&sort=date%3Adesc' \
     --header 'accept: application/json'

{
  "totalCount": 1,
  "results": [
    {
      "id": "1",
      "username": "admin",
      "date": "2019-02-04T21:09:31.661Z",
      "note": "Sso settings update",
      "details": "Is SSO Enabled false\\nSelected SSO Provider"
    }
  ]
}

-----

Add a Note to Client Check-In History
post
https://yourServer.jamfcloud.com/api/v3/check-in/history

Adds Client Check-In history object notes

Body Params
history notes to create

note
string
required
Responses

201
Notes of Client Check-In history were added

Response body
object
id
string
length ≥ 1
href
string

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v3/check-in/history \
     --header 'accept: application/json' \
     --header 'content-type: application/json'

{
  "id": "1",
  "href": "https://yourJamfProUrl.jamf/api/v1/resource/1"
}