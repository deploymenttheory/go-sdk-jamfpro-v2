Retrieve the current Single Sign On configuration settings
get
https://yourServer.jamfcloud.com/api/v3/sso

Retrieves the current Single Sign On configuration settings

Response

200
Successful response

Response body
object
configurationType
string
enum
required
SAML OIDC OIDC_WITH_SAML

oidcSettings
object
required
userMapping
string
enum
required
USERNAME EMAIL

jamfIdAuthenticationEnabled
boolean
Defaults to true
usernameAttributeClaimMapping
string
enum
USERNAME EMAIL

samlSettings
object
required
tokenExpirationDisabled
boolean
Defaults to false
userAttributeEnabled
boolean
Defaults to false
userAttributeName
string
Defaults to
userMapping
string
enum
USERNAME EMAIL

groupAttributeName
string
Defaults to http://schemas.xmlsoap.org/claims/Group
groupRdnKey
string
Defaults to
idpProviderType
string
enum
ADFS OKTA GOOGLE SHIBBOLETH ONELOGIN PING CENTRIFY AZURE OTHER

idpUrl
string
entityId
string
metadataFileName
string
otherProviderTypeName
string
Defaults to
federationMetadataFile
string
metadataSource
string
enum
URL FILE UNKNOWN

sessionTimeout
int32
Defaults to 480
ssoForEnrollmentEnabled
boolean
required
Defaults to false
ssoBypassAllowed
boolean
required
Defaults to false
ssoEnabled
boolean
required
Defaults to false
ssoForMacOsSelfServiceEnabled
boolean
required
Defaults to false
enrollmentSsoForAccountDrivenEnrollmentEnabled
boolean
required
Defaults to false
enrollmentSsoConfig
object
hosts
array of strings
length ≥ 0
Defaults to
managementHint
string
groupEnrollmentAccessEnabled
boolean
required
Defaults to false
groupEnrollmentAccessName
string
Defaults to

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v3/sso \
     --header 'accept: application/json'

{
  "configurationType": "SAML",
  "oidcSettings": {
    "userMapping": "USERNAME",
    "jamfIdAuthenticationEnabled": true,
    "usernameAttributeClaimMapping": "EMAIL"
  },
  "samlSettings": {
    "tokenExpirationDisabled": false,
    "userAttributeEnabled": false,
    "userAttributeName": " ",
    "userMapping": "USERNAME",
    "groupAttributeName": "http://schemas.xmlsoap.org/claims/Group",
    "groupRdnKey": " ",
    "idpProviderType": "ADFS",
    "idpUrl": "https://example.idp.com/app/id/sso/saml/metadata",
    "entityId": "saml/metadata",
    "metadataFileName": "if MetadataSource is set to URL, remove this field",
    "otherProviderTypeName": " ",
    "federationMetadataFile": "WlhoaGJYQnNaU0J2WmlCaElHSmhjMlUyTkNCbGJtTnZaR1ZrSUhaaGJHbGtJSEF4TWk0Z2EyVjVjM1J2Y21VZ1ptbHNaUT09",
    "metadataSource": "URL",
    "sessionTimeout": 480
  },
  "ssoForEnrollmentEnabled": false,
  "ssoBypassAllowed": false,
  "ssoEnabled": false,
  "ssoForMacOsSelfServiceEnabled": false,
  "enrollmentSsoForAccountDrivenEnrollmentEnabled": false,
  "enrollmentSsoConfig": {
    "hosts": [
      "dev-12324233.okta.com",
      "example.okta.com"
    ],
    "managementHint": ""
  },
  "groupEnrollmentAccessEnabled": false,
  "groupEnrollmentAccessName": " "
}
-----
Updates the current Single Sign On configuration settings
put
https://yourServer.jamfcloud.com/api/v3/sso

Updates the current Single Sign On configuration settings

Body Params
configurationType
string
enum
required

SAML
Allowed:

SAML

OIDC

OIDC_WITH_SAML
oidcSettings
object
required

oidcSettings object
samlSettings
object
required

samlSettings object
ssoForEnrollmentEnabled
boolean
required
Defaults to false

false
ssoBypassAllowed
boolean
required
Defaults to false

false
ssoEnabled
boolean
required
Defaults to false

false
ssoForMacOsSelfServiceEnabled
boolean
required
Defaults to false

false
enrollmentSsoForAccountDrivenEnrollmentEnabled
boolean
required
Defaults to false

false
enrollmentSsoConfig
object

enrollmentSsoConfig object
groupEnrollmentAccessEnabled
boolean
required
Defaults to false

false
groupEnrollmentAccessName
string
Defaults to
 
Responses

200
The update was successful and the newly updated object is returned.

Response body
object
configurationType
string
enum
required
SAML OIDC OIDC_WITH_SAML

oidcSettings
object
required
userMapping
string
enum
required
USERNAME EMAIL

jamfIdAuthenticationEnabled
boolean
Defaults to true
usernameAttributeClaimMapping
string
enum
USERNAME EMAIL

samlSettings
object
required
tokenExpirationDisabled
boolean
Defaults to false
userAttributeEnabled
boolean
Defaults to false
userAttributeName
string
Defaults to
userMapping
string
enum
USERNAME EMAIL

groupAttributeName
string
Defaults to http://schemas.xmlsoap.org/claims/Group
groupRdnKey
string
Defaults to
idpProviderType
string
enum
ADFS OKTA GOOGLE SHIBBOLETH ONELOGIN PING CENTRIFY AZURE OTHER

idpUrl
string
entityId
string
metadataFileName
string
otherProviderTypeName
string
Defaults to
federationMetadataFile
string
metadataSource
string
enum
URL FILE UNKNOWN

sessionTimeout
int32
Defaults to 480
ssoForEnrollmentEnabled
boolean
required
Defaults to false
ssoBypassAllowed
boolean
required
Defaults to false
ssoEnabled
boolean
required
Defaults to false
ssoForMacOsSelfServiceEnabled
boolean
required
Defaults to false
enrollmentSsoForAccountDrivenEnrollmentEnabled
boolean
required
Defaults to false
enrollmentSsoConfig
object
hosts
array of strings
length ≥ 0
Defaults to
managementHint
string
groupEnrollmentAccessEnabled
boolean
required
Defaults to false
groupEnrollmentAccessName
string
Defaults to

curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v3/sso \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "configurationType": "SAML",
  "oidcSettings": {
    "userMapping": "USERNAME",
    "jamfIdAuthenticationEnabled": true
  },
  "samlSettings": {
    "tokenExpirationDisabled": false,
    "userAttributeEnabled": false,
    "userAttributeName": " ",
    "groupAttributeName": "http://schemas.xmlsoap.org/claims/Group",
    "groupRdnKey": " ",
    "otherProviderTypeName": " ",
    "sessionTimeout": 480
  },
  "ssoForEnrollmentEnabled": false,
  "ssoBypassAllowed": false,
  "ssoEnabled": false,
  "ssoForMacOsSelfServiceEnabled": false,
  "enrollmentSsoForAccountDrivenEnrollmentEnabled": false,
  "groupEnrollmentAccessEnabled": false,
  "groupEnrollmentAccessName": " "
}
'

{
  "configurationType": "SAML",
  "oidcSettings": {
    "userMapping": "USERNAME",
    "jamfIdAuthenticationEnabled": true,
    "usernameAttributeClaimMapping": "EMAIL"
  },
  "samlSettings": {
    "tokenExpirationDisabled": false,
    "userAttributeEnabled": false,
    "userAttributeName": " ",
    "userMapping": "USERNAME",
    "groupAttributeName": "http://schemas.xmlsoap.org/claims/Group",
    "groupRdnKey": " ",
    "idpProviderType": "ADFS",
    "idpUrl": "https://example.idp.com/app/id/sso/saml/metadata",
    "entityId": "saml/metadata",
    "metadataFileName": "if MetadataSource is set to URL, remove this field",
    "otherProviderTypeName": " ",
    "federationMetadataFile": "WlhoaGJYQnNaU0J2WmlCaElHSmhjMlUyTkNCbGJtTnZaR1ZrSUhaaGJHbGtJSEF4TWk0Z2EyVjVjM1J2Y21VZ1ptbHNaUT09",
    "metadataSource": "URL",
    "sessionTimeout": 480
  },
  "ssoForEnrollmentEnabled": false,
  "ssoBypassAllowed": false,
  "ssoEnabled": false,
  "ssoForMacOsSelfServiceEnabled": false,
  "enrollmentSsoForAccountDrivenEnrollmentEnabled": false,
  "enrollmentSsoConfig": {
    "hosts": [
      "dev-12324233.okta.com",
      "example.okta.com"
    ],
    "managementHint": ""
  },
  "groupEnrollmentAccessEnabled": false,
  "groupEnrollmentAccessName": " "
}
-----
Retrieve the list of Enrollment Customizations using SSO
get
https://yourServer.jamfcloud.com/api/v3/sso/dependencies

Retrieves the list of Enrollment Customizations using SSO

Response

200
Success

Response body
object
dependencies
array of objects
object
name
string
humanReadableName
string
hyperlink
string

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v3/sso/dependencies \
     --header 'accept: application/json'

{
  "dependencies": [
    {
      "name": "Name",
      "humanReadableName": "Enrollment Customization",
      "hyperlink": "/enrollment-customization/id"
    }
  ]
}
-----
Disable SSO
post
https://yourServer.jamfcloud.com/api/v3/sso/disable

Disable SSO

Responses
202
SSO has been disabled

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v3/sso/disable
-----
Get SSO history object
get
https://yourServer.jamfcloud.com/api/v3/sso/history

Gets SSO history object

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
Defaults to id:desc
Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc


string

id:desc

ADD string
filter
string
Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: username, date, note, details. This param can be combined with paging and sorting. Example: filter=username!=admin and details==disabled and date<2019-12-15

Response

200
Details of SSO history were found

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
     --url 'https://yourserver.jamfcloud.com/api/v3/sso/history?page=0&page-size=100&sort=id%3Adesc' \
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
Add SSO history object notes
post
https://yourServer.jamfcloud.com/api/v3/sso/history

Adds SSO history object notes

Body Params
history notes to create

note
string
required
A generic note can sometimes be useful, but generally not.
Responses

201
Notes of SSO history were added

Response body
object
id
string
length ≥ 1
href
string

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v3/sso/history \
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
Download the Jamf Pro SAML metadata file
get
https://yourServer.jamfcloud.com/api/v3/sso/metadata/download

Download the Jamf Pro SAML metadata file

Response

200
Successful resposne

Response body
file

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v3/sso/metadata/download \
     --header 'accept: text/plain'
-----
