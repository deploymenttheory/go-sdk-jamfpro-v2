Get Jamf Pro account preferences
get
https://yourServer.jamfcloud.com/api/v3/account-preferences

Get Jamf Pro account preferences

Headers
Accept-Language
string
Locale to be used.

Response

200
Successful response - Jamf Pro Account Preferences retrieved

Response body
object
language
string
enum
required
Defaults to en
Language codes supported by Jamf Pro

en de fr es ja zh-hant

dateFormat
string
required
timezone
string
required
resultsPerPage
integer
required
Defaults to 100
userInterfaceDisplayTheme
string
enum
required
MATCH_SYSTEM LIGHT DARK

disableRelativeDates
boolean
required
disablePageLeaveCheck
boolean
required
disableTablePagination
boolean
required
disableShortcutsTooltips
boolean
required
configProfilesSortingMethod
string
required
computerSearchMethod
string
enum
required
EXACT_MATCH STARTS_WITH CONTAINS

computerApplicationSearchMethod
string
enum
required
EXACT_MATCH STARTS_WITH CONTAINS

computerApplicationUsageSearchMethod
string
enum
required
EXACT_MATCH STARTS_WITH CONTAINS

computerSoftwareUpdateSearchMethod
string
enum
EXACT_MATCH STARTS_WITH CONTAINS

computerLocalUserAccountSearchMethod
string
enum
required
EXACT_MATCH STARTS_WITH CONTAINS

computerPackageReceiptSearchMethod
string
enum
required
EXACT_MATCH STARTS_WITH CONTAINS

computerPrinterSearchMethod
string
enum
required
EXACT_MATCH STARTS_WITH CONTAINS

computerPeripheralSearchMethod
string
enum
EXACT_MATCH STARTS_WITH CONTAINS

computerServiceSearchMethod
string
enum
required
EXACT_MATCH STARTS_WITH CONTAINS

mobileDeviceSearchMethod
string
enum
required
EXACT_MATCH STARTS_WITH CONTAINS

mobileDeviceAppSearchMethod
string
enum
required
EXACT_MATCH STARTS_WITH CONTAINS

userSearchMethod
string
enum
required
EXACT_MATCH STARTS_WITH CONTAINS

userAllContentSearchMethod
string
enum
required
EXACT_MATCH STARTS_WITH CONTAINS

userMobileDeviceAppSearchMethod
string
enum
required
EXACT_MATCH STARTS_WITH CONTAINS

userMacAppStoreAppSearchMethod
string
enum
required
EXACT_MATCH STARTS_WITH CONTAINS

userEbookSearchMethod
string
enum
required
EXACT_MATCH STARTS_WITH CONTAINS

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v3/account-preferences \
     --header 'accept: application/json'

{
  "language": "en",
  "dateFormat": "MM/dd/yyyy",
  "timezone": "America/Chicago",
  "resultsPerPage": 50,
  "userInterfaceDisplayTheme": "MATCH_SYSTEM",
  "disableRelativeDates": false,
  "disablePageLeaveCheck": true,
  "disableTablePagination": true,
  "disableShortcutsTooltips": true,
  "configProfilesSortingMethod": "ALPHABETICALLY",
  "computerSearchMethod": "EXACT_MATCH",
  "computerApplicationSearchMethod": "EXACT_MATCH",
  "computerApplicationUsageSearchMethod": "EXACT_MATCH",
  "computerSoftwareUpdateSearchMethod": "EXACT_MATCH",
  "computerLocalUserAccountSearchMethod": "EXACT_MATCH",
  "computerPackageReceiptSearchMethod": "EXACT_MATCH",
  "computerPrinterSearchMethod": "EXACT_MATCH",
  "computerPeripheralSearchMethod": "EXACT_MATCH",
  "computerServiceSearchMethod": "EXACT_MATCH",
  "mobileDeviceSearchMethod": "EXACT_MATCH",
  "mobileDeviceAppSearchMethod": "EXACT_MATCH",
  "userSearchMethod": "EXACT_MATCH",
  "userAllContentSearchMethod": "EXACT_MATCH",
  "userMobileDeviceAppSearchMethod": "EXACT_MATCH",
  "userMacAppStoreAppSearchMethod": "EXACT_MATCH",
  "userEbookSearchMethod": "EXACT_MATCH"
}
-----
Update Jamf Pro account preferences
patch
https://yourServer.jamfcloud.com/api/v3/account-preferences

Update Jamf Pro account preferences

Body Params
language
string
enum
required
Defaults to en
Language codes supported by Jamf Pro


en
Allowed:

en

de

fr

es

ja

zh-hant
dateFormat
string
required
MM/dd/yyyy
timezone
string
required
America/Chicago
resultsPerPage
integer
required
Defaults to 100
100
userInterfaceDisplayTheme
string
enum
required

MATCH_SYSTEM
Allowed:

MATCH_SYSTEM

LIGHT

DARK
disableRelativeDates
boolean
required

true
disablePageLeaveCheck
boolean
required

true
disableTablePagination
boolean
required

true
disableShortcutsTooltips
boolean
required

true
configProfilesSortingMethod
string
required
ALPHABETICALLY
computerSearchMethod
string
enum
required

EXACT_MATCH
Allowed:

EXACT_MATCH

STARTS_WITH

CONTAINS
computerApplicationSearchMethod
string
enum
required

EXACT_MATCH
Allowed:

EXACT_MATCH

STARTS_WITH

CONTAINS
computerApplicationUsageSearchMethod
string
enum
required

EXACT_MATCH
Allowed:

EXACT_MATCH

STARTS_WITH

CONTAINS
computerSoftwareUpdateSearchMethod
string
enum

EXACT_MATCH
Allowed:

EXACT_MATCH

STARTS_WITH

CONTAINS
computerLocalUserAccountSearchMethod
string
enum
required

EXACT_MATCH
Allowed:

EXACT_MATCH

STARTS_WITH

CONTAINS
computerPackageReceiptSearchMethod
string
enum
required

EXACT_MATCH
Allowed:

EXACT_MATCH

STARTS_WITH

CONTAINS
computerPrinterSearchMethod
string
enum
required

EXACT_MATCH
Allowed:

EXACT_MATCH

STARTS_WITH

CONTAINS
computerPeripheralSearchMethod
string
enum

EXACT_MATCH
Allowed:

EXACT_MATCH

STARTS_WITH

CONTAINS
computerServiceSearchMethod
string
enum
required

EXACT_MATCH
Allowed:

EXACT_MATCH

STARTS_WITH

CONTAINS
mobileDeviceSearchMethod
string
enum
required

EXACT_MATCH
Allowed:

EXACT_MATCH

STARTS_WITH

CONTAINS
mobileDeviceAppSearchMethod
string
enum
required

EXACT_MATCH
Allowed:

EXACT_MATCH

STARTS_WITH

CONTAINS
userSearchMethod
string
enum
required

EXACT_MATCH
Allowed:

EXACT_MATCH

STARTS_WITH

CONTAINS
userAllContentSearchMethod
string
enum
required

EXACT_MATCH
Allowed:

EXACT_MATCH

STARTS_WITH

CONTAINS
userMobileDeviceAppSearchMethod
string
enum
required

EXACT_MATCH
Allowed:

EXACT_MATCH

STARTS_WITH

CONTAINS
userMacAppStoreAppSearchMethod
string
enum
required

EXACT_MATCH
Allowed:

EXACT_MATCH

STARTS_WITH

CONTAINS
userEbookSearchMethod
string
enum
required

EXACT_MATCH
Allowed:

EXACT_MATCH

STARTS_WITH

CONTAINS
Cookie Params
JSESSIONID
string
Defaults to null
Session cookie, that's used to determine user session where account preferences should be refreshed

null
Headers
Accept-Language
string
Locale to be used, when user has not defined preferred language.

Responses
204
Successful response - Jamf Pro Account Preferences updated

curl --request PATCH \
     --url https://yourserver.jamfcloud.com/api/v3/account-preferences \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --cookie JSESSIONID=null \
     --data '
{
  "language": "en",
  "resultsPerPage": 100,
  "userInterfaceDisplayTheme": "MATCH_SYSTEM",
  "disableRelativeDates": true,
  "disablePageLeaveCheck": true,
  "disableTablePagination": true,
  "disableShortcutsTooltips": true,
  "computerSearchMethod": "EXACT_MATCH",
  "computerApplicationSearchMethod": "EXACT_MATCH",
  "computerApplicationUsageSearchMethod": "EXACT_MATCH",
  "computerLocalUserAccountSearchMethod": "EXACT_MATCH",
  "computerPackageReceiptSearchMethod": "EXACT_MATCH",
  "computerPrinterSearchMethod": "EXACT_MATCH",
  "computerServiceSearchMethod": "EXACT_MATCH",
  "mobileDeviceSearchMethod": "EXACT_MATCH",
  "mobileDeviceAppSearchMethod": "EXACT_MATCH",
  "userSearchMethod": "EXACT_MATCH",
  "userAllContentSearchMethod": "EXACT_MATCH",
  "userMobileDeviceAppSearchMethod": "EXACT_MATCH",
  "userMacAppStoreAppSearchMethod": "EXACT_MATCH",
  "userEbookSearchMethod": "EXACT_MATCH",
  "dateFormat": "MM/dd/yyyy",
  "timezone": "America/Chicago",
  "configProfilesSortingMethod": "ALPHABETICALLY"
}
'