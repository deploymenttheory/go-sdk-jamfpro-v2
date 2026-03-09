Update device common details (partial update)
patch
https://yourServer.jamfcloud.com/api/v1/mdm-renewal/device-common-details

Partially updates existing device common details. The clientManagementId must be provided in the request body to identify which record to update. Only updates fields that are explicitly provided in the request - missing fields preserve their existing values. Only updates existing records; does not create new ones.

Body Params
clientManagementId
string
required
The client management ID associated with this device (required)

550e8400-e29b-41d4-a716-446655440000
renewMdmProfileStartDate
date-time | null
Timestamp when MDM profile renewal started (ISO 8601 format)

2021-12-31T16:00:00Z
mdmProfileNeedsRenewalDueToCaRenewed
boolean | null
Whether the MDM profile needs renewal due to CA renewal


false
mdmProfileNeedsRenewalDueToDeviceIdentityCertExpiring
boolean | null
Whether the MDM profile needs renewal due to expiring device identity certificate


true
mdmCheckinUrl
string | null
URL for MDM check-in

https://example.jamfcloud.com/mdm/CheckInURL
mdmServerUrl
string | null
URL for MDM server

https://example.jamfcloud.com/mdm/ServerURL
Responses
204
Successfully updated device common details

curl --request PATCH \
     --url https://yourserver.jamfcloud.com/api/v1/mdm-renewal/device-common-details \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "clientManagementId": "550e8400-e29b-41d4-a716-446655440000",
  "renewMdmProfileStartDate": "2021-12-31T16:00:00Z",
  "mdmProfileNeedsRenewalDueToCaRenewed": false,
  "mdmProfileNeedsRenewalDueToDeviceIdentityCertExpiring": true,
  "mdmCheckinUrl": "https://example.jamfcloud.com/mdm/CheckInURL",
  "mdmServerUrl": "https://example.jamfcloud.com/mdm/ServerURL"
}
'
-----
Get device common details for a client management ID
get
https://yourServer.jamfcloud.com/api/v1/mdm-renewal/device-common-details/{clientManagementId}

Retrieves device common details associated with a specific client management ID

Path Params
clientManagementId
string
required
The client management ID to retrieve device common details for

550e8400-e29b-41d4-a716-446655440000
Responses

200
Successfully retrieved device common details

Response body
object
id
string
Unique identifier for the device common details record

clientManagementId
string
required
The client management ID associated with this device

renewMdmProfileStartDate
date-time | null
Timestamp when MDM profile renewal started (ISO 8601 format)

mdmProfileNeedsRenewalDueToCaRenewed
boolean
Whether the MDM profile needs renewal due to CA renewal

mdmProfileNeedsRenewalDueToDeviceIdentityCertExpiring
boolean
Whether the MDM profile needs renewal due to expiring device identity certificate

mdmCheckinUrl
string | null
URL for MDM check-in

mdmServerUrl
string | null
URL for MDM server

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/mdm-renewal/device-common-details/550e8400-e29b-41d4-a716-446655440000 \
     --header 'accept: application/json'

{
  "id": "123",
  "clientManagementId": "550e8400-e29b-41d4-a716-446655440000",
  "renewMdmProfileStartDate": "2021-12-31T16:00:00Z",
  "mdmProfileNeedsRenewalDueToCaRenewed": false,
  "mdmProfileNeedsRenewalDueToDeviceIdentityCertExpiring": true,
  "mdmCheckinUrl": "https://example.jamfcloud.com/mdm/CheckInURL",
  "mdmServerUrl": "https://example.jamfcloud.com/mdm/ServerURL"
}
-----
Get MDM renewal errors and strategies for a client management ID
get
https://yourServer.jamfcloud.com/api/v1/mdm-renewal/renewal-strategies/{clientManagementId}

Retrieves all MDM renewal errors and their associated renewal strategies for a specific client management ID

Path Params
clientManagementId
string
required
The client management ID to retrieve renewal strategies for

550e8400-e29b-41d4-a716-446655440000
Responses

200
Successfully retrieved MDM renewal errors and strategies

Response body
array of objects
object
error
object
required

error object
mdmRenewalErrorId
string
required
Unique identifier for the MDM renewal error

clientManagementId
string
required
The client management ID associated with this error

mdmRenewalErrorType
string
enum
required
Type of MDM renewal error

SERVER_ERROR CHECK_IN_ERROR OTHER

errorTimeStamp
date-time
Timestamp when the error occurred (ISO 8601 format)

failureCount
int32
Number of times this error has occurred

strategies
array of objects
required
List of renewal strategies associated with this error

object
id
string
required
Unique identifier for the renewal strategy

mdmRenewalErrorId
string
required
The MDM renewal error ID this strategy is associated with

mdmRenewalStrategyType
string
enum
required
Type of MDM renewal strategy

RETURN_NO_CHECK_IN_INVITATION RETURN_CHECK_IN_INVITATION_FROM_MDM_INVITATION_TABLE RETURN_CHECK_IN_INVITATION_FROM_ENROLLMENT_USAGE_TABLE RETURN_CHECK_IN_INVITATION_FROM_MDM_PROFILE_PROTOTYPE_TABLE JSS_URL_OVERRIDE PAYLOAD_IDENTIFIER

strategyTimeStamp
date-time
Timestamp when this renewal strategy was created (ISO 8601 format)

mdmRenewalCheckInUrl
string
URL for MDM renewal check-in

mdmRenewalServerUrl
string
URL for MDM renewal server

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/mdm-renewal/renewal-strategies/550e8400-e29b-41d4-a716-446655440000 \
     --header 'accept: application/json'

[
  {
    "error": {
      "mdmRenewalErrorId": "123",
      "clientManagementId": "550e8400-e29b-41d4-a716-446655440000",
      "mdmRenewalErrorType": "SERVER_ERROR",
      "errorTimeStamp": "2021-12-31T22:00:00Z",
      "failureCount": 3
    },
    "strategies": [
      {
        "id": "456",
        "mdmRenewalErrorId": "123",
        "mdmRenewalStrategyType": "JSS_URL_OVERRIDE",
        "strategyTimeStamp": "2021-12-31T22:00:00Z",
        "mdmRenewalCheckInUrl": "https://example.jamfcloud.com/mdm/CheckInURL",
        "mdmRenewalServerUrl": "https://example.jamfcloud.com/mdm/ServerURL"
      }
    ]
  }
]
-----
Delete MDM renewal strategies for a client management ID
delete
https://yourServer.jamfcloud.com/api/v1/mdm-renewal/renewal-strategies/{clientManagementId}

Deletes all MDM renewal strategies and errors associated with the specified client management ID

Path Params
clientManagementId
string
required
The client management ID to delete renewal strategies for

550e8400-e29b-41d4-a716-446655440000
Responses
204
Successfully deleted MDM renewal strategies and errors

curl --request DELETE \
     --url https://yourserver.jamfcloud.com/api/v1/mdm-renewal/renewal-strategies/550e8400-e29b-41d4-a716-446655440000 \
     --header 'accept: application/json'
