Retrieve Packages
get
https://yourServer.jamfcloud.com/api/v1/packages

Retrieves packages

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
Filters results. Use RSQL format for query. Allows for many fields, including ID, name, etc. Can be combined with paging and sorting. Fields allowed in the query: id, fileName, packageName, categoryId, info, notes, manifestFileName, cloudTransferStatus. Default filter is an empty query and returns all results from the requested page.

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
packageName
string
required
fileName
string
required
categoryId
string
required
info
string | null
notes
string | null
priority
integer
required
osRequirements
string | null
fillUserTemplate
boolean
required
indexed
boolean
fillExistingUsers
boolean
swu
boolean
rebootRequired
boolean
required
selfHealNotify
boolean
selfHealingAction
string | null
osInstall
boolean
required
serialNumber
string | null
parentPackageId
string | null
basePath
string | null
suppressUpdates
boolean
required
cloudTransferStatus
string
ignoreConflicts
boolean
suppressFromDock
boolean
required
suppressEula
boolean
required
suppressRegistration
boolean
required
installLanguage
string | null
md5
string | null
sha256
string | null
sha3512
string | null
hashType
string | null
hashValue
string | null
size
string | null
osInstallerVersion
string | null
manifest
string | null
manifestFileName
string | null
format
string | null

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v1/packages?page=0&page-size=100&sort=id%3Aasc' \
     --header 'accept: application/json'

{
  "totalCount": 1,
  "results": [
    {
      "id": "1",
      "packageName": "Google Chrome",
      "fileName": "my-package.pkg",
      "categoryId": "-1",
      "info": "A package that is important to my organization.",
      "notes": "Some notes.",
      "priority": 3,
      "osRequirements": "10.6.8, 10.7.x",
      "fillUserTemplate": false,
      "indexed": false,
      "fillExistingUsers": false,
      "swu": false,
      "rebootRequired": false,
      "selfHealNotify": false,
      "selfHealingAction": "nothing",
      "osInstall": false,
      "serialNumber": "1234",
      "parentPackageId": "3",
      "basePath": "my/path",
      "suppressUpdates": false,
      "cloudTransferStatus": "READY",
      "ignoreConflicts": false,
      "suppressFromDock": false,
      "suppressEula": false,
      "suppressRegistration": false,
      "installLanguage": "en_US",
      "md5": "0cc175b9c0f1b6a831c399e269772661",
      "sha256": "61be55a8e2f6b4e172338bddf184d6dbee29c98853e0a0485ecee7f27b9af0b4",
      "sha3512": "a51a01e63f76c4601a2989575280e2a1c1e83191382277906b0f92a251c53bbbd1d9804dd30edd304cffe75d1c6455440b0c39081ca94c68f149d1f719e15092",
      "hashType": "MD5",
      "hashValue": "0cc175b9c0f1b6a831c399e269772661",
      "size": "234KB",
      "osInstallerVersion": "10.3.x",
      "manifest": "manifest",
      "manifestFileName": "manifest.plist",
      "format": "format"
    }
  ]
}
-----
Create package
post
https://yourServer.jamfcloud.com/api/v1/packages

Create package

Body Params
Package to be created

packageName
string
required
Google Chrome
fileName
string
required
my-package.pkg
categoryId
string
required
-1
info
string | null
thing
notes
string | null
Some notes.
priority
integer
required
3
osRequirements
string | null
10.6.8, 10.7.x
fillUserTemplate
boolean
required

true
fillExistingUsers
boolean

true
swu
boolean

true
rebootRequired
boolean
required

true
selfHealNotify
boolean

true
selfHealingAction
string | null
nothing
osInstall
boolean
required

true
serialNumber
string | null
1234
parentPackageId
string | null
3
basePath
string | null
my/path
suppressUpdates
boolean
required

true
ignoreConflicts
boolean

false
suppressFromDock
boolean
required

true
suppressEula
boolean
required

true
suppressRegistration
boolean
required

true
installLanguage
string | null
en_US
md5
string | null
0cc175b9c0f1b6a831c399e269772661
sha256
string | null
61be55a8e2f6b4e172338bddf184d6dbee29c98853e0a0485ecee7f27b9af0b4
sha3512
string | null
a51a01e63f76c4601a2989575280e2a1c1e83191382277906b0f92a251c53bbbd1d9804dd30edd304cffe75d1c6455440b0c39081ca94c68f149d1f719e15092
hashType
string | null
MD5
hashValue
string | null
0cc175b9c0f1b6a831c399e269772661
osInstallerVersion
string | null
10.3.x
manifest
string | null
manifest
manifestFileName
string | null
manifest.plist
format
string | null
format
Response

201
Package was created

Response body
object
id
string
length ≥ 1
href
string

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/packages \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "fillUserTemplate": true,
  "rebootRequired": true,
  "osInstall": true,
  "suppressUpdates": true,
  "suppressFromDock": true,
  "suppressEula": true,
  "suppressRegistration": true,
  "packageName": "Google Chrome",
  "fileName": "my-package.pkg",
  "categoryId": "-1",
  "info": "thing",
  "notes": "Some notes.",
  "priority": 3,
  "osRequirements": "10.6.8, 10.7.x",
  "fillExistingUsers": true,
  "swu": true,
  "selfHealNotify": true,
  "selfHealingAction": "nothing",
  "serialNumber": "1234",
  "parentPackageId": "3",
  "basePath": "my/path",
  "ignoreConflicts": false,
  "installLanguage": "en_US",
  "md5": "0cc175b9c0f1b6a831c399e269772661",
  "sha256": "61be55a8e2f6b4e172338bddf184d6dbee29c98853e0a0485ecee7f27b9af0b4",
  "sha3512": "a51a01e63f76c4601a2989575280e2a1c1e83191382277906b0f92a251c53bbbd1d9804dd30edd304cffe75d1c6455440b0c39081ca94c68f149d1f719e15092",
  "hashType": "MD5",
  "hashValue": "0cc175b9c0f1b6a831c399e269772661",
  "osInstallerVersion": "10.3.x",
  "manifest": "manifest",
  "manifestFileName": "manifest.plist",
  "format": "format"
}
'

{
  "id": "1",
  "href": "https://yourJamfProUrl.jamf/api/v1/resource/1"
}
-----
Delete multiple packages at once
post
https://yourServer.jamfcloud.com/api/v1/packages/delete-multiple

IDs of the packages to be deleted

Body Params
ids of the package to be deleted

ids
array of strings

string


ADD string
Responses
204
All specified packages deleted successfully

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/packages/delete-multiple \
     --header 'accept: application/json' \
     --header 'content-type: application/json'
-----
Export Packages collection
post
https://yourServer.jamfcloud.com/api/v1/packages/export

Export Packages collection

Query Params
export-fields
array of strings
Defaults to
Export fields parameter, used to change default order or ignore some of the response properties. Default is empty array, which means that all fields of the response entity will be serialized. Example: export-fields=id,username


string


ADD string
export-labels
array of strings
Defaults to
Export labels parameter, used to customize fieldnames/columns in the exported file. Default is empty array, which means that response properties names will be used. Number of the provided labels must match the number of export-fields Example: export-labels=identifier,name with matching: export-fields=id,username


ADD string
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

Body Params
Optional. Can be used to override query parameters so that the URI does not exceed the 2,000 character limit.

page
integer | null
Defaults to 0
0
pageSize
integer | null
Defaults to 100
100
sort
array of strings | null
Defaults to id:desc
Sorting criteria in the format: [[:asc/desc]. Default direction when not stated is ascending.


string

id:desc

ADD string
filter
string | null
fields
array of objects | null
Used to change default order or ignore some of the fields. When null or empty array, all fields will be exported.


ADD object
Headers
accept
string
enum
Defaults to application/json
Generated from available response content types


text/csv
Allowed:

application/json

text/csv
Responses

200
Export successful

Response body
json

curl --request POST \
     --url 'https://yourserver.jamfcloud.com/api/v1/packages/export?page=0&page-size=100&sort=id%3Aasc' \
     --header 'accept: text/csv' \
     --header 'content-type: application/json' \
     --data '
{
  "page": 0,
  "pageSize": 100,
  "sort": [
    "id:desc"
  ]
}
'
-----
Get specified Package object
get
https://yourServer.jamfcloud.com/api/v1/packages/{id}

Gets specified Package object

Path Params
id
string
required
instance id of package

Responses

200
Package details found

Response body
object
id
string
packageName
string
required
fileName
string
required
categoryId
string
required
info
string | null
notes
string | null
priority
integer
required
osRequirements
string | null
fillUserTemplate
boolean
required
indexed
boolean
fillExistingUsers
boolean
swu
boolean
rebootRequired
boolean
required
selfHealNotify
boolean
selfHealingAction
string | null
osInstall
boolean
required
serialNumber
string | null
parentPackageId
string | null
basePath
string | null
suppressUpdates
boolean
required
cloudTransferStatus
string
ignoreConflicts
boolean
suppressFromDock
boolean
required
suppressEula
boolean
required
suppressRegistration
boolean
required
installLanguage
string | null
md5
string | null
sha256
string | null
sha3512
string | null
hashType
string | null
hashValue
string | null
size
string | null
osInstallerVersion
string | null
manifest
string | null
manifestFileName
string | null
format
string | null

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/packages/ \
     --header 'accept: application/json'

{
  "id": "1",
  "packageName": "Google Chrome",
  "fileName": "my-package.pkg",
  "categoryId": "-1",
  "info": "A package that is important to my organization.",
  "notes": "Some notes.",
  "priority": 3,
  "osRequirements": "10.6.8, 10.7.x",
  "fillUserTemplate": false,
  "indexed": false,
  "fillExistingUsers": false,
  "swu": false,
  "rebootRequired": false,
  "selfHealNotify": false,
  "selfHealingAction": "nothing",
  "osInstall": false,
  "serialNumber": "1234",
  "parentPackageId": "3",
  "basePath": "my/path",
  "suppressUpdates": false,
  "cloudTransferStatus": "READY",
  "ignoreConflicts": false,
  "suppressFromDock": false,
  "suppressEula": false,
  "suppressRegistration": false,
  "installLanguage": "en_US",
  "md5": "0cc175b9c0f1b6a831c399e269772661",
  "sha256": "61be55a8e2f6b4e172338bddf184d6dbee29c98853e0a0485ecee7f27b9af0b4",
  "sha3512": "a51a01e63f76c4601a2989575280e2a1c1e83191382277906b0f92a251c53bbbd1d9804dd30edd304cffe75d1c6455440b0c39081ca94c68f149d1f719e15092",
  "hashType": "MD5",
  "hashValue": "0cc175b9c0f1b6a831c399e269772661",
  "size": "234KB",
  "osInstallerVersion": "10.3.x",
  "manifest": "manifest",
  "manifestFileName": "manifest.plist",
  "format": "format"
}
-----
Update specified package object
put
https://yourServer.jamfcloud.com/api/v1/packages/{id}

Update specified package object

Path Params
id
string
required
Instance ID of package

Body Params
Package object to update. IDs defined in this body will be ignored

packageName
string
required
fileName
string
required
categoryId
string
required
info
string | null
notes
string | null
priority
integer
required
osRequirements
string | null
fillUserTemplate
boolean
required

true
fillExistingUsers
boolean

false
swu
boolean

false
rebootRequired
boolean
required

true
selfHealNotify
boolean

false
selfHealingAction
string | null
osInstall
boolean
required

true
serialNumber
string | null
parentPackageId
string | null
basePath
string | null
suppressUpdates
boolean
required

true
ignoreConflicts
boolean

false
suppressFromDock
boolean
required

true
suppressEula
boolean
required

true
suppressRegistration
boolean
required

true
installLanguage
string | null
md5
string | null
sha256
string | null
sha3512
string | null
hashType
string | null
hashValue
string | null
osInstallerVersion
string | null
manifest
string | null
manifestFileName
string | null
format
string | null
Response

200
Package update

Response body
object
id
string
packageName
string
required
fileName
string
required
categoryId
string
required
info
string | null
notes
string | null
priority
integer
required
osRequirements
string | null
fillUserTemplate
boolean
required
indexed
boolean
fillExistingUsers
boolean
swu
boolean
rebootRequired
boolean
required
selfHealNotify
boolean
selfHealingAction
string | null
osInstall
boolean
required
serialNumber
string | null
parentPackageId
string | null
basePath
string | null
suppressUpdates
boolean
required
cloudTransferStatus
string
ignoreConflicts
boolean
suppressFromDock
boolean
required
suppressEula
boolean
required
suppressRegistration
boolean
required
installLanguage
string | null
md5
string | null
sha256
string | null
sha3512
string | null
hashType
string | null
hashValue
string | null
size
string | null
osInstallerVersion
string | null
manifest
string | null
manifestFileName
string | null
format
string | null

curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v1/packages/ \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "fillUserTemplate": true,
  "rebootRequired": true,
  "osInstall": true,
  "suppressUpdates": true,
  "suppressFromDock": true,
  "suppressEula": true,
  "suppressRegistration": true
}
'

{
  "id": "1",
  "packageName": "Google Chrome",
  "fileName": "my-package.pkg",
  "categoryId": "-1",
  "info": "A package that is important to my organization.",
  "notes": "Some notes.",
  "priority": 3,
  "osRequirements": "10.6.8, 10.7.x",
  "fillUserTemplate": false,
  "indexed": false,
  "fillExistingUsers": false,
  "swu": false,
  "rebootRequired": false,
  "selfHealNotify": false,
  "selfHealingAction": "nothing",
  "osInstall": false,
  "serialNumber": "1234",
  "parentPackageId": "3",
  "basePath": "my/path",
  "suppressUpdates": false,
  "cloudTransferStatus": "READY",
  "ignoreConflicts": false,
  "suppressFromDock": false,
  "suppressEula": false,
  "suppressRegistration": false,
  "installLanguage": "en_US",
  "md5": "0cc175b9c0f1b6a831c399e269772661",
  "sha256": "61be55a8e2f6b4e172338bddf184d6dbee29c98853e0a0485ecee7f27b9af0b4",
  "sha3512": "a51a01e63f76c4601a2989575280e2a1c1e83191382277906b0f92a251c53bbbd1d9804dd30edd304cffe75d1c6455440b0c39081ca94c68f149d1f719e15092",
  "hashType": "MD5",
  "hashValue": "0cc175b9c0f1b6a831c399e269772661",
  "size": "234KB",
  "osInstallerVersion": "10.3.x",
  "manifest": "manifest",
  "manifestFileName": "manifest.plist",
  "format": "format"
}
-----
Remove specified package
delete
https://yourServer.jamfcloud.com/api/v1/packages/{id}

Removes specified package

Path Params
id
string
required
Instance ID of package

1
Response
204
Package deleted successfully

curl --request DELETE \
     --url https://yourserver.jamfcloud.com/api/v1/packages/1

-----
Get specified Package History object
get
https://yourServer.jamfcloud.com/api/v1/packages/{id}/history

Gets specified Package history object

Path Params
id
string
required
Instance ID of package history

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
Sorts results by one or more criteria, following the format property:asc/desc. Default sort is ID:asc. If using multiple criteria, separate with commas.


string

date:desc

ADD string
filter
string
Filters results. Use RSQL format for query. Allows for many fields, including ID, name, etc. Can be combined with paging and sorting. Default filter is an empty query and returns all results from the requested page.

Responses

200
Details of package history were found

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
     --url 'https://yourserver.jamfcloud.com/api/v1/packages//history?page=0&page-size=100&sort=date%3Adesc' \
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
Add specified Package history object notes
post
https://yourServer.jamfcloud.com/api/v1/packages/{id}/history

Adds specified Package history object notes

Path Params
id
string
required
Instance ID of package history

1
Body Params
History note to be created

note
string
required
A generic note can sometimes be useful, but generally not.
Responses

201
Package history note created successfully

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
     --url https://yourserver.jamfcloud.com/api/v1/packages/1/history \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "note": "A generic note can sometimes be useful, but generally not."
}
'

{
  "id": 1,
  "username": "admin",
  "date": "2019-02-04T21:09:31.661Z",
  "note": "Sso settings update",
  "details": "Is SSO Enabled false\\nSelected SSO Provider"
}
-----
Export history object collection in specified format for specified Packages
post
https://yourServer.jamfcloud.com/api/v1/packages/{id}/history/export


Export history object collection in specified format for specified Packages

Path Params
id
string
required
Instance ID of package history note

1
Query Params
export-fields
array of strings
Defaults to
Export fields parameter, used to change default order or ignore some of the response properties. Default is empty array, which means that all fields of the response entity will be serialized. Example: export-fields=id,username


ADD string
export-labels
array of strings
Defaults to
Export labels parameter, used to customize fieldnames/columns in the exported file. Default is empty array, which means that response properties names will be used. Number of the provided labels must match the number of export-fields Example: export-labels=identifier,name with matching: export-fields=id,username


ADD string
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
Sorts results by one or more criteria, following the format property:asc/desc. Default sort is ID:asc. If using multiple criteria, separate with commas.


string

date:desc

ADD string
filter
string
Filters results. Use RSQL format for query. Allows for many fields, including ID, name, etc. Can be combined with paging and sorting. Default filter is an empty query and returns all results from the requested page.

Body Params
Optional. Can be used to override query parameters so that the URI does not exceed the 2,000 character limit.

page
integer | null
Defaults to 0
0
pageSize
integer | null
Defaults to 100
100
sort
array of strings | null
Defaults to id:desc
Sorting criteria in the format: [[:asc/desc]. Default direction when not stated is ascending.


string

id:desc

ADD string
filter
string | null
fields
array of objects | null
Used to change default order or ignore some of the fields. When null or empty array, all fields will be exported.


ADD object
Headers
accept
string
enum
Defaults to application/json
Generated from available response content types


text/csv
Allowed:

application/json

text/csv
Responses

200
Export successful

Response body
json

curl --request POST \
     --url 'https://yourserver.jamfcloud.com/api/v1/packages/1/history/export?page=0&page-size=100&sort=date%3Adesc' \
     --header 'accept: text/csv' \
     --header 'content-type: application/json' \
     --data '
{
  "page": 0,
  "pageSize": 100,
  "sort": [
    "id:desc"
  ]
}
'
-----
Add a manifest to a package
post
https://yourServer.jamfcloud.com/api/v1/packages/{id}/manifest


Add a manifest to a package

Path Params
id
string
required
Id of the package the manifest should be assigned to

Body Params
file
file
required
The manifest file to upload

No file chosen
Responses

201
Success

Response body
object
id
string
packageName
string
required
fileName
string
required
categoryId
string
required
info
string | null
notes
string | null
priority
integer
required
osRequirements
string | null
fillUserTemplate
boolean
required
indexed
boolean
fillExistingUsers
boolean
swu
boolean
rebootRequired
boolean
required
selfHealNotify
boolean
selfHealingAction
string | null
osInstall
boolean
required
serialNumber
string | null
parentPackageId
string | null
basePath
string | null
suppressUpdates
boolean
required
cloudTransferStatus
string
ignoreConflicts
boolean
suppressFromDock
boolean
required
suppressEula
boolean
required
suppressRegistration
boolean
required
installLanguage
string | null
md5
string | null
sha256
string | null
sha3512
string | null
hashType
string | null
hashValue
string | null
size
string | null
osInstallerVersion
string | null
manifest
string | null
manifestFileName
string | null
format
string | null

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/packages//manifest \
     --header 'accept: application/json' \
     --header 'content-type: multipart/form-data'

{
  "id": "1",
  "packageName": "Google Chrome",
  "fileName": "my-package.pkg",
  "categoryId": "-1",
  "info": "A package that is important to my organization.",
  "notes": "Some notes.",
  "priority": 3,
  "osRequirements": "10.6.8, 10.7.x",
  "fillUserTemplate": false,
  "indexed": false,
  "fillExistingUsers": false,
  "swu": false,
  "rebootRequired": false,
  "selfHealNotify": false,
  "selfHealingAction": "nothing",
  "osInstall": false,
  "serialNumber": "1234",
  "parentPackageId": "3",
  "basePath": "my/path",
  "suppressUpdates": false,
  "cloudTransferStatus": "READY",
  "ignoreConflicts": false,
  "suppressFromDock": false,
  "suppressEula": false,
  "suppressRegistration": false,
  "installLanguage": "en_US",
  "md5": "0cc175b9c0f1b6a831c399e269772661",
  "sha256": "61be55a8e2f6b4e172338bddf184d6dbee29c98853e0a0485ecee7f27b9af0b4",
  "sha3512": "a51a01e63f76c4601a2989575280e2a1c1e83191382277906b0f92a251c53bbbd1d9804dd30edd304cffe75d1c6455440b0c39081ca94c68f149d1f719e15092",
  "hashType": "MD5",
  "hashValue": "0cc175b9c0f1b6a831c399e269772661",
  "size": "234KB",
  "osInstallerVersion": "10.3.x",
  "manifest": "manifest",
  "manifestFileName": "manifest.plist",
  "format": "format"
}
-----
Delete the manifest for a specified package
delete
https://yourServer.jamfcloud.com/api/v1/packages/{id}/manifest


Delete the manifest file for a specified package

Path Params
id
string
required
Id of the package to delete manifest from

Responses
204
Package manifest deleted successfully

curl --request DELETE \
     --url https://yourserver.jamfcloud.com/api/v1/packages//manifest \
     --header 'accept: application/json'
-----
Upload package
post
https://yourServer.jamfcloud.com/api/v1/packages/{id}/upload


Uploads a package

Path Params
id
string
required
instance id of package

Body Params
file
file
required
The file to upload

No file chosen
Responses

201
Successfully uploaded package

Response body
object
id
string
length ≥ 1
href
string

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/packages//upload \
     --header 'accept: application/json' \
     --header 'content-type: multipart/form-data'

{
  "id": "1",
  "href": "https://yourJamfProUrl.jamf/api/v1/resource/1"
}
-----