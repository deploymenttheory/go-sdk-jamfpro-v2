Get Ebook object
get
https://yourServer.jamfcloud.com/api/v1/ebooks

Gets ebook object

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
Defaults to name:asc
Sorting criteria in the format: property:asc/desc. Default sort is name:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc


string

name:asc

ADD string
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
object
id
string
name
string
kind
string
enum
UNKNOWN PDF EPUB IBOOKS

url
string
free
boolean
version
string
author
string
deployAsManaged
boolean
If true, it will be automatically installed

installAutomatically
boolean
categoryId
string
siteId
string

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v1/ebooks?page=0&page-size=100&sort=name%3Aasc' \
     --header 'accept: application/json'

{
  "totalCount": 3,
  "results": [
    {
      "id": "1",
      "name": "The Neverending API",
      "kind": "IBOOKS",
      "url": "https://jamf.com/ibooks/the_neverending_api.just_kidding",
      "free": true,
      "version": "10.9.0",
      "author": "IT Bob",
      "deployAsManaged": false,
      "installAutomatically": false,
      "categoryId": "-1",
      "siteId": "-1"
    }
  ]
}
-----
Get specified Ebook object
get
https://yourServer.jamfcloud.com/api/v1/ebooks/{id}

Gets specified Ebook object

Path Params
id
string
required
instance id of ebook record

1
Responses

200
Details about ebook were found for given id

Response body
object
id
string
name
string
kind
string
enum
UNKNOWN PDF EPUB IBOOKS

url
string
free
boolean
version
string
author
string
deployAsManaged
boolean
If true, it will be automatically installed

installAutomatically
boolean
categoryId
string
siteId
string

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/ebooks/1 \
     --header 'accept: application/json'

{
  "id": "1",
  "name": "The Neverending API",
  "kind": "IBOOKS",
  "url": "https://jamf.com/ibooks/the_neverending_api.just_kidding",
  "free": true,
  "version": "10.9.0",
  "author": "IT Bob",
  "deployAsManaged": false,
  "installAutomatically": false,
  "categoryId": "-1",
  "siteId": "-1"
}
-----
Get specified scope of Ebook object
get
https://yourServer.jamfcloud.com/api/v1/ebooks/{id}/scope

Gets specified scope of Ebook object

Path Params
id
string
required
instance id of ebook record

1
Responses

200
Details of scope for ebook were found

Response body
object
allComputers
boolean
allMobileDevices
boolean
allUsers
boolean
computerIds
array of strings
computerGroupIds
array of strings
mobileDeviceIds
array of strings
mobileDeviceGroupIds
array of strings
buildingIds
array of strings
departmentIds
array of strings
userIds
array of strings
userGroupIds
array of strings
classroomIds
array of strings
limitations
object
networkSegments
array of strings
users
array of objects
object
name
string
userGroups
array of strings
exclusions
object
computerIds
array of strings
computerGroupIds
array of strings
mobileDeviceIds
array of strings
mobileDeviceGroupIds
array of strings
buildingIds
array of strings
departmentIds
array of strings
userIds
array of strings
userGroupIds
array of strings
limitations
object

limitations object
networkSegments
array of strings
users
array of objects
object
name
string
userGroups
array of strings

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/ebooks/1/scope \
     --header 'accept: application/json'

{
  "allComputers": false,
  "allMobileDevices": false,
  "allUsers": false,
  "computerIds": [
    "-1"
  ],
  "computerGroupIds": [
    "-1"
  ],
  "mobileDeviceIds": [
    "-1"
  ],
  "mobileDeviceGroupIds": [
    "-1"
  ],
  "buildingIds": [
    "-1"
  ],
  "departmentIds": [
    "-1"
  ],
  "userIds": [
    "-1"
  ],
  "userGroupIds": [
    "-1"
  ],
  "classroomIds": [
    "-1"
  ],
  "limitations": {
    "networkSegments": [
      "1"
    ],
    "users": [
      {
        "name": "admin"
      }
    ],
    "userGroups": [
      "1"
    ]
  },
  "exclusions": {
    "computerIds": [
      "-1"
    ],
    "computerGroupIds": [
      "-1"
    ],
    "mobileDeviceIds": [
      "-1"
    ],
    "mobileDeviceGroupIds": [
      "-1"
    ],
    "buildingIds": [
      "-1"
    ],
    "departmentIds": [
      "-1"
    ],
    "userIds": [
      "-1"
    ],
    "userGroupIds": [
      "-1"
    ],
    "limitations": {
      "networkSegments": [
        "1"
      ],
      "users": [
        {
          "name": "admin"
        }
      ],
      "userGroups": [
        "1"
      ]
    }
  }
}
-----