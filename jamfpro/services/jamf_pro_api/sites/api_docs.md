Find all sites
get
https://yourServer.jamfcloud.com/api/v1/sites

Find all sites

Response

200
Successful response

Response body
array of objects
object
id
string
name
string

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/sites \
     --header 'accept: application/json'

[
  {
    "id": "1",
    "name": "Eau Claire"
  }
]
-----

Find and filter site objects for a site ID
get
https://yourServer.jamfcloud.com/api/v1/sites/{id}/objects

Find site objects for Site ID, with the ability to filter out different object types and object IDs for the site ID

Path Params
id
string
required
Site ID to get objects for

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
Defaults to objectType:asc
Sorting criteria in the format: property:asc/desc. Default sort is objectType:asc. Multiple sort criteria are supported and must be separated with a comma.

Example: sort=objectId:asc,objectType:desc.


string

objectType:asc

ADD string
filter
string
Defaults to objectType=="User"
Query in the RSQL format, allowing filter of site object information. Default filter returns all objects for the site ID.

Fields allowed in the query: objectType, objectId

Example: filter=objectType=="User"

List of objectType options (case-insensitive) ["Computer", "Peripheral", "Licensed Software", "Licensed Software Template", "Policy", "macOS Configuration Profile", "Restricted Software", "Managed Preference Profile", "Computer Group", "Mobile Device", "Apple TV", "Android Device", "User Group", "iOS Configuration Profile", "Mobile Device App", "E-book", "Mobile Device Group", "Classroom", "Advanced Computer Search", "Advanced Mobile Search", "Advanced User Search", "Advanced User Content Search", "Computer Invitation", "Mobile Device Invitation", "Mobile Device Enrollment Profile", "Device Enrollment Program Instance", "Mobile Device Prestage", "Computer DEP Prestage", "Enrollment Customization", "VPP Location", "VPP Subscription", "VPP Invitation", "VPP Assignment", "User", "Network Integration", "Mac App", "App Installer", "Self Service Plugin", "Software Title", "Patch Software Title Summary", "Patch Policy", "Patch Software Title Configuration", "Change Password", "Mobile Device Inventory", "Computer Inventory", "Change Management", "Licensed Software License"]

objectType=="User"
Response

200
Successful response

Response body
array of objects
object
siteId
string
required
objectType
string
enum
required
Computer Peripheral Licensed Software Licensed Software Template Policy macOS Configuration Profile Restricted Software Managed Preference Profile Computer Group Mobile Device Apple TV User Group iOS Configuration Profile Mobile Device App E-book Mobile Device Group Classroom Advanced Computer Search Advanced Mobile Search Advanced User Search Advanced User Content Search Computer Invitation Mobile Device Invitation Mobile Device Enrollment Profile Device Enrollment Program Instance Mobile Device Prestage Computer DEP Prestage Enrollment Customization VPP Location VPP Subscription VPP Invitation VPP Assignment User Network Integration Mac App App Installer Self Service Plugin Software Title Patch Software Title Summary Patch Policy Patch Software Title Configuration Change Password Mobile Device Inventory Computer Inventory Change Management Licensed Software License Unknown

objectId
string
required

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v1/sites//objects?page=0&page-size=100&sort=objectType%3Aasc&filter=objectType%3D%3D%22User%22' \
     --header 'accept: application/json'

[
  {
    "siteId": "1",
    "objectType": "User",
    "objectId": "1"
  }
]
-----
