Retrieves all settings for device communication
get
https://yourServer.jamfcloud.com/api/v1/device-communication-settings


Retrieves all device communication settings, including automatic renewal of the MDM profile.

Response

200
Successful response - Device Communication Settings retrieved

Response body
object
autoRenewMobileDeviceMdmProfileWhenCaRenewed
boolean
autoRenewMobileDeviceMdmProfileWhenDeviceIdentityCertExpiring
boolean
autoRenewComputerMdmProfileWhenCaRenewed
boolean
autoRenewComputerMdmProfileWhenDeviceIdentityCertExpiring
boolean
mdmProfileMobileDeviceExpirationLimitInDays
integer
enum
Defaults to 180
90 120 180

mdmProfileComputerExpirationLimitInDays
integer
enum
Defaults to 180
90 120 180

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/device-communication-settings \
     --header 'accept: application/json'

{
  "autoRenewMobileDeviceMdmProfileWhenCaRenewed": true,
  "autoRenewMobileDeviceMdmProfileWhenDeviceIdentityCertExpiring": true,
  "autoRenewComputerMdmProfileWhenCaRenewed": true,
  "autoRenewComputerMdmProfileWhenDeviceIdentityCertExpiring": true,
  "mdmProfileMobileDeviceExpirationLimitInDays": 180,
  "mdmProfileComputerExpirationLimitInDays": 180
}
-----
Update device communication settings
put
https://yourServer.jamfcloud.com/api/v1/device-communication-settings


Update device communication settings

Body Params
autoRenewMobileDeviceMdmProfileWhenCaRenewed
boolean

true
autoRenewMobileDeviceMdmProfileWhenDeviceIdentityCertExpiring
boolean

true
autoRenewComputerMdmProfileWhenCaRenewed
boolean

true
autoRenewComputerMdmProfileWhenDeviceIdentityCertExpiring
boolean

true
mdmProfileMobileDeviceExpirationLimitInDays
integer
enum
Defaults to 180

180
Allowed:

90

120

180
mdmProfileComputerExpirationLimitInDays
integer
enum
Defaults to 180

180
Allowed:

90

120

180
Response

200
Successful response - Device communication Settings updated

Response body
object
autoRenewMobileDeviceMdmProfileWhenCaRenewed
boolean
autoRenewMobileDeviceMdmProfileWhenDeviceIdentityCertExpiring
boolean
autoRenewComputerMdmProfileWhenCaRenewed
boolean
autoRenewComputerMdmProfileWhenDeviceIdentityCertExpiring
boolean
mdmProfileMobileDeviceExpirationLimitInDays
integer
enum
Defaults to 180
90 120 180

mdmProfileComputerExpirationLimitInDays
integer
enum
Defaults to 180
90 120 180

curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v1/device-communication-settings \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "mdmProfileMobileDeviceExpirationLimitInDays": 180,
  "mdmProfileComputerExpirationLimitInDays": 180
}
'

{
  "autoRenewMobileDeviceMdmProfileWhenCaRenewed": true,
  "autoRenewMobileDeviceMdmProfileWhenDeviceIdentityCertExpiring": true,
  "autoRenewComputerMdmProfileWhenCaRenewed": true,
  "autoRenewComputerMdmProfileWhenDeviceIdentityCertExpiring": true,
  "mdmProfileMobileDeviceExpirationLimitInDays": 180,
  "mdmProfileComputerExpirationLimitInDays": 180
}
-----
Get Device Communication settings history
get
https://yourServer.jamfcloud.com/api/v1/device-communication-settings/history


Gets Device Communication settings history

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
Defaults to date:desc
Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc


string

date:desc

ADD string
filter
string
Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: username, date, note, details. This param can be combined with paging and sorting. Example: filter=username!=admin and details==disabled and date<2019-12-15

Response

200
Details of Device Communication Settings history were found

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
     --url 'https://yourserver.jamfcloud.com/api/v1/device-communication-settings/history?page=0&page-size=100&sort=date%3Adesc' \
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
Add Device Communication Settings history notes
post
https://yourServer.jamfcloud.com/api/v1/device-communication-settings/history


Adds Device Communication Settings history notes

Body Params
history notes to create

note
string
required
A generic note can sometimes be useful, but generally not.
Responses

201
Notes to Device Communication Settings history were added

Response body
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

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/device-communication-settings/history \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "note": "A generic note can sometimes be useful, but generally not."
}
'

{
  "id": 1,
  "username": "admin",
  "date": "2019-02-04T21:09:31.661Z",
  "note": "Sso settings update",
  "details": "Is SSO Enabled false\\nSelected SSO Provider"
}
-----