Get all the Authorization details associated with the current api
get
https://yourServer.jamfcloud.com/api/v1/auth

Get all the authorization details associated with the current api token

Additional information about authentication, including a Postman collection, can be found in the Jamf Pro API Overview.

Response

200
Current authorization details.

Response body
object
account
object
id
string
length ≥ 1
username
string
realName
string
email
string
preferences
object

preferences object
language
string
dateFormat
string
region
string
timezone
string
disableRelativeDates
boolean
multiSiteAdmin
boolean
accessLevel
string
enum
FullAccess SiteAccess GroupBasedAccess

privilegeSet
string
enum
ADMINISTRATOR AUDITOR ENROLLMENT CUSTOM

privilegesBySite
object

privilegesBySite object
array of strings
array of strings
array of strings
array of strings
array of strings

View Additional Properties
groupIds
array of strings
currentSiteId
string
accountGroups
array of objects
object
accessLevel
string
enum
FullAccess SiteAccess GroupBasedAccess

privilegeSet
string
enum
ADMINISTRATOR AUDITOR ENROLLMENT CUSTOM

siteId
integer
privileges
array of strings
memberUserIds
array of integers
sites
array of objects
object
id
string
name
string
authenticationType
string
enum
JSS LDAP SAML INVITE NATIVE_APP_API_INTEGRATION DEVICE_SIGNATURE CLOUD_CONNECTOR SYSTEM_ACCOUNT USER_ENROLLMENT CLIENT_CREDENTIALS OIDC M2M

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/auth \
     --header 'accept: application/json'

{
  "account": {
    "id": "1",
    "username": "admin",
    "realName": "IT Bob",
    "email": "ITBob@Jamf.com",
    "preferences": {
      "language": "en",
      "dateFormat": "MM/dd/yyyy",
      "region": "Europe",
      "timezone": "Etc/GMT",
      "disableRelativeDates": false
    },
    "multiSiteAdmin": true,
    "accessLevel": "FullAccess",
    "privilegeSet": "CUSTOM",
    "privilegesBySite": {
      "1": [
        "Read SSO Settings",
        "Delete eBooks"
      ]
    },
    "groupIds": [
      1,
      3
    ],
    "currentSiteId": "1"
  },
  "accountGroups": [
    {
      "accessLevel": "FullAccess",
      "privilegeSet": "CUSTOM",
      "siteId": 1,
      "privileges": "Read SSO Settings",
      "memberUserIds": [
        1,
        3
      ]
    }
  ],
  "sites": [
    {
      "id": "1",
      "name": "Eau Claire"
    }
  ],
  "authenticationType": "JSS"
}

-----