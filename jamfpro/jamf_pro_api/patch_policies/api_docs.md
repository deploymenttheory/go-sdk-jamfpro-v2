Retrieve Patch Policies
get
https://yourServer.jamfcloud.com/api/v2/patch-policies

Retrieves a list of patch policies.

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
Defaults to id:asc
Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma.


string

id:asc

ADD string
filter
string
Query in the RSQL format, allowing to filter Patch Policy collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: id, policyName, policyEnabled, policyTargetVersion, policyDeploymentMethod, softwareTitle, softwareTitleConfigurationId, pending, completed, deferred, and failed. This param can be combined with paging and sorting.

Responses

200
Success

Response body
object
totalCount
integer
results
array of objects
object
id
string
policyName
string
policyEnabled
boolean
policyTargetVersion
string
policyDeploymentMethod
string
softwareTitle
string
softwareTitleConfigurationId
string
pending
integer
completed
integer
deferred
integer
failed
integer

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v2/patch-policies?page=0&page-size=100&sort=id%3Aasc' \
     --header 'accept: application/json'

{
  "totalCount": 1,
  "results": [
    {
      "id": "1",
      "policyName": "Policy name",
      "policyEnabled": false,
      "policyTargetVersion": "v1",
      "policyDeploymentMethod": "automatically",
      "softwareTitle": "Software title",
      "softwareTitleConfigurationId": "1",
      "pending": 0,
      "completed": 0,
      "deferred": 0,
      "failed": 0
    }
  ]
}
-----
curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v2/patch-policies?page=0&page-size=100&sort=id%3Aasc' \
     --header 'accept: application/json'
{
  "totalCount": 1,
  "results": [
    {
      "id": "1",
      "policyName": "Policy name",
      "policyEnabled": false,
      "policyTargetVersion": "v1",
      "policyDeploymentMethod": "automatically",
      "softwareTitle": "Software title",
      "softwareTitleConfigurationId": "1",
      "pending": 0,
      "completed": 0,
      "deferred": 0,
      "failed": 0
    }
  ]
}
-----
Retrieve Patch Policies
get
https://yourServer.jamfcloud.com/api/v2/patch-policies/policy-details


Retrieves a list of patch policies.

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
Defaults to id:asc
Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma.


string

id:asc

ADD string
filter
string
Query in the RSQL format, allowing to filter Patch Policy collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: id, name, enabled, targetPatchVersion, deploymentMethod, softwareTitleId, softwareTitleConfigurationId, killAppsDelayMinutes, killAppsMessage, isDowngrade, isPatchUnknownVersion, notificationHeader, selfServiceEnforceDeadline, selfServiceDeadline, installButtonText, selfServiceDescription, iconId, reminderFrequency, reminderEnabled. This param can be combined with paging and sorting.

Responses

200
Success

Response body
object
totalCount
integer
results
array of objects
object
id
string
name
string
enabled
boolean
targetPatchVersion
string
deploymentMethod
string
softwareTitleId
string
softwareTitleConfigurationId
string
killAppsDelayMinutes
integer
killAppsMessage
string
downgrade
boolean
patchUnknownVersion
boolean
notificationHeader
string
selfServiceEnforceDeadline
boolean
selfServiceDeadline
integer
installButtonText
string
selfServiceDescription
string
iconId
string
reminderFrequency
integer
reminderEnabled
boolean

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v2/patch-policies/policy-details?page=0&page-size=100&sort=id%3Aasc' \
     --header 'accept: application/json'

{
  "totalCount": 1,
  "results": [
    {
      "id": "1",
      "name": "Policy name",
      "enabled": false,
      "targetPatchVersion": "v1",
      "deploymentMethod": "automatically",
      "softwareTitleId": "1",
      "softwareTitleConfigurationId": "1",
      "killAppsDelayMinutes": 5,
      "killAppsMessage": "message",
      "downgrade": false,
      "patchUnknownVersion": false,
      "notificationHeader": "notification header",
      "selfServiceEnforceDeadline": false,
      "selfServiceDeadline": 1,
      "installButtonText": "install",
      "selfServiceDescription": "description",
      "iconId": "1",
      "reminderFrequency": 1,
      "reminderEnabled": false
    }
  ]
}
-----
Return whether or not the requested patch policy is on the dashboard
get
https://yourServer.jamfcloud.com/api/v2/patch-policies/{id}/dashboard


Returns whether or not the requested patch policy is on the dashboard

Path Params
id
string
required
patch policy id

Responses

200
Whether the Patch Policy is on the Dashboard.

Response body
object
onDashboard
boolean

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v2/patch-policies//dashboard \
     --header 'accept: application/json'

{
  "onDashboard": true
}
-----
Add a patch policy to the dashboard
post
https://yourServer.jamfcloud.com/api/v2/patch-policies/{id}/dashboard


Adds a patch policy to the dashboard.

Path Params
id
string
required
patch policy id

Response
204
OK

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v2/patch-policies//dashboard
-----

Remove a patch policy from the dashboard
delete
https://yourServer.jamfcloud.com/api/v2/patch-policies/{id}/dashboard


Removes a patch policy from the dashboard.

Path Params
id
string
required
patch policy id

1
Response
204
OK

curl --request DELETE \
     --url https://yourserver.jamfcloud.com/api/v2/patch-policies/1/dashboard
-----
