Retrieve the current failover settings
get
https://yourServer.jamfcloud.com/api/v1/sso/failover

Retrieve the current failover settings

Responses

200
Successful response

Response body
object
failoverUrl
string
generationTime
int64
≥ 0
Generation time of failover key

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/sso/failover \
     --header 'accept: application/json'

{
  "failoverUrl": "https://jamf.jamfcloud.com/?failover=0123456789ABCDEF",
  "generationTime": 1674133253000
}
-----
Regenerates failover url
post
https://yourServer.jamfcloud.com/api/v1/sso/failover/generate

Regenerates failover url, by changing failover key to new one, and returns new failover settings

Response

200
The generation was successful and generated failover settings are returned.

Response body
object
failoverUrl
string
generationTime
int64
≥ 0
Generation time of failover key

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/sso/failover/generate \
     --header 'accept: application/json'

{
  "failoverUrl": "https://jamf.jamfcloud.com/?failover=0123456789ABCDEF",
  "generationTime": 1674133253000
}