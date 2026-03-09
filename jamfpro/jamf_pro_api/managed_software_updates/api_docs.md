Retrieve available macOS and iOS Managed Software Updates
get
https://yourServer.jamfcloud.com/api/v1/managed-software-updates/available-updates

Retrieves available macOS and iOS Managed Software Updates

Response

200
Success

Response body
object
availableUpdates
object
macOS
array of strings
iOS
array of strings

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/managed-software-updates/available-updates \
     --header 'accept: application/json'

{
  "availableUpdates": {
    "macOS": [
      "12.0.1",
      "11.6.1",
      "11.6",
      "11.5.2"
    ],
    "iOS": [
      "16.1.1",
      "16.0.1",
      "15.7.1",
      "9.1"
    ]
  }
}
-----
Retrieve Managed Software Update Plans
get
https://yourServer.jamfcloud.com/api/v1/managed-software-updates/plans

Retrieve Managed Software Update Plans

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
Defaults to planUuid:asc
Sorting criteria in the format: property:asc/desc. Default sort is planUuid:asc. Multiple sort criteria are supported and must be separated with a comma.


string

planUuid:asc

ADD string
filter
string
Query in the RSQL format, allowing to filter Managed Software Updates collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: planUuid, device.deviceId, device.objectType, updateAction, versionType, specificVersion, maxDeferrals, recipeId, forceInstallLocalDateTime, state.

Responses

200
Success

Response body
object
totalCount
integer
≥ 0
results
array of objects
length ≥ 0
object
planUuid
string
required
length ≥ 1
device
object
required

device object
updateAction
string
enum
required
DOWNLOAD_ONLY DOWNLOAD_INSTALL DOWNLOAD_INSTALL_ALLOW_DEFERRAL DOWNLOAD_INSTALL_RESTART DOWNLOAD_INSTALL_SCHEDULE UNKNOWN

versionType
string
enum
required
LATEST_MAJOR LATEST_MINOR LATEST_ANY SPECIFIC_VERSION CUSTOM_VERSION UNKNOWN

specificVersion
string
Defaults to NO_SPECIFIC_VERSION
Optional. Indicates the specific version to update to. Only available when the version type is set to specific version or custom version, otherwise defaults to NO_SPECIFIC_VERSION.

buildVersion
string | null
Optional. Indicates the build version to update to. Only available when the version type is set to custom version.

maxDeferrals
integer
required
≥ 0
Not applicable to all managed software update plans

forceInstallLocalDateTime
string | null
Optional. Indicates the local date and time of the device to force update by.

recipeId
string
Defaults to -1
The id of the recipe that was used to generate the plan.

status
object
required

status object
state
string
enum
Init PendingPlanValidation AcceptingPlan RejectingPlan ProcessingPlanType ProcessingPlanTypeMdm StartingPlan PlanFailed SchedulingScanForOSUpdates ProcessingScheduleOSUpdateScanResponse WaitingForScheduledOSUpdateScanToComplete CollectingAvailableOSUpdates ProcessingAvailableOSUpdatesResponse ProcessingSchedulingType SchedulingDDM DDMPlanScheduled WaitingToStartDDMUpdate ProcessingDDMStatusResponse CollectingDDMStatus SchedulingMDM MDMPlanScheduled SchedulingOSUpdate ProcessingScheduleOSUpdateResponse CollectingOSUpdateStatus ProcessingOSUpdateStatusResponse WaitingToCollectOSUpdateStatus VerifyingInstallation ProcessingInstallationVerification PlanCompleted PlanCanceled PlanException Unknown

errorReasons
array of strings | null

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v1/managed-software-updates/plans?page=0&page-size=100&sort=planUuid%3Aasc' \
     --header 'accept: application/json'

{
  "totalCount": 1,
  "results": [
    {
      "planUuid": "6E47EF55-5318-494F-A09E-70F613E0AFD1",
      "device": {
        "deviceId": "1",
        "objectType": "COMPUTER",
        "href": "/v1/computers-inventory/1"
      },
      "updateAction": "DOWNLOAD_INSTALL",
      "versionType": "SPECIFIC_VERSION",
      "specificVersion": "12.6.1",
      "buildVersion": "21F79",
      "maxDeferrals": 5,
      "forceInstallLocalDateTime": "2023-12-25T21:09:31",
      "recipeId": "1",
      "status": {
        "state": "SchedulingScanForOSUpdates",
        "errorReasons": [
          "NO_UPDATES_AVAILABLE",
          "NOT_SUPERVISED"
        ]
      }
    }
  ]
}
-----
Create a Managed Software Update Plan
post
https://yourServer.jamfcloud.com/api/v1/managed-software-updates/plans

Creates a Managed Software Update Plan.

Body Params
Managed Software Update Plan to create

devices
array of objects
required
length ≥ 1

object

deviceId
string
required
length ≥ 1
1
objectType
string
enum
required

COMPUTER
Allowed:

COMPUTER

MOBILE_DEVICE

APPLE_TV

object

deviceId
string
required
length ≥ 1
1
objectType
string
enum
required

COMPUTER
Allowed:

COMPUTER

MOBILE_DEVICE

APPLE_TV

ADD object
config
object
required

config object
updateAction
string
enum
required

DOWNLOAD_ONLY
Allowed:

DOWNLOAD_ONLY

DOWNLOAD_INSTALL

DOWNLOAD_INSTALL_ALLOW_DEFERRAL

DOWNLOAD_INSTALL_RESTART

DOWNLOAD_INSTALL_SCHEDULE

UNKNOWN
versionType
string
enum
required

LATEST_MAJOR
Allowed:

LATEST_MAJOR

LATEST_MINOR

LATEST_ANY

SPECIFIC_VERSION

CUSTOM_VERSION

UNKNOWN
specificVersion
string
length ≥ 0
Defaults to NO_SPECIFIC_VERSION
Optional. Indicates the specific version to update to. Only available when the version type is set to specific version or custom version, otherwise defaults to NO_SPECIFIC_VERSION.

NO_SPECIFIC_VERSION
buildVersion
string | null
Optional. Indicates the build version to update to. Only available when the version type is set to custom version.

21F79
maxDeferrals
integer
≥ 0
Required when the provided updateAction is DOWNLOAD_INSTALL_ALLOW_DEFERRAL, not applicable to all managed software update plans

5
forceInstallLocalDateTime
string | null
Optional. Indicates the local date and time of the device to force update by.

2023-12-25T21:09:31
Responses

201
Managed Software Update Plan was created

Response body
object
plans
array of objects
length ≥ 0
object
device
object

device object
planId
string
length ≥ 1
href
string

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/managed-software-updates/plans \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "devices": [
    {
      "objectType": "COMPUTER",
      "deviceId": "1"
    },
    {
      "objectType": "COMPUTER",
      "deviceId": "1"
    }
  ],
  "config": {
    "updateAction": "DOWNLOAD_ONLY",
    "versionType": "LATEST_MAJOR",
    "specificVersion": "NO_SPECIFIC_VERSION",
    "buildVersion": "21F79",
    "maxDeferrals": 5,
    "forceInstallLocalDateTime": "2023-12-25T21:09:31"
  }
}
'

{
  "plans": [
    {
      "device": {
        "deviceId": "1",
        "objectType": "COMPUTER",
        "href": "/v1/computers-inventory/1"
      },
      "planId": "6E47EF55-5318-494F-A09E-70F613E0AFD1",
      "href": "/v1/managed-software-updates/plans/6E47EF55-5318-494F-A09E-70F613E0AFD1"
    }
  ]
}
-----
Retrieve current value of the Feature Toggle
get
https://yourServer.jamfcloud.com/api/v1/managed-software-updates/plans/feature-toggle

Retrieves current value of the Feature Toggle

Response

200
Success

Response body
object
toggle
boolean
required
forceInstallLocalDateEnabled
boolean
customVersionEnabled
boolean
dssEnabled
boolean
recipeEnabled
boolean

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/managed-software-updates/plans/feature-toggle \
     --header 'accept: application/json'

{
  "toggle": false,
  "forceInstallLocalDateEnabled": false,
  "customVersionEnabled": false,
  "dssEnabled": false,
  "recipeEnabled": false
}
-----
Updates Feature Toggle Value
put
https://yourServer.jamfcloud.com/api/v1/managed-software-updates/plans/feature-toggle

Updates the value of the Feature Toggle - This endpoint is asynchronous, the provided value will not be immediately updated. Please use the following endpoint to track the status of your toggle request. /v1/managed-software-updates/plans/feature-toggle/status:

Body Params
toggle
boolean
required

true
Responses

200
Managed Software Update Plan Feature Toggle was updated

Response body
object
toggle
boolean
required
forceInstallLocalDateEnabled
boolean
customVersionEnabled
boolean
dssEnabled
boolean
recipeEnabled
boolean

curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v1/managed-software-updates/plans/feature-toggle \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '{"toggle":true}'

{
  "toggle": false,
  "forceInstallLocalDateEnabled": false,
  "customVersionEnabled": false,
  "dssEnabled": false,
  "recipeEnabled": false
}
-----
Force stops any ongoing or stalled feature-toggle processes
post
https://yourServer.jamfcloud.com/api/v1/managed-software-updates/plans/feature-toggle/abandon

"Break Glass" endpoint, not for nominal usage. Use this endpoint to forcefully abandon the feature-toggle background process if the status of the feature-toggle is 'stuck' or has reached an non-restartable failed state. Usage of this endpoint under nominal conditions is undefined and unsupported.

Responses
204
Managed Software Update Plan Feature Toggle abandon request was received

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/managed-software-updates/plans/feature-toggle/abandon \
     --header 'accept: application/json'

Retrieves background status of the Feature Toggle
get
https://yourServer.jamfcloud.com/api/v1/managed-software-updates/plans/feature-toggle/status

Retrieves background status of the Feature Toggle

Responses

200
Managed Software Update Plan Feature Toggle Status was retrieved

Response body
object
toggleOn
object
required
startTime
string | null
The local server time when the toggle was initiated. Null if state is NEVER_RAN

endTime
string | null
The local server time when the toggle was completed. Null if state is NEVER_RAN

elapsedTime
integer | null
Duration in seconds between the start time and end time. "Now" is used when end time is null. Null if state is NEVER_RAN

state
string
enum
The current state of the toggle

NOT_RUNNING RUNNING NEVER_RAN

totalRecords
int64
The total number of records that will be deleted

processedRecords
int64
The total number of records that have been deleted

percentComplete
double
The percentage between total and completed records.

formattedPercentComplete
string
Pretty print of total, processed, and percentage complete

exitState
string
enum
Troubleshooting - The exit status code from the toggle processing job. "Unknown" will return when the toggle is running.

UNKNOWN EXECUTING COMPLETED NOOP FAILED STOPPED

exitMessage
string
Troubleshooting - The exit message of the toggle job if it encounters an exception while running. Nominal return is an empty string

toggleOff
object
required
startTime
string | null
The local server time when the toggle was initiated. Null if state is NEVER_RAN

endTime
string | null
The local server time when the toggle was completed. Null if state is NEVER_RAN

elapsedTime
integer | null
Duration in seconds between the start time and end time. "Now" is used when end time is null. Null if state is NEVER_RAN

state
string
enum
The current state of the toggle

NOT_RUNNING RUNNING NEVER_RAN

totalRecords
int64
The total number of records that will be deleted

processedRecords
int64
The total number of records that have been deleted

percentComplete
double
The percentage between total and completed records.

formattedPercentComplete
string
Pretty print of total, processed, and percentage complete

exitState
string
enum
Troubleshooting - The exit status code from the toggle processing job. "Unknown" will return when the toggle is running.

UNKNOWN EXECUTING COMPLETED NOOP FAILED STOPPED

exitMessage
string
Troubleshooting - The exit message of the toggle job if it encounters an exception while running. Nominal return is an empty string

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/managed-software-updates/plans/feature-toggle/status \
     --header 'accept: application/json'

{
  "toggleOn": {
    "startTime": "2023-12-25T21:09:31",
    "endTime": "2023-12-25T21:09:31",
    "elapsedTime": 17,
    "state": "NOT_RUNNING",
    "totalRecords": 1000,
    "processedRecords": 500,
    "percentComplete": 61.2,
    "formattedPercentComplete": "100 / 500 -> 20% Complete",
    "exitState": "UNKNOWN",
    "exitMessage": "Error"
  },
  "toggleOff": {
    "startTime": "2023-12-25T21:09:31",
    "endTime": "2023-12-25T21:09:31",
    "elapsedTime": 17,
    "state": "NOT_RUNNING",
    "totalRecords": 1000,
    "processedRecords": 500,
    "percentComplete": 61.2,
    "formattedPercentComplete": "100 / 500 -> 20% Complete",
    "exitState": "UNKNOWN",
    "exitMessage": "Error"
  }
}
-----
Create Managed Software Update Plans for a Group
post
https://yourServer.jamfcloud.com/api/v1/managed-software-updates/plans/group

Creates Managed Software Update Plans for a Group

Body Params
Managed Software Update Plan to create for Group

group
object
required

group object
groupId
string
required
length ≥ 1
objectType
string
enum
required

COMPUTER_GROUP
Allowed:

COMPUTER_GROUP

MOBILE_DEVICE_GROUP
config
object
required

config object
updateAction
string
enum
required

DOWNLOAD_ONLY
Allowed:

DOWNLOAD_ONLY

DOWNLOAD_INSTALL

DOWNLOAD_INSTALL_ALLOW_DEFERRAL

DOWNLOAD_INSTALL_RESTART

DOWNLOAD_INSTALL_SCHEDULE

UNKNOWN
versionType
string
enum
required

LATEST_MAJOR
Allowed:

LATEST_MAJOR

LATEST_MINOR

LATEST_ANY

SPECIFIC_VERSION

CUSTOM_VERSION

UNKNOWN
specificVersion
string
length ≥ 0
Defaults to NO_SPECIFIC_VERSION
Optional. Indicates the specific version to update to. Only available when the version type is set to specific version or custom version, otherwise defaults to NO_SPECIFIC_VERSION.

NO_SPECIFIC_VERSION
buildVersion
string | null
Optional. Indicates the build version to update to. Only available when the version type is set to custom version.

maxDeferrals
integer
≥ 0
Required when the provided updateAction is DOWNLOAD_INSTALL_ALLOW_DEFERRAL, not applicable to all managed software update plans

forceInstallLocalDateTime
string | null
Optional. Indicates the local date and time of the device to force update by.

Responses

201
Managed Software Update Plan was created

Response body
object
plans
array of objects
length ≥ 0
object
device
object

device object
deviceId
string
length ≥ 1
objectType
string
enum
COMPUTER MOBILE_DEVICE APPLE_TV

href
string
planId
string
length ≥ 1
href
string

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/managed-software-updates/plans/group \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "group": {
    "objectType": "COMPUTER_GROUP"
  },
  "config": {
    "updateAction": "DOWNLOAD_ONLY",
    "versionType": "LATEST_MAJOR",
    "specificVersion": "NO_SPECIFIC_VERSION"
  }
}
'

{
  "plans": [
    {
      "device": {
        "deviceId": "1",
        "objectType": "COMPUTER",
        "href": "/v1/computers-inventory/1"
      },
      "planId": "6E47EF55-5318-494F-A09E-70F613E0AFD1",
      "href": "/v1/managed-software-updates/plans/6E47EF55-5318-494F-A09E-70F613E0AFD1"
    }
  ]
}
-----

Retrieve Managed Software Update Plans for a Group
get
https://yourServer.jamfcloud.com/api/v1/managed-software-updates/plans/group/{id}

Retrieves Managed Software Update Plans for a Group

Path Params
id
string
required
Managed Software Update Group Id

Query Params
group-type
string
enum
required
Managed Software Update Group Type, Available options are "COMPUTER_GROUP" or "MOBILE_DEVICE_GROUP"


COMPUTER_GROUP
Allowed:

COMPUTER_GROUP

MOBILE_DEVICE_GROUP
Responses

200
Success

Response body
object
totalCount
integer
≥ 0
results
array of objects
length ≥ 0
object
planUuid
string
required
length ≥ 1
device
object
required

device object
updateAction
string
enum
required
DOWNLOAD_ONLY DOWNLOAD_INSTALL DOWNLOAD_INSTALL_ALLOW_DEFERRAL DOWNLOAD_INSTALL_RESTART DOWNLOAD_INSTALL_SCHEDULE UNKNOWN

versionType
string
enum
required
LATEST_MAJOR LATEST_MINOR LATEST_ANY SPECIFIC_VERSION CUSTOM_VERSION UNKNOWN

specificVersion
string
Defaults to NO_SPECIFIC_VERSION
Optional. Indicates the specific version to update to. Only available when the version type is set to specific version or custom version, otherwise defaults to NO_SPECIFIC_VERSION.

buildVersion
string | null
Optional. Indicates the build version to update to. Only available when the version type is set to custom version.

maxDeferrals
integer
required
≥ 0
Not applicable to all managed software update plans

forceInstallLocalDateTime
string | null
Optional. Indicates the local date and time of the device to force update by.

recipeId
string
Defaults to -1
The id of the recipe that was used to generate the plan.

status
object
required

status object
state
string
enum
Init PendingPlanValidation AcceptingPlan RejectingPlan ProcessingPlanType ProcessingPlanTypeMdm StartingPlan PlanFailed SchedulingScanForOSUpdates ProcessingScheduleOSUpdateScanResponse WaitingForScheduledOSUpdateScanToComplete CollectingAvailableOSUpdates ProcessingAvailableOSUpdatesResponse ProcessingSchedulingType SchedulingDDM DDMPlanScheduled WaitingToStartDDMUpdate ProcessingDDMStatusResponse CollectingDDMStatus SchedulingMDM MDMPlanScheduled SchedulingOSUpdate ProcessingScheduleOSUpdateResponse CollectingOSUpdateStatus ProcessingOSUpdateStatusResponse WaitingToCollectOSUpdateStatus VerifyingInstallation ProcessingInstallationVerification PlanCompleted PlanCanceled PlanException Unknown

errorReasons
array of strings | null

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v1/managed-software-updates/plans/group/?group-type=COMPUTER_GROUP' \
     --header 'accept: application/json'

{
  "totalCount": 1,
  "results": [
    {
      "planUuid": "6E47EF55-5318-494F-A09E-70F613E0AFD1",
      "device": {
        "deviceId": "1",
        "objectType": "COMPUTER",
        "href": "/v1/computers-inventory/1"
      },
      "updateAction": "DOWNLOAD_INSTALL",
      "versionType": "SPECIFIC_VERSION",
      "specificVersion": "12.6.1",
      "buildVersion": "21F79",
      "maxDeferrals": 5,
      "forceInstallLocalDateTime": "2023-12-25T21:09:31",
      "recipeId": "1",
      "status": {
        "state": "SchedulingScanForOSUpdates",
        "errorReasons": [
          "NO_UPDATES_AVAILABLE",
          "NOT_SUPERVISED"
        ]
      }
    }
  ]
}
-----
Retrieve a Managed Software Update Plan
get
https://yourServer.jamfcloud.com/api/v1/managed-software-updates/plans/{id}

Retrieves a Managed Software Update Plan

Path Params
id
string
required
Managed Software Update Plan Uuid

Responses

200
Success

Response body
object
planUuid
string
required
length ≥ 1
device
object
required
deviceId
string
length ≥ 1
objectType
string
enum
COMPUTER MOBILE_DEVICE APPLE_TV

href
string
updateAction
string
enum
required
DOWNLOAD_ONLY DOWNLOAD_INSTALL DOWNLOAD_INSTALL_ALLOW_DEFERRAL DOWNLOAD_INSTALL_RESTART DOWNLOAD_INSTALL_SCHEDULE UNKNOWN

versionType
string
enum
required
LATEST_MAJOR LATEST_MINOR LATEST_ANY SPECIFIC_VERSION CUSTOM_VERSION UNKNOWN

specificVersion
string
Defaults to NO_SPECIFIC_VERSION
Optional. Indicates the specific version to update to. Only available when the version type is set to specific version or custom version, otherwise defaults to NO_SPECIFIC_VERSION.

buildVersion
string | null
Optional. Indicates the build version to update to. Only available when the version type is set to custom version.

maxDeferrals
integer
required
≥ 0
Not applicable to all managed software update plans

forceInstallLocalDateTime
string | null
Optional. Indicates the local date and time of the device to force update by.

recipeId
string
Defaults to -1
The id of the recipe that was used to generate the plan.

status
object
required
state
string
enum
Init PendingPlanValidation AcceptingPlan RejectingPlan ProcessingPlanType ProcessingPlanTypeMdm StartingPlan PlanFailed SchedulingScanForOSUpdates ProcessingScheduleOSUpdateScanResponse WaitingForScheduledOSUpdateScanToComplete CollectingAvailableOSUpdates ProcessingAvailableOSUpdatesResponse ProcessingSchedulingType SchedulingDDM DDMPlanScheduled WaitingToStartDDMUpdate ProcessingDDMStatusResponse CollectingDDMStatus SchedulingMDM MDMPlanScheduled SchedulingOSUpdate ProcessingScheduleOSUpdateResponse CollectingOSUpdateStatus ProcessingOSUpdateStatusResponse WaitingToCollectOSUpdateStatus VerifyingInstallation ProcessingInstallationVerification PlanCompleted PlanCanceled PlanException Unknown

errorReasons
array of strings | null

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/managed-software-updates/plans/ \
     --header 'accept: application/json'
  
{
  "planUuid": "6E47EF55-5318-494F-A09E-70F613E0AFD1",
  "device": {
    "deviceId": "1",
    "objectType": "COMPUTER",
    "href": "/v1/computers-inventory/1"
  },
  "updateAction": "DOWNLOAD_INSTALL",
  "versionType": "SPECIFIC_VERSION",
  "specificVersion": "12.6.1",
  "buildVersion": "21F79",
  "maxDeferrals": 5,
  "forceInstallLocalDateTime": "2023-12-25T21:09:31",
  "recipeId": "1",
  "status": {
    "state": "SchedulingScanForOSUpdates",
    "errorReasons": [
      "NO_UPDATES_AVAILABLE",
      "NOT_SUPERVISED"
    ]
  }
}
-----
Retrieve all Declarations associated with a Managed Software Update Plan
get
https://yourServer.jamfcloud.com/api/v1/managed-software-updates/plans/{id}/declarations

Retrieves all Declarations associated with a Managed Software Update Plan

Path Params
id
string
required
Managed Software Update Plan Uuid

Responses

200
Success

Response body
object
declarations
array of objects
Defaults to
object
uuid
string
payloadJson
string | null
type
string | null
group
string | null
enum
ACTIVATION ASSET CONFIGURATION MANAGEMENT

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/managed-software-updates/plans//declarations \
     --header 'accept: application/json'

{
  "declarations": [
    {
      "uuid": "72676372-af55-432f-acd8-12984522e472",
      "payloadJson": {},
      "type": "com.apple.configuration.management.status-subscriptions",
      "group": "activation"
    }
  ]
}
-----
Retrieve a Managed Software Update Plan Event Store
get
https://yourServer.jamfcloud.com/api/v1/managed-software-updates/plans/{id}/events

Retrieves a Managed Software Update Plan Event Store

Path Params
id
string
required
Managed Software Update Plan Uuid

Responses

200
Success

Response body
object
events
string

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/managed-software-updates/plans//events \
     --header 'accept: application/json'

{
  "events": ""
}
-----
Retrieve Managed Software Update Statuses
get
https://yourServer.jamfcloud.com/api/v1/managed-software-updates/update-statuses

Retrieve Managed Software Update Statuses

Query Params
filter
string
Query in the RSQL format, allowing to filter Managed Software Updates collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: osUpdatesStatusId, device.deviceId, device.objectType, downloaded, downloadPercentComplete, productKey, status, deferralsRemaining, maxDeferrals, nextScheduledInstall, created and updated.

Response

200
Success

Response body
object
totalCount
integer
results
array of objects
object
osUpdatesStatusId
string
device
object

device object
downloadPercentComplete
number
downloaded
boolean
productKey
string
status
string
enum
DOWNLOADING IDLE INSTALLING INSTALLED ERROR DOWNLOAD_FAILED DOWNLOAD_REQUIRES_COMPUTER DOWNLOAD_INSUFFICIENT_SPACE DOWNLOAD_INSUFFICIENT_POWER DOWNLOAD_INSUFFICIENT_NETWORK INSTALL_INSUFFICIENT_SPACE INSTALL_INSUFFICIENT_POWER INSTALL_PHONE_CALL_IN_PROGRESS INSTALL_FAILED UNKNOWN

deferralsRemaining
integer
not applicable to all managed software update statuses

maxDeferrals
integer
not applicable to all managed software update statuses

nextScheduledInstall
date-time
not applicable to all managed software update statuses

pastNotifications
array of date-times
not applicable to all managed software update statuses

created
date-time
updated
date-time

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/managed-software-updates/update-statuses \
     --header 'accept: application/json'

{
  "results": [
    {
      "osUpdatesStatusId": "1",
      "device": {
        "deviceId": 1,
        "objectType": "COMPUTER",
        "href": "/v1/computers-inventory/1"
      },
      "downloadPercentComplete": 0.8,
      "downloaded": true,
      "productKey": "macOSUpdate19F77",
      "status": "DOWNLOADING",
      "deferralsRemaining": 5,
      "maxDeferrals": 5,
      "nextScheduledInstall": "2022-12-25T21:09:31.661Z",
      "pastNotifications": [
        "2022-12-22T21:09:31.661Z",
        "2022-12-23T21:09:31.661Z",
        "2022-12-24T21:09:31.661Z"
      ],
      "created": "2022-12-22T21:09:00.661Z",
      "updated": "2022-12-24T21:09:31.661Z"
    }
  ],
  "totalCount": 1
}
-----
Retrieve Managed Software Update Statuses for Computer Groups
get
https://yourServer.jamfcloud.com/api/v1/managed-software-updates/update-statuses/computer-groups/{id}

Retrieve Managed Software Update Statuses for Computer Groups

Path Params
id
string
required
Computer Group identifier

Response

200
Success

Response body
object
totalCount
integer
results
array of objects
object
osUpdatesStatusId
string
device
object

device object
deviceId
string
objectType
string
enum
COMPUTER MOBILE_DEVICE APPLE_TV

href
string
downloadPercentComplete
number
downloaded
boolean
productKey
string
status
string
enum
DOWNLOADING IDLE INSTALLING INSTALLED ERROR DOWNLOAD_FAILED DOWNLOAD_REQUIRES_COMPUTER DOWNLOAD_INSUFFICIENT_SPACE DOWNLOAD_INSUFFICIENT_POWER DOWNLOAD_INSUFFICIENT_NETWORK INSTALL_INSUFFICIENT_SPACE INSTALL_INSUFFICIENT_POWER INSTALL_PHONE_CALL_IN_PROGRESS INSTALL_FAILED UNKNOWN

deferralsRemaining
integer
not applicable to all managed software update statuses

maxDeferrals
integer
not applicable to all managed software update statuses

nextScheduledInstall
date-time
not applicable to all managed software update statuses

pastNotifications
array of date-times
not applicable to all managed software update statuses

created
date-time
updated
date-time

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/managed-software-updates/update-statuses/computer-groups/ \
     --header 'accept: application/json'

{
  "results": [
    {
      "osUpdatesStatusId": "1",
      "device": {
        "deviceId": 1,
        "objectType": "COMPUTER",
        "href": "/v1/computers-inventory/1"
      },
      "downloadPercentComplete": 0.8,
      "downloaded": true,
      "productKey": "macOSUpdate19F77",
      "status": "DOWNLOADING",
      "deferralsRemaining": 5,
      "maxDeferrals": 5,
      "nextScheduledInstall": "2022-12-25T21:09:31.661Z",
      "pastNotifications": [
        "2022-12-22T21:09:31.661Z",
        "2022-12-23T21:09:31.661Z",
        "2022-12-24T21:09:31.661Z"
      ],
      "created": "2022-12-22T21:09:00.661Z",
      "updated": "2022-12-24T21:09:31.661Z"
    }
  ],
  "totalCount": 1
}
-----
Retrieve Managed Software Update Statuses for Computers
get
https://yourServer.jamfcloud.com/api/v1/managed-software-updates/update-statuses/computers/{id}

Retrieve Managed Software Update Statuses for Computers

Path Params
id
string
required
Computer identifier

Response

200
Success

Response body
object
totalCount
integer
results
array of objects
object
osUpdatesStatusId
string
device
object

device object
deviceId
string
objectType
string
enum
COMPUTER MOBILE_DEVICE APPLE_TV

href
string
downloadPercentComplete
number
downloaded
boolean
productKey
string
status
string
enum
DOWNLOADING IDLE INSTALLING INSTALLED ERROR DOWNLOAD_FAILED DOWNLOAD_REQUIRES_COMPUTER DOWNLOAD_INSUFFICIENT_SPACE DOWNLOAD_INSUFFICIENT_POWER DOWNLOAD_INSUFFICIENT_NETWORK INSTALL_INSUFFICIENT_SPACE INSTALL_INSUFFICIENT_POWER INSTALL_PHONE_CALL_IN_PROGRESS INSTALL_FAILED UNKNOWN

deferralsRemaining
integer
not applicable to all managed software update statuses

maxDeferrals
integer
not applicable to all managed software update statuses

nextScheduledInstall
date-time
not applicable to all managed software update statuses

pastNotifications
array of date-times
not applicable to all managed software update statuses

created
date-time
updated
date-time

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/managed-software-updates/update-statuses/computers/ \
     --header 'accept: application/json'

{
  "results": [
    {
      "osUpdatesStatusId": "1",
      "device": {
        "deviceId": 1,
        "objectType": "COMPUTER",
        "href": "/v1/computers-inventory/1"
      },
      "downloadPercentComplete": 0.8,
      "downloaded": true,
      "productKey": "macOSUpdate19F77",
      "status": "DOWNLOADING",
      "deferralsRemaining": 5,
      "maxDeferrals": 5,
      "nextScheduledInstall": "2022-12-25T21:09:31.661Z",
      "pastNotifications": [
        "2022-12-22T21:09:31.661Z",
        "2022-12-23T21:09:31.661Z",
        "2022-12-24T21:09:31.661Z"
      ],
      "created": "2022-12-22T21:09:00.661Z",
      "updated": "2022-12-24T21:09:31.661Z"
    }
  ],
  "totalCount": 1
}
-----
Retrieve Managed Software Update Statuses for Mobile Device Groups
get
https://yourServer.jamfcloud.com/api/v1/managed-software-updates/update-statuses/mobile-device-groups/{id}

Retrieve Managed Software Update Statuses for Mobile Device Groups

Path Params
id
string
required
Mobile Device Group identifier

Response

200
Success

Response body
object
totalCount
integer
results
array of objects
object
osUpdatesStatusId
string
device
object

device object
downloadPercentComplete
number
downloaded
boolean
productKey
string
status
string
enum
DOWNLOADING IDLE INSTALLING INSTALLED ERROR DOWNLOAD_FAILED DOWNLOAD_REQUIRES_COMPUTER DOWNLOAD_INSUFFICIENT_SPACE DOWNLOAD_INSUFFICIENT_POWER DOWNLOAD_INSUFFICIENT_NETWORK INSTALL_INSUFFICIENT_SPACE INSTALL_INSUFFICIENT_POWER INSTALL_PHONE_CALL_IN_PROGRESS INSTALL_FAILED UNKNOWN

deferralsRemaining
integer
not applicable to all managed software update statuses

maxDeferrals
integer
not applicable to all managed software update statuses

nextScheduledInstall
date-time
not applicable to all managed software update statuses

pastNotifications
array of date-times
not applicable to all managed software update statuses

created
date-time
updated
date-time

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/managed-software-updates/update-statuses/mobile-device-groups/ \
     --header 'accept: application/json'

{
  "results": [
    {
      "osUpdatesStatusId": "1",
      "device": {
        "deviceId": 1,
        "objectType": "MOBILE_DEVICE",
        "href": "/v2/mobile-devices/1"
      },
      "downloadPercentComplete": 0.8,
      "downloaded": true,
      "productKey": "iOSUpdate19F77",
      "status": "DOWNLOADING",
      "created": "2022-12-22T21:09:00.661Z",
      "updated": "2022-12-22T21:09:00.661Z"
    }
  ],
  "totalCount": 1
}
-----
Retrieve Managed Software Update Statuses for Mobile Devices
get
https://yourServer.jamfcloud.com/api/v1/managed-software-updates/update-statuses/mobile-devices/{id}

Retrieve Managed Software Update Statuses for Mobile Devices

Path Params
id
string
required
Mobile Device identifier

Response

200
Success

Response body
object
totalCount
integer
results
array of objects
object
osUpdatesStatusId
string
device
object

device object
downloadPercentComplete
number
downloaded
boolean
productKey
string
status
string
enum
DOWNLOADING IDLE INSTALLING INSTALLED ERROR DOWNLOAD_FAILED DOWNLOAD_REQUIRES_COMPUTER DOWNLOAD_INSUFFICIENT_SPACE DOWNLOAD_INSUFFICIENT_POWER DOWNLOAD_INSUFFICIENT_NETWORK INSTALL_INSUFFICIENT_SPACE INSTALL_INSUFFICIENT_POWER INSTALL_PHONE_CALL_IN_PROGRESS INSTALL_FAILED UNKNOWN

deferralsRemaining
integer
not applicable to all managed software update statuses

maxDeferrals
integer
not applicable to all managed software update statuses

nextScheduledInstall
date-time
not applicable to all managed software update statuses

pastNotifications
array of date-times
not applicable to all managed software update statuses

created
date-time
updated
date-time

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/managed-software-updates/update-statuses/mobile-devices/ \
     --header 'accept: application/json'

{
  "results": [
    {
      "osUpdatesStatusId": "1",
      "device": {
        "deviceId": 1,
        "objectType": "MOBILE_DEVICE",
        "href": "/v2/mobile-devices/1"
      },
      "downloadPercentComplete": 0.8,
      "downloaded": true,
      "productKey": "iOSUpdate19F77",
      "status": "DOWNLOADING",
      "created": "2022-12-22T21:09:00.661Z",
      "updated": "2022-12-22T21:09:00.661Z"
    }
  ],
  "totalCount": 1
}