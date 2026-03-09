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