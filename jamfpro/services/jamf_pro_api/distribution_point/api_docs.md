Finds all Distribution Points
get
https://yourServer.jamfcloud.com/api/v1/distribution-points

Finds all Distribution Points

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
Sorts results by one or more criteria, following the format property:asc/desc. Default sort is id:asc. If using multiple criteria, separate with commas. Allows fields such as - name, serverName


string

id:asc

ADD string
filter
string
Filters results. Use RSQL format for query. Allows fields such as - name, serverName, principal, fileSharingConnectionType, and httpsEnabled Can be combined with paging and sorting. Default filter is an empty query and returns all results from the requested page.

Response

200
Success

Response body
object
totalCount
integer
results
array of objects
length ≥ 0
object
id
string
name
string
required
serverName
string
required
principal
boolean
Defaults to false
backupDistributionPointId
string
Defaults to -1
sshUsername
string
localPathToShare
string
fileSharingConnectionType
string
enum
required
Defaults to NONE
Specify the type of connection , Either of fileSharingConnectionType (or) https connection type needs to be enabled using httpsEnabled for a distribution point to be created.

AFP SMB NONE

shareName
string
Required if fileSharingConnectionType is either AFP (or) SMB

workgroup
string
port
integer
Defaults to 139
Required if fileSharingConnectionType is either AFP (or) SMB

readWriteUsername
string
Required if fileSharingConnectionType is either AFP (or) SMB

readOnlyUsername
string
Required if fileSharingConnectionType is either AFP (or) SMB

httpsEnabled
boolean
Defaults to false
Allow downloads over HTTPS - requires installation of a valid SSL certificate

httpsPort
integer
Defaults to 443
Port number of the server - required if HTTPS enabled

httpsContext
string
Path to the share (e.g. if the share is accessible at http://192.168.10.10/JamfShare, the context is "JamfShare") - required if HTTPS enabled

httpsSecurityType
string
enum
Defaults to NONE
Type of authentication required to download files from the distribution point - required if HTTPS enabled

USERNAME_PASSWORD NONE

httpsUsername
string
Required if httpsSecurityType is USERNAME_PASSWORD

enableLoadBalancing
boolean
Defaults to false
This is used to configure load balancing on the backup distribution point. Cannot be enabled when the backup distribution point configured is cloud.

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v1/distribution-points?page=0&page-size=100&sort=id%3Aasc' \
     --header 'accept: application/json'

{
  "totalCount": 1,
  "results": [
    {
      "id": "1",
      "name": "My distribution point",
      "serverName": "My Server",
      "principal": false,
      "backupDistributionPointId": "-1",
      "sshUsername": "john.doe",
      "localPathToShare": "/",
      "fileSharingConnectionType": "AFP",
      "shareName": "My Share",
      "workgroup": "WORKGROUP1",
      "port": 139,
      "readWriteUsername": "john.doe",
      "readOnlyUsername": "john.doe",
      "httpsEnabled": false,
      "httpsPort": 443,
      "httpsContext": "JamfShare",
      "httpsSecurityType": "NONE",
      "httpsUsername": "admin",
      "enableLoadBalancing": false
    }
  ]
}
-----
Create distribution point
post
https://yourServer.jamfcloud.com/api/v1/distribution-points

Create distribution point

Body Params
distribution point to be created

name
string
required
serverName
string
required
principal
boolean
Defaults to false

false
backupDistributionPointId
string
Defaults to -1
-1
sshUsername
string
sshPassword
password
localPathToShare
string
fileSharingConnectionType
string
enum
required
Defaults to NONE
Specify the type of connection , Either of fileSharingConnectionType (or) https connection type needs to be enabled using httpsEnabled for a distribution point to be created.


NONE
Allowed:

AFP

SMB

NONE
shareName
string
Required if fileSharingConnectionType is either AFP (or) SMB

workgroup
string
port
integer
Defaults to 139
Required if fileSharingConnectionType is either AFP (or) SMB

139
readWriteUsername
string
Required if fileSharingConnectionType is either AFP (or) SMB

readWritePassword
password
Required if fileSharingConnectionType is either AFP (or) SMB

readOnlyUsername
string
Required if fileSharingConnectionType is either AFP (or) SMB

readOnlyPassword
password
Required if fileSharingConnectionType is either AFP (or) SMB

httpsEnabled
boolean
Defaults to false
Allow downloads over HTTPS - requires installation of a valid SSL certificate


false
httpsPort
integer
Defaults to 443
Port number of the server - required if HTTPS enabled

443
httpsContext
string
Path to the share (e.g. if the share is accessible at http://192.168.10.10/JamfShare, the context is "JamfShare") - required if HTTPS enabled

httpsSecurityType
string
enum
Defaults to NONE
Type of authentication required to download files from the distribution point - required if HTTPS enabled


NONE
Allowed:

USERNAME_PASSWORD

NONE
httpsUsername
string
Required if httpsSecurityType is USERNAME_PASSWORD

httpsPassword
password
Required if httpsSecurityType is USERNAME_PASSWORD

enableLoadBalancing
boolean
Defaults to false
This is used to configure load balancing on the backup distribution point. Cannot be enabled when the backup distribution point configured is cloud.


false
Response

201
distribution point was created

Response body
object
id
string
length ≥ 1
href
string

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/distribution-points \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "principal": false,
  "backupDistributionPointId": "-1",
  "fileSharingConnectionType": "NONE",
  "port": 139,
  "httpsEnabled": false,
  "httpsPort": 443,
  "httpsSecurityType": "NONE",
  "enableLoadBalancing": false
}
'

{
  "id": "1",
  "href": "https://yourJamfProUrl.jamf/api/v1/resource/1"
}
-----

Delete multiple distribution points at once
post
https://yourServer.jamfcloud.com/api/v1/distribution-points/delete-multiple


Delete multiple distribution points at once

Body Params
ids of the distribution points to be deleted

ids
array of strings

string

1,2

ADD string
Responses
204
All specified distribution points deleted successfully

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/distribution-points/delete-multiple \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '{"ids":["1,2"]}'

-----
Get specified distribution point
get
https://yourServer.jamfcloud.com/api/v1/distribution-points/{id}


Get specified distribution point

Path Params
id
string
required
instance id of distribution point

Responses

200
distribution point found

Response body
object
id
string
name
string
required
serverName
string
required
principal
boolean
Defaults to false
backupDistributionPointId
string
Defaults to -1
sshUsername
string
localPathToShare
string
fileSharingConnectionType
string
enum
required
Defaults to NONE
Specify the type of connection , Either of fileSharingConnectionType (or) https connection type needs to be enabled using httpsEnabled for a distribution point to be created.

AFP SMB NONE

shareName
string
Required if fileSharingConnectionType is either AFP (or) SMB

workgroup
string
port
integer
Defaults to 139
Required if fileSharingConnectionType is either AFP (or) SMB

readWriteUsername
string
Required if fileSharingConnectionType is either AFP (or) SMB

readOnlyUsername
string
Required if fileSharingConnectionType is either AFP (or) SMB

httpsEnabled
boolean
Defaults to false
Allow downloads over HTTPS - requires installation of a valid SSL certificate

httpsPort
integer
Defaults to 443
Port number of the server - required if HTTPS enabled

httpsContext
string
Path to the share (e.g. if the share is accessible at http://192.168.10.10/JamfShare, the context is "JamfShare") - required if HTTPS enabled

httpsSecurityType
string
enum
Defaults to NONE
Type of authentication required to download files from the distribution point - required if HTTPS enabled

USERNAME_PASSWORD NONE

httpsUsername
string
Required if httpsSecurityType is USERNAME_PASSWORD

enableLoadBalancing
boolean
Defaults to false
This is used to configure load balancing on the backup distribution point. Cannot be enabled when the backup distribution point configured is cloud.

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/distribution-points/ \
     --header 'accept: application/json'

{
  "id": "1",
  "name": "My distribution point",
  "serverName": "My Server",
  "principal": false,
  "backupDistributionPointId": "-1",
  "sshUsername": "john.doe",
  "localPathToShare": "/",
  "fileSharingConnectionType": "AFP",
  "shareName": "My Share",
  "workgroup": "WORKGROUP1",
  "port": 139,
  "readWriteUsername": "john.doe",
  "readOnlyUsername": "john.doe",
  "httpsEnabled": false,
  "httpsPort": 443,
  "httpsContext": "JamfShare",
  "httpsSecurityType": "NONE",
  "httpsUsername": "admin",
  "enableLoadBalancing": false
}
-----

Update specified distribution point object
put
https://yourServer.jamfcloud.com/api/v1/distribution-points/{id}


Update specified distribution point object

Path Params
id
string
required
Instance id of distribution point

Body Params
distribution point object to update. ids defined in this body will be ignored

name
string
required
serverName
string
required
principal
boolean
Defaults to false

false
backupDistributionPointId
string
Defaults to -1
-1
sshUsername
string
sshPassword
password
localPathToShare
string
fileSharingConnectionType
string
enum
required
Defaults to NONE
Specify the type of connection , Either of fileSharingConnectionType (or) https connection type needs to be enabled using httpsEnabled for a distribution point to be created.


NONE
Allowed:

AFP

SMB

NONE
shareName
string
Required if fileSharingConnectionType is either AFP (or) SMB

workgroup
string
port
integer
Defaults to 139
Required if fileSharingConnectionType is either AFP (or) SMB

139
readWriteUsername
string
Required if fileSharingConnectionType is either AFP (or) SMB

readWritePassword
password
Required if fileSharingConnectionType is either AFP (or) SMB

readOnlyUsername
string
Required if fileSharingConnectionType is either AFP (or) SMB

readOnlyPassword
password
Required if fileSharingConnectionType is either AFP (or) SMB

httpsEnabled
boolean
Defaults to false
Allow downloads over HTTPS - requires installation of a valid SSL certificate


false
httpsPort
integer
Defaults to 443
Port number of the server - required if HTTPS enabled

443
httpsContext
string
Path to the share (e.g. if the share is accessible at http://192.168.10.10/JamfShare, the context is "JamfShare") - required if HTTPS enabled

httpsSecurityType
string
enum
Defaults to NONE
Type of authentication required to download files from the distribution point - required if HTTPS enabled


NONE
Allowed:

USERNAME_PASSWORD

NONE
httpsUsername
string
Required if httpsSecurityType is USERNAME_PASSWORD

httpsPassword
password
Required if httpsSecurityType is USERNAME_PASSWORD

enableLoadBalancing
boolean
Defaults to false
This is used to configure load balancing on the backup distribution point. Cannot be enabled when the backup distribution point configured is cloud.


false
Response

202
distribution point updated

Response body
object
id
string
name
string
required
serverName
string
required
principal
boolean
Defaults to false
backupDistributionPointId
string
Defaults to -1
sshUsername
string
localPathToShare
string
fileSharingConnectionType
string
enum
required
Defaults to NONE
Specify the type of connection , Either of fileSharingConnectionType (or) https connection type needs to be enabled using httpsEnabled for a distribution point to be created.

AFP SMB NONE

shareName
string
Required if fileSharingConnectionType is either AFP (or) SMB

workgroup
string
port
integer
Defaults to 139
Required if fileSharingConnectionType is either AFP (or) SMB

readWriteUsername
string
Required if fileSharingConnectionType is either AFP (or) SMB

readOnlyUsername
string
Required if fileSharingConnectionType is either AFP (or) SMB

httpsEnabled
boolean
Defaults to false
Allow downloads over HTTPS - requires installation of a valid SSL certificate

httpsPort
integer
Defaults to 443
Port number of the server - required if HTTPS enabled

httpsContext
string
Path to the share (e.g. if the share is accessible at http://192.168.10.10/JamfShare, the context is "JamfShare") - required if HTTPS enabled

httpsSecurityType
string
enum
Defaults to NONE
Type of authentication required to download files from the distribution point - required if HTTPS enabled

USERNAME_PASSWORD NONE

httpsUsername
string
Required if httpsSecurityType is USERNAME_PASSWORD

enableLoadBalancing
boolean
Defaults to false
This is used to configure load balancing on the backup distribution point. Cannot be enabled when the backup distribution point configured is cloud.

curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v1/distribution-points/ \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "principal": false,
  "backupDistributionPointId": "-1",
  "fileSharingConnectionType": "NONE",
  "port": 139,
  "httpsEnabled": false,
  "httpsPort": 443,
  "httpsSecurityType": "NONE",
  "enableLoadBalancing": false
}
'

{
  "id": "1",
  "name": "My distribution point",
  "serverName": "My Server",
  "principal": false,
  "backupDistributionPointId": "-1",
  "sshUsername": "john.doe",
  "localPathToShare": "/",
  "fileSharingConnectionType": "AFP",
  "shareName": "My Share",
  "workgroup": "WORKGROUP1",
  "port": 139,
  "readWriteUsername": "john.doe",
  "readOnlyUsername": "john.doe",
  "httpsEnabled": false,
  "httpsPort": 443,
  "httpsContext": "JamfShare",
  "httpsSecurityType": "NONE",
  "httpsUsername": "admin",
  "enableLoadBalancing": false
}
-----

Remove specified distribution point
delete
https://yourServer.jamfcloud.com/api/v1/distribution-points/{id}


Removes specified distribution point

Path Params
id
string
required
Instance id of distribution point

Response
204
distribution point deleted successfully

curl --request DELETE \
     --url https://yourserver.jamfcloud.com/api/v1/distribution-points/

-----

Update specified distribution point object
patch
https://yourServer.jamfcloud.com/api/v1/distribution-points/{id}


Updates the specified object configuration of a File Share Distribution Point in Jamf Pro. ID path parameter is mandatory and other fields can be updated as a whole or with minimal object.

Path Params
id
string
required
Instance id of distribution point

Body Params
distribution point object to update.

name
string
required
serverName
string
required
principal
boolean
Defaults to false

false
backupDistributionPointId
string
Defaults to -1
-1
sshUsername
string
sshPassword
password
localPathToShare
string
fileSharingConnectionType
string
enum
required
Defaults to NONE
Specify the type of connection , Either of fileSharingConnectionType (or) https connection type needs to be enabled using httpsEnabled for a distribution point to be created.


NONE
Allowed:

AFP

SMB

NONE
shareName
string
Required if fileSharingConnectionType is either AFP (or) SMB

workgroup
string
port
integer
Defaults to 139
Required if fileSharingConnectionType is either AFP (or) SMB

139
readWriteUsername
string
Required if fileSharingConnectionType is either AFP (or) SMB

readWritePassword
password
Required if fileSharingConnectionType is either AFP (or) SMB

readOnlyUsername
string
Required if fileSharingConnectionType is either AFP (or) SMB

readOnlyPassword
password
Required if fileSharingConnectionType is either AFP (or) SMB

httpsEnabled
boolean
Defaults to false
Allow downloads over HTTPS - requires installation of a valid SSL certificate


false
httpsPort
integer
Defaults to 443
Port number of the server - required if HTTPS enabled

443
httpsContext
string
Path to the share (e.g. if the share is accessible at http://192.168.10.10/JamfShare, the context is "JamfShare") - required if HTTPS enabled

httpsSecurityType
string
enum
Defaults to NONE
Type of authentication required to download files from the distribution point - required if HTTPS enabled


NONE
Allowed:

USERNAME_PASSWORD

NONE
httpsUsername
string
Required if httpsSecurityType is USERNAME_PASSWORD

httpsPassword
password
Required if httpsSecurityType is USERNAME_PASSWORD

enableLoadBalancing
boolean
Defaults to false
This is used to configure load balancing on the backup distribution point. Cannot be enabled when the backup distribution point configured is cloud.


false
Responses

202
distribution point updated

Response body
object
id
string
name
string
required
serverName
string
required
principal
boolean
Defaults to false
backupDistributionPointId
string
Defaults to -1
sshUsername
string
localPathToShare
string
fileSharingConnectionType
string
enum
required
Defaults to NONE
Specify the type of connection , Either of fileSharingConnectionType (or) https connection type needs to be enabled using httpsEnabled for a distribution point to be created.

AFP SMB NONE

shareName
string
Required if fileSharingConnectionType is either AFP (or) SMB

workgroup
string
port
integer
Defaults to 139
Required if fileSharingConnectionType is either AFP (or) SMB

readWriteUsername
string
Required if fileSharingConnectionType is either AFP (or) SMB

readOnlyUsername
string
Required if fileSharingConnectionType is either AFP (or) SMB

httpsEnabled
boolean
Defaults to false
Allow downloads over HTTPS - requires installation of a valid SSL certificate

httpsPort
integer
Defaults to 443
Port number of the server - required if HTTPS enabled

httpsContext
string
Path to the share (e.g. if the share is accessible at http://192.168.10.10/JamfShare, the context is "JamfShare") - required if HTTPS enabled

httpsSecurityType
string
enum
Defaults to NONE
Type of authentication required to download files from the distribution point - required if HTTPS enabled

USERNAME_PASSWORD NONE

httpsUsername
string
Required if httpsSecurityType is USERNAME_PASSWORD

enableLoadBalancing
boolean
Defaults to false
This is used to configure load balancing on the backup distribution point. Cannot be enabled when the backup distribution point configured is cloud.

curl --request PATCH \
     --url https://yourserver.jamfcloud.com/api/v1/distribution-points/ \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "principal": false,
  "backupDistributionPointId": "-1",
  "fileSharingConnectionType": "NONE",
  "port": 139,
  "httpsEnabled": false,
  "httpsPort": 443,
  "httpsSecurityType": "NONE",
  "enableLoadBalancing": false
}
'

{
  "id": "1",
  "name": "My distribution point",
  "serverName": "My Server",
  "principal": false,
  "backupDistributionPointId": "-1",
  "sshUsername": "john.doe",
  "localPathToShare": "/",
  "fileSharingConnectionType": "AFP",
  "shareName": "My Share",
  "workgroup": "WORKGROUP1",
  "port": 139,
  "readWriteUsername": "john.doe",
  "readOnlyUsername": "john.doe",
  "httpsEnabled": false,
  "httpsPort": 443,
  "httpsContext": "JamfShare",
  "httpsSecurityType": "NONE",
  "httpsUsername": "admin",
  "enableLoadBalancing": false
}

-----

Get specified distribution point History object
get
https://yourServer.jamfcloud.com/api/v1/distribution-points/{id}/history


Gets specified distribution point history object

Path Params
id
string
required
Instance id of distribution point history

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
Sorts results by one or more criteria, following the format property:asc/desc. Default sort is id:asc. If using multiple criteria, separate with commas.


string

date:desc

ADD string
filter
string
Filters results. Use RSQL format for query. Allows for many fields, including id, name, etc. Can be combined with paging and sorting. Default filter is an empty query and returns all results from the requested page.

Responses

200
Details of distribution point history were found

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
     --url 'https://yourserver.jamfcloud.com/api/v1/distribution-points//history?page=0&page-size=100&sort=date%3Adesc' \
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

Add specified distribution point History object notes
post
https://yourServer.jamfcloud.com/api/v1/distribution-points/{id}/history


Adds specified distribution point History object notes

Path Params
id
string
required
Instance id of distribution point history

Body Params
History note to be created

note
string
required
Responses

201
distribution point History note created successfully

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
     --url https://yourserver.jamfcloud.com/api/v1/distribution-points//history \
     --header 'accept: application/json' \
     --header 'content-type: application/json'

{
  "id": 1,
  "username": "admin",
  "date": "2019-02-04T21:09:31.661Z",
  "note": "Sso settings update",
  "details": "Is SSO Enabled false\\nSelected SSO Provider"
}