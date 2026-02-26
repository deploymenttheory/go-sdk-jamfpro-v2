Get Jamf Pro Server URL settings
get
https://yourServer.jamfcloud.com/api/v1/jamf-pro-server-url

Get Jamf Pro Server URL settings

Response

200
Successful response

Response body
object
url
string
required

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/jamf-pro-server-url \
     --header 'accept: application/json'

{
  "url": "https://example.com:8443"
}
-----
Update Jamf Pro Server URL settings
put
https://yourServer.jamfcloud.com/api/v1/jamf-pro-server-url

Update Jamf Pro Server URL settings

Body Params
Jamf Pro Server URL settings object

url
string
required
https://example.com:8443
Responses

200
Jamf Pro Server URL settings updated

Response body
object
url
string
required

curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v1/jamf-pro-server-url \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "url": "https://example.com:8443"
}
'

{
  "url": "https://example.com:8443"
}
-----
Get Jamf Pro Server URL settings history
get
https://yourServer.jamfcloud.com/api/v1/jamf-pro-server-url/history


Gets Jamf Pro Server URL settings history

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
string
Defaults to date:desc
Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc

date:desc
Show Deprecated
Response

200
Details of Jamf Pro Server URL settings history were found

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
integer
≥ 1
username
string
date
string
note
string
details
string | null

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v1/jamf-pro-server-url/history?page=0&page-size=100&sort=date%3Adesc' \
     --header 'accept: application/json'

{
  "totalCount": 1,
  "results": [
    {
      "id": 1,
      "username": "admin",
      "date": "2019-02-04T21:09:31.661Z",
      "note": "Sso settings update",
      "details": "Is SSO Enabled false\\nSelected SSO Provider"
    }
  ]
}
-----
Add Jamf Pro Server URL settings history notes
post
https://yourServer.jamfcloud.com/api/v1/jamf-pro-server-url/history


Adds Jamf Pro Server URL settings history notes

Body Params
History notes to create

note
string
required
A generic note can sometimes be useful, but generally not.
Responses

201
Notes to Jamf Pro Server URL settings history were added

Response body
object
id
integer
≥ 1
username
string
date
string
note
string
details
string | null

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/jamf-pro-server-url/history \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "note": "A generic note can sometimes be useful, but generally not."
}
'
{
  "id": 1,
  "username": "admin",
  "date": "2019-02-04T21:09:31.661Z",
  "note": "Sso settings update",
  "details": "Is SSO Enabled false\\nSelected SSO Provider"
}
-----