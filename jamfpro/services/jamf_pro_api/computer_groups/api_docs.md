Returns the list of all computer groups
get
https://yourServer.jamfcloud.com/api/v1/computer-groups

Use it to get the list of all computer groups.

Response

200
Success

Response body
array of objects
object
id
string
length ≥ 1
name
string
description
string
smartGroup
boolean

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/computer-groups \
     --header 'accept: application/json'

[
  {
    "id": "1",
    "name": "All Managed Computers",
    "description": "A group containing all managed computers",
    "smartGroup": true
  }
]
-----
Get the membership of a Smart Computer Group
get
https://yourServer.jamfcloud.com/api/v2/computer-groups/smart-group-membership/{id}


Gets the membership of a Smart Computer Group

Path Params
id
string
required
id of the Smart Computer Group

Responses

200
Successful response - Smart Computer Group membership retrieved

Response body
object
members
array of integers

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v2/computer-groups/smart-group-membership/ \
     --header 'accept: application/json'

{
  "members": [
    1,
    2,
    3
  ]
}
-----
Search for Smart Computer Groups
get
https://yourServer.jamfcloud.com/api/v2/computer-groups/smart-groups


Search for Smart Computer Groups

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
Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=name:asc


string

id:asc

ADD string
filter
string
Query in the RSQL format, allowing to filter smart computer group collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: id, name, siteId. The siteId field can only be filtered by admins with full access. Any sited admin will have siteId filtered automatically. Example: name=="group"

Response

200
Successful response

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
string
length ≥ 1
siteId
string
length ≥ 1
name
string
length ≥ 1
description
string
length ≥ 0
membershipCount
integer

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v2/computer-groups/smart-groups?page=0&page-size=100&sort=id%3Aasc' \
     --header 'accept: application/json'

{
  "totalCount": 3,
  "results": [
    {
      "id": "1",
      "siteId": "1",
      "name": "My computer group",
      "description": "My computer group description",
      "membershipCount": 231
    }
  ]
}

Create a Smart Computer Group
post
https://yourServer.jamfcloud.com/api/v2/computer-groups/smart-groups


Creates a Smart Computer Group

Query Params
platform
boolean
Defaults to false
Optional. Return platform identifiers instead of internal identifiers when set to true.


false
Body Params
name
string
required
description
string
criteria
array of objects

ADD object
siteId
string | null
Response

201
Successful response - Smart Computer Group created

Response body
object
id
string
length ≥ 1
href
string

curl --request POST \
     --url 'https://yourserver.jamfcloud.com/api/v2/computer-groups/smart-groups?platform=false' \
     --header 'accept: application/json' \
     --header 'content-type: application/json'

{
  "id": "1",
  "href": "https://yourJamfProUrl.jamf/api/v1/resource/1"
}
-----
Update a Smart Computer Group
put
https://yourServer.jamfcloud.com/api/v2/computer-groups/smart-groups/{id}


Updates a Smart Computer Group

Path Params
id
string
required
id of target Smart Computer Group

Body Params
name
string
required
description
string
criteria
array of objects

ADD object
siteId
string | null
Responses

202
Successful response - Smart Computer Group updated

Response body
object
name
string
required
description
string
criteria
array of objects
object
name
string
required
priority
integer
andOr
string
required
searchType
string
required
value
string
required
openingParen
boolean
closingParen
boolean
siteId
string | null

curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v2/computer-groups/smart-groups/ \
     --header 'accept: application/json' \
     --header 'content-type: application/json'

{
  "name": "New Group Name",
  "description": "New Group Description",
  "criteria": [
    {
      "name": "Account",
      "priority": 0,
      "andOr": "and",
      "searchType": "is",
      "value": "test",
      "openingParen": false,
      "closingParen": false
    }
  ],
  "siteId": "-1"
}
-----
Remove specified Smart Computer Group
delete
https://yourServer.jamfcloud.com/api/v2/computer-groups/smart-groups/{id}


Remove specified Smart Computer Group

Path Params
id
string
required
id of target Smart Computer Group

1
Responses
204
Successful response - Smart Computer Group removed

curl --request DELETE \
     --url https://yourserver.jamfcloud.com/api/v2/computer-groups/smart-groups/1 \
     --header 'accept: application/json'
-----
Get Smart Computer Group by Id
get
https://yourServer.jamfcloud.com/api/v2/computer-groups/smart-groups/{id}


Get Smart Computer Group by Id

Path Params
id
string
required
instance id of smart computer group

Responses

200
Successful response

Response body
object
name
string
required
description
string
criteria
array of objects
object
name
string
required
priority
integer
andOr
string
required
searchType
string
required
value
string
required
openingParen
boolean
closingParen
boolean
siteId
string | null

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v2/computer-groups/smart-groups/ \
     --header 'accept: application/json'

{
  "name": "New Group Name",
  "description": "New Group Description",
  "criteria": [
    {
      "name": "Account",
      "priority": 0,
      "andOr": "and",
      "searchType": "is",
      "value": "test",
      "openingParen": false,
      "closingParen": false
    }
  ],
  "siteId": "-1"
}
_____
Search for Static Computer Groups
get
https://yourServer.jamfcloud.com/api/v2/computer-groups/static-groups


Search for Static Computer Groups

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
Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=name:asc


string

id:asc

ADD string
filter
string
Query in the RSQL format, allowing to filter static computer group collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: id, name, siteId. The siteId field can only be filtered by admins with full access. Any sited admin will have siteId filtered automatically. Example: name=="group"

Response

200
Successful response

Response body
object
totalCount
integer
required
≥ 0
results
array of objects
required
object
id
string
required
length ≥ 1
name
string
required
length ≥ 1
description
string | null
siteId
string | null
count
integer
≥ 0

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v2/computer-groups/static-groups?page=0&page-size=100&sort=id%3Aasc' \
     --header 'accept: application/json'

{
  "totalCount": 2,
  "results": [
    {
      "id": "1",
      "name": "Test Static Computer Group",
      "description": "A test static computer group",
      "siteId": "1",
      "count": 5
    }
  ]
}
-----
Create membership of a static computer group.
post
https://yourServer.jamfcloud.com/api/v2/computer-groups/static-groups


Create membership of a static computer group.

Query Params
platform
boolean
Defaults to false
Optional. Return platform identifiers instead of internal identifiers when set to true.


false
Body Params
name
string
required
length ≥ 1
description
string | null
siteId
string | null
assignments
array of strings
Array of computer IDs to assign to the static group


ADD string
Responses

201
Static computer group created successfully

Response body
object
id
string
length ≥ 1
href
string

curl --request POST \
     --url 'https://yourserver.jamfcloud.com/api/v2/computer-groups/static-groups?platform=false' \
     --header 'accept: application/json' \
     --header 'content-type: application/json'

{
  "id": "1",
  "href": "https://yourJamfProUrl.jamf/api/v1/resource/1"
}
-----
et Static Computer Group by Id
get
https://yourServer.jamfcloud.com/api/v2/computer-groups/static-groups/{id}


Get Static Computer Group by Id

Path Params
id
string
required
instance id of static computer group

Responses

200
Successful response

Response body
object
id
string
required
length ≥ 1
name
string
required
length ≥ 1
description
string | null
siteId
string | null

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v2/computer-groups/static-groups/ \
     --header 'accept: application/json'

{
  "id": "1",
  "name": "Test Static Computer Group",
  "description": "A test static computer group",
  "siteId": "1"
}
-----
Update membership of a static computer group.
put
https://yourServer.jamfcloud.com/api/v2/computer-groups/static-groups/{id}


Update membership of a static computer group.

Path Params
id
string
required
instance id of a static computer group

Body Params
name
string
required
length ≥ 1
description
string | null
siteId
string | null
assignments
array of strings
Array of computer IDs to assign to the static group


ADD string
Responses

202
Successful response

Response body
object
id
string
name
string
required
length ≥ 1
description
string | null
siteId
string | null
assignments
array of strings
Array of computer IDs to assign to the static group

curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v2/computer-groups/static-groups/ \
     --header 'accept: application/json' \
     --header 'content-type: application/json'

{
  "id": "1",
  "name": "Test Static Computer Group",
  "description": "A test static computer group",
  "siteId": "1",
  "assignments": [
    "1"
  ]
}
-----
Remove Static Computer Group by Id
delete
https://yourServer.jamfcloud.com/api/v2/computer-groups/static-groups/{id}


Remove Static Computer Group by Id

Path Params
id
string
required
instance id of static computer group

1
Responses
204
Static Computer Group successfully removed

curl --request DELETE \
     --url https://yourserver.jamfcloud.com/api/v2/computer-groups/static-groups/1 \
     --header 'accept: application/json'

-----
Create Computer Inventory record
post
https://yourServer.jamfcloud.com/api/v3/computers-inventory

Creates Computer Inventory record

Body Params
udid
string | null
general
object

general object
name
string
required
length ≥ 1
Boalime
lastIpAddress
string
247.185.82.186
lastReportedIp
string
247.185.82.186
jamfBinaryVersion
string
9.27
platform
string
enum

MAC
Allowed:

WINDOWS

MAC

NONE
barcode1
string
5 12345 678900
barcode2
string
5 12345 678900
assetTag
string
304822
remoteManagement
object

remoteManagement object
managed
boolean | null

true
supervised
boolean

true
mdmCapable
boolean

true
reportDate
date-time
2018-10-31T18:04:13Z
lastContactTime
date-time
2018-10-31T18:04:13Z
lastCloudBackupDate
date-time
2018-10-31T18:04:13Z
lastEnrolledDate
date-time
2018-10-31T18:04:13Z
distributionPointId
string
1
siteId
string
1
itunesStoreAccountActive
boolean

true
enrolledViaAutomatedDeviceEnrollment
boolean

false
userApprovedMdm
boolean

true
declarativeDeviceManagementEnabled
boolean

true
applications
array of objects | null

object

name
string | null
Microsoft Word
path
string | null
/usr/local/app
version
string | null
1.0.0

object

name
string | null
Microsoft Word
path
string | null
/usr/local/app
version
string | null
1.0.0

ADD object
storage
object

storage object
disks
array of objects | null

object

device
string | null
disk0
model
string | null
APPLE HDD TOSHIBA MK5065GSXF
revision
string | null
5
serialNumber
string | null
a8598f013366
sizeMegabytes
int64 | null
Disk Size in MB.

262144
smartStatus
string | null
S.M.A.R.T Status

OK
type
string | null
Connection type attribute.

false
partitions
array of objects | null

object

name
string | null
Foo
sizeMegabytes
int64 | null
Partition Size in MB.

262144
availableMegabytes
int64 | null
Available space in MB.

131072
partitionType
string | null
enum

RECOVERY
Allowed:

BOOT

RECOVERY

OTHER
percentUsed
integer | null
0 to 100
Percentage of space used.

25
fileVault2State
string
enum

INELIGIBLE

UNKNOWN

UNENCRYPTED

INELIGIBLE

DECRYPTED

DECRYPTING

ENCRYPTED

ENCRYPTING

RESTART_NEEDED

OPTIMIZING

DECRYPTING_PAUSED

ENCRYPTING_PAUSED
fileVault2ProgressPercent
integer | null
Percentage progress of current FileVault 2 operation.

45
lvmManaged
boolean | null

false

ADD object

ADD object
security
object

security object
sipStatus
string | null
enum

NOT_AVAILABLE
Allowed:

NOT_COLLECTED

NOT_AVAILABLE

DISABLED

ENABLED
gatekeeperStatus
string | null
enum

NOT_COLLECTED
Allowed:

NOT_COLLECTED

DISABLED

APP_STORE_AND_IDENTIFIED_DEVELOPERS

APP_STORE
xprotectVersion
string | null
1.2.3
activationLockEnabled
boolean | null

false
recoveryLockEnabled
boolean | null

true
firewallEnabled
boolean | null

true
secureBootLevel
string | null
enum

MEDIUM_SECURITY
Allowed:

NO_SECURITY

MEDIUM_SECURITY

FULL_SECURITY

NOT_SUPPORTED

UNKNOWN
externalBootLevel
string | null
enum

DISALLOW_BOOTING_FROM_EXTERNAL_MEDIA
Allowed:

ALLOW_BOOTING_FROM_EXTERNAL_MEDIA

DISALLOW_BOOTING_FROM_EXTERNAL_MEDIA

NOT_SUPPORTED

UNKNOWN
configurationProfiles
array of objects | null

object

id
string | null
1
username
string | null
username
lastInstalled
date-time | null
2018-10-31T18:04:13Z
removable
boolean | null

true
displayName
string | null
Displayed profile
profileIdentifier
string | null
0ae590fe-9b30-11ea-bb37-0242ac130002

object

id
string | null
1
username
string | null
username
lastInstalled
date-time | null
2018-10-31T18:04:13Z
removable
boolean | null

true
displayName
string | null
Displayed profile
profileIdentifier
string | null
0ae590fe-9b30-11ea-bb37-0242ac130002

ADD object
printers
array of objects | null

object

name
string | null
My Printer
type
string | null
XYZ 1122
uri
string | null
ipp://10.0.0.5
location
string | null
7th floor

ADD object
services
array of objects | null

object

name
string | null
SomeService

object

name
string | null
SomeService

ADD object
localUserAccounts
array of objects | null

object

uid
string | null
501
userGuid
string | null
844F1177-0CF5-40C6-901F-38EDD9969C1C
username
string | null
jamf
fullName
string | null
j
admin
boolean | null

true
homeDirectory
string | null
/Users/jamf
homeDirectorySizeMb
int64 | null
Home directory size in MB.

131072
fileVault2Enabled
boolean | null

true
userAccountType
string | null
enum

LOCAL
Allowed:

LOCAL

MOBILE

UNKNOWN
passwordMinLength
integer | null
4
passwordMaxAge
integer | null
5
passwordMinComplexCharacters
integer | null
5
passwordHistoryDepth
integer | null
5
passwordRequireAlphanumeric
boolean | null

true
computerAzureActiveDirectoryId
string | null
1
userAzureActiveDirectoryId
string | null
1

ADD object
certificates
array of objects | null

object

commonName
string | null
jamf.com
identity
boolean | null

true
username
string | null
test

ADD object
packageReceipts
object
All package receipts are listed by their package name


packageReceipts object
installedByJamfPro
array of strings | null

string

com.jamf.protect.JamfProtect

string

com.jamf.protect.JamfProtect2

ADD string
installedByInstallerSwu
array of strings | null

string

com.thing

string


ADD string
cached
array of strings | null

string


ADD string
softwareUpdates
array of objects | null

object

name
string | null
BEdit
version
string | null
1.15.2
packageName
string | null
com.apple.pkg.AdditionalEssentials

ADD object
purchasing
object

purchasing object
leased
boolean | null

true
purchased
boolean | null

true
poNumber
string | null
53-1
poDate
date | null
2019-01-01
vendor
string | null
Example Vendor
warrantyDate
date | null
2019-01-01
appleCareId
string | null
abcd
leaseDate
date | null
2019-01-01
purchasePrice
string | null
$500
lifeExpectancy
integer | null
5
purchasingAccount
string | null
admin
purchasingContact
string | null
true
userAndLocation
object

userAndLocation object
username
string | null
Madison Anderson
realname
string | null
13-inch MacBook
email
string | null
email@com.pl
position
string | null
IT Team Lead
phone
string | null
123-456-789
departmentId
string | null
1
buildingId
string | null
1
room
string | null
5
hardware
object

hardware object
make
string
Apple
model
string
13-inch MacBook Pro (Mid 2012)
modelIdentifier
string
MacBookPro9,2
serialNumber
string
C02ZC2QYLVDL
processorSpeedMhz
int64
Processor Speed in MHz.

2100
processorCount
integer
2
coreCount
integer
2
processorType
string
Intel Core i5
processorArchitecture
string
i386
busSpeedMhz
int64
2133
cacheSizeKilobytes
int64
Cache Size in KB.

3072
networkAdapterType
string
Foo
macAddress
string
6A:2C:4B:B7:65:B5
altNetworkAdapterType
string
Bar
altMacAddress
string
82:45:58:44:dc:01
totalRamMegabytes
int64
Total RAM Size in MB.

4096
openRamSlots
integer
Available RAM slots.

0
batteryCapacityPercent
integer
0 to 100
Remaining percentage of battery power.

85
batteryHealth
string
enum
Defaults to UNKNOWN
NON_GENUINE: The battery isnâ€™t a genuine Apple battery.
NORMAL: The battery is operating normally.
SERVICE_RECOMMENDED: The system recommends battery service.
UNKNOWN: The system couldnâ€™t determine battery health information.
UNSUPPORTED: The device doesnâ€™t support battery health reporting.

SERVICE_RECOMMENDED
Allowed:

NON_GENUINE

NORMAL

SERVICE_RECOMMENDED

UNKNOWN

UNSUPPORTED
smcVersion
string
2.2f38
nicSpeed
string
N/A
opticalDrive
string
MATSHITA DVD-R UJ-8A8
bootRom
string
MBP91.00D3.B08
bleCapable
boolean

true
supportsIosAppInstalls
boolean

true
appleSilicon
boolean

true
operatingSystem
object

operatingSystem object
name
string | null
Mac OS X
version
string | null
10.9.5
build
string | null
13A603
supplementalBuildVersion
string | null
13A953
rapidSecurityResponse
string | null
(a)
activeDirectoryStatus
string | null
Not Bound
softwareUpdateDeviceId
string | null
J132AP

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v3/computers-inventory \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "general": {
    "remoteManagement": {
      "managed": true
    },
    "name": "Boalime",
    "lastIpAddress": "247.185.82.186",
    "lastReportedIp": "247.185.82.186",
    "jamfBinaryVersion": "9.27",
    "platform": "MAC",
    "barcode1": "5 12345 678900",
    "barcode2": "5 12345 678900",
    "assetTag": "304822",
    "supervised": true,
    "mdmCapable": true,
    "reportDate": "2018-10-31T18:04:13Z",
    "lastContactTime": "2018-10-31T18:04:13Z",
    "lastCloudBackupDate": "2018-10-31T18:04:13Z",
    "lastEnrolledDate": "2018-10-31T18:04:13Z",
    "distributionPointId": "1",
    "siteId": "1",
    "itunesStoreAccountActive": true,
    "enrolledViaAutomatedDeviceEnrollment": false,
    "userApprovedMdm": true,
    "declarativeDeviceManagementEnabled": true
  },
  "storage": {
    "disks": [
      {
        "device": "disk0",
        "model": "APPLE HDD TOSHIBA MK5065GSXF",
        "revision": "5",
        "serialNumber": "a8598f013366",
        "sizeMegabytes": 262144,
        "smartStatus": "OK",
        "type": "false",
        "partitions": [
          {
            "name": "Foo",
            "sizeMegabytes": 262144,
            "availableMegabytes": 131072,
            "partitionType": "RECOVERY",
            "percentUsed": 25,
            "fileVault2State": "INELIGIBLE",
            "fileVault2ProgressPercent": 45,
            "lvmManaged": false
          }
        ]
      }
    ]
  },
  "security": {
    "sipStatus": "NOT_AVAILABLE",
    "gatekeeperStatus": "NOT_COLLECTED",
    "xprotectVersion": "1.2.3",
    "activationLockEnabled": false,
    "recoveryLockEnabled": true,
    "firewallEnabled": true,
    "secureBootLevel": "MEDIUM_SECURITY",
    "externalBootLevel": "DISALLOW_BOOTING_FROM_EXTERNAL_MEDIA"
  },
  "packageReceipts": {
    "installedByJamfPro": [
      "com.jamf.protect.JamfProtect",
      "com.jamf.protect.JamfProtect2"
    ],
    "installedByInstallerSwu": [
      "com.thing"
    ]
  },
  "purchasing": {
    "leased": true,
    "purchased": true,
    "poNumber": "53-1",
    "poDate": "2019-01-01",
    "vendor": "Example Vendor",
    "warrantyDate": "2019-01-01",
    "appleCareId": "abcd",
    "leaseDate": "2019-01-01",
    "purchasePrice": "$500",
    "lifeExpectancy": 5,
    "purchasingAccount": "admin",
    "purchasingContact": "true"
  },
  "userAndLocation": {
    "username": "Madison Anderson",
    "realname": "13-inch MacBook",
    "email": "email@com.pl",
    "position": "IT Team Lead",
    "phone": "123-456-789",
    "departmentId": "1",
    "buildingId": "1",
    "room": "5"
  },
  "hardware": {
    "batteryHealth": "SERVICE_RECOMMENDED",
    "make": "Apple",
    "model": "13-inch MacBook Pro (Mid 2012)",
    "modelIdentifier": "MacBookPro9,2",
    "serialNumber": "C02ZC2QYLVDL",
    "processorSpeedMhz": 2100,
    "processorCount": 2,
    "coreCount": 2,
    "processorType": "Intel Core i5",
    "processorArchitecture": "i386",
    "busSpeedMhz": 2133,
    "cacheSizeKilobytes": 3072,
    "networkAdapterType": "Foo",
    "macAddress": "6A:2C:4B:B7:65:B5",
    "altNetworkAdapterType": "Bar",
    "altMacAddress": "82:45:58:44:dc:01",
    "totalRamMegabytes": 4096,
    "openRamSlots": 0,
    "batteryCapacityPercent": 85,
    "smcVersion": "2.2f38",
    "nicSpeed": "N/A",
    "opticalDrive": "MATSHITA DVD-R UJ-8A8",
    "bootRom": "MBP91.00D3.B08",
    "bleCapable": true,
    "supportsIosAppInstalls": true,
    "appleSilicon": true
  },
  "operatingSystem": {
    "name": "Mac OS X",
    "version": "10.9.5",
    "build": "13A603",
    "supplementalBuildVersion": "13A953",
    "rapidSecurityResponse": "(a)",
    "activeDirectoryStatus": "Not Bound",
    "softwareUpdateDeviceId": "J132AP"
  },
  "applications": [
    {
      "name": "Microsoft Word",
      "path": "/usr/local/app",
      "version": "1.0.0"
    },
    {
      "name": "Microsoft Word",
      "path": "/usr/local/app",
      "version": "1.0.0"
    }
  ],
  "configurationProfiles": [
    {
      "id": "1",
      "username": "username",
      "lastInstalled": "2018-10-31T18:04:13Z",
      "removable": true,
      "displayName": "Displayed profile",
      "profileIdentifier": "0ae590fe-9b30-11ea-bb37-0242ac130002"
    },
    {
      "id": "1",
      "username": "username",
      "lastInstalled": "2018-10-31T18:04:13Z",
      "removable": true,
      "displayName": "Displayed profile",
      "profileIdentifier": "0ae590fe-9b30-11ea-bb37-0242ac130002"
    }
  ],
  "printers": [
    {
      "name": "My Printer",
      "type": "XYZ 1122",
      "uri": "ipp://10.0.0.5",
      "location": "7th floor"
    }
  ],
  "services": [
    {
      "name": "SomeService"
    },
    {
      "name": "SomeService"
    }
  ],
  "localUserAccounts": [
    {
      "uid": "501",
      "userGuid": "844F1177-0CF5-40C6-901F-38EDD9969C1C",
      "username": "jamf",
      "fullName": "j",
      "admin": true,
      "homeDirectory": "/Users/jamf",
      "homeDirectorySizeMb": 131072,
      "fileVault2Enabled": true,
      "userAccountType": "LOCAL",
      "passwordMinLength": 4,
      "passwordMaxAge": 5,
      "passwordMinComplexCharacters": 5,
      "passwordHistoryDepth": 5,
      "passwordRequireAlphanumeric": true,
      "computerAzureActiveDirectoryId": "1",
      "userAzureActiveDirectoryId": "1"
    }
  ],
  "certificates": [
    {
      "commonName": "jamf.com",
      "identity": true,
      "username": "test"
    }
  ],
  "softwareUpdates": [
    {
      "name": "BEdit",
      "version": "1.15.2",
      "packageName": "com.apple.pkg.AdditionalEssentials"
    }
  ]
}
'

{
  "id": "1",
  "href": "https://yourJamfProUrl.jamf/api/v1/resource/1"
}
-----