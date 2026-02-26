Create a PKI configuration in Jamf Pro for Venafi
post
https://yourServer.jamfcloud.com/api/v1/pki/venafi

Creates a Venafi PKI configuration in Jamf Pro, which can be used to issue certificates

Body Params
name
string
required
Venafi Certificate Authority
proxyAddress
string
localhost:9443
revocationEnabled
boolean

true
clientId
string
jamf-pro
refreshToken
string
qdkP4SrCFKd7tefAVM6N
Responses

201
Successful response creates a Venafi PKI configuration

Response body
object
id
string
length ≥ 1
href
string

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/pki/venafi \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "name": "Venafi Certificate Authority",
  "proxyAddress": "localhost:9443",
  "revocationEnabled": true,
  "clientId": "jamf-pro",
  "refreshToken": "qdkP4SrCFKd7tefAVM6N"
}
'

{
  "id": "1",
  "href": "https://yourJamfProUrl.jamf/api/v1/resource/1"
}
-----
Retrieve a Venafi PKI configuration from Jamf Pro
get
https://yourServer.jamfcloud.com/api/v1/pki/venafi/{id}

Retrieve a Venafi PKI configuration from Jamf Pro

Path Params
id
string
required
ID of the Venafi configuration

1
Responses

200
Successful response returns a Venafi PKI configuration

Response body
object
id
integer
name
string
required
proxyAddress
string
revocationEnabled
boolean
clientId
string
refreshTokenConfigured
boolean

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/pki/venafi/1 \
     --header 'accept: application/json'

{
  "id": 4,
  "name": "Venafi Certificate Authority",
  "proxyAddress": "localhost:9443",
  "revocationEnabled": true,
  "clientId": "jamf-pro",
  "refreshTokenConfigured": true
}
-----
Delete a Venafi PKI configuration from Jamf Pro
delete
https://yourServer.jamfcloud.com/api/v1/pki/venafi/{id}

Delete a Venafi PKI configuration from Jamf Pro

Path Params
id
string
required
ID of the Venafi configuration

1
Responses
204
Successful response deletes Venafi PKI configuration from Jamf Pro

curl --request DELETE \
     --url https://yourserver.jamfcloud.com/api/v1/pki/venafi/1 \
     --header 'accept: application/json'
-----
Update a Venafi PKI configuration in Jamf Pro
patch
https://yourServer.jamfcloud.com/api/v1/pki/venafi/{id}

Update a Venafi PKI configuration in Jamf Pro

Path Params
id
string
required
ID of the Venafi configuration

1
Body Params
name
string
required
Venafi Certificate Authority
proxyAddress
string
localhost:9443
revocationEnabled
boolean

true
clientId
string
jamf-pro
refreshToken
string
qdkP4SrCFKd7tefAVM6N
Responses

200
Successful response returns a Venafi PKI configuration

Response body
object
id
integer
name
string
required
proxyAddress
string
revocationEnabled
boolean
clientId
string
refreshTokenConfigured
boolean

curl --request PATCH \
     --url https://yourserver.jamfcloud.com/api/v1/pki/venafi/1 \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "name": "Venafi Certificate Authority",
  "proxyAddress": "localhost:9443",
  "revocationEnabled": true,
  "clientId": "jamf-pro",
  "refreshToken": "qdkP4SrCFKd7tefAVM6N"
}
'

{
  "id": 4,
  "name": "Venafi Certificate Authority",
  "proxyAddress": "localhost:9443",
  "revocationEnabled": true,
  "clientId": "jamf-pro",
  "refreshTokenConfigured": true
}
-----
Tests the communication between Jamf Pro and a Jamf Pro PKI Proxy Server
get
https://yourServer.jamfcloud.com/api/v1/pki/venafi/{id}/connection-status

Tests the communication between Jamf Pro and a Jamf Pro PKI Proxy Server

Path Params
id
string
required
ID of the Venafi configuration

Responses

200
Successfully connected to Venafi

Response body
object
status
string

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/pki/venafi//connection-status \
     --header 'accept: application/json'

{
  "status": "Successfully connected"
}
-----
Get configuration profile data using specified Venafi CA object
get
https://yourServer.jamfcloud.com/api/v1/pki/venafi/{id}/dependent-profiles

Get configuration profile data using specified Venafi CA object

Path Params
id
string
required
ID of the Venafi configuration

Response

200
Successfully returns a list of configuration profile data connected to the Venafi CA

Response body
object
totalCount
integer
≥ 0
results
array of objects
object
urlPath
string
name
string

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/pki/venafi//dependent-profiles \
     --header 'accept: application/json'

{
  "totalCount": 3,
  "results": [
    {
      "urlPath": "OSXConfigurationProfile.html?id=1",
      "name": "Configuration Profile Name"
    }
  ]
}
-----
Get specified Venafi CA history object
get
https://yourServer.jamfcloud.com/api/v1/pki/venafi/{id}/history

Get specified Venafi CA history object

Path Params
id
string
required
ID of the Venafi configuration

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
Details of Venafi CA history was found

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
     --url 'https://yourserver.jamfcloud.com/api/v1/pki/venafi//history?page=0&page-size=100&sort=date%3Adesc' \
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
Add specified Venafi CA Object Note
post
https://yourServer.jamfcloud.com/api/v1/pki/venafi/{id}/history

Adds specified Venafi CA Object Note

Path Params
id
string
required
instance id of Venafi CA history record

Body Params
venafi ca history notes to create

note
string
required
A generic note can sometimes be useful, but generally not.
Responses

201
Notes of Venafi CA history were added

Response body
object
id
string
length ≥ 1
href
string

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/pki/venafi//history \
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
Downloads a certificate used to secure communication between Jamf Pro and a Jamf Pro PKI Proxy Server
get
https://yourServer.jamfcloud.com/api/v1/pki/venafi/{id}/jamf-public-key

Downloads a certificate for an existing Venafi configuration that can be used to secure communication between Jamf Pro and a Jamf Pro PKI Proxy Server

Path Params
id
string
required
ID of the Venafi configuration

Headers
accept
string
enum
Defaults to application/json
Generated from available response content types


application/pem-certificate-chain
Allowed:

application/json

application/pem-certificate-chain
Responses

200
Successful response downloads the certificate

Response body
file

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/pki/venafi//jamf-public-key \
     --header 'accept: application/pem-certificate-chain'
-----
Regenerates a certificate used to secure communication between Jamf Pro and a Jamf Pro PKI Proxy Server
post
https://yourServer.jamfcloud.com/api/v1/pki/venafi/{id}/jamf-public-key/regenerate

Regenerates a certificate for an existing Venafi configuration that can be used to secure communication between Jamf Pro and a Jamf Pro PKI Proxy Server

Path Params
id
string
required
ID of the Venafi configuration

1
Responses
204
Successful response regenerates the certificate

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/pki/venafi/1/jamf-public-key/regenerate \
     --header 'accept: application/json'
-----
Downloads the PKI Proxy Server public key to secure communication between Jamf Pro and a Jamf Pro PKI Proxy Server
get
https://yourServer.jamfcloud.com/api/v1/pki/venafi/{id}/proxy-trust-store

Downloads the uploaded PKI Proxy Server public key to do basic TLS certificate validation between Jamf Pro and a Jamf Pro PKI Proxy Server

Path Params
id
string
required
ID of the Venafi configuration

1
Headers
accept
string
enum
Defaults to application/json
Generated from available response content types


application/pem-certificate-chain
Allowed:

application/json

application/pem-certificate-chain
Responses

200
Successful response downloads the certificate

Response body
file

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/pki/venafi/1/proxy-trust-store \
     --header 'accept: application/pem-certificate-chain'
-----
Uploads the PKI Proxy Server public key to secure communication between Jamf Pro and a Jamf Pro PKI Proxy Server
post
https://yourServer.jamfcloud.com/api/v1/pki/venafi/{id}/proxy-trust-store

Uploads the PKI Proxy Server public key to do basic TLS certificate validation between Jamf Pro and a Jamf Pro PKI Proxy Server

Path Params
id
string
required
ID of the Venafi configuration

Body Params
No file chosen
Responses
204
Successful response replaces or removes public key

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/pki/venafi/1/proxy-trust-store \
     --header 'accept: application/json' \
     --header 'content-type: application/pem-certificate-chain'
-----
Removes the PKI Proxy Server public key used to secure communication between Jamf Pro and a Jamf Pro PKI Proxy Server
delete
https://yourServer.jamfcloud.com/api/v1/pki/venafi/{id}/proxy-trust-store

Removes the uploaded PKI Proxy Server public key to do basic TLS certificate validation between Jamf Pro and a Jamf Pro PKI Proxy Server

Path Params
id
string
required
ID of the Venafi configuration

1
Responses
204
Successful removes public key

curl --request DELETE \
     --url https://yourserver.jamfcloud.com/api/v1/pki/venafi/1/proxy-trust-store \
     --header 'accept: application/json'
-----