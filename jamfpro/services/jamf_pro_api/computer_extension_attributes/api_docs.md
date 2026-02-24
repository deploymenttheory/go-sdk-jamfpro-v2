Retrieve Computer Extension Attributes.
get
https://yourServer.jamfcloud.com/api/v1/computer-extension-attributes

Retrieves All Computer Extension Attributes Configuration.

Query Params
page
integer
Defaults to 0
1
page-size
integer
Defaults to 100
100
sort
array of strings
Defaults to name.asc
Sorts results by one or more criteria, following the format property:asc/desc.
Default sort is name:asc.
If using multiple criteria, separate with commas. Allows sort for id and name.


string

name.asc

ADD string
filter
string
Filters results. Use RSQL format for query. Allows for many fields, including ID, name, etc.
Can be combined with paging and sorting.
Fields allowed in the query: id, name
Default filter is an empty query and returns all results from the requested page.

Responses

200
Computer Extension Attribute objects has been fetched successfully.

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
Unique Id for Mobile Device Extension Attribute.

name
string
required
Display name for the extension attribute.

description
string
Description for the extension attribute.

dataType
string
enum
required
Defaults to STRING
Type of data being collected.

INTEGER STRING DATE

enabled
boolean
Defaults to true
Enabled by default, but for inputType Script we can disable it as well.
Possible values are:
false
true

inventoryDisplayType
string
enum
required
Defaults to GENERAL
Category in which to display the extension attribute in Jamf Pro.

GENERAL HARDWARE OPERATING_SYSTEM USER_AND_LOCATION PURCHASING EXTENSION_ATTRIBUTES

inputType
string
enum
required
Defaults to TEXT
Extension attributes collect inventory data by using an input type.The type of the Input used to populate the extension attribute.

SCRIPT TEXT POPUP DIRECTORY_SERVICE_ATTRIBUTE_MAPPING

scriptContents
string | null
When we run this script it returns a data value each time a computer submits inventory to Jamf Pro. Provide scriptContents only when inputType is 'SCRIPT'.

popupMenuChoices
array of strings
length ≥ 0
When added with list of choices while creating computer extension attributes these Pop-up menu can be displayed in inventory information. User can choose a value from the pop-up menu list when enrolling a computer any time using Jamf Pro. Provide popupMenuChoices only when inputType is 'POPUP'.

ldapAttributeMapping
string
Directory Service attribute use to populate the extension attribute.
Required when inputType is "DIRECTORY_SERVICE_ATTRIBUTE_MAPPING"

ldapExtensionAttributeAllowed
boolean
Defaults to false
Collect multiple values for this extension attribute. ldapExtensionAttributeAllowed is disabled by default, only for inputType 'DIRECTORY_SERVICE_ATTRIBUTE_MAPPING' it can be enabled. It's value cannot be modified during edit operation.
Possible values are:
false
true

manageExistingData
string
enum
It is used to specify to either delete or retain the extension attributes values when inputType is Script and enabled is false.

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v1/computer-extension-attributes?page=1&page-size=100&sort=name.asc' \
     --header 'accept: application/json'

{
  "totalCount": 1,
  "results": [
    {
      "id": "1",
      "name": "MobileDeviceExtensionAttribute",
      "description": "Mobile Device Extension Attribute",
      "dataType": "STRING",
      "enabled": true,
      "inventoryDisplayType": "GENERAL",
      "inputType": "TEXT",
      "scriptContents": "scriptContent",
      "popupMenuChoices": [
        "Test",
        "Popup"
      ],
      "ldapAttributeMapping": "ldapAttributeMapping",
      "ldapExtensionAttributeAllowed": false,
      "manageExistingData": "RETAIN"
    }
  ]
}
-----
Create Computer Extension Attribute.
post
https://yourServer.jamfcloud.com/api/v1/computer-extension-attributes

Create Computer Extension Attribute to collect extra inventory information.

Body Params
Computer Extension Attribute to be created.

name
string
required
Display name for the extension attribute.

MobileDeviceExtensionAttribute
description
string
Description for the extension attribute.

Mobile Device Extension Attribute
dataType
string
enum
required
Defaults to STRING
Type of data being collected.


STRING
Allowed:

INTEGER

STRING

DATE
enabled
boolean
Defaults to true
Enabled by default, but for inputType Script we can disable it as well.
Possible values are:
false
true


true
inventoryDisplayType
string
enum
required
Defaults to GENERAL
Category in which to display the extension attribute in Jamf Pro.


GENERAL
Allowed:

GENERAL

HARDWARE

OPERATING_SYSTEM

USER_AND_LOCATION

PURCHASING

EXTENSION_ATTRIBUTES
inputType
string
enum
required
Defaults to TEXT
Extension attributes collect inventory data by using an input type.The type of the Input used to populate the extension attribute.


TEXT
Allowed:

SCRIPT

TEXT

POPUP

DIRECTORY_SERVICE_ATTRIBUTE_MAPPING
scriptContents
string | null
When we run this script it returns a data value each time a computer submits inventory to Jamf Pro. Provide scriptContents only when inputType is 'SCRIPT'.

popupMenuChoices
array of strings
length ≥ 0
When added with list of choices while creating computer extension attributes these Pop-up menu can be displayed in inventory information. User can choose a value from the pop-up menu list when enrolling a computer any time using Jamf Pro. Provide popupMenuChoices only when inputType is 'POPUP'.


ADD string
ldapAttributeMapping
string
Directory Service attribute use to populate the extension attribute.
Required when inputType is "DIRECTORY_SERVICE_ATTRIBUTE_MAPPING"

ldapExtensionAttributeAllowed
boolean
Defaults to false
Collect multiple values for this extension attribute. ldapExtensionAttributeAllowed is disabled by default, only for inputType 'DIRECTORY_SERVICE_ATTRIBUTE_MAPPING' it can be enabled. It's value cannot be modified during edit operation.
Possible values are:
false
true


false
manageExistingData
string
enum
It is used to specify to either delete or retain the extension attributes values when inputType is Script and enabled is false.


RETAIN
Allowed:

RETAIN

DELETE
Responses

201
Computer Extension Attribute was successfully created.

Response body
object
id
string
length ≥ 1
href
string

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/computer-extension-attributes \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "dataType": "STRING",
  "enabled": true,
  "inventoryDisplayType": "GENERAL",
  "inputType": "TEXT",
  "ldapExtensionAttributeAllowed": false,
  "name": "MobileDeviceExtensionAttribute",
  "description": "Mobile Device Extension Attribute"
}
'

{
  "id": "1",
  "href": "https://yourJamfProUrl.jamf/api/v1/resource/1"
}
-----
Delete multiple Computer Extension Attribute at once.
post
https://yourServer.jamfcloud.com/api/v1/computer-extension-attributes/delete-multiple

IDs of the Computer Extension Attribute to be deleted.

Body Params
IDs of the Computer Extension Attribute to be deleted

ids
array of strings

string

1,2

ADD string
Responses
204
All specified Computer Extension Attribute deleted successfully.

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/computer-extension-attributes/delete-multiple \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '{"ids":["1,2"]}'

-----
Retrieve All Computer Extension Attributes Templates.
get
https://yourServer.jamfcloud.com/api/v1/computer-extension-attributes/templates

Retrieves All Computer Extension Attributes Templates.

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
Defaults to templateName.asc
Sorts results by one or more criteria, following the format property:asc/desc.
Default sort is templateName:asc.
If using multiple criteria, separate with commas. Allows sort for templateName and templateCategory.


string

templateName.asc

ADD string
filter
string
Filters results. Use RSQL format for queries. which allows filtering by multiple fields such as templateName, templateCategoryName.
Can be combined with paging and sorting.
Fields allowed in the query: templateName, templateCategoryName
Default filter is an empty query and returns all results from the requested page.

Response

200
Computer Extension Attribute templates objects has been fetched successfully.

Response body
object
totalCount
integer
required
results
array of objects
length ≥ 0
object
templateId
string
Unique Id for Computer Extension Attribute Template.

templateName
string
required
Template display name of the extension attribute.

templateCategoryName
string
required
Template category name of the extension attribute.

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v1/computer-extension-attributes/templates?page=0&page-size=100&sort=templateName.asc' \
     --header 'accept: application/json'

{
  "totalCount": 1,
  "results": [
    {
      "templateId": "1",
      "templateName": "ClamXav - Virus Definition Date",
      "templateCategoryName": "AntiVirus"
    }
  ]
}
-----
Get specified Computer Extension Attribute Template object.
get
https://yourServer.jamfcloud.com/api/v1/computer-extension-attributes/templates/{id}

Gets specified Computer Extension Attribute Template object.

Path Params
id
string
required
Unique Id of the Template.

Responses

200
Computer Extension Attribute Template object has been fetched successfully.

Response body
object
id
string
Unique Id for Mobile Device Extension Attribute.

name
string
required
Display name for the extension attribute.

description
string
Description for the extension attribute.

dataType
string
enum
required
Defaults to STRING
Type of data being collected.

INTEGER STRING DATE

enabled
boolean
Defaults to true
Enabled by default, but for inputType Script we can disable it as well.
Possible values are:
false
true

inventoryDisplayType
string
enum
required
Defaults to GENERAL
Category in which to display the extension attribute in Jamf Pro.

GENERAL HARDWARE OPERATING_SYSTEM USER_AND_LOCATION PURCHASING EXTENSION_ATTRIBUTES

inputType
string
enum
required
Defaults to TEXT
Extension attributes collect inventory data by using an input type.The type of the Input used to populate the extension attribute.

SCRIPT TEXT POPUP DIRECTORY_SERVICE_ATTRIBUTE_MAPPING

scriptContents
string | null
When we run this script it returns a data value each time a computer submits inventory to Jamf Pro. Provide scriptContents only when inputType is 'SCRIPT'.

popupMenuChoices
array of strings
length ≥ 0
When added with list of choices while creating computer extension attributes these Pop-up menu can be displayed in inventory information. User can choose a value from the pop-up menu list when enrolling a computer any time using Jamf Pro. Provide popupMenuChoices only when inputType is 'POPUP'.

ldapAttributeMapping
string
Directory Service attribute use to populate the extension attribute.
Required when inputType is "DIRECTORY_SERVICE_ATTRIBUTE_MAPPING"

ldapExtensionAttributeAllowed
boolean
Defaults to false
Collect multiple values for this extension attribute. ldapExtensionAttributeAllowed is disabled by default, only for inputType 'DIRECTORY_SERVICE_ATTRIBUTE_MAPPING' it can be enabled. It's value cannot be modified during edit operation.
Possible values are:
false
true

manageExistingData
string
enum
It is used to specify to either delete or retain the extension attributes values when inputType is Script and enabled is false.

RETAIN DELETE

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/computer-extension-attributes/templates/ \
     --header 'accept: application/json'

{
  "id": "1",
  "name": "MobileDeviceExtensionAttribute",
  "description": "Mobile Device Extension Attribute",
  "dataType": "STRING",
  "enabled": true,
  "inventoryDisplayType": "GENERAL",
  "inputType": "TEXT",
  "scriptContents": "scriptContent",
  "popupMenuChoices": [
    "Test",
    "Popup"
  ],
  "ldapAttributeMapping": "ldapAttributeMapping",
  "ldapExtensionAttributeAllowed": false,
  "manageExistingData": "RETAIN"
}
-----
Upload Computer Extension Attribute.
post
https://yourServer.jamfcloud.com/api/v1/computer-extension-attributes/upload

Uploads a Computer Extension Attribute.

Body Params
file
file
required
The file to upload

No file chosen
Responses

200
Computer Extension Attribute object has been fetched successfully.

Response body
object
id
string
Unique Id for Mobile Device Extension Attribute.

name
string
required
Display name for the extension attribute.

description
string
Description for the extension attribute.

dataType
string
enum
required
Defaults to STRING
Type of data being collected.

INTEGER STRING DATE

enabled
boolean
Defaults to true
Enabled by default, but for inputType Script we can disable it as well.
Possible values are:
false
true

inventoryDisplayType
string
enum
required
Defaults to GENERAL
Category in which to display the extension attribute in Jamf Pro.

GENERAL HARDWARE OPERATING_SYSTEM USER_AND_LOCATION PURCHASING EXTENSION_ATTRIBUTES

inputType
string
enum
required
Defaults to TEXT
Extension attributes collect inventory data by using an input type.The type of the Input used to populate the extension attribute.

SCRIPT TEXT POPUP DIRECTORY_SERVICE_ATTRIBUTE_MAPPING

scriptContents
string | null
When we run this script it returns a data value each time a computer submits inventory to Jamf Pro. Provide scriptContents only when inputType is 'SCRIPT'.

popupMenuChoices
array of strings
length ≥ 0
When added with list of choices while creating computer extension attributes these Pop-up menu can be displayed in inventory information. User can choose a value from the pop-up menu list when enrolling a computer any time using Jamf Pro. Provide popupMenuChoices only when inputType is 'POPUP'.

ldapAttributeMapping
string
Directory Service attribute use to populate the extension attribute.
Required when inputType is "DIRECTORY_SERVICE_ATTRIBUTE_MAPPING"

ldapExtensionAttributeAllowed
boolean
Defaults to false
Collect multiple values for this extension attribute. ldapExtensionAttributeAllowed is disabled by default, only for inputType 'DIRECTORY_SERVICE_ATTRIBUTE_MAPPING' it can be enabled. It's value cannot be modified during edit operation.
Possible values are:
false
true

manageExistingData
string
enum
It is used to specify to either delete or retain the extension attributes values when inputType is Script and enabled is false.

RETAIN DELETE

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/computer-extension-attributes/upload \
     --header 'accept: application/json' \
     --header 'content-type: multipart/form-data'

{
  "id": "1",
  "name": "MobileDeviceExtensionAttribute",
  "description": "Mobile Device Extension Attribute",
  "dataType": "STRING",
  "enabled": true,
  "inventoryDisplayType": "GENERAL",
  "inputType": "TEXT",
  "scriptContents": "scriptContent",
  "popupMenuChoices": [
    "Test",
    "Popup"
  ],
  "ldapAttributeMapping": "ldapAttributeMapping",
  "ldapExtensionAttributeAllowed": false,
  "manageExistingData": "RETAIN"
}
-----
Get specified Computer Extension Attribute object.
get
https://yourServer.jamfcloud.com/api/v1/computer-extension-attributes/{id}

Gets specified Computer Extension Attribute object.

Path Params
id
string
required
Unique ID of Computer Extension Attribute.

1
Responses

200
Computer Extension Attribute object has been fetched successfully.

Response body
object
id
string
Unique Id for Mobile Device Extension Attribute.

name
string
required
Display name for the extension attribute.

description
string
Description for the extension attribute.

dataType
string
enum
required
Defaults to STRING
Type of data being collected.

INTEGER STRING DATE

enabled
boolean
Defaults to true
Enabled by default, but for inputType Script we can disable it as well.
Possible values are:
false
true

inventoryDisplayType
string
enum
required
Defaults to GENERAL
Category in which to display the extension attribute in Jamf Pro.

GENERAL HARDWARE OPERATING_SYSTEM USER_AND_LOCATION PURCHASING EXTENSION_ATTRIBUTES

inputType
string
enum
required
Defaults to TEXT
Extension attributes collect inventory data by using an input type.The type of the Input used to populate the extension attribute.

SCRIPT TEXT POPUP DIRECTORY_SERVICE_ATTRIBUTE_MAPPING

scriptContents
string | null
When we run this script it returns a data value each time a computer submits inventory to Jamf Pro. Provide scriptContents only when inputType is 'SCRIPT'.

popupMenuChoices
array of strings
length ≥ 0
When added with list of choices while creating computer extension attributes these Pop-up menu can be displayed in inventory information. User can choose a value from the pop-up menu list when enrolling a computer any time using Jamf Pro. Provide popupMenuChoices only when inputType is 'POPUP'.

ldapAttributeMapping
string
Directory Service attribute use to populate the extension attribute.
Required when inputType is "DIRECTORY_SERVICE_ATTRIBUTE_MAPPING"

ldapExtensionAttributeAllowed
boolean
Defaults to false
Collect multiple values for this extension attribute. ldapExtensionAttributeAllowed is disabled by default, only for inputType 'DIRECTORY_SERVICE_ATTRIBUTE_MAPPING' it can be enabled. It's value cannot be modified during edit operation.
Possible values are:
false
true

manageExistingData
string
enum
It is used to specify to either delete or retain the extension attributes values when inputType is Script and enabled is false.

RETAIN DELETE

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/computer-extension-attributes/1 \
     --header 'accept: application/json'

{
  "id": "1",
  "name": "MobileDeviceExtensionAttribute",
  "description": "Mobile Device Extension Attribute",
  "dataType": "STRING",
  "enabled": true,
  "inventoryDisplayType": "GENERAL",
  "inputType": "TEXT",
  "scriptContents": "scriptContent",
  "popupMenuChoices": [
    "Test",
    "Popup"
  ],
  "ldapAttributeMapping": "ldapAttributeMapping",
  "ldapExtensionAttributeAllowed": false,
  "manageExistingData": "RETAIN"
}
-----
Update specified Computer Extension Attribute object.
put
https://yourServer.jamfcloud.com/api/v1/computer-extension-attributes/{id}

Update specified Computer Extension Attribute object.

Path Params
id
string
required
Unique ID of Computer Extension Attribute.

Body Params
Computer Extension Attribute object to be updated. IDs defined in this body will be ignored.

name
string
required
Display name for the extension attribute.

MobileDeviceExtensionAttribute
description
string
Description for the extension attribute.

Mobile Device Extension Attribute
dataType
string
enum
required
Defaults to STRING
Type of data being collected.


STRING
Allowed:

INTEGER

STRING

DATE
enabled
boolean
Defaults to true
Enabled by default, but for inputType Script we can disable it as well.
Possible values are:
false
true


true
inventoryDisplayType
string
enum
required
Defaults to GENERAL
Category in which to display the extension attribute in Jamf Pro.


GENERAL
Allowed:

GENERAL

HARDWARE

OPERATING_SYSTEM

USER_AND_LOCATION

PURCHASING

EXTENSION_ATTRIBUTES
inputType
string
enum
required
Defaults to TEXT
Extension attributes collect inventory data by using an input type.The type of the Input used to populate the extension attribute.


TEXT
Allowed:

SCRIPT

TEXT

POPUP

DIRECTORY_SERVICE_ATTRIBUTE_MAPPING
scriptContents
string | null
When we run this script it returns a data value each time a computer submits inventory to Jamf Pro. Provide scriptContents only when inputType is 'SCRIPT'.

popupMenuChoices
array of strings
length ≥ 0
When added with list of choices while creating computer extension attributes these Pop-up menu can be displayed in inventory information. User can choose a value from the pop-up menu list when enrolling a computer any time using Jamf Pro. Provide popupMenuChoices only when inputType is 'POPUP'.


ADD string
ldapAttributeMapping
string
Directory Service attribute use to populate the extension attribute.
Required when inputType is "DIRECTORY_SERVICE_ATTRIBUTE_MAPPING"

ldapAttributeMapping
ldapExtensionAttributeAllowed
boolean
Defaults to false
Collect multiple values for this extension attribute. ldapExtensionAttributeAllowed is disabled by default, only for inputType 'DIRECTORY_SERVICE_ATTRIBUTE_MAPPING' it can be enabled. It's value cannot be modified during edit operation.
Possible values are:
false
true


false
manageExistingData
string
enum
It is used to specify to either delete or retain the extension attributes values when inputType is Script and enabled is false.


RETAIN
Allowed:

RETAIN

DELETE
Responses

202
Computer Extension Attribute updated successfully.

Response body
object
id
string
Unique Id for Mobile Device Extension Attribute.

name
string
required
Display name for the extension attribute.

description
string
Description for the extension attribute.

dataType
string
enum
required
Defaults to STRING
Type of data being collected.

INTEGER STRING DATE

enabled
boolean
Defaults to true
Enabled by default, but for inputType Script we can disable it as well.
Possible values are:
false
true

inventoryDisplayType
string
enum
required
Defaults to GENERAL
Category in which to display the extension attribute in Jamf Pro.

GENERAL HARDWARE OPERATING_SYSTEM USER_AND_LOCATION PURCHASING EXTENSION_ATTRIBUTES

inputType
string
enum
required
Defaults to TEXT
Extension attributes collect inventory data by using an input type.The type of the Input used to populate the extension attribute.

SCRIPT TEXT POPUP DIRECTORY_SERVICE_ATTRIBUTE_MAPPING

scriptContents
string | null
When we run this script it returns a data value each time a computer submits inventory to Jamf Pro. Provide scriptContents only when inputType is 'SCRIPT'.

popupMenuChoices
array of strings
length ≥ 0
When added with list of choices while creating computer extension attributes these Pop-up menu can be displayed in inventory information. User can choose a value from the pop-up menu list when enrolling a computer any time using Jamf Pro. Provide popupMenuChoices only when inputType is 'POPUP'.

ldapAttributeMapping
string
Directory Service attribute use to populate the extension attribute.
Required when inputType is "DIRECTORY_SERVICE_ATTRIBUTE_MAPPING"

ldapExtensionAttributeAllowed
boolean
Defaults to false
Collect multiple values for this extension attribute. ldapExtensionAttributeAllowed is disabled by default, only for inputType 'DIRECTORY_SERVICE_ATTRIBUTE_MAPPING' it can be enabled. It's value cannot be modified during edit operation.
Possible values are:
false
true

manageExistingData
string
enum
It is used to specify to either delete or retain the extension attributes values when inputType is Script and enabled is false.

RETAIN DELETE

curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v1/computer-extension-attributes/ \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "dataType": "STRING",
  "enabled": true,
  "inventoryDisplayType": "GENERAL",
  "inputType": "TEXT",
  "ldapExtensionAttributeAllowed": false,
  "name": "MobileDeviceExtensionAttribute",
  "description": "Mobile Device Extension Attribute",
  "ldapAttributeMapping": "ldapAttributeMapping",
  "manageExistingData": "RETAIN"
}
'

-----
Remove specified Computer Extension Attribute.
delete
https://yourServer.jamfcloud.com/api/v1/computer-extension-attributes/{id}

ID of the Computer Extension Attribute to be deleted.

Path Params
id
string
required
Unique ID of Computer Extension Attribute.

curl --request DELETE \
     --url https://yourserver.jamfcloud.com/api/v1/computer-extension-attributes/1
-----
Get smart group/advance search dependent objects for a specified computer extension attribute
get
https://yourServer.jamfcloud.com/api/v1/computer-extension-attributes/{id}/data-dependency

Get smart group/advance search dependent objects for a specified computer extension attribute

Path Params
id
string
required
Unique ID of computer extension attribute.

1
Responses

200
Fetches list of dependent objects for a specified computer extension attribute.

Response body
object
totalCount
integer
results
array of objects
length ≥ 0
object
id
integer
Unique Id for Dependency Object.

objectId
integer
Object Type Id of the dependency object.

nameLocalization
string
Name of localization which display dependent object.

identifiableName
string
Name of the dependent object.

hyperlink
string
Link to dependent object or to page with list of dependent objects.

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/computer-extension-attributes/1/data-dependency \
     --header 'accept: application/json'

{
  "totalCount": 1,
  "results": [
    {
      "id": 1,
      "objectId": 1,
      "nameLocalization": "SMART_COMPUTER_GROUPS",
      "identifiableName": "identifiableName",
      "hyperlink": "/smartMobileDeviceGroups.html"
    }
  ]
}
-----
Download the specified Computer Extension Attribute.
get
https://yourServer.jamfcloud.com/api/v1/computer-extension-attributes/{id}/download

Retrieves the specified Computer Extension Attribute in XML format based on the provided unique ID.

Path Params
id
string
required
The unique ID of the Computer Extension Attribute to be downloaded.

1
Headers
accept
string
enum
Defaults to application/json
Generated from available response content types


application/xml
Allowed:

application/json

application/xml
Responses

200
The requested Computer Extension Attribute XML file was fetched successfully.

Response body
string

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/computer-extension-attributes/1/download \
     --header 'accept: application/xml'
-----
Get specified Computer Extension Attribute History object
get
https://yourServer.jamfcloud.com/api/v1/computer-extension-attributes/{id}/history

Get specified Computer Extension Attribute history object

Path Params
id
string
required
Instance ID of Computer Extension Attribute history

1
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
Details of Computer Extension Attribute history were found.

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
     --url 'https://yourserver.jamfcloud.com/api/v1/computer-extension-attributes/1/history?page=0&page-size=100&sort=id%3Aasc' \
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
Add specified Computer Extension Attribute history object notes
post
https://yourServer.jamfcloud.com/api/v1/computer-extension-attributes/{id}/history

Add specified Computer Extension Attribute history object notes

Path Params
id
string
required
Instance ID of Computer Extension Attribute history

Body Params
History note to be created

note
string
required
Responses

201
Computer Extension Attribute history note created successfully.

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
     --url https://yourserver.jamfcloud.com/api/v1/computer-extension-attributes//history \
     --header 'accept: application/json' \
     --header 'content-type: application/json'

{
  "id": 1,
  "username": "admin",
  "date": "2019-02-04T21:09:31.661Z",
  "note": "Sso settings update",
  "details": "Is SSO Enabled false\\nSelected SSO Provider"
}