Get the cloud distribution point Details.
get
https://yourServer.jamfcloud.com/api/v1/cloud-distribution-point


Retrieves the details of the cloud distribution point. The distribution point exists only when a content delivery network (CDN) is configured, such as Jamf Cloud(JAMF_CLOUD), Rackspace Cloud Files(RACKSPACE_CLOUD_FILES), Amazon Web Services(AMAZON_S3) or Akamai(AKAMAI). If the cdnType is NONE the response will be NONE empty CDP object, indicating no distribution point is set up.

Responses

200
The request was successful, and the details of the cloud distribution point (CDP) have been retrieved. The response contains the current configuration of the CDP, including the CDN service being used (e.g., Jamf Cloud(JAMF_CLOUD), Rackspace Cloud Files(RACKSPACE_CLOUD_FILES), Amazon Web Services(AMAZON_S3), Akamai(AKAMAI)), the CDN URL, and the status of the test connection.

Response body
object
hasConnectionSucceeded
boolean
required
Defaults to false
Indicates whether the connection to the cloud distribution point was successful. If true, the connection was successful. If false, the connection failed.
Possible values are:
false
true

message
string
required
A message detailing the result of the connection test. This could be a success message or an error message if the connection failed.

inventoryId
string
Defaults to 0
The unique identifier (inventoryId) that links the cloud distribution point to its inventory data. By default, its value is 0, and it increments by +1 based on the existing inventory ID present in the table for each new cloud distribution point configuration. If the cdnType is set to NONE in the next configuration, the ID resets and starts from 1.

cdnType
string
enum
required
Defaults to NONE
Specifies the content delivery network (CDN) used to distribute content for the cloud distribution point.

NONE JAMF_CLOUD RACKSPACE_CLOUD_FILES AMAZON_S3 AKAMAI

master
boolean
Defaults to false
Use as principal distribution point. Use as the authoritative source for all files.
Possible values are:
false
true

username
string
required
The username or access key used for authenticating with the selected content delivery network (CDN). This field is required when the cdnType is set to Rackspace Cloud Files(RACKSPACE_CLOUD_FILES), Amazon Web Services(AMAZON_S3), or Akamai(AKAMAI), as it is used to authenticate and authorize access to the respective cloud services.

For Rackspace Cloud Files(RACKSPACE_CLOUD_FILES), this is typically the username associated with your Rackspace cloud account.
For Amazon Web Services(AMAZON_S3), this corresponds to the Access Key ID used to interact with Amazon Web Services(AMAZON_S3) resources.
For Akamai(AKAMAI), this is the username used for API authentication to access Akamai's content delivery services. If the cdnType is None, this field is not applicable.
directory
string
The directory or path for content delivery in Akamai. This field is required when the cdnType is set to Akamai(AKAMAI) and specifies where content is stored within Akamai's system.

cdnUrl
string
The CDN URL for the cloud distribution point. The URL format varies depending on the selected CDN provider:

Rackspace Cloud Files(RACKSPACE_CLOUD_FILES)
Amazon Web Services(AMAZON_S3)
Akamai(AKAMAI)
The cdnUrl should point to the content distribution location where software or other content is stored and made available for distribution.

uploadUrl
string
The URL used to upload files to Akamai's NetStorage. This field is required when the cdnType is set to Akamai(AKAMAI). It specifies where content should be uploaded to Akamaiâ€™s cloud storage before being distributed via their CDN. The upload typically uses FTP or SFTP.

downloadUrl
string
The URL used to access and download content from Akamai's EdgeSuite. This field is required when the cdnType is set to Akamai(AKAMAI). It specifies the endpoint from which files are retrieved by devices or users.

secondaryAuthRequired
boolean
Defaults to false
Enable Remote Authentication.Authorize requests for files stored on the distribution point. This field is required when the cdnType is set to Akamai(AKAMAI).
Possible values are:
false
true

secondaryAuthStatusCode
integer
Defaults to 200
Secondary Auth Status Code. Configure the HTTP response code that will be returned by Jamf Pro during remote authentication. This field is required when the cdnType is set to Akamai(AKAMAI) and secondaryAuthRequired is true.

secondaryAuthTimeToLive
integer
≥ 1
Defaults to 3600
Secondary Auth Time To Live. Number of seconds before the authorization token expires. This field is required when the cdnType is set to Akamai(AKAMAI) and secondaryAuthRequired is true.

requireSignedUrls
boolean
Defaults to false
Amazon Sign Url. It restrict access to requests that use a signed URL. This field is required when the cdnType is set to Amazon Web Services(AMAZON_S3).
Possible values are:
false
true

keyPairId
string
The CloudFront Access Key ID (keyPairId) is part of the credentials used to generate signed URLs for secure access to content in a CloudFront distribution. When using AWS, this key is paired with the CloudFront Secret Access Key to create the signed URL, ensuring that only authorized users can access specific content within a specified timeframe. This field is required when the cdnType is set to Amazon Web Services(AMAZON_S3) and requireSignedUrls is true.

expirationSeconds
integer
≥ 1
Defaults to 3600
Signed URL Expiration. Number of seconds before the signed URL expires, This field is required when the cdnType is set to Amazon Web Services(AMAZON_S3) and requireSignedUrls is true.

privateKey
string
The CloudFront Private Key file is required when the cdnType is set to Amazon Web Services(AMAZON_S3) and requireSignedUrls parameter is enabled. This private key is used for signing URLs for restricted access to CloudFront-distributed content. The private key allows secure URL generation for signed URLs, ensuring that only authorized users can access certain content. The key must be uploaded in one of the following formats:

.pem: A Privacy-Enhanced Mail (PEM) file containing the private key in base64 encoded format.
.der: A Distinguished Encoding Rules (DER) encoded file, which is a binary format for the private key. The uploaded file should be kept secure, as it provides the ability to generate signed URLs with access to protected content.

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/cloud-distribution-point \
     --header 'accept: application/json'

{
  "hasConnectionSucceeded": false,
  "message": "Cannot contact JCDS",
  "inventoryId": "0",
  "cdnType": "NONE",
  "master": false,
  "username": "Admin",
  "directory": "123456",
  "cdnUrl": "https://example-cdn-url.com/content-path",
  "uploadUrl": "ftp://mycompany.upload.akamai.com",
  "downloadUrl": "https://download.mycompany.com",
  "secondaryAuthRequired": false,
  "secondaryAuthStatusCode": 200,
  "secondaryAuthTimeToLive": 3600,
  "requireSignedUrls": false,
  "keyPairId": "K1R8C5EXAMPLE",
  "expirationSeconds": 3600,
  "privateKey": "string"
}
-----

Create cloud distribution point
post
https://yourServer.jamfcloud.com/api/v1/cloud-distribution-point


Creates cloud distribution point. This operation is triggered when the content delivery network (CDN) settings change, specifically when the network type is updated from "None" to any other supported type. Upon successful creation,the API returns the updated details of the cloud distribution point.

Body Params
cdnType
string
enum
required
Defaults to NONE
Specifies the content delivery network (CDN) used to distribute content for the cloud distribution point.


NONE
Allowed:

NONE

JAMF_CLOUD

RACKSPACE_CLOUD_FILES

AMAZON_S3

AKAMAI
master
boolean
Defaults to false
Use as principal distribution point. Use as the authoritative source for all files.
Possible values are:
false
true


false
username
string
required
The username or access key used for authenticating with the selected content delivery network (CDN). This field is required when the cdnType is set to Rackspace Cloud Files(RACKSPACE_CLOUD_FILES), Amazon Web Services(AMAZON_S3), or Akamai(AKAMAI), as it is used to authenticate and authorize access to the respective cloud services.

For Rackspace Cloud Files(RACKSPACE_CLOUD_FILES), this is typically the username associated with your Rackspace cloud account.
For Amazon Web Services(AMAZON_S3), this corresponds to the Access Key ID used to interact with Amazon Web Services(AMAZON_S3) resources.
For Akamai(AKAMAI), this is the username used for API authentication to access Akamai's content delivery services. If the cdnType is None, this field is not applicable.
password
password
required
The password or authentication key used for connecting to the selected content delivery network (CDN). This field is required when the cdnType is set to Rackspace Cloud Files(RACKSPACE_CLOUD_FILES), Amazon Web Services(AMAZON_S3), or Akamai(AKAMAI), and is used to authenticate and authorize access to the respective cloud services.

For Rackspace Cloud Files(RACKSPACE_CLOUD_FILES), this refers to the API Key that is used in conjunction with the username for authenticating API requests.
For Amazon Web Services(AMAZON_S3), this corresponds to the Secret Access Key associated with your AWS account, used to securely sign requests to AWS services.
For Akamai(AKAMAI), this is the password used for API authentication to access Akamai's content delivery services. If the cdnType is None, this field is not applicable.
directory
string
The directory or path for content delivery in Akamai. This field is required when the cdnType is set to Akamai(AKAMAI) and specifies where content is stored within Akamai's system.

uploadUrl
string
The URL used to upload files to Akamai's NetStorage. This field is required when the cdnType is set to Akamai(AKAMAI). It specifies where content should be uploaded to Akamaiâ€™s cloud storage before being distributed via their CDN. The upload typically uses FTP or SFTP.

downloadUrl
string
The URL used to access and download content from Akamai's EdgeSuite. This field is required when the cdnType is set to Akamai(AKAMAI). It specifies the endpoint from which files are retrieved by devices or users.

secondaryAuthRequired
boolean
Defaults to false
Enable Remote Authentication.Authorize requests for files stored on the distribution point. This field is required when the cdnType is set to Akamai(AKAMAI).
Possible values are:
false
true


false
secondaryAuthStatusCode
integer
Defaults to 200
Secondary Auth Status Code. Configure the HTTP response code that will be returned by Jamf Pro during remote authentication. This field is required when the cdnType is set to Akamai(AKAMAI) and secondaryAuthRequired is true.

200
secondaryAuthTimeToLive
integer
≥ 1
Defaults to 3600
Secondary Auth Time To Live. Number of seconds before the authorization token expires. This field is required when the cdnType is set to Akamai(AKAMAI) and secondaryAuthRequired is true.

3600
requireSignedUrls
boolean
Defaults to false
Amazon Sign Url. It restrict access to requests that use a signed URL. This field is required when the cdnType is set to Amazon Web Services(AMAZON_S3).
Possible values are:
false
true


false
keyPairId
string
The CloudFront Access Key ID (keyPairId) is part of the credentials used to generate signed URLs for secure access to content in a CloudFront distribution. When using AWS, this key is paired with the CloudFront Secret Access Key to create the signed URL, ensuring that only authorized users can access specific content within a specified timeframe. This field is required when the cdnType is set to Amazon Web Services(AMAZON_S3) and requireSignedUrls is true.

expirationSeconds
integer
≥ 1
Defaults to 3600
Signed URL Expiration. Number of seconds before the signed URL expires, This field is required when the cdnType is set to Amazon Web Services(AMAZON_S3) and requireSignedUrls is true.

3600
privateKey
string
The CloudFront Private Key file is required when the cdnType is set to Amazon Web Services(AMAZON_S3) and requireSignedUrls parameter is enabled. This private key is used for signing URLs for restricted access to CloudFront-distributed content. The private key allows secure URL generation for signed URLs, ensuring that only authorized users can access certain content. The key must be uploaded in one of the following formats:

.pem: A Privacy-Enhanced Mail (PEM) file containing the private key in base64 encoded format.
.der: A Distinguished Encoding Rules (DER) encoded file, which is a binary format for the private key. The uploaded file should be kept secure, as it provides the ability to generate signed URLs with access to protected content.
Responses

201
The cloud distribution point was successfully created. After creation, a connection test was performed to ensure that the distribution point is properly configured and accessible. If the connection test passes, files will be uploaded to the distribution point. The response returns the details of the newly created cloud distribution point, including: - The unique identifier assigned to the cloud distribution point, along with the configured data. - The CDN type and URL where content will be distributed. - The connection test result (success or failure).

Response body
object
hasConnectionSucceeded
boolean
required
Defaults to false
Indicates whether the connection to the cloud distribution point was successful. If true, the connection was successful. If false, the connection failed.
Possible values are:
false
true

message
string
required
A message detailing the result of the connection test. This could be a success message or an error message if the connection failed.

inventoryId
string
Defaults to 0
The unique identifier (inventoryId) that links the cloud distribution point to its inventory data. By default, its value is 0, and it increments by +1 based on the existing inventory ID present in the table for each new cloud distribution point configuration. If the cdnType is set to NONE in the next configuration, the ID resets and starts from 1.

cdnType
string
enum
required
Defaults to NONE
Specifies the content delivery network (CDN) used to distribute content for the cloud distribution point.

NONE JAMF_CLOUD RACKSPACE_CLOUD_FILES AMAZON_S3 AKAMAI

master
boolean
Defaults to false
Use as principal distribution point. Use as the authoritative source for all files.
Possible values are:
false
true

username
string
required
The username or access key used for authenticating with the selected content delivery network (CDN). This field is required when the cdnType is set to Rackspace Cloud Files(RACKSPACE_CLOUD_FILES), Amazon Web Services(AMAZON_S3), or Akamai(AKAMAI), as it is used to authenticate and authorize access to the respective cloud services.

For Rackspace Cloud Files(RACKSPACE_CLOUD_FILES), this is typically the username associated with your Rackspace cloud account.
For Amazon Web Services(AMAZON_S3), this corresponds to the Access Key ID used to interact with Amazon Web Services(AMAZON_S3) resources.
For Akamai(AKAMAI), this is the username used for API authentication to access Akamai's content delivery services. If the cdnType is None, this field is not applicable.
directory
string
The directory or path for content delivery in Akamai. This field is required when the cdnType is set to Akamai(AKAMAI) and specifies where content is stored within Akamai's system.

cdnUrl
string
The CDN URL for the cloud distribution point. The URL format varies depending on the selected CDN provider:

Rackspace Cloud Files(RACKSPACE_CLOUD_FILES)
Amazon Web Services(AMAZON_S3)
Akamai(AKAMAI)
The cdnUrl should point to the content distribution location where software or other content is stored and made available for distribution.

uploadUrl
string
The URL used to upload files to Akamai's NetStorage. This field is required when the cdnType is set to Akamai(AKAMAI). It specifies where content should be uploaded to Akamaiâ€™s cloud storage before being distributed via their CDN. The upload typically uses FTP or SFTP.

downloadUrl
string
The URL used to access and download content from Akamai's EdgeSuite. This field is required when the cdnType is set to Akamai(AKAMAI). It specifies the endpoint from which files are retrieved by devices or users.

secondaryAuthRequired
boolean
Defaults to false
Enable Remote Authentication.Authorize requests for files stored on the distribution point. This field is required when the cdnType is set to Akamai(AKAMAI).
Possible values are:
false
true

secondaryAuthStatusCode
integer
Defaults to 200
Secondary Auth Status Code. Configure the HTTP response code that will be returned by Jamf Pro during remote authentication. This field is required when the cdnType is set to Akamai(AKAMAI) and secondaryAuthRequired is true.

secondaryAuthTimeToLive
integer
≥ 1
Defaults to 3600
Secondary Auth Time To Live. Number of seconds before the authorization token expires. This field is required when the cdnType is set to Akamai(AKAMAI) and secondaryAuthRequired is true.

requireSignedUrls
boolean
Defaults to false
Amazon Sign Url. It restrict access to requests that use a signed URL. This field is required when the cdnType is set to Amazon Web Services(AMAZON_S3).
Possible values are:
false
true

keyPairId
string
The CloudFront Access Key ID (keyPairId) is part of the credentials used to generate signed URLs for secure access to content in a CloudFront distribution. When using AWS, this key is paired with the CloudFront Secret Access Key to create the signed URL, ensuring that only authorized users can access specific content within a specified timeframe. This field is required when the cdnType is set to Amazon Web Services(AMAZON_S3) and requireSignedUrls is true.

expirationSeconds
integer
≥ 1
Defaults to 3600
Signed URL Expiration. Number of seconds before the signed URL expires, This field is required when the cdnType is set to Amazon Web Services(AMAZON_S3) and requireSignedUrls is true.

privateKey
string
The CloudFront Private Key file is required when the cdnType is set to Amazon Web Services(AMAZON_S3) and requireSignedUrls parameter is enabled. This private key is used for signing URLs for restricted access to CloudFront-distributed content. The private key allows secure URL generation for signed URLs, ensuring that only authorized users can access certain content. The key must be uploaded in one of the following formats:

.pem: A Privacy-Enhanced Mail (PEM) file containing the private key in base64 encoded format.
.der: A Distinguished Encoding Rules (DER) encoded file, which is a binary format for the private key. The uploaded file should be kept secure, as it provides the ability to generate signed URLs with access to protected content.

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/cloud-distribution-point \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "cdnType": "NONE",
  "master": false,
  "secondaryAuthRequired": false,
  "secondaryAuthStatusCode": 200,
  "secondaryAuthTimeToLive": 3600,
  "requireSignedUrls": false,
  "expirationSeconds": 3600
}
'

{
  "hasConnectionSucceeded": false,
  "message": "Cannot contact JCDS",
  "inventoryId": "0",
  "cdnType": "NONE",
  "master": false,
  "username": "Admin",
  "directory": "123456",
  "cdnUrl": "https://example-cdn-url.com/content-path",
  "uploadUrl": "ftp://mycompany.upload.akamai.com",
  "downloadUrl": "https://download.mycompany.com",
  "secondaryAuthRequired": false,
  "secondaryAuthStatusCode": 200,
  "secondaryAuthTimeToLive": 3600,
  "requireSignedUrls": false,
  "keyPairId": "K1R8C5EXAMPLE",
  "expirationSeconds": 3600,
  "privateKey": "string"
}
-----

Delete cloud distribution point.
delete
https://yourServer.jamfcloud.com/api/v1/cloud-distribution-point


The cloud distribution point and inventory details to be deleted.

Responses
204
cloud distribution point deleted successfully.

curl --request DELETE \
     --url https://yourserver.jamfcloud.com/api/v1/cloud-distribution-point

-----

Update specific fields on a cloud distribution point
patch
https://yourServer.jamfcloud.com/api/v1/cloud-distribution-point


Update specific fields on a cloud distribution point, then return the updated cloud distribution point details object.

Body Params
cdnType
string
enum
required
Defaults to NONE
Specifies the content delivery network (CDN) used to distribute content for the cloud distribution point.


NONE
Allowed:

NONE

JAMF_CLOUD

RACKSPACE_CLOUD_FILES

AMAZON_S3

AKAMAI
master
boolean
Defaults to false
Use as principal distribution point. Use as the authoritative source for all files.
Possible values are:
false
true


false
username
string
required
The username or access key used for authenticating with the selected content delivery network (CDN). This field is required when the cdnType is set to Rackspace Cloud Files(RACKSPACE_CLOUD_FILES), Amazon Web Services(AMAZON_S3), or Akamai(AKAMAI), as it is used to authenticate and authorize access to the respective cloud services.

For Rackspace Cloud Files(RACKSPACE_CLOUD_FILES), this is typically the username associated with your Rackspace cloud account.
For Amazon Web Services(AMAZON_S3), this corresponds to the Access Key ID used to interact with Amazon Web Services(AMAZON_S3) resources.
For Akamai(AKAMAI), this is the username used for API authentication to access Akamai's content delivery services. If the cdnType is None, this field is not applicable.
password
password
required
The password or authentication key used for connecting to the selected content delivery network (CDN). This field is required when the cdnType is set to Rackspace Cloud Files(RACKSPACE_CLOUD_FILES), Amazon Web Services(AMAZON_S3), or Akamai(AKAMAI), and is used to authenticate and authorize access to the respective cloud services.

For Rackspace Cloud Files(RACKSPACE_CLOUD_FILES), this refers to the API Key that is used in conjunction with the username for authenticating API requests.
For Amazon Web Services(AMAZON_S3), this corresponds to the Secret Access Key associated with your AWS account, used to securely sign requests to AWS services.
For Akamai(AKAMAI), this is the password used for API authentication to access Akamai's content delivery services. If the cdnType is None, this field is not applicable.
directory
string
The directory or path for content delivery in Akamai. This field is required when the cdnType is set to Akamai(AKAMAI) and specifies where content is stored within Akamai's system.

uploadUrl
string
The URL used to upload files to Akamai's NetStorage. This field is required when the cdnType is set to Akamai(AKAMAI). It specifies where content should be uploaded to Akamaiâ€™s cloud storage before being distributed via their CDN. The upload typically uses FTP or SFTP.

downloadUrl
string
The URL used to access and download content from Akamai's EdgeSuite. This field is required when the cdnType is set to Akamai(AKAMAI). It specifies the endpoint from which files are retrieved by devices or users.

secondaryAuthRequired
boolean
Defaults to false
Enable Remote Authentication.Authorize requests for files stored on the distribution point. This field is required when the cdnType is set to Akamai(AKAMAI).
Possible values are:
false
true


false
secondaryAuthStatusCode
integer
Defaults to 200
Secondary Auth Status Code. Configure the HTTP response code that will be returned by Jamf Pro during remote authentication. This field is required when the cdnType is set to Akamai(AKAMAI) and secondaryAuthRequired is true.

200
secondaryAuthTimeToLive
integer
≥ 1
Defaults to 3600
Secondary Auth Time To Live. Number of seconds before the authorization token expires. This field is required when the cdnType is set to Akamai(AKAMAI) and secondaryAuthRequired is true.

3600
requireSignedUrls
boolean
Defaults to false
Amazon Sign Url. It restrict access to requests that use a signed URL. This field is required when the cdnType is set to Amazon Web Services(AMAZON_S3).
Possible values are:
false
true


false
keyPairId
string
The CloudFront Access Key ID (keyPairId) is part of the credentials used to generate signed URLs for secure access to content in a CloudFront distribution. When using AWS, this key is paired with the CloudFront Secret Access Key to create the signed URL, ensuring that only authorized users can access specific content within a specified timeframe. This field is required when the cdnType is set to Amazon Web Services(AMAZON_S3) and requireSignedUrls is true.

expirationSeconds
integer
≥ 1
Defaults to 3600
Signed URL Expiration. Number of seconds before the signed URL expires, This field is required when the cdnType is set to Amazon Web Services(AMAZON_S3) and requireSignedUrls is true.

3600
privateKey
string
The CloudFront Private Key file is required when the cdnType is set to Amazon Web Services(AMAZON_S3) and requireSignedUrls parameter is enabled. This private key is used for signing URLs for restricted access to CloudFront-distributed content. The private key allows secure URL generation for signed URLs, ensuring that only authorized users can access certain content. The key must be uploaded in one of the following formats:

.pem: A Privacy-Enhanced Mail (PEM) file containing the private key in base64 encoded format.
.der: A Distinguished Encoding Rules (DER) encoded file, which is a binary format for the private key. The uploaded file should be kept secure, as it provides the ability to generate signed URLs with access to protected content.
Responses

200
The cloud distribution point was successfully updated. After the update, a connection test was performed to verify the configuration, and files were successfully uploaded to the distribution point. The response body will return the most up-to-date information about the cloud distribution point, including: - Updated configuration fields. - Connection test status (success or failure).

Response body
object
hasConnectionSucceeded
boolean
required
Defaults to false
Indicates whether the connection to the cloud distribution point was successful. If true, the connection was successful. If false, the connection failed.
Possible values are:
false
true

message
string
required
A message detailing the result of the connection test. This could be a success message or an error message if the connection failed.

inventoryId
string
Defaults to 0
The unique identifier (inventoryId) that links the cloud distribution point to its inventory data. By default, its value is 0, and it increments by +1 based on the existing inventory ID present in the table for each new cloud distribution point configuration. If the cdnType is set to NONE in the next configuration, the ID resets and starts from 1.

cdnType
string
enum
required
Defaults to NONE
Specifies the content delivery network (CDN) used to distribute content for the cloud distribution point.

NONE JAMF_CLOUD RACKSPACE_CLOUD_FILES AMAZON_S3 AKAMAI

master
boolean
Defaults to false
Use as principal distribution point. Use as the authoritative source for all files.
Possible values are:
false
true

username
string
required
The username or access key used for authenticating with the selected content delivery network (CDN). This field is required when the cdnType is set to Rackspace Cloud Files(RACKSPACE_CLOUD_FILES), Amazon Web Services(AMAZON_S3), or Akamai(AKAMAI), as it is used to authenticate and authorize access to the respective cloud services.

For Rackspace Cloud Files(RACKSPACE_CLOUD_FILES), this is typically the username associated with your Rackspace cloud account.
For Amazon Web Services(AMAZON_S3), this corresponds to the Access Key ID used to interact with Amazon Web Services(AMAZON_S3) resources.
For Akamai(AKAMAI), this is the username used for API authentication to access Akamai's content delivery services. If the cdnType is None, this field is not applicable.
directory
string
The directory or path for content delivery in Akamai. This field is required when the cdnType is set to Akamai(AKAMAI) and specifies where content is stored within Akamai's system.

cdnUrl
string
The CDN URL for the cloud distribution point. The URL format varies depending on the selected CDN provider:

Rackspace Cloud Files(RACKSPACE_CLOUD_FILES)
Amazon Web Services(AMAZON_S3)
Akamai(AKAMAI)
The cdnUrl should point to the content distribution location where software or other content is stored and made available for distribution.

uploadUrl
string
The URL used to upload files to Akamai's NetStorage. This field is required when the cdnType is set to Akamai(AKAMAI). It specifies where content should be uploaded to Akamaiâ€™s cloud storage before being distributed via their CDN. The upload typically uses FTP or SFTP.

downloadUrl
string
The URL used to access and download content from Akamai's EdgeSuite. This field is required when the cdnType is set to Akamai(AKAMAI). It specifies the endpoint from which files are retrieved by devices or users.

secondaryAuthRequired
boolean
Defaults to false
Enable Remote Authentication.Authorize requests for files stored on the distribution point. This field is required when the cdnType is set to Akamai(AKAMAI).
Possible values are:
false
true

secondaryAuthStatusCode
integer
Defaults to 200
Secondary Auth Status Code. Configure the HTTP response code that will be returned by Jamf Pro during remote authentication. This field is required when the cdnType is set to Akamai(AKAMAI) and secondaryAuthRequired is true.

secondaryAuthTimeToLive
integer
≥ 1
Defaults to 3600
Secondary Auth Time To Live. Number of seconds before the authorization token expires. This field is required when the cdnType is set to Akamai(AKAMAI) and secondaryAuthRequired is true.

requireSignedUrls
boolean
Defaults to false
Amazon Sign Url. It restrict access to requests that use a signed URL. This field is required when the cdnType is set to Amazon Web Services(AMAZON_S3).
Possible values are:
false
true

keyPairId
string
The CloudFront Access Key ID (keyPairId) is part of the credentials used to generate signed URLs for secure access to content in a CloudFront distribution. When using AWS, this key is paired with the CloudFront Secret Access Key to create the signed URL, ensuring that only authorized users can access specific content within a specified timeframe. This field is required when the cdnType is set to Amazon Web Services(AMAZON_S3) and requireSignedUrls is true.

expirationSeconds
integer
≥ 1
Defaults to 3600
Signed URL Expiration. Number of seconds before the signed URL expires, This field is required when the cdnType is set to Amazon Web Services(AMAZON_S3) and requireSignedUrls is true.

privateKey
string
The CloudFront Private Key file is required when the cdnType is set to Amazon Web Services(AMAZON_S3) and requireSignedUrls parameter is enabled. This private key is used for signing URLs for restricted access to CloudFront-distributed content. The private key allows secure URL generation for signed URLs, ensuring that only authorized users can access certain content. The key must be uploaded in one of the following formats:

.pem: A Privacy-Enhanced Mail (PEM) file containing the private key in base64 encoded format.
.der: A Distinguished Encoding Rules (DER) encoded file, which is a binary format for the private key. The uploaded file should be kept secure, as it provides the ability to generate signed URLs with access to protected content.

curl --request PATCH \
     --url https://yourserver.jamfcloud.com/api/v1/cloud-distribution-point \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "cdnType": "NONE",
  "master": false,
  "secondaryAuthRequired": false,
  "secondaryAuthStatusCode": 200,
  "secondaryAuthTimeToLive": 3600,
  "requireSignedUrls": false,
  "expirationSeconds": 3600
}
'

{
  "hasConnectionSucceeded": false,
  "message": "Cannot contact JCDS",
  "inventoryId": "0",
  "cdnType": "NONE",
  "master": false,
  "username": "Admin",
  "directory": "123456",
  "cdnUrl": "https://example-cdn-url.com/content-path",
  "uploadUrl": "ftp://mycompany.upload.akamai.com",
  "downloadUrl": "https://download.mycompany.com",
  "secondaryAuthRequired": false,
  "secondaryAuthStatusCode": 200,
  "secondaryAuthTimeToLive": 3600,
  "requireSignedUrls": false,
  "keyPairId": "K1R8C5EXAMPLE",
  "expirationSeconds": 3600,
  "privateKey": "string"
}

-----

Get cloud distribution point history details
get
https://yourServer.jamfcloud.com/api/v1/cloud-distribution-point/history


Get cloud distribution point history details

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
Sorts results by one or more criteria, following the format property:asc/desc. Default sort is ID:asc. If using multiple criteria, separate with commas.


string

id:asc

ADD string
filter
string
Filters results. Use RSQL format for query. Allows for many fields, including ID, name, etc. Can be combined with paging and sorting. Default filter is an empty query and returns all results from the requested page.

Responses

200
Details of cloud distribution point history were found.

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
     --url 'https://yourserver.jamfcloud.com/api/v1/cloud-distribution-point/history?page=0&page-size=100&sort=id%3Aasc' \
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

Get the cloud distribution point Inventory files details
get
https://yourServer.jamfcloud.com/api/v1/cloud-distribution-point/files


Retrieves the details of the inventory files associated with a cloud distribution point.This includes information about the files used for content distribution, such as their type, status, and categorization.The response provides a comprehensive list of inventory files, which may include packages, ebooks, or mobile device apps, allowing users to view the current state and metadata for each file in the distribution system.

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
Defaults to id.asc
Sorts results by one or more criteria, following the format property:asc/desc.
Default sort is id:asc.
If using multiple criteria, separate with commas. Allows sort for id, fileName, inventoryId and type etc.


string

id.asc

ADD string
filter
string
Filters results. Use RSQL format for query. Allows for many fields, including fileName and type
Can be combined with paging and sorting.
Fields allowed in the query: fileName, inventoryId and type
Default filter is an empty query and returns all results from the requested page.

Responses

200
cloud distribution point Inventory objects has been fetched successfully.

Response body
object
totalCount
integer
required
results
array of objects
length ≥ 0
object
id
string
A unique identifier for the cloud distribution point inventory file table.

inventoryId
string
A unique identifier for the cloud distribution point inventory file and cloud distribution point tables.This ID is used to reference a specific inventory resource within the system.

type
string
enum
required
The type of the inventory file. This field indicates whether the file is related to a package, mobile device app, or an ebook.

NONE PACKAGE EBOOK MOBILE_DEVICE_APP SCRIPT

fileName
string
required
The name of the inventory file. This could be the name of a package, a mobile device app, or an ebook file, depending on the file type. The name should match the actual file or package name as stored in the cloud distribution system.

fileObjectId
string
required
A unique identifier for each file type (package, ebook, or mobile device app). This ID is used to construct the URL for accessing or navigating to the specific resource related to the file type.

category
string
The category assigned to the inventory file (package, ebook, or mobile device app) during creation. This helps group and organize files based on their type or purpose, such as security software or productivity tools.

status
string
enum
required
The current status of the inventory file, indicating the progress or outcome of the file's upload process.It reflects whether the file is ready for use, still being processed, or has encountered an error.

READY PENDING ERROR

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v1/cloud-distribution-point/files?page=0&page-size=100&sort=id.asc' \
     --header 'accept: application/json'

{
  "totalCount": 1,
  "results": [
    {
      "id": "1",
      "inventoryId": "1",
      "type": "PACKAGE",
      "fileName": "SettingPackage",
      "fileObjectId": "1",
      "category": "Category1",
      "status": "READY"
    }
  ]
}

-----

Add specified cloud distribution point history object notes
post
https://yourServer.jamfcloud.com/api/v1/cloud-distribution-point/history

Add specified cloud distribution point history object notes

Body Params
History note to be created

note
string
required
Responses

201
cloud distribution point history note created successfully.

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
     --url https://yourserver.jamfcloud.com/api/v1/cloud-distribution-point/history \
     --header 'accept: application/json' \
     --header 'content-type: application/json'

{
  "id": 1,
  "username": "admin",
  "date": "2019-02-04T21:09:31.661Z",
  "note": "Sso settings update",
  "details": "Is SSO Enabled false\\nSelected SSO Provider"
}
-----

Get the cloud distribution point test connection details.
get
https://yourServer.jamfcloud.com/api/v1/cloud-distribution-point/test-connection

Verifies the connection to the cloud distribution point after updating its configuration. This endpoint returns the connection status and a message indicating whether the connection is successful or failed.

Response

200
This response indicates the test connection of the cloud distribution point (CDP). The response includes the following details:

hasConnectionSucceeded: A status indicating whether the connection test was successful or failed. This will be either true or false, based on the outcome of the connection test.
message: A detailed message providing additional context regarding the test result. For example, if the connection failed, the message might contain an error description.
Response body
object
hasConnectionSucceeded
boolean
required
Defaults to false
Indicates whether the connection to the cloud distribution point was successful. If true, the connection was successful. If false, the connection failed.
Possible values are:
false
true

message
string
required
A message detailing the result of the connection test. This could be a success message or an error message if the connection failed.

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/cloud-distribution-point/test-connection \
     --header 'accept: application/json'

{
  "hasConnectionSucceeded": false,
  "message": "Cannot contact JCDS"
}
-----

Finds specific information for the currently configured cloud distribution point.
get
https://yourServer.jamfcloud.com/api/v1/cloud-distribution-point/upload-capability

Finds a variety of values based on the currently configured cloud distribution point.

Response

200
The request was successful, and the upload capability details for the currently configured cloud distribution point (CDP) have been retrieved. This response provides information about whether the cloud distribution point supports direct uploads and the principal distribution technology being used.

Response body
object
principalDistributionTechnology
boolean
Defaults to false
directUploadCapable
boolean
Defaults to false

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/cloud-distribution-point/upload-capability \
     --header 'accept: application/json'

{
  "principalDistributionTechnology": false,
  "directUploadCapable": false
}

-----

Marks a specific file upload as failed for the currently configured cloud distribution point.
post
https://yourServer.jamfcloud.com/api/v1/cloud-distribution-point/fail-upload/{id}

Marks a specific file upload as failed for the currently configured cloud distribution point.

Path Params
id
string
required
The identifier of the inventory file to be marked as failed. The type and ID will make a unique identifier for the file.

Query Params
file-name
string
required
length ≤ 255
Name of the file to mark failure for.

filename.pkg
type
string
required
length ≤ 50
Type of file to mark failure for. Possible values are PACKAGE, EBOOK, MOBILE_DEVICE_APP.

PACKAGE
Responses
204
Successfully marked the specified file upload as failed for the currently configured cloud distribution point.

curl --request POST \
     --url 'https://yourserver.jamfcloud.com/api/v1/cloud-distribution-point/fail-upload/?file-name=filename.pkg&type=PACKAGE' \
     --header 'accept: application/json'

-----

Updates inventory data for the currently configured cloud distribution point.
post
https://yourServer.jamfcloud.com/api/v1/cloud-distribution-point/refresh-inventory

Updates inventory data for the currently configured cloud distribution point.

Query Params
file-name
string
length ≤ 255
Name of the file to check the availability of. If available, the inventory and status will be updated in Jamf Pro. If no file is specified, it will force an immediate inventory refresh at a rate-limit of once every 15 seconds.

Response
204
Successfully updated inventory data for the currently configured cloud distribution point.

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/cloud-distribution-point/refresh-inventory