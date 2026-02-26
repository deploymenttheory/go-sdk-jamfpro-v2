Retrieve the URL to directly login to the IdP
get
https://yourServer.jamfcloud.com/api/v1/oidc/direct-idp-login-url

Retrieve the URL to directly login to the IdP

Responses

200
Successful response

Response body
object
url
string
required
Direct IdP login URL to skip unified login page.

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/oidc/direct-idp-login-url \
     --header 'accept: application/json'

{
  "url": "https://jamf-pro-server/oauth2/authorization/idp-region-domain.com"
}
-----
Generate a new keystore used for signing OIDC messages
post
https://yourServer.jamfcloud.com/api/v1/oidc/generate-certificate

Generates a new certificate used for signing OIDC messages

Responses
201
Successful response

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/oidc/generate-certificate \
     --header 'accept: application/json'
-----
Get the public key of the keystore used for signing OIDC messages as a JWT
get
https://yourServer.jamfcloud.com/api/v1/oidc/public-key

Gets the public key of the keystore used for signing OIDC messages as a JWT

Responses

200
Successful response

Response body
object
keys
array of objects
object
kty
string
e
string
use
string
kid
string
alg
string
iat
integer
n
string

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/oidc/public-key \
     --header 'accept: application/json'

{
  "keys": [
    {
      "kty": "RSA",
      "e": "AQAB",
      "use": "sig",
      "kid": "695eccd8c1918105c0f57fcac412651245d778cd301fca44a5ee6f7e0c5320ce",
      "alg": "RS256",
      "iat": 1717518913,
      "n": "yvoRJjwm_JR3z0_MQS5zU8RQCav82PbdZOG8DYcDzRrbklVJZXAGOH16m6egW5B-hb_zxncsSCZTvwmXYiUsWcYE8Lu-CZ0-O2WmpOsk-rc5mklQlEGlD2N3u6MZWIpOofyrbM2AhVssmzgFRHR_22O05KDsNKMKy7ZQaUvvIbqBrnvGRKJMZ4GqzeWyo3rmBpdhmY1wLoPQBAQgwO-TTy5xLhTmLLimAV6ckkLtxjUUdwu4pSKiiJGkJAlDEFXJA5-Y9GE3CZOywj1sJsuHvfHGcgWbEDYVVqdD6PB4NRKmuvj_yJ1BACObaFSLNzfw693FBvuyH1s9D-zGT58Kvw"
    }
  ]
}
-----
Get the public features of the OIDC configuration
get
https://yourServer.jamfcloud.com/api/v1/oidc/public-features

Retrieves public OIDC configuration features.

Response

200
Successful response

Response body
object
jamfIdAuthenticationEnabled
boolean
required
Indicates whether Jamf ID authentication is enabled for this instance. When true, users can authenticate using Jamf ID credentials. When false, Jamf ID login option is not available.

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/oidc/public-features \
     --header 'accept: application/json'

{
  "jamfIdAuthenticationEnabled": true
}
-----
Provide the url to redirect for OIDC login
post
https://yourServer.jamfcloud.com/api/v2/oidc/dispatch

Provide the url to redirect for OIDC login based on email

Body Params
originalUrl
string
required
Original Url

aHR0cHM6Ly9qYW1mLXByby11cmwuY29tL2xvZ2dpbmcuaHRtbA==
emailAddress
string
required
User email address

admin@domain.name
Responses

200
Successful response

Response body
object
idpRedirects
array of objects
object
redirectUrl
string
idpName
string
idpType
string

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v2/oidc/dispatch \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "originalUrl": "aHR0cHM6Ly9qYW1mLXByby11cmwuY29tL2xvZ2dpbmcuaHRtbA==",
  "emailAddress": "admin@domain.name"
}
'

{
  "idpRedirects": [
    {
      "redirectUrl": "/oauth2/authorization/jamf-account-domain.name?original_url=aHR0cHM6Ly9qYW1mLXByby11cmwuY29tL2xvZ2dpbmcuaHRtbA==&login_hint=YWRtaW5AZG9tYWluLm5hbWU=",
      "idpName": "SomeIdp",
      "idpType": "GENERIC_OIDC"
    }
  ]
}
-----