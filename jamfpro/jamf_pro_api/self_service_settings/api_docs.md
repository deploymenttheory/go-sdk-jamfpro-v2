Get an object representation of Self Service settings
get
https://yourServer.jamfcloud.com/api/v1/self-service/settings

gets an object representation of Self Service settings

Response

200
successful GET

Response body
object
installSettings
object
object representation of Self Service settings regarding installation

installAutomatically
boolean
Defaults to false
true if Self Service is installed automatically, false if not

installLocation
string
required
path at which Self Service is installed. Required if installAutomatically is true

loginSettings
object
object representation of Self Service settings regarding login

userLoginLevel
string
enum
required
login setting to tell clients how to let users log in

NotRequired Anonymous Required

allowRememberMe
boolean
Defaults to false
true if remember me functionality is allowed, false if not

useFido2
boolean
Defaults to false
true if use FIDO2 functionality is allowed, false if not

authType
string
enum
required
login type to be used when asking users to log in

Basic Saml

configurationSettings
object
object representation of Self Service settings regarding user interaction

notificationsEnabled
boolean
Defaults to false
global Self Service setting for if notifications are on or off

alertUserApprovedMdm
boolean
Defaults to true
whether users should be notified they need to approve organization's MDM profile

defaultLandingPage
string
enum
Defaults to HOME
the default landing page in Self Service

HOME BROWSE HISTORY NOTIFICATIONS

defaultHomeCategoryId
int32
≥ -4
Defaults to -1
id for the default home category in Self Service

bookmarksName
string
required
renamed string for bookmarks if the admin wishes

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/self-service/settings \
     --header 'accept: application/json'

{
  "installSettings": {
    "installAutomatically": false,
    "installLocation": "/Applications"
  },
  "loginSettings": {
    "userLoginLevel": "NotRequired",
    "allowRememberMe": false,
    "useFido2": false,
    "authType": "Basic"
  },
  "configurationSettings": {
    "notificationsEnabled": false,
    "alertUserApprovedMdm": true,
    "defaultLandingPage": "HOME",
    "defaultHomeCategoryId": -1,
    "bookmarksName": "Bookmarks"
  }
}

Put an object representation of Self Service settings
put
https://yourServer.jamfcloud.com/api/v1/self-service/settings

puts an object representation of Self Service settings

Body Params
object that contains all editable global fields to alter Self Service settings

installSettings
object
object representation of Self Service settings regarding installation


installSettings object
loginSettings
object
object representation of Self Service settings regarding login


loginSettings object
configurationSettings
object
object representation of Self Service settings regarding user interaction


configurationSettings object
Response

200
successful PUT

Response body
object
installSettings
object
object representation of Self Service settings regarding installation

installAutomatically
boolean
Defaults to false
true if Self Service is installed automatically, false if not

installLocation
string
required
path at which Self Service is installed. Required if installAutomatically is true

loginSettings
object
object representation of Self Service settings regarding login

userLoginLevel
string
enum
required
login setting to tell clients how to let users log in

NotRequired Anonymous Required

allowRememberMe
boolean
Defaults to false
true if remember me functionality is allowed, false if not

useFido2
boolean
Defaults to false
true if use FIDO2 functionality is allowed, false if not

authType
string
enum
required
login type to be used when asking users to log in

Basic Saml

configurationSettings
object
object representation of Self Service settings regarding user interaction

notificationsEnabled
boolean
Defaults to false
global Self Service setting for if notifications are on or off

alertUserApprovedMdm
boolean
Defaults to true
whether users should be notified they need to approve organization's MDM profile

defaultLandingPage
string
enum
Defaults to HOME
the default landing page in Self Service

HOME BROWSE HISTORY NOTIFICATIONS

defaultHomeCategoryId
int32
≥ -4
Defaults to -1
id for the default home category in Self Service

bookmarksName
string
required
renamed string for bookmarks if the admin wishes


curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v1/self-service/settings \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "installSettings": {
    "installAutomatically": false
  },
  "loginSettings": {
    "userLoginLevel": "NotRequired",
    "allowRememberMe": false,
    "useFido2": false,
    "authType": "Basic"
  },
  "configurationSettings": {
    "notificationsEnabled": false,
    "alertUserApprovedMdm": true,
    "defaultLandingPage": "HOME",
    "defaultHomeCategoryId": -1
  }
}
'

{
  "installSettings": {
    "installAutomatically": false,
    "installLocation": "/Applications"
  },
  "loginSettings": {
    "userLoginLevel": "NotRequired",
    "allowRememberMe": false,
    "useFido2": false,
    "authType": "Basic"
  },
  "configurationSettings": {
    "notificationsEnabled": false,
    "alertUserApprovedMdm": true,
    "defaultLandingPage": "HOME",
    "defaultHomeCategoryId": -1,
    "bookmarksName": "Bookmarks"
  }
}
-----
Get a page of Self Service settings history
get
https://yourServer.jamfcloud.com/api/v1/self-service/settings/history


Get a page of Self Service settings history

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
Details of Self Service settings history were found

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
     --url 'https://yourserver.jamfcloud.com/api/v1/self-service/settings/history?page=0&page-size=100' \
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
Add Self Service settings history notes
post
https://yourServer.jamfcloud.com/api/v1/self-service/settings/history


Add Self Service settings history notes

Body Params
history notes to create

note
string
required
Response

201
Notes to Self Service settings history were added

Response body
object
id
string
length ≥ 1
href
string

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/self-service/settings/history \
     --header 'accept: application/json' \
     --header 'content-type: application/json'

{
  "id": "1",
  "href": "https://yourJamfProUrl.jamf/api/v1/resource/1"
}
-----