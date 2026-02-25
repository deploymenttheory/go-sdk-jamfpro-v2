Get compliance information for a single computer device
get
https://yourServer.jamfcloud.com/api/v1/conditional-access/device-compliance-information/computer/{deviceId}


Return basic compliance information for the given computer device

Path Params
deviceId
string
required
ID of the device the query pertains

Response

200
Array of device compliance information records

Response body
array of objects
object
deviceId
string
ID of the device

applicable
boolean
If device is applicable for compliance calculation

complianceState
string
enum
Device compliance state. Possible values are:

UNKNOWN for unknow compliance state, this usually means that the compliance state is being calculated,
NON_COMPLIANT for non compliant state,
COMPLIANT for compliant state
UNKNOWN NON_COMPLIANT COMPLIANT

complianceVendor
string
Name of the compliance vendor

complianceVendorDeviceInformation
object
Additional, compliance vendor specific device details.


complianceVendorDeviceInformation object
deviceIds
array of strings
Vendor's device IDs. Currently provided only for Intune.

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/conditional-access/device-compliance-information/computer/ \
     --header 'accept: application/json'

[
  {
    "deviceId": "1",
    "applicable": true,
    "complianceState": "COMPLIANT",
    "complianceVendor": "Vendor A",
    "complianceVendorDeviceInformation": [
      {
        "deviceIds": [
          "device-id-1",
          "device-id-2"
        ]
      }
    ]
  },
  {
    "deviceId": "1",
    "applicable": false,
    "complianceState": "NON_COMPLIANT",
    "complianceVendor": "Vendor B",
    "complianceVendorDeviceInformation": [
      {
        "deviceIds": []
      }
    ]
  }
]
-----
Get compliance information for a single mobile device
get
https://yourServer.jamfcloud.com/api/v1/conditional-access/device-compliance-information/mobile/{deviceId}


Return basic compliance information for the given mobile device

Path Params
deviceId
string
required
ID of the device the query pertains

Response

200
Array of device compliance information records

Response body
array of objects
object
deviceId
string
ID of the device

applicable
boolean
If device is applicable for compliance calculation

complianceState
string
enum
Device compliance state. Possible values are:

UNKNOWN for unknow compliance state, this usually means that the compliance state is being calculated,
NON_COMPLIANT for non compliant state,
COMPLIANT for compliant state
UNKNOWN NON_COMPLIANT COMPLIANT

complianceVendor
string
Name of the compliance vendor

complianceVendorDeviceInformation
object
Additional, compliance vendor specific device details.


complianceVendorDeviceInformation object
deviceIds
array of strings
Vendor's device IDs. Currently provided only for Intune.

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/conditional-access/device-compliance-information/mobile/ \
     --header 'accept: application/json'

[
  {
    "deviceId": "1",
    "applicable": true,
    "complianceState": "COMPLIANT",
    "complianceVendor": "Vendor A",
    "complianceVendorDeviceInformation": [
      {
        "deviceIds": [
          "device-id-1",
          "device-id-2"
        ]
      }
    ]
  },
  {
    "deviceId": "1",
    "applicable": false,
    "complianceState": "NON_COMPLIANT",
    "complianceVendor": "Vendor B",
    "complianceVendorDeviceInformation": [
      {
        "deviceIds": []
      }
    ]
  }
]
-----
Retrieves Status of the Feature Toggle
get
https://yourServer.jamfcloud.com/api/v1/conditional-access/device-compliance/feature-toggle


Retrieves Status of the Feature Toggle

Response

200
Success

Response body
object
sharedDeviceFeatureEnabled
boolean
required

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/conditional-access/device-compliance/feature-toggle \
     --header 'accept: application/json'

{
  "sharedDeviceFeatureEnabled": false
}