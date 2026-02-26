
Retrieve Volume Purchasing Locations
get
https://yourServer.jamfcloud.com/api/v1/volume-purchasing-locations

Retrieves Volume Purchasing Locations

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
Query in the RSQL format, allowing to filter Volume Purchasing Location collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: id, name, appleId, email, organizationName, tokenExpiration, countryCode, locationName, automaticallyPopulatePurchasedContent, sendNotificationWhenNoLongerAssigned, siteId and siteName. This param can be combined with paging and sorting.

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
name
string
totalPurchasedLicenses
integer
totalUsedLicenses
integer
id
string
appleId
string
email
string
organizationName
string
tokenExpiration
string
countryCode
string
The two-letter ISO 3166-1 code that designates the country where the Volume Purchasing account is located.

locationName
string
clientContextMismatch
boolean
If this is "true", the clientContext used by this server does not match the clientContext returned by the Volume Purchasing API.

automaticallyPopulatePurchasedContent
boolean
sendNotificationWhenNoLongerAssigned
boolean
autoRegisterManagedUsers
boolean
siteId
string
siteName
string
lastSyncTime
string

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v1/volume-purchasing-locations?page=0&page-size=100&sort=id%3Aasc' \
     --header 'accept: application/json'
{
  "totalCount": 1,
  "results": [
    {
      "name": "Example Location",
      "totalPurchasedLicenses": 1,
      "totalUsedLicenses": 1,
      "id": "1",
      "appleId": "testUser@appleId.com",
      "email": "testUser@email.com",
      "organizationName": "Jamf",
      "tokenExpiration": "2022-04-25T21:09:31.661Z",
      "countryCode": "US",
      "locationName": "Example Location",
      "clientContextMismatch": false,
      "automaticallyPopulatePurchasedContent": false,
      "sendNotificationWhenNoLongerAssigned": false,
      "autoRegisterManagedUsers": false,
      "siteId": "1",
      "siteName": "Example Name",
      "lastSyncTime": "2022-09-25T21:09:31.661Z"
    }
  ]
}

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v1/volume-purchasing-locations?page=0&page-size=100&sort=id%3Aasc' \
     --header 'accept: application/json'

{
  "totalCount": 1,
  "results": [
    {
      "name": "Example Location",
      "totalPurchasedLicenses": 1,
      "totalUsedLicenses": 1,
      "id": "1",
      "appleId": "testUser@appleId.com",
      "email": "testUser@email.com",
      "organizationName": "Jamf",
      "tokenExpiration": "2022-04-25T21:09:31.661Z",
      "countryCode": "US",
      "locationName": "Example Location",
      "clientContextMismatch": false,
      "automaticallyPopulatePurchasedContent": false,
      "sendNotificationWhenNoLongerAssigned": false,
      "autoRegisterManagedUsers": false,
      "siteId": "1",
      "siteName": "Example Name",
      "lastSyncTime": "2022-09-25T21:09:31.661Z"
    }
  ]
}
-----
Create a Volume Purchasing Location
post
https://yourServer.jamfcloud.com/api/v1/volume-purchasing-locations

Creates a Volume Purchasing Location using an sToken

Body Params
Volume Purchasing Location to create

name
string
If no value is provided when creating a VolumePurchasingLocation object, the 'name' will default to the 'locationName' value

Example Location
automaticallyPopulatePurchasedContent
boolean
Defaults to false

false
sendNotificationWhenNoLongerAssigned
boolean
Defaults to false

false
autoRegisterManagedUsers
boolean
Defaults to false

false
siteId
string
Defaults to -1
-1
serviceToken
string
required
eyJleHBEYXRlIjoiMjAyMi0wMy0yOVQxNTozNjoyNiswMDAwIiwidG9rZW4iOiJWR2hwY3lCcGN5QnViM1FnWVNCMGIydGxiaTRnU0c5d1pXWjFiR3g1SUdsMElHeHZiMnR6SUd4cGEyVWdZU0IwYjJ0bGJpd2dZblYwSUdsMEozTWdibTkwTGc9PSIsIm9yZ05hbWUiOiJFeGFtcGxlIE9yZyJ9
Responses

201
Volume Purchasing Location was created

Response body
object
id
string
length ≥ 1
href
string

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/volume-purchasing-locations \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "automaticallyPopulatePurchasedContent": false,
  "sendNotificationWhenNoLongerAssigned": false,
  "autoRegisterManagedUsers": false,
  "siteId": "-1",
  "name": "Example Location",
  "serviceToken": "eyJleHBEYXRlIjoiMjAyMi0wMy0yOVQxNTozNjoyNiswMDAwIiwidG9rZW4iOiJWR2hwY3lCcGN5QnViM1FnWVNCMGIydGxiaTRnU0c5d1pXWjFiR3g1SUdsMElHeHZiMnR6SUd4cGEyVWdZU0IwYjJ0bGJpd2dZblYwSUdsMEozTWdibTkwTGc9PSIsIm9yZ05hbWUiOiJFeGFtcGxlIE9yZyJ9"
}
'

{
  "id": "1",
  "href": "https://yourJamfProUrl.jamf/api/v1/resource/1"
}
-----
Retrieve a Volume Purchasing Location with the supplied id
get
https://yourServer.jamfcloud.com/api/v1/volume-purchasing-locations/{id}

Retrieves a Volume Purchasing Location with the supplied id

Path Params
id
string
required
Volume Purchasing Location identifier

1
Responses

200
Success

Response body
object
name
string
totalPurchasedLicenses
integer
totalUsedLicenses
integer
id
string
appleId
string
email
string
organizationName
string
tokenExpiration
string
countryCode
string
The two-letter ISO 3166-1 code that designates the country where the Volume Purchasing account is located.

locationName
string
clientContextMismatch
boolean
If this is "true", the clientContext used by this server does not match the clientContext returned by the Volume Purchasing API.

automaticallyPopulatePurchasedContent
boolean
sendNotificationWhenNoLongerAssigned
boolean
autoRegisterManagedUsers
boolean
siteId
string
siteName
string
lastSyncTime
string
content
array of objects
object
name
string
licenseCountTotal
integer
licenseCountInUse
integer
licenseCountReported
integer
iconUrl
string
deviceTypes
array of strings
contentType
string
enum
IOS_APP MAC_APP BOOK UNKNOWN

pricingParam
string
enum
STDQ PLUS Unknown

adamId
string

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/volume-purchasing-locations/1 \
     --header 'accept: application/json'

{
  "name": "Example Location",
  "totalPurchasedLicenses": 1,
  "totalUsedLicenses": 1,
  "id": "1",
  "appleId": "testUser@appleId.com",
  "email": "testUser@email.com",
  "organizationName": "Jamf",
  "tokenExpiration": "2022-04-25T21:09:31.661Z",
  "countryCode": "US",
  "locationName": "Example Location",
  "clientContextMismatch": false,
  "automaticallyPopulatePurchasedContent": false,
  "sendNotificationWhenNoLongerAssigned": false,
  "autoRegisterManagedUsers": false,
  "siteId": "1",
  "siteName": "Example Name",
  "lastSyncTime": "2022-09-25T21:09:31.661Z",
  "content": [
    {
      "name": "Example Content",
      "licenseCountTotal": 1,
      "licenseCountInUse": 1,
      "licenseCountReported": 1,
      "iconUrl": "https://is4-ssl.mzstatic.com/image/thumb/Purple113/v4/73/d4/73/73d47332-fefc-d350-2984-5b4a4755a502/AppIcon-0-1x_U007emarketing-0-0-GLES2_U002c0-512MB-sRGB-0-0-0-85-220-0-0-0-6.png/360x216bb.png",
      "deviceTypes": [
        "IOS"
      ],
      "contentType": "IOS_APP",
      "pricingParam": "STDQ",
      "adamId": "748057890"
    }
  ]
}
-----
Delete a Volume Purchasing Location with the supplied id
delete
https://yourServer.jamfcloud.com/api/v1/volume-purchasing-locations/{id}

Deletes a Volume Purchasing Location with the supplied id

Path Params
id
string
required
Volume Purchasing Location identifier

1
Responses
204
Success

curl --request DELETE \
     --url https://yourserver.jamfcloud.com/api/v1/volume-purchasing-locations/1 \
     --header 'accept: application/json'
-----
Update a Volume Purchasing Location
patch
https://yourServer.jamfcloud.com/api/v1/volume-purchasing-locations/{id}

Updates a Volume Purchasing Location

Path Params
id
string
required
Volume Purchasing Location identifier

Body Params
Volume Purchasing Location to update

name
string
Example Location
automaticallyPopulatePurchasedContent
boolean

true
sendNotificationWhenNoLongerAssigned
boolean

true
autoRegisterManagedUsers
boolean
Defaults to false

false
siteId
string
1
serviceToken
string
eyJleHBEYXRlIjoiMjAyMi0wMy0yOVQxNTozNjoyNiswMDAwIiwidG9rZW4iOiJWR2hwY3lCcGN5QnViM1FnWVNCMGIydGxiaTRnU0c5d1pXWjFiR3g1SUdsMElHeHZiMnR6SUd4cGEyVWdZU0IwYjJ0bGJpd2dZblYwSUdsMEozTWdibTkwTGc9PSIsIm9yZ05hbWUiOiJFeGFtcGxlIE9yZyJ9
Responses

200
Success

Response body
object
name
string
totalPurchasedLicenses
integer
totalUsedLicenses
integer
id
string
appleId
string
email
string
organizationName
string
tokenExpiration
string
countryCode
string
The two-letter ISO 3166-1 code that designates the country where the Volume Purchasing account is located.

locationName
string
clientContextMismatch
boolean
If this is "true", the clientContext used by this server does not match the clientContext returned by the Volume Purchasing API.

automaticallyPopulatePurchasedContent
boolean
sendNotificationWhenNoLongerAssigned
boolean
autoRegisterManagedUsers
boolean
siteId
string
siteName
string
lastSyncTime
string
content
array of objects
object
name
string
licenseCountTotal
integer
licenseCountInUse
integer
licenseCountReported
integer
iconUrl
string
deviceTypes
array of strings
contentType
string
enum
IOS_APP MAC_APP BOOK UNKNOWN

pricingParam
string
enum
STDQ PLUS Unknown

adamId
string

curl --request PATCH \
     --url https://yourserver.jamfcloud.com/api/v1/volume-purchasing-locations/ \
     --header 'accept: application/json' \
     --header 'content-type: application/merge-patch+json' \
     --data '
{
  "autoRegisterManagedUsers": false,
  "name": "Example Location",
  "automaticallyPopulatePurchasedContent": true,
  "sendNotificationWhenNoLongerAssigned": true,
  "siteId": "1",
  "serviceToken": "eyJleHBEYXRlIjoiMjAyMi0wMy0yOVQxNTozNjoyNiswMDAwIiwidG9rZW4iOiJWR2hwY3lCcGN5QnViM1FnWVNCMGIydGxiaTRnU0c5d1pXWjFiR3g1SUdsMElHeHZiMnR6SUd4cGEyVWdZU0IwYjJ0bGJpd2dZblYwSUdsMEozTWdibTkwTGc9PSIsIm9yZ05hbWUiOiJFeGFtcGxlIE9yZyJ9"
}
'

{
  "name": "Example Location",
  "totalPurchasedLicenses": 1,
  "totalUsedLicenses": 1,
  "id": "1",
  "appleId": "testUser@appleId.com",
  "email": "testUser@email.com",
  "organizationName": "Jamf",
  "tokenExpiration": "2022-04-25T21:09:31.661Z",
  "countryCode": "US",
  "locationName": "Example Location",
  "clientContextMismatch": false,
  "automaticallyPopulatePurchasedContent": false,
  "sendNotificationWhenNoLongerAssigned": false,
  "autoRegisterManagedUsers": false,
  "siteId": "1",
  "siteName": "Example Name",
  "lastSyncTime": "2022-09-25T21:09:31.661Z",
  "content": [
    {
      "name": "Example Content",
      "licenseCountTotal": 1,
      "licenseCountInUse": 1,
      "licenseCountReported": 1,
      "iconUrl": "https://is4-ssl.mzstatic.com/image/thumb/Purple113/v4/73/d4/73/73d47332-fefc-d350-2984-5b4a4755a502/AppIcon-0-1x_U007emarketing-0-0-GLES2_U002c0-512MB-sRGB-0-0-0-85-220-0-0-0-6.png/360x216bb.png",
      "deviceTypes": [
        "IOS"
      ],
      "contentType": "IOS_APP",
      "pricingParam": "STDQ",
      "adamId": "748057890"
    }
  ]
}
-----
Retrieve the Volume Purchasing Content for the Volume Purchasing Location with the supplied id
get
https://yourServer.jamfcloud.com/api/v1/volume-purchasing-locations/{id}/content

Retrieves the Volume Purchasing Content for the Volume Purchasing Location with the supplied id

Path Params
id
string
required
Volume Purchasing Location identifier

1
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
Sorting criteria in the format: property:asc/desc. Default sort is name:asc. Multiple sort criteria are supported and must be separated with a comma.


string

id:asc

ADD string
filter
string
Query in the RSQL format, allowing to filter Volume Purchasing Content collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: name, licenseCountTotal, licenseCountInUse, licenseCountReported, contentType, and pricingParam. This param can be combined with paging and sorting.

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
name
string
licenseCountTotal
integer
licenseCountInUse
integer
licenseCountReported
integer
iconUrl
string
deviceTypes
array of strings
contentType
string
enum
IOS_APP MAC_APP BOOK UNKNOWN

pricingParam
string
enum
STDQ PLUS Unknown

adamId
string

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v1/volume-purchasing-locations/1/content?page=0&page-size=100&sort=id%3Aasc' \
     --header 'accept: application/json'

{
  "totalCount": 1,
  "results": [
    {
      "name": "Example Content",
      "licenseCountTotal": 1,
      "licenseCountInUse": 1,
      "licenseCountReported": 1,
      "iconUrl": "https://is4-ssl.mzstatic.com/image/thumb/Purple113/v4/73/d4/73/73d47332-fefc-d350-2984-5b4a4755a502/AppIcon-0-1x_U007emarketing-0-0-GLES2_U002c0-512MB-sRGB-0-0-0-85-220-0-0-0-6.png/360x216bb.png",
      "deviceTypes": [
        "IOS"
      ],
      "contentType": "IOS_APP",
      "pricingParam": "STDQ",
      "adamId": "748057890"
    }
  ]
}
-----
Get specified Volume Purchasing Location history object
get
https://yourServer.jamfcloud.com/api/v1/volume-purchasing-locations/{id}/history

Gets specified Volume Purchasing Location history object

Path Params
id
string
required
instance id of Volume Purchasing Location history record

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
Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma.


string

date:desc

ADD string
filter
string
Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: username, date, note, details. This param can be combined with paging and sorting. Example: filter=username!=admin and details==disabled and date<2019-12-15

Responses

200
Details of Volume Purchasing Location history were found

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
     --url 'https://yourserver.jamfcloud.com/api/v1/volume-purchasing-locations//history?page=0&page-size=100&sort=date%3Adesc' \
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
Add specified Volume Purchasing Location history object notes
post
https://yourServer.jamfcloud.com/api/v1/volume-purchasing-locations/{id}/history

Adds specified Volume Purchasing Location history object notes

Path Params
id
string
required
instance id of Volume Purchasing Location history record

1
Body Params
history note to create

note
string
required
A generic note can sometimes be useful, but generally not.
Responses

201
Notes of Volume Purchase Location history were added

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
     --url https://yourserver.jamfcloud.com/api/v1/volume-purchasing-locations/1/history \
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

Reclaim a Volume Purchasing Location with the supplied id
post
https://yourServer.jamfcloud.com/api/v1/volume-purchasing-locations/{id}/reclaim

Reclaims a Volume Purchasing Location with the supplied id

Path Params
id
string
required
Volume Purchasing Location identifier

Response
202
Request accepted

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/volume-purchasing-locations//reclaim
-----
Revoke licenses for a Volume Purchasing Location with the supplied id
post
https://yourServer.jamfcloud.com/api/v1/volume-purchasing-locations/{id}/revoke-licenses

Revokes licenses for a Volume Purchasing Location with the supplied id. The licenses must be revokable - any asset whose licenses are irrevocable will not be revoked.

Path Params
id
string
required
Volume Purchasing Location identifier

1
Responses
202
Request accepted

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/volume-purchasing-locations/1/revoke-licenses \
     --header 'accept: application/json'