etrieve the certificate currently configured for use with SSO
get
https://yourServer.jamfcloud.com/api/v2/sso/cert

Retrieves the certificate currently configured for use with SSO.

Response

200
Successful operation.

Response body
object
keystore
object
key
string
Defaults to
keys
array of objects
object
id
string
valid
boolean
type
string
enum
PKCS12 JKS NONE

keystoreSetupType
string
enum
NONE UPLOADED GENERATED

keystoreFileName
string
keystoreDetails
object
keys
array of strings
serialNumber
integer
subject
string
issuer
string
expiration
string

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v2/sso/cert \
     --header 'accept: application/json'

{
  "keystore": {
    "key": " ",
    "keys": [
      {
        "id": "1",
        "valid": true
      }
    ],
    "type": "PKCS12",
    "keystoreSetupType": "UPLOADED",
    "keystoreFileName": "keystore.p12"
  },
  "keystoreDetails": {
    "keys": [
      ""
    ],
    "serialNumber": 2322472237,
    "subject": "CN=SSO:jamf.com, OU=JSS, O=JAMF Software, L=Minneapolis, ST=MN, C=US",
    "issuer": "CN= Jamf Pro JSS Built-in Certificate Authority",
    "expiration": "2030-02-24T12:18:32.000"
  }
}
-----
Update the certificate used by Jamf Pro to sign SSO requests to the identify provider
put
https://yourServer.jamfcloud.com/api/v2/sso/cert

Update the certificate used by Jamf Pro to sign SSO requests to the identify provider.

Body Params
keystorePassword
password
required
•••
keystoreFile
string
required
WlhoaGJYQnNaU0J2WmlCaElHSmhjMlUyTkNCbGJtTnZaR1ZrSUhaaGJHbGtJSEF4TWk0Z2EyVjVjM1J2Y21VZ1ptbHNaUT09
keystoreFileName
string
required
keystore.p12
keys
array of objects

object

id
string
1
valid
boolean

false

ADD object
key
string
required
Defaults to
 sadasd
password
password
required
•••••
type
string
enum
required

PKCS12
Allowed:

PKCS12

JKS

NONE
keystoreSetupType
string
enum

UPLOADED
Allowed:

NONE

UPLOADED

GENERATED
Responses

200
Successfully changed the keystore.

Response body
object
keystore
object
key
string
Defaults to
keys
array of objects
object
id
string
valid
boolean
type
string
enum
PKCS12 JKS NONE

keystoreSetupType
string
enum
NONE UPLOADED GENERATED

keystoreFileName
string
keystoreDetails
object
keys
array of strings
serialNumber
integer
subject
string
issuer
string
expiration
string

curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v2/sso/cert \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "key": " sadasd",
  "type": "PKCS12",
  "keystorePassword": "***",
  "keystoreFile": "WlhoaGJYQnNaU0J2WmlCaElHSmhjMlUyTkNCbGJtTnZaR1ZrSUhaaGJHbGtJSEF4TWk0Z2EyVjVjM1J2Y21VZ1ptbHNaUT09",
  "keystoreFileName": "keystore.p12",
  "keys": [
    {
      "id": "1",
      "valid": false
    }
  ],
  "password": "asdas",
  "keystoreSetupType": "UPLOADED"
}
'

{
  "keystore": {
    "key": " ",
    "keys": [
      {
        "id": "1",
        "valid": true
      }
    ],
    "type": "PKCS12",
    "keystoreSetupType": "UPLOADED",
    "keystoreFileName": "keystore.p12"
  },
  "keystoreDetails": {
    "keys": [
      ""
    ],
    "serialNumber": 2322472237,
    "subject": "CN=SSO:jamf.com, OU=JSS, O=JAMF Software, L=Minneapolis, ST=MN, C=US",
    "issuer": "CN= Jamf Pro JSS Built-in Certificate Authority",
    "expiration": "2030-02-24T12:18:32.000"
  }
}
-----
Jamf Pro will generate a new certificate and use it to sign SSO
post
https://yourServer.jamfcloud.com/api/v2/sso/cert

Jamf Pro will generate a new certificate and use it to sign SSO requests to the identity provider.

Response

200
Newly generated will be set and returned.

Response body
object
keystore
object
key
string
Defaults to
keys
array of objects
object
id
string
valid
boolean
type
string
enum
PKCS12 JKS NONE

keystoreSetupType
string
enum
NONE UPLOADED GENERATED

keystoreFileName
string
keystoreDetails
object
keys
array of strings
serialNumber
integer
subject
string
issuer
string
expiration
string

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v2/sso/cert \
     --header 'accept: application/json'

{
  "keystore": {
    "key": " ",
    "keys": [
      {
        "id": "1",
        "valid": true
      }
    ],
    "type": "PKCS12",
    "keystoreSetupType": "UPLOADED",
    "keystoreFileName": "keystore.p12"
  },
  "keystoreDetails": {
    "keys": [
      ""
    ],
    "serialNumber": 2322472237,
    "subject": "CN=SSO:jamf.com, OU=JSS, O=JAMF Software, L=Minneapolis, ST=MN, C=US",
    "issuer": "CN= Jamf Pro JSS Built-in Certificate Authority",
    "expiration": "2030-02-24T12:18:32.000"
  }
}
-----
Delete the currently configured certificate used by SSO
delete
https://yourServer.jamfcloud.com/api/v2/sso/cert

Deletes the currently configured certificate used by SSO.

Response
204
Operation successful.

curl --request DELETE \
     --url https://yourserver.jamfcloud.com/api/v2/sso/cert

-----
Download the certificate currently configured for use with Jamf Pro's SSO configuration
get
https://yourServer.jamfcloud.com/api/v2/sso/cert/download

Downloads the certificate currently configured for use with Jamf Pro's SSO configuration

Response

200
Request successful

Response body
file

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v2/sso/cert/download \
     --header 'accept: text/plain'
-----
Parse the certificate to get details about certificate type and keys needed to upload certificate file
post
https://yourServer.jamfcloud.com/api/v2/sso/cert/parse

Parse the certificate to get details about certificate type and keys needed to upload certificate file.

Body Params

SsoKeystore
Responses

200
Successfully parsed the certificate.

Response body
object
key
string
Defaults to
keys
array of objects
object
id
string
valid
boolean
type
string
enum
PKCS12 JKS NONE

keystoreSetupType
string
enum
NONE UPLOADED GENERATED

keystoreFile
array of strings
keystoreFileName
string

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v2/sso/cert/parse \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "key": " ",
  "type": "PKCS12"
}
'

{
  "key": " ",
  "keys": [
    {
      "id": "1",
      "valid": true
    }
  ],
  "type": "PKCS12",
  "keystoreSetupType": "UPLOADED",
  "keystoreFile": "ZXhhbXBsZSBvZiBhIGJhc2U2NCBlbmNvZGVkIHZhbGlkIHAxMi4ga2V5c3RvcmUgZmlsZQ==",
  "keystoreFileName": "keystore.p12"
}
-----