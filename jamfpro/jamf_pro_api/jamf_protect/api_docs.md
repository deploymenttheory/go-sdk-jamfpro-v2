Jamf Protect integration settings
get
https://yourServer.jamfcloud.com/api/v1/jamf-protect

Jamf Protect integration settings

Responses

200
Success.

Response body
object
id
string
apiClientId
string
apiClientName
string
display name used when creating the API Client in the Jamf Protect web console

registrationId
string
ID used when making requests to identify this particular Protect registration.

protectUrl
string
platformPlanSync
boolean
determines whether Protect Platform Plan syncing is enabled

lastSyncTime
string
syncStatus
string
enum
IN_PROGRESS COMPLETED ERROR UNKNOWN

autoInstall
boolean
determines whether the Jamf Protect agent will be automatically installed on client computers

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/jamf-protect \
     --header 'accept: application/json'

{
  "id": "1",
  "apiClientId": "esgzYzYBqN7wCImwyusbQcXob2qalGMN",
  "apiClientName": "Jamf Pro integration",
  "registrationId": "6f250316-2cfb-4521-8cb7-bfaf46497bc5",
  "protectUrl": "https://examplejamfprotect.jamfcloud.com/graphql",
  "platformPlanSync": true,
  "lastSyncTime": "2003-01-05T18:00:14.885Z",
  "syncStatus": "COMPLETED",
  "autoInstall": true
}
-----
Jamf Protect integration settings
put
https://yourServer.jamfcloud.com/api/v1/jamf-protect

Jamf Protect integration settings

Body Params
Updatable Jamf Protect Settings

autoInstall
boolean
determines whether the Jamf Protect agent will be automatically installed on client computers


true
Response

200
Success.

Response body
object
id
string
apiClientId
string
apiClientName
string
display name used when creating the API Client in the Jamf Protect web console

registrationId
string
ID used when making requests to identify this particular Protect registration.

protectUrl
string
platformPlanSync
boolean
determines whether Protect Platform Plan syncing is enabled

lastSyncTime
string
syncStatus
string
enum
IN_PROGRESS COMPLETED ERROR UNKNOWN

autoInstall
boolean
determines whether the Jamf Protect agent will be automatically installed on client computers

curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v1/jamf-protect \
     --header 'accept: application/json' \
     --header 'content-type: application/json'

{
  "id": "1",
  "apiClientId": "esgzYzYBqN7wCImwyusbQcXob2qalGMN",
  "apiClientName": "Jamf Pro integration",
  "registrationId": "6f250316-2cfb-4521-8cb7-bfaf46497bc5",
  "protectUrl": "https://examplejamfprotect.jamfcloud.com/graphql",
  "platformPlanSync": true,
  "lastSyncTime": "2003-01-05T18:00:14.885Z",
  "syncStatus": "COMPLETED",
  "autoInstall": true
}
-----
Delete Jamf Protect API registration.
delete
https://yourServer.jamfcloud.com/api/v1/jamf-protect

Deletes an existing Jamf Protect API registration if present. Jamf Protect API integration will be disabled.

Responses
204
Success, No Content

curl --request DELETE \
     --url https://yourserver.jamfcloud.com/api/v1/jamf-protect \
     --header 'accept: application/json'

-----
Search for deployment tasks for a config profile linked to Jamf Protect
get
https://yourServer.jamfcloud.com/api/v1/jamf-protect/deployments/{id}/tasks


Search for config profiles linked to Jamf Protect

Path Params
id
string
required
the UUID of the Jamf Protect deployment

24a7bb2a-9871-4895-9009-d1be07ed31b1
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
Defaults to
Sorting criteria in the format: property:asc/desc. Default sort order is descending. Multiple sort criteria are supported and must be entered on separate lines in Swagger UI. In the URI the 'sort' query param is not duplicated for each sort criterion, e.g., ...&sort=name:asc,date:desc. Fields that can be sorted: status, updated


ADD string
filter
string
Query in the RSQL format, allowing to filter results. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: status, updated, version This param can be combined with paging and sorting. Example: filter=username!=admin and details==disabled and date<2019-12-15

Response

200
Success.

Response body
object
totalCount
integer
results
array of objects
object
id
string
computerId
string
computerName
string
version
string
updated
date-time
status
string
enum
Status of this Jamf Connect deployment task. "Command" below refers to an InstallEnterpriseApplication command. Tasks that are not finished (i.e., COMPLETE or GAVE_UP) are evaluated once every thirty minutes, so the status value for a device may lag behind a successful Jamf Connect package install up to thirty minutes. * COMMAND_QUEUED - command has been queued * NO_COMMAND - command has not yet been queued * PENDING_MANIFEST - task is waiting to obtain a valid package manifest before a command can be queued * COMPLETE - command has been completed successfully * GAVE_UP - the command failed with an error or the device did not process it in a reasonable amount of time * UNKNOWN - unknown; tasks in this state will be evaluated

COMMAND_QUEUED NO_COMMAND PENDING_MANIFEST COMPLETE GAVE_UP UNKNOWN

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v1/jamf-protect/deployments/24a7bb2a-9871-4895-9009-d1be07ed31b1/tasks?page=0&page-size=100' \
     --header 'accept: application/json'

{
  "totalCount": 13,
  "results": [
    {
      "id": "82",
      "computerId": "111",
      "computerName": "Polka dot Stratocaster",
      "version": "2.3.4",
      "updated": "2021-05-07T21:20:34.35Z",
      "status": "COMPLETE"
    }
  ]
}
-----
Request a retry of Protect install tasks
post
https://yourServer.jamfcloud.com/api/v1/jamf-protect/deployments/{id}/tasks/retry


Request a retry of Protect install tasks

Path Params
id
string
required
the UUID of the deployment associated with the retry

24a7bb2a-9871-4895-9009-d1be07ed31b1
Body Params
task IDs to retry

ids
array of strings

string

1

ADD string
Responses
204
Success, no content.

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/jamf-protect/deployments/24a7bb2a-9871-4895-9009-d1be07ed31b1/tasks/retry \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '{"ids":["1"]}'
-----
Get Jamf Protect history
get
https://yourServer.jamfcloud.com/api/v1/jamf-protect/history


Get Jamf Protect history

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
Defaults to
Sorting criteria in the format: property:asc/desc. Default sort order is descending. Multiple sort criteria are supported and must be entered on separate lines in Swagger UI. In the URI the 'sort' query param is not duplicated for each sort criterion, e.g., ...&sort=name:asc,date:desc. Fields that can be sorted: status, updated


ADD string
filter
string
Query in the RSQL format, allowing to filter results. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: status, updated, version This param can be combined with paging and sorting. Example: filter=username!=admin and details==disabled and date<2019-12-15

Response

200
Details of Jamf Protect history were found

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
     --url 'https://yourserver.jamfcloud.com/api/v1/jamf-protect/history?page=0&page-size=100' \
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
Add Jamf Protect history notes
post
https://yourServer.jamfcloud.com/api/v1/jamf-protect/history


Add Jamf Protect history notes

Body Params
history notes to create

note
string
required
A generic note can sometimes be useful, but generally not.
Responses

201
Notes to Jamf Protect history were added

Response body
object
id
string
length ≥ 1
href
string

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/jamf-protect/history \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "note": "A generic note can sometimes be useful, but generally not."
}
'
{
  "id": "1",
  "href": "https://yourJamfProUrl.jamf/api/v1/resource/1"
}
-----
Get all of the previously synced Jamf Protect Plans with information about their associated configuration profile
get
https://yourServer.jamfcloud.com/api/v1/jamf-protect/plans

Get all of the previously synced Jamf Protect Plans with information about their associated configuration profile

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
Defaults to
Sorting criteria in the format: property:asc/desc. Default sort order is descending. Multiple sort criteria are supported and must be entered on separate lines in Swagger UI. In the URI the 'sort' query param is not duplicated for each sort criterion, e.g., ...&sort=name:asc,date:desc. Fields that can be sorted: status, updated


ADD string
filter
string
Query in the RSQL format, allowing to filter results. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: status, updated, version This param can be combined with paging and sorting. Example: filter=username!=admin and details==disabled and date<2019-12-15

Response

200
Success.

Response body
object
totalCount
integer
results
array of objects
object
uuid
string
id
string
name
string
description
string
profileId
integer
profileName
string
profileVersion
integer
scopeDescription
string
siteId
string

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v1/jamf-protect/plans?page=0&page-size=100' \
     --header 'accept: application/json'

{
  "totalCount": 3,
  "results": [
    {
      "uuid": "b999dfd3-8fde-4bf2-aa56-b167c8b68071",
      "id": "d265dfd3-8fde-4bf2-aa56-b167c8b68069",
      "name": "Main Plan",
      "description": "Plan for the majority of the company's computers",
      "profileId": 12,
      "profileName": "Main Plan (Jamf Protect)",
      "profileVersion": 1,
      "scopeDescription": "All Computers",
      "siteId": "-1"
    }
  ]
}
-----
Sync Plans with Jamf Protect
post
https://yourServer.jamfcloud.com/api/v1/jamf-protect/plans/sync


Sync Plans with Jamf Protect. Configuration profiles associated with new plans will be imported to Jamf Pro.

Responses
204
Success.

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/jamf-protect/plans/sync \
     --header 'accept: application/json'
-----

Register a Jamf Protect API configuration with Jamf Pro
post
https://yourServer.jamfcloud.com/api/v1/jamf-protect/register


Register a Jamf Protect API configuration with Jamf Pro

Body Params
Jamf Protect API connection information

protectUrl
string
required
clientId
string
required
password
password
required
Responses

201
Successful registration.

Response body
object
id
string
apiClientId
string
apiClientName
string
display name used when creating the API Client in the Jamf Protect web console

registrationId
string
ID used when making requests to identify this particular Protect registration.

protectUrl
string
platformPlanSync
boolean
determines whether Protect Platform Plan syncing is enabled

lastSyncTime
string
syncStatus
string
enum
IN_PROGRESS COMPLETED ERROR UNKNOWN

autoInstall
boolean
determines whether the Jamf Protect agent will be automatically installed on client computers

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/jamf-protect/register \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "protectUrl": "https://examplejamfprotect.jamfcloud.com/graphql",
  "clientId": "uzPJXlArmzTAmPRQtZEnQ2OFtNw8qQV",
  "password": "7fyP6BphUUQ5B_zoLrkYhM5j1HTcf-4PxshettZbK0ZcnzV57gyHwF23U3F96F"
}
'

{
  "id": "1",
  "apiClientId": "esgzYzYBqN7wCImwyusbQcXob2qalGMN",
  "apiClientName": "Jamf Pro integration",
  "registrationId": "6f250316-2cfb-4521-8cb7-bfaf46497bc5",
  "protectUrl": "https://examplejamfprotect.jamfcloud.com/graphql",
  "platformPlanSync": true,
  "lastSyncTime": "2003-01-05T18:00:14.885Z",
  "syncStatus": "COMPLETED",
  "autoInstall": true
}
-----