Validate keystore for Cloud Identity Provider secure connection
post
https://yourServer.jamfcloud.com/api/v1/ldap-keystore/verify

Validate keystore for Cloud Identity Provider secure connection

Body Params
password
password
required
•••
fileBytes
string
required
WlhoaGJYQnNaU0J2WmlCaElHSmhjMlUyTkNCbGJtTnZaR1ZrSUhaaGJHbGtJSEF4TWk0Z2EyVjVjM1J2Y21VZ1ptbHNaUT09
fileName
string
required
keystore.p12
Responses

200
Keystore verified.

Response body
object
type
string
expirationDate
date-time
subject
string
fileName
string

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/ldap-keystore/verify \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "password": "***",
  "fileBytes": "WlhoaGJYQnNaU0J2WmlCaElHSmhjMlUyTkNCbGJtTnZaR1ZrSUhaaGJHbGtJSEF4TWk0Z2EyVjVjM1J2Y21VZ1ptbHNaUT09",
  "fileName": "keystore.p12"
}
'

{
  "type": "PKCS12",
  "expirationDate": "2030-02-21T12:05:47.244Z",
  "subject": "ST=California, C=US, OU=GSuite, CN=LDAP Client, L=Mountain View, O=Google Inc.",
  "fileName": "keystore.p12"
}

-----
Create Cloud Identity Provider configuration
post
https://yourServer.jamfcloud.com/api/v2/cloud-ldaps

Create new Cloud Identity Provider configuration with unique display name. If mappings not provided, then defaults will be generated instead.

Body Params
Cloud Identity Provider configuration to create

cloudIdPCommon
object
required
A Cloud Identity Provider information for request


cloudIdPCommon object
displayName
string
required
Cloud Identity Provider
providerName
string
enum
required

GOOGLE
Allowed:

GOOGLE

AZURE
server
object
required
A Cloud Identity Provider LDAP server configuration for requests


server object
serverUrl
string
required
ldap.google.com
enabled
boolean
required

true
domainName
string
required
jamf.com
port
integer
required
1 to 65535
636
keystore
object
required
Request with the Base64-encoded keystore file


keystore object
password
password
required
•••
fileBytes
string
required
WlhoaGJYQnNaU0J2WmlCaElHSmhjMlUyTkNCbGJtTnZaR1ZrSUhaaGJHbGtJSEF4TWk0Z2EyVjVjM1J2Y21VZ1ptbHNaUT09
fileName
string
required
keystore.p12
connectionTimeout
integer
required
5 to 600
15
searchTimeout
integer
required
5 to 600
60
useWildcards
boolean
required

true
connectionType
string
enum
required

LDAPS
Allowed:

LDAPS

START_TLS
membershipCalculationOptimizationEnabled
boolean

true
mappings
object
Mappings configurations request for Ldap Cloud Identity Provider configuration


mappings object
userMappings
object
required
Cloud Identity Provider user mappings configuration


userMappings object
objectClassLimitation
string
enum
required

ANY_OBJECT_CLASSES
Allowed:

ANY_OBJECT_CLASSES

ALL_OBJECT_CLASSES
objectClasses
string
required
inetOrgPerson
searchBase
string
required
ou=Users
searchScope
string
enum
required

ALL_SUBTREES
Allowed:

ALL_SUBTREES

FIRST_LEVEL_ONLY
additionalSearchBase
string
userID
string
required
mail
username
string
required
uid
realName
string
required
displayName
emailAddress
string
required
mail
department
string
required
departmentNumber
building
string
required
thing
room
string
required
thing
phone
string
required
thing
position
string
required
title
userUuid
string
required
uid
groupMappings
object
required
Cloud Identity Provider user group mappings configuration


groupMappings object
objectClassLimitation
string
enum
required

ANY_OBJECT_CLASSES
Allowed:

ANY_OBJECT_CLASSES

ALL_OBJECT_CLASSES
objectClasses
string
required
groupOfNames
searchBase
string
required
ou=Groups
searchScope
string
enum
required

ALL_SUBTREES
Allowed:

ALL_SUBTREES

FIRST_LEVEL_ONLY
groupID
string
required
cn
groupName
string
required
cn
groupUuid
string
required
gidNumber
membershipMappings
object
required
Cloud Identity Provider user group membership mappings configuration


membershipMappings object
groupMembershipMapping
string
required
memberOf
Responses

201
Cloud Identity Provider configuration created

Response body
object
id
string
length ≥ 1
href
string

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v2/cloud-ldaps \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "cloudIdPCommon": {
    "providerName": "GOOGLE",
    "displayName": "Cloud Identity Provider"
  },
  "server": {
    "enabled": true,
    "keystore": {
      "password": "***",
      "fileBytes": "WlhoaGJYQnNaU0J2WmlCaElHSmhjMlUyTkNCbGJtTnZaR1ZrSUhaaGJHbGtJSEF4TWk0Z2EyVjVjM1J2Y21VZ1ptbHNaUT09",
      "fileName": "keystore.p12"
    },
    "useWildcards": true,
    "connectionType": "LDAPS",
    "serverUrl": "ldap.google.com",
    "domainName": "jamf.com",
    "port": 636,
    "connectionTimeout": 15,
    "searchTimeout": 60,
    "membershipCalculationOptimizationEnabled": true
  },
  "mappings": {
    "userMappings": {
      "objectClassLimitation": "ANY_OBJECT_CLASSES",
      "searchScope": "ALL_SUBTREES",
      "objectClasses": "inetOrgPerson",
      "searchBase": "ou=Users",
      "userID": "mail",
      "username": "uid",
      "realName": "displayName",
      "emailAddress": "mail",
      "department": "departmentNumber",
      "building": "thing",
      "room": "thing",
      "phone": "thing",
      "position": "title",
      "userUuid": "uid"
    },
    "groupMappings": {
      "objectClassLimitation": "ANY_OBJECT_CLASSES",
      "searchScope": "ALL_SUBTREES",
      "objectClasses": "groupOfNames",
      "searchBase": "ou=Groups",
      "groupID": "cn",
      "groupName": "cn",
      "groupUuid": "gidNumber"
    },
    "membershipMappings": {
      "groupMembershipMapping": "memberOf"
    }
  }
}
'

{
  "id": "1",
  "href": "https://yourJamfProUrl.jamf/api/v1/resource/1"
}
-----
Get default mappings
get
https://yourServer.jamfcloud.com/api/v2/cloud-ldaps/defaults/{provider}/mappings

Get default mappings for Cloud Identity Provider Provider.

Path Params
provider
string
required
Cloud Identity Provider name

Responses

200
Default mappings returned.

Response body
object
userMappings
object
Cloud Identity Provider user mappings configuration

objectClassLimitation
string
enum
required
ANY_OBJECT_CLASSES ALL_OBJECT_CLASSES

objectClasses
string
required
searchBase
string
required
searchScope
string
enum
required
ALL_SUBTREES FIRST_LEVEL_ONLY

additionalSearchBase
string
userID
string
required
username
string
required
realName
string
required
emailAddress
string
required
department
string
required
building
string
required
room
string
required
phone
string
required
position
string
required
userUuid
string
required
groupMappings
object
Cloud Identity Provider user group mappings configuration

objectClassLimitation
string
enum
required
ANY_OBJECT_CLASSES ALL_OBJECT_CLASSES

objectClasses
string
required
searchBase
string
required
searchScope
string
enum
required
ALL_SUBTREES FIRST_LEVEL_ONLY

groupID
string
required
groupName
string
required
groupUuid
string
required
membershipMappings
object
Cloud Identity Provider user group membership mappings configuration

groupMembershipMapping
string
required

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v2/cloud-ldaps/defaults//mappings \
     --header 'accept: application/json'

{
  "userMappings": {
    "objectClassLimitation": "ANY_OBJECT_CLASSES",
    "objectClasses": "inetOrgPerson",
    "searchBase": "ou=Users",
    "searchScope": "ALL_SUBTREES",
    "additionalSearchBase": "",
    "userID": "mail",
    "username": "uid",
    "realName": "displayName",
    "emailAddress": "mail",
    "department": "departmentNumber",
    "building": "",
    "room": "",
    "phone": "",
    "position": "title",
    "userUuid": "uid"
  },
  "groupMappings": {
    "objectClassLimitation": "ANY_OBJECT_CLASSES",
    "objectClasses": "groupOfNames",
    "searchBase": "ou=Groups",
    "searchScope": "ALL_SUBTREES",
    "groupID": "cn",
    "groupName": "cn",
    "groupUuid": "gidNumber"
  },
  "membershipMappings": {
    "groupMembershipMapping": "memberOf"
  }
}
-----
Get default server configuration
get
https://yourServer.jamfcloud.com/api/v2/cloud-ldaps/defaults/{provider}/server-configuration

Get default server configuration for Cloud Identity Provider Identity Provider.

Path Params
provider
string
required
Cloud Identity Provider name

Responses

200
Default server configuration returned.

Response body
object
id
string
enabled
boolean
serverUrl
string
domainName
string
port
integer
1 to 65535
keystore
object
Response with keystore information

type
string
expirationDate
date-time
subject
string
fileName
string
connectionTimeout
integer
≥ 5
searchTimeout
integer
≥ 5
useWildcards
boolean
connectionType
string
enum
LDAPS START_TLS

membershipCalculationOptimizationEnabled
boolean

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v2/cloud-ldaps/defaults//server-configuration \
     --header 'accept: application/json'

{
  "id": "1001",
  "enabled": true,
  "serverUrl": "ldap.google.com",
  "domainName": "jamf.com",
  "port": 636,
  "keystore": {
    "type": "PKCS12",
    "expirationDate": "2030-02-21T12:05:47.244Z",
    "subject": "ST=California, C=US, OU=GSuite, CN=LDAP Client, L=Mountain View, O=Google Inc.",
    "fileName": "keystore.p12"
  },
  "connectionTimeout": 15,
  "searchTimeout": 60,
  "useWildcards": true,
  "connectionType": "LDAPS",
  "membershipCalculationOptimizationEnabled": true
}
-----

Get Cloud Identity Provider configuration with given id.
get
https://yourServer.jamfcloud.com/api/v2/cloud-ldaps/{id}

Get Cloud Identity Provider configuration with given id.

Path Params
id
string
required
Cloud Identity Provider identifier

Responses

200
Cloud Identity Provider configuration returned.

Response body
object
cloudIdPCommon
object
required
A Cloud Identity Provider information

id
string
required
displayName
string
required
providerName
string
enum
required
GOOGLE AZURE

server
object
required
A Cloud Identity Provider LDAP server configuration for responses

id
string
enabled
boolean
serverUrl
string
domainName
string
port
integer
1 to 65535
keystore
object
Response with keystore information


keystore object
connectionTimeout
integer
≥ 5
searchTimeout
integer
≥ 5
useWildcards
boolean
connectionType
string
enum
LDAPS START_TLS

membershipCalculationOptimizationEnabled
boolean
mappings
object
Mappings configuration response for Ldap Cloud Identity Provider configuration

userMappings
object
Cloud Identity Provider user mappings configuration


userMappings object
objectClassLimitation
string
enum
required
ANY_OBJECT_CLASSES ALL_OBJECT_CLASSES

objectClasses
string
required
searchBase
string
required
searchScope
string
enum
required
ALL_SUBTREES FIRST_LEVEL_ONLY

additionalSearchBase
string
userID
string
required
username
string
required
realName
string
required
emailAddress
string
required
department
string
required
building
string
required
room
string
required
phone
string
required
position
string
required
userUuid
string
required
groupMappings
object
Cloud Identity Provider user group mappings configuration


groupMappings object
objectClassLimitation
string
enum
required
ANY_OBJECT_CLASSES ALL_OBJECT_CLASSES

objectClasses
string
required
searchBase
string
required
searchScope
string
enum
required
ALL_SUBTREES FIRST_LEVEL_ONLY

groupID
string
required
groupName
string
required
groupUuid
string
required
membershipMappings
object
Cloud Identity Provider user group membership mappings configuration


membershipMappings object
groupMembershipMapping
string
required

{
  "cloudIdPCommon": {
    "id": "1001",
    "displayName": "Cloud Identity Provider",
    "providerName": "PROVIDER"
  },
  "server": {
    "id": "1001",
    "enabled": true,
    "serverUrl": "ldap.google.com",
    "domainName": "jamf.com",
    "port": 636,
    "keystore": {
      "type": "PKCS12",
      "expirationDate": "2030-02-21T12:05:47.244Z",
      "subject": "ST=California, C=US, OU=GSuite, CN=LDAP Client, L=Mountain View, O=Google Inc.",
      "fileName": "keystore.p12"
    },
    "connectionTimeout": 15,
    "searchTimeout": 60,
    "useWildcards": true,
    "connectionType": "LDAPS",
    "membershipCalculationOptimizationEnabled": true
  },
  "mappings": {
    "userMappings": {
      "objectClassLimitation": "ANY_OBJECT_CLASSES",
      "objectClasses": "inetOrgPerson",
      "searchBase": "ou=Users",
      "searchScope": "ALL_SUBTREES",
      "additionalSearchBase": "",
      "userID": "mail",
      "username": "uid",
      "realName": "displayName",
      "emailAddress": "mail",
      "department": "departmentNumber",
      "building": "",
      "room": "",
      "phone": "",
      "position": "title",
      "userUuid": "uid"
    },
    "groupMappings": {
      "objectClassLimitation": "ANY_OBJECT_CLASSES",
      "objectClasses": "groupOfNames",
      "searchBase": "ou=Groups",
      "searchScope": "ALL_SUBTREES",
      "groupID": "cn",
      "groupName": "cn",
      "groupUuid": "gidNumber"
    },
    "membershipMappings": {
      "groupMembershipMapping": "memberOf"
    }
  }
}
-----
Update Cloud Identity Provider configuration
put
https://yourServer.jamfcloud.com/api/v2/cloud-ldaps/{id}

Update Cloud Identity Provider configuration. Cannot be used for partial updates, all content body must be sent.

Path Params
id
string
required
Cloud Identity Provider identifier

Body Params
Cloud Identity Provider configuration to update

cloudIdPCommon
object
required
A Cloud Identity Provider information


cloudIdPCommon object
server
object
required
A Cloud Identity Provider LDAP server configuration for updates


server object
mappings
object
Mappings configurations request for Ldap Cloud Identity Provider configuration


mappings object
Responses

200
Cloud Identity Provider configuration updated

Response body
object
cloudIdPCommon
object
required
A Cloud Identity Provider information

id
string
required
displayName
string
required
providerName
string
enum
required
GOOGLE AZURE

server
object
required
A Cloud Identity Provider LDAP server configuration for responses

id
string
enabled
boolean
serverUrl
string
domainName
string
port
integer
1 to 65535
keystore
object
Response with keystore information


keystore object
type
string
expirationDate
date-time
subject
string
fileName
string
connectionTimeout
integer
≥ 5
searchTimeout
integer
≥ 5
useWildcards
boolean
connectionType
string
enum
LDAPS START_TLS

membershipCalculationOptimizationEnabled
boolean
mappings
object
Mappings configuration response for Ldap Cloud Identity Provider configuration

userMappings
object
Cloud Identity Provider user mappings configuration


userMappings object
objectClassLimitation
string
enum
required
ANY_OBJECT_CLASSES ALL_OBJECT_CLASSES

objectClasses
string
required
searchBase
string
required
searchScope
string
enum
required
ALL_SUBTREES FIRST_LEVEL_ONLY

additionalSearchBase
string
userID
string
required
username
string
required
realName
string
required
emailAddress
string
required
department
string
required
building
string
required
room
string
required
phone
string
required
position
string
required
userUuid
string
required
groupMappings
object
Cloud Identity Provider user group mappings configuration


groupMappings object
objectClassLimitation
string
enum
required
ANY_OBJECT_CLASSES ALL_OBJECT_CLASSES

objectClasses
string
required
searchBase
string
required
searchScope
string
enum
required
ALL_SUBTREES FIRST_LEVEL_ONLY

groupID
string
required
groupName
string
required
groupUuid
string
required
membershipMappings
object
Cloud Identity Provider user group membership mappings configuration


membershipMappings object
groupMembershipMapping
string
required

curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v2/cloud-ldaps/ \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "cloudIdPCommon": {
    "providerName": "GOOGLE"
  },
  "server": {
    "enabled": true,
    "useWildcards": true,
    "connectionType": "LDAPS"
  },
  "mappings": {
    "userMappings": {
      "objectClassLimitation": "ANY_OBJECT_CLASSES",
      "searchScope": "ALL_SUBTREES"
    },
    "groupMappings": {
      "objectClassLimitation": "ANY_OBJECT_CLASSES",
      "searchScope": "ALL_SUBTREES"
    }
  }
}
'

{
  "cloudIdPCommon": {
    "id": "1001",
    "displayName": "Cloud Identity Provider",
    "providerName": "PROVIDER"
  },
  "server": {
    "id": "1001",
    "enabled": true,
    "serverUrl": "ldap.google.com",
    "domainName": "jamf.com",
    "port": 636,
    "keystore": {
      "type": "PKCS12",
      "expirationDate": "2030-02-21T12:05:47.244Z",
      "subject": "ST=California, C=US, OU=GSuite, CN=LDAP Client, L=Mountain View, O=Google Inc.",
      "fileName": "keystore.p12"
    },
    "connectionTimeout": 15,
    "searchTimeout": 60,
    "useWildcards": true,
    "connectionType": "LDAPS",
    "membershipCalculationOptimizationEnabled": true
  },
  "mappings": {
    "userMappings": {
      "objectClassLimitation": "ANY_OBJECT_CLASSES",
      "objectClasses": "inetOrgPerson",
      "searchBase": "ou=Users",
      "searchScope": "ALL_SUBTREES",
      "additionalSearchBase": "",
      "userID": "mail",
      "username": "uid",
      "realName": "displayName",
      "emailAddress": "mail",
      "department": "departmentNumber",
      "building": "",
      "room": "",
      "phone": "",
      "position": "title",
      "userUuid": "uid"
    },
    "groupMappings": {
      "objectClassLimitation": "ANY_OBJECT_CLASSES",
      "objectClasses": "groupOfNames",
      "searchBase": "ou=Groups",
      "searchScope": "ALL_SUBTREES",
      "groupID": "cn",
      "groupName": "cn",
      "groupUuid": "gidNumber"
    },
    "membershipMappings": {
      "groupMembershipMapping": "memberOf"
    }
  }
}
-----
Delete Cloud Identity Provider configuration.
delete
https://yourServer.jamfcloud.com/api/v2/cloud-ldaps/{id}

Delete Cloud Identity Provider configuration.

Path Params
id
string
required
Cloud Identity Provider identifier

1
Responses
204
Cloud Identity Provider configuration deleted.

curl --request DELETE \
     --url https://yourserver.jamfcloud.com/api/v2/cloud-ldaps/1 \
     --header 'accept: application/json'

-----
Get bind connection pool statistics
get
https://yourServer.jamfcloud.com/api/v2/cloud-ldaps/{id}/connection/bind

Get all search connection pool for chosen Cloud Identity Provider. numConnectionsClosedDefunct - The number of connections that have been closed as defunct. numConnectionsClosedExpired - The number of connections that have been closed because they were expired. numConnectionsClosedUnneeded - The number of connections that have been closed because they were no longer needed. numFailedCheckouts - The number of failed attempts to check out a connection from the pool. numFailedConnectionAttempts - The number of failed attempts to create a connection for use in the pool. numReleasedValid - The number of valid connections released back to the pool. numSuccessfulCheckouts - The number of successful attempts to check out a connection from the pool. numSuccessfulCheckoutsNewConnection - The number of successful checkout attempts that had to create a new connection because none were available. numSuccessfulConnectionAttempts - The number successful attempts to create a connection for use in the pool. maximumAvailableConnections - The maximum number of connections that may be available in the pool at any time. numSuccessfulCheckoutsWithoutWait - The number of successful checkout attempts that were able to take an existing connection without waiting. numSuccessfulCheckoutsAfterWait - The number of successful checkout attempts that retrieved a connection from the pool after waiting for it to become available. numAvailableConnections - The number of connections currently available for use in the pool.

Path Params
id
string
required
Cloud Identity Provider identifier

Responses

200
Cloud Identity Provider bind connection pool statistics returned.

Response body
object
numConnectionsClosedDefunct
int64
numConnectionsClosedExpired
int64
numConnectionsClosedUnneeded
int64
numFailedCheckouts
int64
numFailedConnectionAttempts
int64
numReleasedValid
int64
numSuccessfulCheckouts
int64
numSuccessfulCheckoutsNewConnection
int64
numSuccessfulConnectionAttempts
int64
maximumAvailableConnections
int64
numSuccessfulCheckoutsWithoutWaiting
int64
numSuccessfulCheckoutsAfterWaiting
int64
numAvailableConnections
int64

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v2/cloud-ldaps//connection/bind \
     --header 'accept: application/json'

{
  "numConnectionsClosedDefunct": 1,
  "numConnectionsClosedExpired": 1,
  "numConnectionsClosedUnneeded": 1,
  "numFailedCheckouts": 1,
  "numFailedConnectionAttempts": 1,
  "numReleasedValid": 1,
  "numSuccessfulCheckouts": 1,
  "numSuccessfulCheckoutsNewConnection": 1,
  "numSuccessfulConnectionAttempts": 1,
  "maximumAvailableConnections": 1,
  "numSuccessfulCheckoutsWithoutWaiting": 1,
  "numSuccessfulCheckoutsAfterWaiting": 1,
  "numAvailableConnections": 1
}
-----
Get search connection pool statistics
get
https://yourServer.jamfcloud.com/api/v2/cloud-ldaps/{id}/connection/search

Get all search connection pool for chosen Cloud Identity Provider. numConnectionsClosedDefunct - The number of connections that have been closed as defunct. numConnectionsClosedExpired - The number of connections that have been closed because they were expired. numConnectionsClosedUnneeded - The number of connections that have been closed because they were no longer needed. numFailedCheckouts - The number of failed attempts to check out a connection from the pool. numFailedConnectionAttempts - The number of failed attempts to create a connection for use in the pool. numReleasedValid - The number of valid connections released back to the pool. numSuccessfulCheckouts - The number of successful attempts to check out a connection from the pool. numSuccessfulCheckoutsNewConnection - The number of successful checkout attempts that had to create a new connection because none were available. numSuccessfulConnectionAttempts - The number successful attempts to create a connection for use in the pool. maximumAvailableConnections - The maximum number of connections that may be available in the pool at any time. numSuccessfulCheckoutsWithoutWait - The number of successful checkout attempts that were able to take an existing connection without waiting. numSuccessfulCheckoutsAfterWait - The number of successful checkout attempts that retrieved a connection from the pool after waiting for it to become available. numAvailableConnections - The number of connections currently available for use in the pool.

Path Params
id
string
required
Cloud Identity Provider identifier

Responses

200
Cloud Identity Provider search connection pool statistics returned.

Response body
object
numConnectionsClosedDefunct
int64
numConnectionsClosedExpired
int64
numConnectionsClosedUnneeded
int64
numFailedCheckouts
int64
numFailedConnectionAttempts
int64
numReleasedValid
int64
numSuccessfulCheckouts
int64
numSuccessfulCheckoutsNewConnection
int64
numSuccessfulConnectionAttempts
int64
maximumAvailableConnections
int64
numSuccessfulCheckoutsWithoutWaiting
int64
numSuccessfulCheckoutsAfterWaiting
int64
numAvailableConnections
int64

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v2/cloud-ldaps//connection/search \
     --header 'accept: application/json'

{
  "numConnectionsClosedDefunct": 1,
  "numConnectionsClosedExpired": 1,
  "numConnectionsClosedUnneeded": 1,
  "numFailedCheckouts": 1,
  "numFailedConnectionAttempts": 1,
  "numReleasedValid": 1,
  "numSuccessfulCheckouts": 1,
  "numSuccessfulCheckoutsNewConnection": 1,
  "numSuccessfulConnectionAttempts": 1,
  "maximumAvailableConnections": 1,
  "numSuccessfulCheckoutsWithoutWaiting": 1,
  "numSuccessfulCheckoutsAfterWaiting": 1,
  "numAvailableConnections": 1
}
-----
Tests the communication with the specified cloud connection
get
https://yourServer.jamfcloud.com/api/v2/cloud-ldaps/{id}/connection/status

Tests the communication with the specified cloud connection

Path Params
id
string
required
Cloud Identity Provider identifier

Responses

200
Successfully connected to the specified cloud connection

Response body
object
status
string
curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v2/cloud-ldaps//connection/status \
     --header 'accept: application/json'

{
  "status": "Successfully connected"
}
-----
Get mappings configurations for Cloud Identity Providers server configuration.
get
https://yourServer.jamfcloud.com/api/v2/cloud-ldaps/{id}/mappings

Get all mappings configurations for Cloud Identity Providers server configuration.

Path Params
id
string
required
Cloud Identity Provider identifier

Responses

200
Cloud Identity Provider mappings configuration returned.

Response body
object
userMappings
object
Cloud Identity Provider user mappings configuration

objectClassLimitation
string
enum
required
ANY_OBJECT_CLASSES ALL_OBJECT_CLASSES

objectClasses
string
required
searchBase
string
required
searchScope
string
enum
required
ALL_SUBTREES FIRST_LEVEL_ONLY

additionalSearchBase
string
userID
string
required
username
string
required
realName
string
required
emailAddress
string
required
department
string
required
building
string
required
room
string
required
phone
string
required
position
string
required
userUuid
string
required
groupMappings
object
Cloud Identity Provider user group mappings configuration

objectClassLimitation
string
enum
required
ANY_OBJECT_CLASSES ALL_OBJECT_CLASSES

objectClasses
string
required
searchBase
string
required
searchScope
string
enum
required
ALL_SUBTREES FIRST_LEVEL_ONLY

groupID
string
required
groupName
string
required
groupUuid
string
required
membershipMappings
object
Cloud Identity Provider user group membership mappings configuration

groupMembershipMapping
string
required

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v2/cloud-ldaps//mappings \
     --header 'accept: application/json'

{
  "userMappings": {
    "objectClassLimitation": "ANY_OBJECT_CLASSES",
    "objectClasses": "inetOrgPerson",
    "searchBase": "ou=Users",
    "searchScope": "ALL_SUBTREES",
    "additionalSearchBase": "",
    "userID": "mail",
    "username": "uid",
    "realName": "displayName",
    "emailAddress": "mail",
    "department": "departmentNumber",
    "building": "",
    "room": "",
    "phone": "",
    "position": "title",
    "userUuid": "uid"
  },
  "groupMappings": {
    "objectClassLimitation": "ANY_OBJECT_CLASSES",
    "objectClasses": "groupOfNames",
    "searchBase": "ou=Groups",
    "searchScope": "ALL_SUBTREES",
    "groupID": "cn",
    "groupName": "cn",
    "groupUuid": "gidNumber"
  },
  "membershipMappings": {
    "groupMembershipMapping": "memberOf"
  }
}
-----
Update Cloud Identity Provider mappings configuration.
put
https://yourServer.jamfcloud.com/api/v2/cloud-ldaps/{id}/mappings

Update Cloud Identity Provider mappings configuration. Cannot be used for partial updates, all content body must be sent.

Path Params
id
string
required
Cloud Identity Provider identifier

Body Params
Cloud Identity Provider mappings to update.

userMappings
object
required
Cloud Identity Provider user mappings configuration


userMappings object
objectClassLimitation
string
enum
required

ANY_OBJECT_CLASSES
Allowed:

ANY_OBJECT_CLASSES

ALL_OBJECT_CLASSES
objectClasses
string
required
searchBase
string
required
searchScope
string
enum
required

ALL_SUBTREES
Allowed:

ALL_SUBTREES

FIRST_LEVEL_ONLY
additionalSearchBase
string
userID
string
required
username
string
required
realName
string
required
emailAddress
string
required
department
string
required
building
string
required
room
string
required
phone
string
required
position
string
required
userUuid
string
required
groupMappings
object
required
Cloud Identity Provider user group mappings configuration


groupMappings object
objectClassLimitation
string
enum
required

ANY_OBJECT_CLASSES
Allowed:

ANY_OBJECT_CLASSES

ALL_OBJECT_CLASSES
objectClasses
string
required
searchBase
string
required
searchScope
string
enum
required

ALL_SUBTREES
Allowed:

ALL_SUBTREES

FIRST_LEVEL_ONLY
groupID
string
required
groupName
string
required
groupUuid
string
required
membershipMappings
object
required
Cloud Identity Provider user group membership mappings configuration


membershipMappings object
groupMembershipMapping
string
required
Responses

200
Cloud Identity Provider mappings configuration updated.

Response body
object
userMappings
object
Cloud Identity Provider user mappings configuration

objectClassLimitation
string
enum
required
ANY_OBJECT_CLASSES ALL_OBJECT_CLASSES

objectClasses
string
required
searchBase
string
required
searchScope
string
enum
required
ALL_SUBTREES FIRST_LEVEL_ONLY

additionalSearchBase
string
userID
string
required
username
string
required
realName
string
required
emailAddress
string
required
department
string
required
building
string
required
room
string
required
phone
string
required
position
string
required
userUuid
string
required
groupMappings
object
Cloud Identity Provider user group mappings configuration

objectClassLimitation
string
enum
required
ANY_OBJECT_CLASSES ALL_OBJECT_CLASSES

objectClasses
string
required
searchBase
string
required
searchScope
string
enum
required
ALL_SUBTREES FIRST_LEVEL_ONLY

groupID
string
required
groupName
string
required
groupUuid
string
required
membershipMappings
object
Cloud Identity Provider user group membership mappings configuration

groupMembershipMapping
string
required

curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v2/cloud-ldaps//mappings \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "userMappings": {
    "objectClassLimitation": "ANY_OBJECT_CLASSES",
    "searchScope": "ALL_SUBTREES"
  },
  "groupMappings": {
    "objectClassLimitation": "ANY_OBJECT_CLASSES",
    "searchScope": "ALL_SUBTREES"
  }
}
'

{
  "userMappings": {
    "objectClassLimitation": "ANY_OBJECT_CLASSES",
    "objectClasses": "inetOrgPerson",
    "searchBase": "ou=Users",
    "searchScope": "ALL_SUBTREES",
    "additionalSearchBase": "",
    "userID": "mail",
    "username": "uid",
    "realName": "displayName",
    "emailAddress": "mail",
    "department": "departmentNumber",
    "building": "",
    "room": "",
    "phone": "",
    "position": "title",
    "userUuid": "uid"
  },
  "groupMappings": {
    "objectClassLimitation": "ANY_OBJECT_CLASSES",
    "objectClasses": "groupOfNames",
    "searchBase": "ou=Groups",
    "searchScope": "ALL_SUBTREES",
    "groupID": "cn",
    "groupName": "cn",
    "groupUuid": "gidNumber"
  },
  "membershipMappings": {
    "groupMembershipMapping": "memberOf"
  }
}