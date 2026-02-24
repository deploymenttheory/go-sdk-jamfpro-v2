curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/accounts \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "ldapServerId": -1,
  "siteId": -1,
  "accessLevel": "FullAccess",
  "privilegeLevel": "ADMINISTRATOR",
  "accountStatus": "Enabled",
  "accountType": "DEFAULT"
}
'


Adds new account.
post
https://yourServer.jamfcloud.com/api/v1/accounts

Adds the user account provided.

Body Params
plainPassword
password
username
string
realname
string
email
string
phone
string
ldapServerId
integer
Defaults to -1
-1
distinguishedName
string
siteId
integer
Defaults to -1
-1
accessLevel
string
enum
Defaults to FullAccess
Access level for the account


FullAccess
Allowed:

FullAccess

SiteAccess

GroupBasedAccess
privilegeLevel
string
enum
Defaults to ADMINISTRATOR
Privilege level for the account


ADMINISTRATOR
Allowed:

ADMINISTRATOR

AUDITOR

ENROLLMENT

CUSTOM
changePasswordOnNextLogin
boolean

true
accountStatus
string
enum
Defaults to Enabled
Status of the account


Enabled
Allowed:

Enabled

Disabled
accountType
string
enum
Defaults to DEFAULT
Type of the account


DEFAULT
Allowed:

DEFAULT

FEDERATED
Response

201
Successful response - Jamf Pro user account added

Response body
object
id
string
username
string
realname
string
email
string
phone
string
ldapServerId
integer
Defaults to -1
distinguishedName
string
siteId
integer
Defaults to -1
accessLevel
string
enum
Defaults to FullAccess
Access level for the account

FullAccess SiteAccess GroupBasedAccess

privilegeLevel
string
enum
Defaults to ADMINISTRATOR
Privilege level for the account

ADMINISTRATOR AUDITOR ENROLLMENT CUSTOM

lastPasswordChange
date-time
changePasswordOnNextLogin
boolean
failedLoginAttempts
integer
accountStatus
string
enum
Defaults to Enabled
Status of the account

Enabled Disabled

accountType
string
enum
Defaults to DEFAULT
Type of the account

DEFAULT FEDERATED

{
  "id": "1",
  "username": "testusername",
  "realname": "Bob Jones",
  "email": "bob@jamf.com",
  "phone": "715-999-9999",
  "ldapServerId": -1,
  "distinguishedName": "",
  "siteId": -1,
  "accessLevel": "FullAccess",
  "privilegeLevel": "ADMINISTRATOR",
  "lastPasswordChange": "2017-07-21T17:32:28Z",
  "changePasswordOnNextLogin": true,
  "failedLoginAttempts": 0,
  "accountStatus": "Enabled",
  "accountType": "DEFAULT"
}

-----

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v1/accounts?page=0&page-size=100&sort=username%3Aasc' \
     --header 'accept: application/json'

Get user accounts
get
https://yourServer.jamfcloud.com/api/v1/accounts

Get all user accounts with pagination, sorting, and filtering support.

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
Defaults to username:asc
Sorting criteria in the format: property:asc/desc. Default sort is username:desc. Multiple sort criteria are supported and must be separated with a comma. Accepts fields: id, lastPasswordChange, failedLoginAttempts, username, realname, email, phone, ldapServerId, distinguishedName, siteId, privilegeLevel, changePasswordOnNextLogin, accountStatus. If any other field is passed it will be ignored in sorting operation and/or create unpredictable results.


string

username:asc

ADD string
filter
string
Query in the RSQL format to filter user accounts collection. An empty query returns all results for the requested page. Supported fields: id, lastPasswordChange, failedLoginAttempts, username, realname, email, phone, ldapServerId, distinguishedName, siteId, privilegeLevel, changePasswordOnNextLogin, accountStatus. Multiple conditions can be combined using logical operators. This parameter can be used with paging and sorting parameters. Example: username=="admin" and accountStatus==Enabled and failedLoginAttempts==0

Response

200
User accounts retrieved successfully

Response body
object
totalCount
integer
≥ 0
Total number of user accounts matching the filter criteria

results
array of objects
length ≥ 0
The collection of user accounts for the requested page

object
id
string
username
string
realname
string
email
string
phone
string
ldapServerId
integer
Defaults to -1
distinguishedName
string
siteId
integer
Defaults to -1
accessLevel
string
enum
Defaults to FullAccess
Access level for the account

FullAccess SiteAccess GroupBasedAccess

privilegeLevel
string
enum
Defaults to ADMINISTRATOR
Privilege level for the account

ADMINISTRATOR AUDITOR ENROLLMENT CUSTOM

lastPasswordChange
date-time
changePasswordOnNextLogin
boolean
failedLoginAttempts
integer
accountStatus
string
enum
Defaults to Enabled
Status of the account

Enabled Disabled

accountType
string
enum
Defaults to DEFAULT
Type of the account

DEFAULT FEDERATED

{
  "totalCount": 1,
  "results": [
    {
      "id": "1",
      "username": "testusername",
      "realname": "Bob Jones",
      "email": "bob@jamf.com",
      "phone": "715-999-9999",
      "ldapServerId": -1,
      "distinguishedName": "",
      "siteId": -1,
      "accessLevel": "FullAccess",
      "privilegeLevel": "ADMINISTRATOR",
      "lastPasswordChange": "2017-07-21T17:32:28Z",
      "changePasswordOnNextLogin": true,
      "failedLoginAttempts": 0,
      "accountStatus": "Enabled",
      "accountType": "DEFAULT"
    }
  ]
}

-----

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/accounts/ \
     --header 'accept: application/json'

Gets the user account.
get
https://yourServer.jamfcloud.com/api/v1/accounts/{id}

Gets the user account for the given id.

Path Params
id
string
required
id of target account

Response

200
Successful response - Jamf Pro user account retrieved

Response body
object
id
string
username
string
realname
string
email
string
phone
string
ldapServerId
integer
Defaults to -1
distinguishedName
string
siteId
integer
Defaults to -1
accessLevel
string
enum
Defaults to FullAccess
Access level for the account

FullAccess SiteAccess GroupBasedAccess

privilegeLevel
string
enum
Defaults to ADMINISTRATOR
Privilege level for the account

ADMINISTRATOR AUDITOR ENROLLMENT CUSTOM

lastPasswordChange
date-time
changePasswordOnNextLogin
boolean
failedLoginAttempts
integer
accountStatus
string
enum
Defaults to Enabled
Status of the account

Enabled Disabled

accountType
string
enum
Defaults to DEFAULT
Type of the account

DEFAULT FEDERATED

{
  "id": "1",
  "username": "testusername",
  "realname": "Bob Jones",
  "email": "bob@jamf.com",
  "phone": "715-999-9999",
  "ldapServerId": -1,
  "distinguishedName": "",
  "siteId": -1,
  "accessLevel": "FullAccess",
  "privilegeLevel": "ADMINISTRATOR",
  "lastPasswordChange": "2017-07-21T17:32:28Z",
  "changePasswordOnNextLogin": true,
  "failedLoginAttempts": 0,
  "accountStatus": "Enabled",
  "accountType": "DEFAULT"
}

-----

curl --request DELETE \
     --url https://yourserver.jamfcloud.com/api/v1/accounts/

Deletes the user account.
delete
https://yourServer.jamfcloud.com/api/v1/accounts/{id}

Deletes the user account for the given id.

Path Params
id
string
required
id of target account

Response
