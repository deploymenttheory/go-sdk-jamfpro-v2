Get log flushing settings
get
https://yourServer.jamfcloud.com/api/v1/log-flushing

Get all log flushing and retention policy settings

Response

200
Success

Response body
object
retentionPolicies
array of objects
object
displayName
string
qualifier
string
retentionPeriod
integer
retentionPeriodUnit
string
The unit of the retention period (eg: DAY, WEEK, MONTH, YEAR)

hourOfDay
integer
0 to 23
Defaults to 0

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/log-flushing \
     --header 'accept: application/json'

{
  "retentionPolicies": [
    {
      "displayName": "Policy Logs",
      "qualifier": "policy",
      "retentionPeriod": 3,
      "retentionPeriodUnit": "MONTH"
    }
  ],
  "hourOfDay": 0
}
-----
Get log flushing tasks
get
https://yourServer.jamfcloud.com/api/v1/log-flushing/task

Get a list of all log flushing tasks and their statuses

Response

200
Success

Response body
array of objects
object
id
string
The unique identifier of the log flushing task

qualifier
string
required
The qualifier of the retention policy

retentionPeriod
integer
required
The period beyond which data will be flushed

retentionPeriodUnit
string
required
The unit of the retention period (eg: DAY, WEEK, MONTH, YEAR)

state
string
The state of the task (eg: RUNNING, SUCCESS, FAILED, CANCELLED)

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/log-flushing/task \
     --header 'accept: application/json'

[
  {
    "id": "8c47858a-0029-450b-8720-cf5ae4bc27c7",
    "qualifier": "policy",
    "retentionPeriod": 3,
    "retentionPeriodUnit": "MONTH",
    "state": "SUCCESS"
  }
]
-----
Queue a log flushing task
post
https://yourServer.jamfcloud.com/api/v1/log-flushing/task

Queue a log flushing task

Body Params
The manual log flushing settings

qualifier
string
required
The qualifier of the retention policy

retentionPeriod
integer
required
The period beyond which data will be flushed

retentionPeriodUnit
string
required
The unit of the retention period (eg: DAY, WEEK, MONTH, YEAR)

Responses

202
The log flushing request was queued

Response body
object
id
string
length ≥ 1
href
string

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/log-flushing/task \
     --header 'accept: application/json' \
     --header 'content-type: application/json'

{
  "id": "1",
  "href": "https://yourJamfProUrl.jamf/api/v1/resource/1"
}
-----
Get log flushing task
get
https://yourServer.jamfcloud.com/api/v1/log-flushing/task/{id}

Get the log flushing task by the specified ID

Path Params
id
string
required
The identifier of the log flushing task

Response

200
Success

Response body
object
id
string
The unique identifier of the log flushing task

qualifier
string
required
The qualifier of the retention policy

retentionPeriod
integer
required
The period beyond which data will be flushed

retentionPeriodUnit
string
required
The unit of the retention period (eg: DAY, WEEK, MONTH, YEAR)

state
string
The state of the task (eg: RUNNING, SUCCESS, FAILED, CANCELLED)

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/log-flushing/task/ \
     --header 'accept: application/json'

{
  "id": "8c47858a-0029-450b-8720-cf5ae4bc27c7",
  "qualifier": "policy",
  "retentionPeriod": 3,
  "retentionPeriodUnit": "MONTH",
  "state": "SUCCESS"
}
-----
Cancels a log flushing task
delete
https://yourServer.jamfcloud.com/api/v1/log-flushing/task/{id}

Cancels a log flushing task by ID

Path Params
id
string
required
The identifier of the log flushing task

1
Responses
204
Log flushing task successfully canceled

curl --request DELETE \
     --url https://yourserver.jamfcloud.com/api/v1/log-flushing/task/1