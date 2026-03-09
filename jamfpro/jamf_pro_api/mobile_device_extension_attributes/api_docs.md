Retrieve Mobile Device Extension Attributes.
get
https://yourServer.jamfcloud.com/api/v1/mobile-device-extension-attributes

Retrieves all mobile device extension attributes configuration.

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
Mobile Device Extension Attribute objects has been fetched successfully.

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

inventoryDisplayType
string
enum
required
Defaults to GENERAL
Category in which to display the extension attribute in Jamf Pro.

GENERAL HARDWARE USER_AND_LOCATION PURCHASING EXTENSION_ATTRIBUTES

inputType
string
enum
required
Defaults to TEXT
Extension attributes collect inventory data by using an input type.The type of the Input used to populate the extension attribute.

TEXT POPUP DIRECTORY_SERVICE_ATTRIBUTE_MAPPING

popupMenuChoices
array of strings
length ≥ 0
When added with list of choices while creating mobile device extension attributes these Pop-up menu can be displayed in inventory information. User can choose a value from the pop-up menu list when enrolling a mobile device any time using Jamf Pro. Provide popupMenuChoices only when inputType is 'POPUP'.

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

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v1/mobile-device-extension-attributes?page=0&page-size=100&sort=name.asc' \
     --header 'accept: application/json'

{
  "totalCount": 1,
  "results": [
    {
      "id": "1",
      "name": "MobileDeviceExtensionAttribute",
      "description": "Mobile Device Extension Attribute",
      "dataType": "STRING",
      "inventoryDisplayType": "GENERAL",
      "inputType": "TEXT",
      "popupMenuChoices": [
        "Test",
        "Popup"
      ],
      "ldapAttributeMapping": "ldapAttributeMapping",
      "ldapExtensionAttributeAllowed": false
    }
  ]
}
-----
Get specified Mobile Device Extension Attribute object.
get
https://yourServer.jamfcloud.com/api/v1/mobile-device-extension-attributes/{id}

Gets specified Mobile Device Extension Attribute object.

Path Params
id
string
required
Unique ID of Mobile Device Extension Attribute.

Responses

200
Mobile Device Extension Attribute object has been fetched successfully.

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

inventoryDisplayType
string
enum
required
Defaults to GENERAL
Category in which to display the extension attribute in Jamf Pro.

GENERAL HARDWARE USER_AND_LOCATION PURCHASING EXTENSION_ATTRIBUTES

inputType
string
enum
required
Defaults to TEXT
Extension attributes collect inventory data by using an input type.The type of the Input used to populate the extension attribute.

TEXT POPUP DIRECTORY_SERVICE_ATTRIBUTE_MAPPING

popupMenuChoices
array of strings
length ≥ 0
When added with list of choices while creating mobile device extension attributes these Pop-up menu can be displayed in inventory information. User can choose a value from the pop-up menu list when enrolling a mobile device any time using Jamf Pro. Provide popupMenuChoices only when inputType is 'POPUP'.

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

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/mobile-device-extension-attributes/ \
     --header 'accept: application/json'

{
  "id": "1",
  "name": "MobileDeviceExtensionAttribute",
  "description": "Mobile Device Extension Attribute",
  "dataType": "STRING",
  "inventoryDisplayType": "GENERAL",
  "inputType": "TEXT",
  "popupMenuChoices": [
    "Test",
    "Popup"
  ],
  "ldapAttributeMapping": "ldapAttributeMapping",
  "ldapExtensionAttributeAllowed": false
}
-----
Create Mobile Device Extension Attribute.
post
https://yourServer.jamfcloud.com/api/v1/mobile-device-extension-attributes

Create Mobile Device Extension Attribute to collect extra inventory information.

Body Params
Mobile Device Extension Attribute to be created.

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

TEXT

POPUP

DIRECTORY_SERVICE_ATTRIBUTE_MAPPING
popupMenuChoices
array of strings
length ≥ 0
When added with list of choices while creating mobile device extension attributes these Pop-up menu can be displayed in inventory information. User can choose a value from the pop-up menu list when enrolling a mobile device any time using Jamf Pro. Provide popupMenuChoices only when inputType is 'POPUP'.


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
Responses

201
Mobile Device Extension Attribute was successfully created.

Response body
object
id
string
length ≥ 1
href
string

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/mobile-device-extension-attributes \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "dataType": "STRING",
  "inventoryDisplayType": "GENERAL",
  "inputType": "TEXT",
  "ldapExtensionAttributeAllowed": false,
  "name": "MobileDeviceExtensionAttribute",
  "description": "Mobile Device Extension Attribute",
  "ldapAttributeMapping": "ldapAttributeMapping"
}
'

{
  "id": "1",
  "href": "https://yourJamfProUrl.jamf/api/v1/resource/1"
}
-----
Get smart group dependent object for a specified mobile device extension attribute
get
https://yourServer.jamfcloud.com/api/v1/mobile-device-extension-attributes/{id}/data-dependency

Get smart group dependent object for a specified mobile device extension attribute

Path Params
id
string
required
Unique ID of mobile device extension attribute.

1
Responses

200
Fetches list of dependent objects for a specified mobile device extension attribute.

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
     --url https://yourserver.jamfcloud.com/api/v1/mobile-device-extension-attributes/1/data-dependency \
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

Get specified Mobile Device Extension Attribute History object
get
https://yourServer.jamfcloud.com/api/v1/mobile-device-extension-attributes/{id}/history

Get specified Mobile Device Extension Attribute history object

Path Params
id
string
required
Instance ID of Mobile Device Extension Attribute

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
Details of Mobile Device Extension Attribute history were found.

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
     --url 'https://yourserver.jamfcloud.com/api/v1/mobile-device-extension-attributes//history?page=0&page-size=100&sort=id%3Aasc' \
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
Update specified Mobile Device Extension Attribute object.
put
https://yourServer.jamfcloud.com/api/v1/mobile-device-extension-attributes/{id}

Update specified Mobile Device Extension Attribute object.

Path Params
id
string
required
Unique ID of Mobile Device Extension Attribute.

Body Params
Mobile Device Extension Attribute object to be updated. IDs defined in this body will be ignored.

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

TEXT

POPUP

DIRECTORY_SERVICE_ATTRIBUTE_MAPPING
popupMenuChoices
array of strings
length ≥ 0
When added with list of choices while creating mobile device extension attributes these Pop-up menu can be displayed in inventory information. User can choose a value from the pop-up menu list when enrolling a mobile device any time using Jamf Pro. Provide popupMenuChoices only when inputType is 'POPUP'.


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

curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v1/mobile-device-extension-attributes/ \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "dataType": "STRING",
  "inventoryDisplayType": "GENERAL",
  "inputType": "TEXT",
  "ldapExtensionAttributeAllowed": false,
  "name": "MobileDeviceExtensionAttribute",
  "description": "Mobile Device Extension Attribute"
}
'

{
  "id": "1",
  "name": "MobileDeviceExtensionAttribute",
  "description": "Mobile Device Extension Attribute",
  "dataType": "STRING",
  "inventoryDisplayType": "GENERAL",
  "inputType": "TEXT",
  "popupMenuChoices": [
    "Test",
    "Popup"
  ],
  "ldapAttributeMapping": "ldapAttributeMapping",
  "ldapExtensionAttributeAllowed": false
}
-----
Add specified Mobile Device Extension Attribute history object notes
post
https://yourServer.jamfcloud.com/api/v1/mobile-device-extension-attributes/{id}/history

Add specified Mobile Device Extension Attribute history object notes

Path Params
id
string
required
Instance ID of Mobile Device Extension Attribute

1
Body Params
History note to be created

note
string
required
A generic note can sometimes be useful, but generally not.
Responses

201
Mobile Device Extension Attribute history note created successfully.

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
     --url https://yourserver.jamfcloud.com/api/v1/mobile-device-extension-attributes/1/history \
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
Delete a Mobile Device Extension Attribute by ID.
delete
https://yourServer.jamfcloud.com/api/v1/mobile-device-extension-attributes/{id}

Deletes the Mobile Device Extension Attribute identified by the provided ID.<\br> In addition to removing the attribute itself, this operation will also delete any related dependent data, including:<\br>

Associated popup menu choices.<\br>
Fields used in saved search displays.<\br>
User preferences that reference the attribute.
Path Params
id
string
required
Unique ID of Mobile Device Extension Attribute.

1
Responses
204
The Mobile Device Extension was successfully deleted.

curl --request DELETE \
     --url https://yourserver.jamfcloud.com/api/v1/mobile-device-extension-attributes/1
-----