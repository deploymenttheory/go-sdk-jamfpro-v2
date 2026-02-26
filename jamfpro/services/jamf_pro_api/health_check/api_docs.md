Get Jamf Pro API status
get
https://yourServer.jamfcloud.com/api/v1/health-check

Get Jamf Pro API status. Which response codes might be returned in error states will depend on the specific state encountered.

Response
204
The Jamf Pro API is working properly.

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/health-check
-----
Retrieve request acceptance ratios for this Jamf Pro node
get
https://yourServer.jamfcloud.com/api/v1/health-status

Returns metrics representing the request acceptance ratio for each concurrency group and time window on this Jamf Pro node. The acceptance ratio is a decimal value between 0 and 1, where 1 means all requests were accepted and 0 means all were denied. Health status metrics are only available in Jamf Cloud. This API will return a 404 if the Jamf Pro node does not support health status metrics.

Responses

200
Success

Response body
object
api
object
required
thirtySeconds
number
Percentage of accepted requests out of total requests for the last 30 seconds

oneMinute
number
Percentage of accepted requests out of total requests for the last 1 minute

fiveMinutes
number
Percentage of accepted requests out of total requests for the last 5 minutes

fifteenMinutes
number
Percentage of accepted requests out of total requests for the last 15 minutes

thirtyMinutes
number
Percentage of accepted requests out of total requests for the last 30 minutes

ui
object
required
thirtySeconds
number
Percentage of accepted requests out of total requests for the last 30 seconds

oneMinute
number
Percentage of accepted requests out of total requests for the last 1 minute

fiveMinutes
number
Percentage of accepted requests out of total requests for the last 5 minutes

fifteenMinutes
number
Percentage of accepted requests out of total requests for the last 15 minutes

thirtyMinutes
number
Percentage of accepted requests out of total requests for the last 30 minutes

enrollment
object
required
thirtySeconds
number
Percentage of accepted requests out of total requests for the last 30 seconds

oneMinute
number
Percentage of accepted requests out of total requests for the last 1 minute

fiveMinutes
number
Percentage of accepted requests out of total requests for the last 5 minutes

fifteenMinutes
number
Percentage of accepted requests out of total requests for the last 15 minutes

thirtyMinutes
number
Percentage of accepted requests out of total requests for the last 30 minutes

device
object
required
thirtySeconds
number
Percentage of accepted requests out of total requests for the last 30 seconds

oneMinute
number
Percentage of accepted requests out of total requests for the last 1 minute

fiveMinutes
number
Percentage of accepted requests out of total requests for the last 5 minutes

fifteenMinutes
number
Percentage of accepted requests out of total requests for the last 15 minutes

thirtyMinutes
number
Percentage of accepted requests out of total requests for the last 30 minutes

default
object
required
thirtySeconds
number
Percentage of accepted requests out of total requests for the last 30 seconds

oneMinute
number
Percentage of accepted requests out of total requests for the last 1 minute

fiveMinutes
number
Percentage of accepted requests out of total requests for the last 5 minutes

fifteenMinutes
number
Percentage of accepted requests out of total requests for the last 15 minutes

thirtyMinutes
number
Percentage of accepted requests out of total requests for the last 30 minutes

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/health-status \
     --header 'accept: application/json'

{
  "api": {
    "thirtySeconds": 0.75,
    "oneMinute": 0.75,
    "fiveMinutes": 0.75,
    "fifteenMinutes": 0.75,
    "thirtyMinutes": 0.75
  },
  "ui": {
    "thirtySeconds": 0.75,
    "oneMinute": 0.75,
    "fiveMinutes": 0.75,
    "fifteenMinutes": 0.75,
    "thirtyMinutes": 0.75
  },
  "enrollment": {
    "thirtySeconds": 0.75,
    "oneMinute": 0.75,
    "fiveMinutes": 0.75,
    "fifteenMinutes": 0.75,
    "thirtyMinutes": 0.75
  },
  "device": {
    "thirtySeconds": 0.75,
    "oneMinute": 0.75,
    "fiveMinutes": 0.75,
    "fifteenMinutes": 0.75,
    "thirtyMinutes": 0.75
  },
  "default": {
    "thirtySeconds": 0.75,
    "oneMinute": 0.75,
    "fiveMinutes": 0.75,
    "fifteenMinutes": 0.75,
    "thirtyMinutes": 0.75
  }
}
-----