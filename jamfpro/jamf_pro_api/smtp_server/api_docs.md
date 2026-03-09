Get specified SMTP Server history object
get
https://yourServer.jamfcloud.com/api/v1/smtp-server/history

Get specified SMTP Server history object

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
Sorting criteria in the format: property:asc/desc. Default sort is name:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,username:asc


string

date:desc

ADD string
filter
string
Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: username, date, note, details. This param can be combined with paging and sorting. Example: filter=username!=admin and details==disabled and date<2019-12-15

Responses

200
Details of SMTP Server history were found

Response body
object
totalCount
integer
≥ 0
results
array of objects
object
id
string
length ≥ 1
username
string
date
string
note
string
details
string | null

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v1/smtp-server/history?page=0&page-size=100&sort=date%3Adesc' \
     --header 'accept: application/json'

{
  "totalCount": 1,
  "results": [
    {
      "id": "1",
      "username": "admin",
      "date": "2019-02-04T21:09:31.661Z",
      "note": "Sso settings update",
      "details": "Is SSO Enabled false\\nSelected SSO Provider"
    }
  ]
}

-----
Add SMTP Server history object notes
post
https://yourServer.jamfcloud.com/api/v1/smtp-server/history

Adds SMTP Server history object notes

Body Params
History notes to create

note
string
required
A generic note can sometimes be useful, but generally not.
Responses

201
Notes of SMTP Server history were added

Response body
object
id
string
length ≥ 1
href
string

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/smtp-server/history \
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
Test functionality of an SMTP Server
post
https://yourServer.jamfcloud.com/api/v1/smtp-server/test

Test functionality of an SMTP Server

Body Params
Recipient email to test SMTP Server

recipientEmail
string
required
exampleEmail@example.com
Responses
202
Success

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/smtp-server/test \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "recipientEmail": "exampleEmail@example.com"
}
'
-----
Finds the Jamf Pro SMTP Server information
get
https://yourServer.jamfcloud.com/api/v2/smtp-server

Finds the Jamf Pro SMTP Server information

Response

200
Success

Response body
object
enabled
boolean
required
Defaults to false
authenticationType
string
enum
required
NONE BASIC GRAPH_API GOOGLE_MAIL

connectionSettings
object
host
string
required
port
integer
required
Defaults to 25
encryptionType
string
enum
required
Defaults to NONE
NONE SSL TLS_1_2 TLS_1_1 TLS_1 TLS_1_3

connectionTimeout
integer
required
Defaults to 5
senderSettings
object
required
displayName
string
Defaults to Jamf Pro Server
emailAddress
string
required
basicAuthCredentials
object
username
string
required
graphApiCredentials
object
tenantId
string
required
clientId
string
required
googleMailCredentials
object
clientId
string
required
authentications
array of objects
object
emailAddress
string
required
status
string
enum
required
FAILED UNAUTHENTICATED AUTHENTICATED

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v2/smtp-server \
     --header 'accept: application/json'

{
  "enabled": true,
  "authenticationType": "NONE",
  "connectionSettings": {
    "host": "abcd.server.com",
    "port": 25,
    "encryptionType": "SSL",
    "connectionTimeout": 5
  },
  "senderSettings": {
    "displayName": "Jamf Pro Server",
    "emailAddress": "exampleEmail@example.com"
  },
  "basicAuthCredentials": {
    "username": "exampleEmail@example.com"
  },
  "graphApiCredentials": {
    "tenantId": "c84b7b82-c277-411b-975d-7431b4ce40ac",
    "clientId": "5294f9d1-f723-419c-93db-ff040bf7c947"
  },
  "googleMailCredentials": {
    "clientId": "012345678901-abcdefghijklmnopqrstuvwxyz123456.apps.googleusercontent.com",
    "authentications": [
      {
        "emailAddress": "exampleEmail@example.com",
        "status": "AUTHENTICATED"
      }
    ]
  }
}
-----
Updates Jamf Pro SMTP Server information
put
https://yourServer.jamfcloud.com/api/v2/smtp-server

Updates Jamf Pro SMTP Server information. If requiresAuthentication is set to true, a username and password must be provided

Query Params
oauth-state
string
The OAuth state that was last used to authorize a Google Mail account. This is only required when the authentication type is Google Mail and new accounts are being added.

Body Params
SMTP Server to update

enabled
boolean
required
Defaults to false

false
authenticationType
string
enum
required

GRAPH_API
Allowed:

NONE

BASIC

GRAPH_API

GOOGLE_MAIL
connectionSettings
object

connectionSettings object
host
string
required
abcd.server.com
port
integer
required
Defaults to 25
25
encryptionType
string
enum
required
Defaults to NONE

TLS_1_2
Allowed:

NONE

SSL

TLS_1_2

TLS_1_1

TLS_1

TLS_1_3
connectionTimeout
integer
required
Defaults to 5
5
senderSettings
object
required

senderSettings object
displayName
string
Defaults to Jamf Pro Server
Jamf Pro Server
emailAddress
string
required
exampleEmail@example.com
basicAuthCredentials
object

basicAuthCredentials object
username
string
required
exampleEmail@example.com
password
password
required
•••••
graphApiCredentials
object

graphApiCredentials object
tenantId
string
required
c84b7b82-c277-411b-975d-7431b4ce40ac
clientId
string
required
5294f9d1-f723-419c-93db-ff040bf7c947
clientSecret
password
required
••••••
googleMailCredentials
object

googleMailCredentials object
clientId
string
required
012345678901-abcdefghijklmnopqrstuvwxyz123456.apps.googleusercontent.com
clientSecret
password
required
••••••••
authentications
array of objects

object

emailAddress
string
required
exampleEmail@example.com

ADD object
Responses

200
Success

Response body
object
enabled
boolean
required
Defaults to false
authenticationType
string
enum
required
NONE BASIC GRAPH_API GOOGLE_MAIL

connectionSettings
object
host
string
required
port
integer
required
Defaults to 25
encryptionType
string
enum
required
Defaults to NONE
NONE SSL TLS_1_2 TLS_1_1 TLS_1 TLS_1_3

connectionTimeout
integer
required
Defaults to 5
senderSettings
object
required
displayName
string
Defaults to Jamf Pro Server
emailAddress
string
required
basicAuthCredentials
object
username
string
required
graphApiCredentials
object
tenantId
string
required
clientId
string
required
googleMailCredentials
object
clientId
string
required
authentications
array of objects
object
emailAddress
string
required
status
string
enum
required
FAILED UNAUTHENTICATED AUTHENTICATED

curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v2/smtp-server \
     --header 'accept: application/json' \
     --header 'content-type: application/merge-patch+json' \
     --data '
{
  "enabled": false,
  "authenticationType": "GRAPH_API",
  "connectionSettings": {
    "port": 25,
    "encryptionType": "TLS_1_2",
    "connectionTimeout": 5,
    "host": "abcd.server.com"
  },
  "senderSettings": {
    "displayName": "Jamf Pro Server",
    "emailAddress": "exampleEmail@example.com"
  },
  "basicAuthCredentials": {
    "username": "exampleEmail@example.com",
    "password": "qeqwe"
  },
  "graphApiCredentials": {
    "tenantId": "c84b7b82-c277-411b-975d-7431b4ce40ac",
    "clientId": "5294f9d1-f723-419c-93db-ff040bf7c947",
    "clientSecret": "qweqwe"
  },
  "googleMailCredentials": {
    "clientId": "012345678901-abcdefghijklmnopqrstuvwxyz123456.apps.googleusercontent.com",
    "clientSecret": "qweqweqw",
    "authentications": [
      {
        "emailAddress": "exampleEmail@example.com"
      }
    ]
  }
}
'

{
  "enabled": true,
  "authenticationType": "NONE",
  "connectionSettings": {
    "host": "abcd.server.com",
    "port": 25,
    "encryptionType": "SSL",
    "connectionTimeout": 5
  },
  "senderSettings": {
    "displayName": "Jamf Pro Server",
    "emailAddress": "exampleEmail@example.com"
  },
  "basicAuthCredentials": {
    "username": "exampleEmail@example.com"
  },
  "graphApiCredentials": {
    "tenantId": "c84b7b82-c277-411b-975d-7431b4ce40ac",
    "clientId": "5294f9d1-f723-419c-93db-ff040bf7c947"
  },
  "googleMailCredentials": {
    "clientId": "012345678901-abcdefghijklmnopqrstuvwxyz123456.apps.googleusercontent.com",
    "authentications": [
      {
        "emailAddress": "exampleEmail@example.com",
        "status": "AUTHENTICATED"
      }
    ]
  }
}