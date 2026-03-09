Create one or more new Inventory Preload records using CSV
post
https://yourServer.jamfcloud.com/api/v2/inventory-preload/csv


Create one or more new Inventory Preload records using CSV. A CSV template can be downloaded from /v2/inventory-preload/csv-template. Serial number and device type are required. All other fields are optional. When a matching serial number exists in the Inventory Preload data, the record will be overwritten with the CSV data. If the CSV file contains a new username and an email address is provided, the new user is created in Jamf Pro. If the CSV file contains an existing username, the following user-related fields are updated in Jamf Pro. Full Name, Email Address, Phone Number, Position. This endpoint does not do full validation of each record in the CSV data. To do full validation, use the /v2/inventory-preload/csv-validate endpoint first.

Body Params
file
string
required
The CSV file to upload

serial number,device type,full nameserial123,Computer,Test Name
Responses

201
Created

Response body
array of objects
object
id
string
length ≥ 1
href
string

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v2/inventory-preload/csv \
     --header 'accept: application/json' \
     --header 'content-type: multipart/form-data' \
     --form 'file=serial number,device type,full name
serial123,Computer,Test Name
'

[
  {
    "id": "1",
    "href": "https://yourJamfProUrl.jamf/api/v1/resource/1"
  }
]
-----
Download the Inventory Preload CSV template
get
https://yourServer.jamfcloud.com/api/v2/inventory-preload/csv-template


Retrieves the Inventory Preload CSV file template.

Response

200
OK

Response body
string
Headers
object
Content-Disposition
string
A header field named Content-Disposition is returned with the file name contained in its value, which is always inventory-preload-template.csv.

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v2/inventory-preload/csv-template \
     --header 'accept: text/csv'

Serial Number,Device Type,Username,Full Name,Email Address
-----
Validate a given CSV file
post
https://yourServer.jamfcloud.com/api/v2/inventory-preload/csv-validate


Validate a given CSV file. Serial number and device type are required. All other fields are optional. A CSV template can be downloaded from /v2/inventory-preload/csv-template.

Body Params
Inventory Preload records to be validated. A CSV template can be downloaded from /v2/inventory-preload/csv-template.

file
string
required
The CSV file to upload

Responses

200
Ok

Response body
object
recordCount
integer

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v2/inventory-preload/csv-validate \
     --header 'accept: application/json' \
     --header 'content-type: multipart/form-data'

{
  "recordCount": 10
}
-----
Retrieve a list of extension attribute columns
get
https://yourServer.jamfcloud.com/api/v2/inventory-preload/ea-columns


Retrieve a list of extension attribute columns currently associated with inventory preload records

Response

200
Ok

Response body
object
totalCount
integer
≥ 0
results
array of objects
object
name
string
fullName
string

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v2/inventory-preload/ea-columns \
     --header 'accept: application/json'

{
  "totalCount": 10,
  "results": [
    {
      "name": "eaColumn1",
      "fullName": "Column 1"
    }
  ]
}
-----
Export a collection of inventory preload records
post
https://yourServer.jamfcloud.com/api/v2/inventory-preload/export


Export a collection of inventory preload records

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
Defaults to id:asc
Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. All inventory preload fields are supported, however fields added by extension attributes are not supported. If sorting by deviceType, use 0 for Computer and 1 for Mobile Device.

Example: sort=date:desc,name:asc.


string

id:asc

ADD string
filter
string
Allowing to filter inventory preload records. Default search is empty query - returning all results for the requested page. All inventory preload fields are supported, however fields added by extension attributes are not supported. If filtering by deviceType, use 0 for Computer and 1 for Mobile Device.

Query in the RSQL format, allowing ==, !=, >, <, and =in=.

Example: filter=categoryName=="Category"

Body Params
Optional. Override query parameters since they can make URI exceed 2,000 character limit.

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
     --url 'https://yourserver.jamfcloud.com/api/v2/inventory-preload/export?page=0&page-size=100&sort=id%3Aasc' \
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

Column 1,Column 2,Column 3
serial123,data,user
serial1234,additional data,user last
-----
Get Inventory Preload history entries
get
https://yourServer.jamfcloud.com/api/v2/inventory-preload/history


Gets Inventory Preload history entries.

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
Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma.

Example: sort=date:desc,name:asc.


string

date:desc

ADD string
filter
string
Allows filtering inventory preload history records. Default search is empty query - returning all results for the requested page. All inventory preload history fields are supported.

Query in the RSQL format, allowing ==, !=, >, <, and =in=.

Example: filter=username=="admin"

Response

200
OK

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
     --url 'https://yourserver.jamfcloud.com/api/v2/inventory-preload/history?page=0&page-size=100&sort=date%3Adesc' \
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
Add Inventory Preload history object notes
post
https://yourServer.jamfcloud.com/api/v2/inventory-preload/history


Adds Inventory Preload history object notes.

Body Params
History notes to create

note
string
required
Responses

201
Notes of Inventory Preload history were added

Response body
object
id
string
length ≥ 1
href
string

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v2/inventory-preload/history \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "note": "A generic note can sometimes be useful, but generally not."
}
'

{
  "id": "1",
  "href": "https://yourJamfProUrl.jamf/api/v1/resource/1"
}
-----
Return all Inventory Preload records
get
https://yourServer.jamfcloud.com/api/v2/inventory-preload/records


Returns all Inventory Preload records.

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
Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. All inventory preload fields are supported, however fields added by extension attributes are not supported. If sorting by deviceType, use 0 for Computer and 1 for Mobile Device.

Example: sort=date:desc,name:asc.


string

id:asc

ADD string
filter
string
Allowing to filter inventory preload records. Default search is empty query - returning all results for the requested page. All inventory preload fields are supported, however fields added by extension attributes are not supported. If filtering by deviceType, use 0 for Computer and 1 for Mobile Device.

Query in the RSQL format, allowing ==, !=, >, <, and =in=.

Example: filter=categoryName=="Category"

Response

200
OK

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
serialNumber
string
required
deviceType
string
enum
required
Computer Mobile Device Unknown

username
string | null
fullName
string | null
emailAddress
string | null
phoneNumber
string | null
position
string | null
department
string | null
building
string | null
room
string | null
poNumber
string | null
poDate
string | null
warrantyExpiration
string | null
appleCareId
string | null
lifeExpectancy
string | null
purchasePrice
string | null
purchasingContact
string | null
purchasingAccount
string | null
leaseExpiration
string | null
barCode1
string | null
barCode2
string | null
assetTag
string | null
vendor
string | null
extensionAttributes
array of objects
object
name
string
required
value
string

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v2/inventory-preload/records?page=0&page-size=100&sort=id%3Aasc' \
     --header 'accept: application/json'

{
  "totalCount": 10,
  "results": [
    {
      "id": "1",
      "serialNumber": "C02L29ECF8J1",
      "deviceType": "Computer",
      "username": "admin",
      "fullName": "Name",
      "emailAddress": "ITBob@jamf.com",
      "phoneNumber": "555-555-5555",
      "position": "IT Team Lead",
      "department": "IT",
      "building": "Eau Claire",
      "room": "4th Floor - Quad 3",
      "poNumber": "8675309",
      "poDate": "2019-02-04T21:09:31.661Z",
      "warrantyExpiration": "2012-07-21",
      "appleCareId": "5678",
      "lifeExpectancy": "5 years",
      "purchasePrice": "$399",
      "purchasingContact": "Nick in IT",
      "purchasingAccount": "IT Budget",
      "leaseExpiration": "2015-06-19",
      "barCode1": "123456789",
      "barCode2": "123456789",
      "assetTag": "ABCDEFG12345",
      "vendor": "Apple",
      "extensionAttributes": [
        {
          "name": "foo",
          "value": "42"
        }
      ]
    }
  ]
}
-----
Create a new Inventory Preload record using JSON
post
https://yourServer.jamfcloud.com/api/v2/inventory-preload/records


Create a new Inventory Preload record using JSON.

Body Params
Inventory Preload record to be created.

serialNumber
string
required
C02L29ECF8J1
deviceType
string
enum
required

Mobile Device
Allowed:

Computer

Mobile Device

Unknown
username
string | null
admin
fullName
string | null
Name
emailAddress
string | null
ITBob@jamf.com
phoneNumber
string | null
555-555-5555
position
string | null
IT Team Lead
department
string | null
IT
building
string | null
Eau Claire
room
string | null
4th Floor - Quad 3
poNumber
string | null
8675309
poDate
string | null
2019-02-04T21:09:31.661Z
warrantyExpiration
string | null
2012-07-21
appleCareId
string | null
5678
lifeExpectancy
string | null
5 years
purchasePrice
string | null
$399
purchasingContact
string | null
Nick in IT
purchasingAccount
string | null
IT Budget
leaseExpiration
string | null
2015-06-19
barCode1
string | null
123456789
barCode2
string | null
123456789
assetTag
string | null
ABCDEFG12345
vendor
string | null
Apple
extensionAttributes
array of objects

object

name
string
required
foo
value
string
42

ADD object
Responses

201
Created

Response body
object
id
string
length ≥ 1
href
string

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v2/inventory-preload/records \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "deviceType": "Mobile Device",
  "serialNumber": "C02L29ECF8J1",
  "username": "admin",
  "fullName": "Name",
  "emailAddress": "ITBob@jamf.com",
  "phoneNumber": "555-555-5555",
  "position": "IT Team Lead",
  "department": "IT",
  "building": "Eau Claire",
  "room": "4th Floor - Quad 3",
  "poNumber": "8675309",
  "poDate": "2019-02-04T21:09:31.661Z",
  "warrantyExpiration": "2012-07-21",
  "appleCareId": "5678",
  "lifeExpectancy": "5 years",
  "purchasePrice": "$399",
  "purchasingContact": "Nick in IT",
  "purchasingAccount": "IT Budget",
  "leaseExpiration": "2015-06-19",
  "barCode1": "123456789",
  "barCode2": "123456789",
  "assetTag": "ABCDEFG12345",
  "vendor": "Apple",
  "extensionAttributes": [
    {
      "name": "foo",
      "value": "42"
    }
  ]
}
'

{
  "id": "1",
  "href": "https://yourJamfProUrl.jamf/api/v1/resource/1"
}
-----
Delete all Inventory Preload records
post
https://yourServer.jamfcloud.com/api/v2/inventory-preload/records/delete-all


Deletes all Inventory Preload records.

Response
204
OK

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v2/inventory-preload/records/delete-all

-----
Get an Inventory Preload record
get
https://yourServer.jamfcloud.com/api/v2/inventory-preload/records/{id}


Retrieves an Inventory Preload record.

Path Params
id
string
required
Inventory Preload identifier

Responses

200
OK

Response body
object
id
string
serialNumber
string
required
deviceType
string
enum
required
Computer Mobile Device Unknown

username
string | null
fullName
string | null
emailAddress
string | null
phoneNumber
string | null
position
string | null
department
string | null
building
string | null
room
string | null
poNumber
string | null
poDate
string | null
warrantyExpiration
string | null
appleCareId
string | null
lifeExpectancy
string | null
purchasePrice
string | null
purchasingContact
string | null
purchasingAccount
string | null
leaseExpiration
string | null
barCode1
string | null
barCode2
string | null
assetTag
string | null
vendor
string | null
extensionAttributes
array of objects
object
name
string
required
value
string

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v2/inventory-preload/records/ \
     --header 'accept: application/json'

{
  "id": "1",
  "serialNumber": "C02L29ECF8J1",
  "deviceType": "Computer",
  "username": "admin",
  "fullName": "Name",
  "emailAddress": "ITBob@jamf.com",
  "phoneNumber": "555-555-5555",
  "position": "IT Team Lead",
  "department": "IT",
  "building": "Eau Claire",
  "room": "4th Floor - Quad 3",
  "poNumber": "8675309",
  "poDate": "2019-02-04T21:09:31.661Z",
  "warrantyExpiration": "2012-07-21",
  "appleCareId": "5678",
  "lifeExpectancy": "5 years",
  "purchasePrice": "$399",
  "purchasingContact": "Nick in IT",
  "purchasingAccount": "IT Budget",
  "leaseExpiration": "2015-06-19",
  "barCode1": "123456789",
  "barCode2": "123456789",
  "assetTag": "ABCDEFG12345",
  "vendor": "Apple",
  "extensionAttributes": [
    {
      "name": "foo",
      "value": "42"
    }
  ]
}
-----
Update an Inventory Preload record
put
https://yourServer.jamfcloud.com/api/v2/inventory-preload/records/{id}


Updates an Inventory Preload record.

Path Params
id
string
required
Inventory Preload identifier

Body Params
Inventory Preload record to update

serialNumber
string
required
C02L29ECF8J1
deviceType
string
enum
required

Mobile Device
Allowed:

Computer

Mobile Device

Unknown
username
string | null
admin
fullName
string | null
Name
emailAddress
string | null
ITBob@jamf.com
phoneNumber
string | null
555-555-5555
position
string | null
IT Team Lead
department
string | null
IT
building
string | null
Eau Claire
room
string | null
4th Floor - Quad 3
poNumber
string | null
8675309
poDate
string | null
2019-02-04T21:09:31.661Z
warrantyExpiration
string | null
2012-07-21
appleCareId
string | null
5678
lifeExpectancy
string | null
5 years
purchasePrice
string | null
$399
purchasingContact
string | null
Nick in IT
purchasingAccount
string | null
IT Budget
leaseExpiration
string | null
2015-06-19
barCode1
string | null
123456789
barCode2
string | null
123456789
assetTag
string | null
ABCDEFG12345
vendor
string | null
Apple
extensionAttributes
array of objects

object

name
string
required
foo
value
string
42

object

name
string
required
foo
value
string
42

ADD object
Responses

200
OK

Response body
object
id
string
serialNumber
string
required
deviceType
string
enum
required
Computer Mobile Device Unknown

username
string | null
fullName
string | null
emailAddress
string | null
phoneNumber
string | null
position
string | null
department
string | null
building
string | null
room
string | null
poNumber
string | null
poDate
string | null
warrantyExpiration
string | null
appleCareId
string | null
lifeExpectancy
string | null
purchasePrice
string | null
purchasingContact
string | null
purchasingAccount
string | null
leaseExpiration
string | null
barCode1
string | null
barCode2
string | null
assetTag
string | null
vendor
string | null
extensionAttributes
array of objects
object
name
string
required
value
string

curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v2/inventory-preload/records/ \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "deviceType": "Mobile Device",
  "serialNumber": "C02L29ECF8J1",
  "username": "admin",
  "fullName": "Name",
  "emailAddress": "ITBob@jamf.com",
  "phoneNumber": "555-555-5555",
  "position": "IT Team Lead",
  "department": "IT",
  "building": "Eau Claire",
  "room": "4th Floor - Quad 3",
  "poNumber": "8675309",
  "poDate": "2019-02-04T21:09:31.661Z",
  "warrantyExpiration": "2012-07-21",
  "appleCareId": "5678",
  "lifeExpectancy": "5 years",
  "purchasePrice": "$399",
  "purchasingContact": "Nick in IT",
  "purchasingAccount": "IT Budget",
  "leaseExpiration": "2015-06-19",
  "barCode1": "123456789",
  "barCode2": "123456789",
  "assetTag": "ABCDEFG12345",
  "vendor": "Apple",
  "extensionAttributes": [
    {
      "name": "foo",
      "value": "42"
    },
    {
      "name": "foo",
      "value": "42"
    }
  ]
}
'

{
  "id": "1",
  "serialNumber": "C02L29ECF8J1",
  "deviceType": "Computer",
  "username": "admin",
  "fullName": "Name",
  "emailAddress": "ITBob@jamf.com",
  "phoneNumber": "555-555-5555",
  "position": "IT Team Lead",
  "department": "IT",
  "building": "Eau Claire",
  "room": "4th Floor - Quad 3",
  "poNumber": "8675309",
  "poDate": "2019-02-04T21:09:31.661Z",
  "warrantyExpiration": "2012-07-21",
  "appleCareId": "5678",
  "lifeExpectancy": "5 years",
  "purchasePrice": "$399",
  "purchasingContact": "Nick in IT",
  "purchasingAccount": "IT Budget",
  "leaseExpiration": "2015-06-19",
  "barCode1": "123456789",
  "barCode2": "123456789",
  "assetTag": "ABCDEFG12345",
  "vendor": "Apple",
  "extensionAttributes": [
    {
      "name": "foo",
      "value": "42"
    }
  ]
}
-----
Delete an Inventory Preload record
delete
https://yourServer.jamfcloud.com/api/v2/inventory-preload/records/{id}


Deletes an Inventory Preload record.

Path Params
id
string
required
Inventory Preload identifier

Responses
204
OK

curl --request DELETE \
     --url https://yourserver.jamfcloud.com/api/v2/inventory-preload/records/ \
     --header 'accept: application/json'

