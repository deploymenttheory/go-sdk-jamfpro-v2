Returns X.509 details of the active Certificate Authority (CA)
get
https://yourServer.jamfcloud.com/api/v1/pki/certificate-authority/active


Returns X.509 details of the active Certificate Authority (CA)

Response

200
Successful response displays the details of the active Certificate Authority (CA)

Response body
object
subjectX500Principal
string
issuerX500Principal
string
serialNumber
string
version
integer
notAfter
integer
notBefore
integer
signature
object
algorithm
string
algorithmOid
string
value
string
keyUsage
array of strings
keyUsageExtended
array of strings
sha1Fingerprint
string
sha256Fingerprint
string

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/pki/certificate-authority/active \
     --header 'accept: application/json'

{
  "subjectX500Principal": "CN=Jamf JSS Built-in Certificate Authority",
  "issuerX500Principal": "CN=Jamf JSS Built-in Certificate Authority",
  "serialNumber": "00bc43bea0",
  "version": 3,
  "notAfter": 1927739379,
  "notBefore": 1612120179,
  "signature": {
    "algorithm": "SHA256withRSA",
    "algorithmOid": "1.2.840.113549.1.1.11",
    "value": "6874f9b8c60a46c25b6b270c14b9a2949c87b25524868394309b20349f95edd624588ada274e5678a921199d26d0fe5b231fc794eb6e020b7a6c4790cf91ec5d9a5bd4da126f141a657cf4af78df6303327964c57721d82b85af1d46379ac7ec45a24ae3418576688f05fd075a1c9c3d137d0fd8831f4c88ce7698af9c747db983a40fb0480ebfb293bf4889e34a949d4a53251b1abef19d895bcea8e0ce590b22244ad1623624319e6a8b7e7d11aea3d94b77be1a94d28fda58e8df2e398c45e2e9c13473dcc81db01acac8f2c6d21cb5c44371c9ebfba632dcb46838a91808d4e82a35500f370dc71f4156528fbce93137c94eb33d83d41d49483d4dcca5e1"
  },
  "keyUsage": [
    "digitalSignature",
    "keyEncipherment",
    "keyCertSign",
    "cRLSign"
  ],
  "keyUsageExtended": [
    "1.3.6.1.5.5.7.3.1"
  ],
  "sha1Fingerprint": "448a7cc4d899d6a1821258133c24c023a5f558d9",
  "sha256Fingerprint": "660958e14891c67491822687d9ac0e3574562664458111ad875b680995ca472b"
}

-----

Returns X.509 of active Certificate Authority (CA) in DER format
get
https://yourServer.jamfcloud.com/api/v1/pki/certificate-authority/active/der


Returns X.509 of active Certificate Authority (CA) in DER format

Response

200
Successful response returns certificate in DER format

Response body
file

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/pki/certificate-authority/active/der \
     --header 'accept: application/pkix-cert'

-----

Returns active Certificate Authority (CA) in PEM format
get
https://yourServer.jamfcloud.com/api/v1/pki/certificate-authority/active/pem


Returns active Certificate Authority (CA) in PEM format

Response

200
Successful response returns certificate in PEM format.

Response body
file

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/pki/certificate-authority/active/pem \
     --header 'accept: application/pem-certificate-chain'
-----

Returns X.509 details of Certificate Authority (CA) with provided ID
get
https://yourServer.jamfcloud.com/api/v1/pki/certificate-authority/{id}


Returns X.509 details of Certificate Authority (CA) with provided ID

Path Params
id
string
required
UUID of the Certificate Authority (CA)

Responses

200
Successful response displays the details of the Certificate Authority (CA) with provided ID

Response body
object
subjectX500Principal
string
issuerX500Principal
string
serialNumber
string
version
integer
notAfter
integer
notBefore
integer
signature
object
algorithm
string
algorithmOid
string
value
string
keyUsage
array of strings
keyUsageExtended
array of strings
sha1Fingerprint
string
sha256Fingerprint
string

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/pki/certificate-authority/ \
     --header 'accept: application/json'

{
  "subjectX500Principal": "CN=Jamf JSS Built-in Certificate Authority",
  "issuerX500Principal": "CN=Jamf JSS Built-in Certificate Authority",
  "serialNumber": "00bc43bea0",
  "version": 3,
  "notAfter": 1927739379,
  "notBefore": 1612120179,
  "signature": {
    "algorithm": "SHA256withRSA",
    "algorithmOid": "1.2.840.113549.1.1.11",
    "value": "6874f9b8c60a46c25b6b270c14b9a2949c87b25524868394309b20349f95edd624588ada274e5678a921199d26d0fe5b231fc794eb6e020b7a6c4790cf91ec5d9a5bd4da126f141a657cf4af78df6303327964c57721d82b85af1d46379ac7ec45a24ae3418576688f05fd075a1c9c3d137d0fd8831f4c88ce7698af9c747db983a40fb0480ebfb293bf4889e34a949d4a53251b1abef19d895bcea8e0ce590b22244ad1623624319e6a8b7e7d11aea3d94b77be1a94d28fda58e8df2e398c45e2e9c13473dcc81db01acac8f2c6d21cb5c44371c9ebfba632dcb46838a91808d4e82a35500f370dc71f4156528fbce93137c94eb33d83d41d49483d4dcca5e1"
  },
  "keyUsage": [
    "digitalSignature",
    "keyEncipherment",
    "keyCertSign",
    "cRLSign"
  ],
  "keyUsageExtended": [
    "1.3.6.1.5.5.7.3.1"
  ],
  "sha1Fingerprint": "448a7cc4d899d6a1821258133c24c023a5f558d9",
  "sha256Fingerprint": "660958e14891c67491822687d9ac0e3574562664458111ad875b680995ca472b"
}

-----

Returns X.509 current Certificate Authority (CA) with provided ID in DER format
get
https://yourServer.jamfcloud.com/api/v1/pki/certificate-authority/{id}/der


Returns X.509 current Certificate Authority (CA) with provided ID in DER format

Path Params
id
string
required
UUID of the Certificate Authority (CA)

Headers
accept
string
enum
Defaults to application/json
Generated from available response content types


application/pkix-cert
Allowed:

application/json

application/pkix-cert
Responses

200
Successful response returns certificate in DER format

Response body
file

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/pki/certificate-authority//der \
     --header 'accept: application/pkix-cert'

-----

Returns current Certificate Authority (CA) with provided ID in PEM format
get
https://yourServer.jamfcloud.com/api/v1/pki/certificate-authority/{id}/pem


Returns current Certificate Authority (CA) with provided ID in PEM format

Path Params
id
string
required
UUID of the Certificate Authority (CA)

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
Successful response returns certificate in PEM format

Response body
file

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/pki/certificate-authority//pem \
     --header 'accept: application/pem-certificate-chain'