Get Policy Properties object
get
https://yourServer.jamfcloud.com/api/v1/policy-properties

Gets Policy Properties object.

Response

200
Successful response

Response body
object
policiesRequireNetworkStateChange
boolean
Defaults to false
This field always returns false.

allowNetworkStateChangeTriggers
boolean
Defaults to true

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/policy-properties \
     --header 'accept: application/json'

{
  "policiesRequireNetworkStateChange": false,
  "allowNetworkStateChangeTriggers": true
}
-----
Update Policy Properties object
put
https://yourServer.jamfcloud.com/api/v1/policy-properties

Update Policy Properties object

Body Params
Policy Properties object to update

policiesRequireNetworkStateChange
boolean
Defaults to false
This field always returns false.


false
allowNetworkStateChangeTriggers
boolean
Defaults to true

true
Response

200
Policy Properties was updated

Response body
object
policiesRequireNetworkStateChange
boolean
Defaults to false
This field always returns false.

allowNetworkStateChangeTriggers
boolean
Defaults to true

curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v1/policy-properties \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "policiesRequireNetworkStateChange": false,
  "allowNetworkStateChangeTriggers": true
}
'
{
  "policiesRequireNetworkStateChange": false,
  "allowNetworkStateChangeTriggers": true
}
-----