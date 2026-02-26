Generate a SSL Certificate using Jamf Certificate Authority
post
https://yourServer.jamfcloud.com/api/settings/issueTomcatSslCertificate

generate a SSL Certificate using Jamf Certificate Authority

Response
204
SSL certificate created successfully

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/settings/issueTomcatSslCertificate