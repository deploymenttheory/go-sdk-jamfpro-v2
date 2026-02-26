Return all Jamf Pro user acounts
get
https://yourServer.jamfcloud.com/api/user

Return all Jamf Pro user acounts.

Response

200
List of all Jamf Pro user acounts.

Response body
array of objects
object
id
integer
username
string
realName
string
email
string
preferences
object

preferences object
isMultiSiteAdmin
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

View Additional Properties
groupIds
array of integers
currentSiteId
integer

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/user \
     --header 'accept: application/json'

[
  {
    "id": 1,
    "username": "admin",
    "realName": "IT Bob",
    "email": "ITBob@Jamf.com",
    "preferences": {
      "language": "en",
      "dateFormat": "MM/dd/yyyy",
      "region": "Europe",
      "timezone": "Etc/GMT",
      "isDisableRelativeDates": false
    },
    "isMultiSiteAdmin": false,
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
    "currentSiteId": 1
  }
]
-----
Update values in the User's current session
post
https://yourServer.jamfcloud.com/api/user/updateSession

Updates values in the user's current session.

Body Params
Values to update in user's current session.

currentSiteId
integer
1
Response

200
The user's current session has been sucessfully updated.

Response body
object
currentSiteId
integer

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/user/updateSession \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '{"currentSiteId":1}'

{
  "currentSiteId": 1
}