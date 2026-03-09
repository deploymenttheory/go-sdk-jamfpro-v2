Create AD CS Settings configuration for either inbound or outbound mode
post
https://yourServer.jamfcloud.com/api/v1/pki/adcs-settings

Create AD CS Settings configuration and initialize renewal monitor. Once set, the configuration cannot change between inbound and outbound modes.

Body Params
displayName
string
length between 1 and 255
caName
string
length between 1 and 255
fqdn
string
length between 1 and 255
adcsUrl
string
length between 1 and 255
serverCert
object

serverCert object
filename
string
required
length between 1 and 255
Server certificate filename should extend .cer or .pem, and client certificate filename should extend .p12 or .pfx.

data
array of strings
required
Must be base-64 encoded data obtainable by openssl base64 < /file/path/filename.pfx | tr -d '\n' | pbcopy in linux terminal, or similar parsing methods.


ADD string
password
password | null
clientCert
object

clientCert object
filename
string
required
length between 1 and 255
Server certificate filename should extend .cer or .pem, and client certificate filename should extend .p12 or .pfx.

data
array of strings
required
Must be base-64 encoded data obtainable by openssl base64 < /file/path/filename.pfx | tr -d '\n' | pbcopy in linux terminal, or similar parsing methods.


ADD string
password
password | null
revocationEnabled
boolean

true
apiClientId
string
outbound
boolean

true
Response

201
Created

Response body
object
id
string
length ≥ 1
href
string

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/pki/adcs-settings \
     --header 'accept: */*' \
     --header 'content-type: application/json'

{
  "id": "1",
  "href": "https://yourJamfProUrl.jamf/api/v1/resource/1"
}

-----

Validate AD CS Settings server certificate
post
https://yourServer.jamfcloud.com/api/v1/pki/adcs-settings/validate-certificate


Validate AD CS Settings server certificate for file format. Must be base64-encoded X.509 file content, obtainable by openssl base64 < /file/path/filename.pfx | tr -d '\n' | pbcopy in linux terminal, or similar parsing methods.

Body Params
filename
string
required
length between 1 and 255
Server certificate filename should extend .cer or .pem, and client certificate filename should extend .p12 or .pfx.

data
array of strings
required
Must be base-64 encoded data obtainable by openssl base64 < /file/path/filename.pfx | tr -d '\n' | pbcopy in linux terminal, or similar parsing methods.


ADD string
password
password | null
Response
204
AD CS Settings certificate for the Jamf Pro server is of valid file format.

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/pki/adcs-settings/validate-certificate \
     --header 'content-type: application/json'

-----

Validate AD CS Settings client certificate
post
https://yourServer.jamfcloud.com/api/v1/pki/adcs-settings/validate-client-certificate


Validate AD CS Settings client certificate for file format and correct password. Must be base64-encoded PKCS#12 file content, obtainable by openssl base64 < /file/path/filename.pfx | tr -d '\n' | pbcopy in linux terminal, or similar parsing methods. This should only contain a single X.509 certificate.

Body Params
filename
string
required
length between 1 and 255
Server certificate filename should extend .cer or .pem, and client certificate filename should extend .p12 or .pfx.

data
array of strings
required
Must be base-64 encoded data obtainable by openssl base64 < /file/path/filename.pfx | tr -d '\n' | pbcopy in linux terminal, or similar parsing methods.


string


ADD string
password
password | null
Response
204
AD CS Settings certificate for client server is of valid file format and has correct password.

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/pki/adcs-settings/validate-client-certificate \
     --header 'content-type: application/json'

-----

Get AD CS Settings configuration for the ID value
get
https://yourServer.jamfcloud.com/api/v1/pki/adcs-settings/{id}


Get AD CS Settings configuration for the ID value including public key information, but not including any password information.

Path Params
id
string
required
ID of the AD CS Settings configuration.

Response

200
OK

Response body
object
id
string
displayName
string
length between 1 and 255
caName
string
length between 1 and 255
fqdn
string
length between 1 and 255
adcsUrl
string
length between 1 and 255
serverCert
object
filename
string
length ≤ 255
Server certificate filename should extend .cer or .pem, and client certificate filename should extend .p12 or .pfx.

serialNumber
string
subject
string
issuer
string
expirationDate
date-time
clientCert
object
filename
string
length ≤ 255
Server certificate filename should extend .cer or .pem, and client certificate filename should extend .p12 or .pfx.

serialNumber
string
subject
string
issuer
string
expirationDate
date-time
revocationEnabled
boolean
apiClientId
string
outbound
boolean
connectorLastCheckInTimestamp
date-time

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/pki/adcs-settings/ \
     --header 'accept: application/json'

{
  "id": "1",
  "displayName": "Example Display Name",
  "caName": "EXAMPLE-SUBCA02-CA",
  "fqdn": "example-subca02.example.com",
  "adcsUrl": "https://<host-name>.example.com",
  "serverCert": {
    "filename": "example.cer",
    "serialNumber": "123456789456123",
    "subject": "CN=SubjectName, OU=OrganizationalUnit, O=Organization, L=Location, ST=State, C=Country",
    "issuer": "IssuerName",
    "expirationDate": "2018-10-15T16:39:56Z"
  },
  "clientCert": {
    "filename": "example.cer",
    "serialNumber": "123456789456123",
    "subject": "CN=SubjectName, OU=OrganizationalUnit, O=Organization, L=Location, ST=State, C=Country",
    "issuer": "IssuerName",
    "expirationDate": "2018-10-15T16:39:56Z"
  },
  "revocationEnabled": true,
  "apiClientId": "A11B43D6-9ED4-4B29-B726-E2DE747D2410",
  "outbound": true,
  "connectorLastCheckInTimestamp": "2018-10-15T16:39:56Z"
}

-----

Delete AD CS Settings configuration by ID
delete
https://yourServer.jamfcloud.com/api/v1/pki/adcs-settings/{id}


Delete AD CS Settings configuration, only if reassignment of Certificate Authority succeeds and no config profiles are using the configuration.

Path Params
id
string
required
ID of the AD CS Settings configuration.

Responses
204
No Content

curl --request DELETE \
     --url https://yourserver.jamfcloud.com/api/v1/pki/adcs-settings/ \
     --header 'accept: application/json'
-----

Update AD CS Settings configuration
patch
https://yourServer.jamfcloud.com/api/v1/pki/adcs-settings/{id}


Update AD CS Settings configuration, where certificate information must be provided in full, or not at all. Cannot change between inbound and outbound modes.

Path Params
id
string
required
ID of the AD CS Settings configuration.

Body Params
displayName
string
length between 1 and 255
caName
string
length between 1 and 255
fqdn
string
length between 1 and 255
adcsUrl
string
length between 1 and 255
serverCert
object

serverCert object
clientCert
object

clientCert object
revocationEnabled
boolean

true
apiClientId
string
outbound
boolean

true
Responses
204
No Content

curl --request PATCH \
     --url https://yourserver.jamfcloud.com/api/v1/pki/adcs-settings/ \
     --header 'accept: application/json' \
     --header 'content-type: application/merge-patch+json'

-----

Retrieve list of AD CS Settings dependencies
get
https://yourServer.jamfcloud.com/api/v1/pki/adcs-settings/{id}/dependencies


Retrieve list of AD CS Settings dependencies

Path Params
id
string
required
AD CS Settings ID

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
configProfileId
integer
configProfileName
string
configProfileType
string
enum
OSX_CONFIGURATION_PROFILE IOS_CONFIGURATION_PROFILE

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/pki/adcs-settings//dependencies \
     --header 'accept: application/json'

{
  "totalCount": 1,
  "results": [
    {
      "configProfileId": 1,
      "configProfileName": "name",
      "configProfileType": "OSX_CONFIGURATION_PROFILE"
    }
  ]
}

-----

Get specified AD CS Settings history object
get
https://yourServer.jamfcloud.com/api/v1/pki/adcs-settings/{id}/history


Get specified AD CS Settings history object.

Path Params
id
string
required
ID of the AD CS Settings configuration.

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

Responses

200
Details of AD CS Settings history was found.

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

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v1/pki/adcs-settings//history?page=0&page-size=100&sort=date%3Adesc' \
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

Add specified AD CS Settings object note
post
https://yourServer.jamfcloud.com/api/v1/pki/adcs-settings/{id}/history


Adds specified AD CS Settings object note.

Path Params
id
string
required
Instance ID of AD CS Settings history record.

Body Params
AD CS Settings history notes to create.

note
string
required
Responses

201
Notes of AD CS Settings history were added.

Response body
object
id
string
length ≥ 1
href
string

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/pki/adcs-settings//history \
     --header 'accept: application/json' \
     --header 'content-type: application/json'

{
  "id": "1",
  "href": "https://yourJamfProUrl.jamf/api/v1/resource/1"
}