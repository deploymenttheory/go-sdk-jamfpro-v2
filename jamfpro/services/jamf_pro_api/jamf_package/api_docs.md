Get the packages for a given Jamf application
get
https://yourServer.jamfcloud.com/api/v1/jamf-package

Get the packages for a given Jamf application.

Query Params
application
string
required
The Jamf Application key. The only supported values are protect and connect.

protect
Responses

200
List of packages for the given application.

Response body
array of objects
object
id
string
filename
string
version
string
created
string
url
string

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v1/jamf-package?application=protect' \
     --header 'accept: application/json'

[
  {
    "id": "B84F9904-868B-4862-A8F5-33706AADED71",
    "filename": "JamfProtect-1.2.2.pkg",
    "version": "1.2.2",
    "created": "1970-01-01T00:00:00Z",
    "url": "https://example.jamf.com/path/to/JamfProtect-1.2.2.pkg"
  }
]
-----
Get the packages for a given Jamf application
get
https://yourServer.jamfcloud.com/api/v2/jamf-package

Get the packages for a given Jamf application.

Query Params
application
string
required
The Jamf Application key. The only supported values are protect and connect.

protect
Responses

200
Properties for the given application.

Response body
object
displayName
string
releaseHistoryUrl
string
artifacts
array of objects
object
id
string
filename
string
version
string
created
string
url
string

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v2/jamf-package?application=protect' \
     --header 'accept: application/json'

{
  "displayName": "Jamf Connect",
  "releaseHistoryUrl": "https://docs.jamf.com/jamf-connect/administrator-guide/Release_History.html",
  "artifacts": [
    {
      "id": "B84F9904-868B-4862-A8F5-33706AADED71",
      "filename": "JamfProtect-1.2.2.pkg",
      "version": "1.2.2",
      "created": "1970-01-01T00:00:00Z",
      "url": "https://example.jamf.com/path/to/JamfProtect-1.2.2.pkg"
    }
  ]
}