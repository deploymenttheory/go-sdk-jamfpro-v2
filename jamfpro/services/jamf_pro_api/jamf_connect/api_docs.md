Get the Jamf Connect settings that you have access to see
get
https://yourServer.jamfcloud.com/api/v1/jamf-connect

Get the Jamf Connect settings that you have access to see.

Responses
204
Success, no content

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/jamf-connect \
     --header 'accept: application/json'
-----
Search for config profiles linked to Jamf Connect
get
https://yourServer.jamfcloud.com/api/v1/jamf-connect/config-profiles


Search for config profiles linked to Jamf Connect

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
profileId
string
profileName
string
profileScopeDescription
string
version
string
Must be a valid Jamf Connect version 2.3.0 or higher. Versions are listed here https://www.jamf.com/resources/product-documentation/jamf-connect-administrators-guide/

autoDeploymentType
string
enum
Defaults to NONE
Determines how the server will behave regarding application updates and installs on the devices that have the configuration profile installed. * PATCH_UPDATES - Server handles initial installation of the application and any patch updates. * MINOR_AND_PATCH_UPDATES - Server handles initial installation of the application and any patch and minor updates. * INITIAL_INSTALLATION_ONLY - Server only handles initial installation of the application. Updates will have to be done manually. * NONE - Server does not handle any installations or updates for the application. Version is ignored for this type.

PATCH_UPDATES MINOR_AND_PATCH_UPDATES INITIAL_INSTALLATION_ONLY NONE

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v1/jamf-connect/config-profiles?page=0&page-size=100' \
     --header 'accept: application/json'
{
  "totalCount": 13,
  "results": [
    {
      "uuid": "d265dfd3-8fde-4bf2-aa56-b167c8b68069",
      "profileId": "d265dfd3-8fde-4bf2-aa56-b167c8b68069",
      "profileName": "Best Plans of Mice",
      "profileScopeDescription": "Plan for all of the mouse computers at our org",
      "version": "2.3.0",
      "autoDeploymentType": "NONE"
    }
  ]
}
-----
Update the way the Jamf Connect app gets updated on computers within scope of the associated configuration profile.
put
https://yourServer.jamfcloud.com/api/v1/jamf-connect/config-profiles/{id}


Update the way the Jamf Connect app gets updated on computers within scope of the associated configuration profile.

Path Params
id
string
required
the UUID of the profile to update

24a7bb2a-9871-4895-9009-d1be07ed31b1
Body Params
Updatable Jamf Connect Settings

version
string
Must be a valid Jamf Connect version 2.3.0 or higher. Versions are listed here https://www.jamf.com/resources/product-documentation/jamf-connect-administrators-guide/

2.3.0
autoDeploymentType
string
enum
Defaults to NONE
Determines how the server will behave regarding application updates and installs on the devices that have the configuration profile installed. * PATCH_UPDATES - Server handles initial installation of the application and any patch updates. * MINOR_AND_PATCH_UPDATES - Server handles initial installation of the application and any patch and minor updates. * INITIAL_INSTALLATION_ONLY - Server only handles initial installation of the application. Updates will have to be done manually. * NONE - Server does not handle any installations or updates for the application. Version is ignored for this type.


INITIAL_INSTALLATION_ONLY
Allowed:

PATCH_UPDATES

MINOR_AND_PATCH_UPDATES

INITIAL_INSTALLATION_ONLY

NONE
Response

200
Success

Response body
object
uuid
string
profileId
string
profileName
string
profileScopeDescription
string
version
string
Must be a valid Jamf Connect version 2.3.0 or higher. Versions are listed here https://www.jamf.com/resources/product-documentation/jamf-connect-administrators-guide/

autoDeploymentType
string
enum
Defaults to NONE
Determines how the server will behave regarding application updates and installs on the devices that have the configuration profile installed. * PATCH_UPDATES - Server handles initial installation of the application and any patch updates. * MINOR_AND_PATCH_UPDATES - Server handles initial installation of the application and any patch and minor updates. * INITIAL_INSTALLATION_ONLY - Server only handles initial installation of the application. Updates will have to be done manually. * NONE - Server does not handle any installations or updates for the application. Version is ignored for this type.

PATCH_UPDATES MINOR_AND_PATCH_UPDATES INITIAL_INSTALLATION_ONLY NONE

curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v1/jamf-connect/config-profiles/24a7bb2a-9871-4895-9009-d1be07ed31b1 \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "autoDeploymentType": "INITIAL_INSTALLATION_ONLY",
  "version": "2.3.0"
}
'

{
  "uuid": "d265dfd3-8fde-4bf2-aa56-b167c8b68069",
  "profileId": "d265dfd3-8fde-4bf2-aa56-b167c8b68069",
  "profileName": "Best Plans of Mice",
  "profileScopeDescription": "Plan for all of the mouse computers at our org",
  "version": "2.3.0",
  "autoDeploymentType": "NONE"
}
-----
Search for deployment tasks for a config profile linked to Jamf Connect
get
https://yourServer.jamfcloud.com/api/v1/jamf-connect/deployments/{id}/tasks


Search for config profiles linked to Jamf Connect

Path Params
id
string
required
the UUID of the Jamf Connect deployment

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


string

1

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
     --url 'https://yourserver.jamfcloud.com/api/v1/jamf-connect/deployments/24a7bb2a-9871-4895-9009-d1be07ed31b1/tasks?page=0&page-size=100&sort=1' \
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
Request a retry of Connect install tasks
post
https://yourServer.jamfcloud.com/api/v1/jamf-connect/deployments/{id}/tasks/retry


Request a retry of Connect install tasks

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
     --url https://yourserver.jamfcloud.com/api/v1/jamf-connect/deployments/24a7bb2a-9871-4895-9009-d1be07ed31b1/tasks/retry \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '{"ids":["1"]}'
-----
Get Jamf Connect history
get
https://yourServer.jamfcloud.com/api/v1/jamf-connect/history


Get Jamf Connect history

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
Details of Jamf Connect history were found

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
     --url 'https://yourserver.jamfcloud.com/api/v1/jamf-connect/history?page=0&page-size=100' \
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
Add Jamf Connect history notes
post
https://yourServer.jamfcloud.com/api/v1/jamf-connect/history


Add Jamf Connect history notes

Body Params
history notes to create

note
string
required
A generic note can sometimes be useful, but generally not.
Responses

201
Notes to Jamf Connect history were added

Response body
object
id
string
length ≥ 1
href
string

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/jamf-connect/history \
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