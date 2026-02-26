Get information about all remote administration configurations.
get
https://yourServer.jamfcloud.com/api/preview/remote-administration-configurations


Remote administration feature creates a secure screen-sharing experience between Jamf Pro administrators and their end-users.

Query Params
page
integer
Defaults to 0
0
page-size
integer
Defaults to 100
100
Responses

200
Remote administration configurations returned.

Response body
object
totalCount
integer
results
array of objects
object
id
string
≥ 1
siteId
string
displayName
string
type
string
enum
team-viewer

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/preview/remote-administration-configurations?page=0&page-size=100' \
     --header 'accept: application/json'

{
  "totalCount": 1,
  "results": [
    {
      "id": "1",
      "siteId": "1",
      "displayName": "Remote administration",
      "type": "team-viewer"
    }
  ]
}