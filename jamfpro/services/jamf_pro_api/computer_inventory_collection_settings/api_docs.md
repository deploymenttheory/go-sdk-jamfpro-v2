Returns computer inventory settings
get
https://yourServer.jamfcloud.com/api/v2/computer-inventory-collection-settings


Returns computer inventory settings

Response

200
Successful response returns the computer inventory settings

Response body
object
computerInventoryCollectionPreferences
object
monitorApplicationUsage
boolean
Defaults to false
includePackages
boolean
Defaults to false
includeSoftwareUpdates
boolean
Defaults to false
includeSoftwareId
boolean
Defaults to false
includeAccounts
boolean
Defaults to false
calculateSizes
boolean
Defaults to false
includeHiddenAccounts
boolean
Defaults to false
includePrinters
boolean
Defaults to false
includeServices
boolean
Defaults to false
collectSyncedMobileDeviceInfo
boolean
Defaults to false
updateLdapInfoOnComputerInventorySubmissions
boolean
Defaults to false
monitorBeacons
boolean
Defaults to false
allowChangingUserAndLocation
boolean
Defaults to true
useUnixUserPaths
boolean
Defaults to true
collectUnmanagedCertificates
boolean
Defaults to true
applicationPaths
array of objects
object
id
string
required
A "-1" id indicates a built-in path that cannot be deleted or modified.

path
string
required

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v2/computer-inventory-collection-settings \
     --header 'accept: application/json'

{
  "computerInventoryCollectionPreferences": {
    "monitorApplicationUsage": true,
    "includePackages": true,
    "includeSoftwareUpdates": true,
    "includeSoftwareId": true,
    "includeAccounts": true,
    "calculateSizes": false,
    "includeHiddenAccounts": true,
    "includePrinters": true,
    "includeServices": true,
    "collectSyncedMobileDeviceInfo": false,
    "updateLdapInfoOnComputerInventorySubmissions": false,
    "monitorBeacons": true,
    "allowChangingUserAndLocation": true,
    "useUnixUserPaths": true,
    "collectUnmanagedCertificates": true
  },
  "applicationPaths": [
    {
      "id": "1",
      "path": "/Example/Path/To/App/"
    }
  ]
}
-----
Update computer inventory settings
patch
https://yourServer.jamfcloud.com/api/v2/computer-inventory-collection-settings


Update computer inventory settings

Body Params
Computer inventory settings to update

computerInventoryCollectionPreferences
object

computerInventoryCollectionPreferences object
monitorApplicationUsage
boolean
Defaults to false

false
includePackages
boolean
Defaults to false

false
includeSoftwareUpdates
boolean
Defaults to false

false
includeSoftwareId
boolean
Defaults to false

false
includeAccounts
boolean
Defaults to false

false
calculateSizes
boolean
Defaults to false

false
includeHiddenAccounts
boolean
Defaults to false

false
includePrinters
boolean
Defaults to false

false
includeServices
boolean
Defaults to false

false
collectSyncedMobileDeviceInfo
boolean
Defaults to false

false
updateLdapInfoOnComputerInventorySubmissions
boolean
Defaults to false

false
monitorBeacons
boolean
Defaults to false

false
allowChangingUserAndLocation
boolean
Defaults to true

true
useUnixUserPaths
boolean
Defaults to true

true
collectUnmanagedCertificates
boolean
Defaults to true

true
applicationPaths
array of objects

object

id
string
required
A "-1" id indicates a built-in path that cannot be deleted or modified.

1
path
string
required
/Example/Path/To/App/

object

id
string
required
A "-1" id indicates a built-in path that cannot be deleted or modified.

2
path
string
required
/Example/Path/To/App/

ADD object
Response
204
Computer inventory settings updated

curl --request PATCH \
     --url https://yourserver.jamfcloud.com/api/v2/computer-inventory-collection-settings \
     --header 'content-type: application/json' \
     --data '
{
  "computerInventoryCollectionPreferences": {
    "monitorApplicationUsage": false,
    "includePackages": false,
    "includeSoftwareUpdates": false,
    "includeSoftwareId": false,
    "includeAccounts": false,
    "calculateSizes": false,
    "includeHiddenAccounts": false,
    "includePrinters": false,
    "includeServices": false,
    "collectSyncedMobileDeviceInfo": false,
    "updateLdapInfoOnComputerInventorySubmissions": false,
    "monitorBeacons": false,
    "allowChangingUserAndLocation": true,
    "useUnixUserPaths": true,
    "collectUnmanagedCertificates": true
  },
  "applicationPaths": [
    {
      "id": "1",
      "path": "/Example/Path/To/App/"
    },
    {
      "id": "2",
      "path": "/Example/Path/To/App/"
    }
  ]
}
'

Create Computer Inventory Collection Settings Custom Path
post
https://yourServer.jamfcloud.com/api/v2/computer-inventory-collection-settings/custom-path


Creates a custom search path to use when collecting applications.

Body Params
Computer inventory settings to update

scope
string
enum
required

APP
Allowed:

APP
path
string
required
Responses

201
Custom path created

Response body
object
id
string
length ≥ 1
href
string

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v2/computer-inventory-collection-settings/custom-path \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '{"scope":"APP"}'

{
  "id": "1",
  "href": "https://yourJamfProUrl.jamf/api/v1/resource/1"
}

Delete Custom Path from Computer Inventory Collection Settings
delete
https://yourServer.jamfcloud.com/api/v2/computer-inventory-collection-settings/custom-path/{id}


Delete Custom Path from Computer Inventory Collection Settings

Path Params
id
string
required
id of Custom Path

Responses
204
Custom path deleted

curl --request DELETE \
     --url https://yourserver.jamfcloud.com/api/v2/computer-inventory-collection-settings/custom-path/ \
     --header 'accept: application/json'