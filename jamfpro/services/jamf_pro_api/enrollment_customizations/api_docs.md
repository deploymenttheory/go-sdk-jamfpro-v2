Retrieve sorted and paged Enrollment Customizations
get
https://yourServer.jamfcloud.com/api/v2/enrollment-customizations

Retrieves sorted and paged Enrollment Customizations

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
Sorting criteria in the format: property:asc/desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc


string

id:asc

ADD string
Response

200
Success

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
siteId
string
required
displayName
string
required
description
string
required
enrollmentCustomizationBrandingSettings
object
required

enrollmentCustomizationBrandingSettings object
textColor
string
required
buttonColor
string
required
buttonTextColor
string
required
backgroundColor
string
required
iconUrl
string
required

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v2/enrollment-customizations?page=0&page-size=100&sort=id%3Aasc' \
     --header 'accept: application/json'

{
  "totalCount": 1,
  "results": [
    {
      "id": "1",
      "siteId": "2",
      "displayName": "Example",
      "description": "Example description",
      "enrollmentCustomizationBrandingSettings": {
        "textColor": "0000FF",
        "buttonColor": "0000FF",
        "buttonTextColor": "0000FF",
        "backgroundColor": "0000FF",
        "iconUrl": "https://jamfUrl/api/v2/enrollment-customizations/images/1"
      }
    }
  ]
}
-----
Create an Enrollment Customization
post
https://yourServer.jamfcloud.com/api/v2/enrollment-customizations

Create an enrollment customization

Body Params
Enrollment customization to create.

siteId
string
required
2
displayName
string
required
Example
description
string
required
Example description
enrollmentCustomizationBrandingSettings
object
required

enrollmentCustomizationBrandingSettings object
textColor
string
required
buttonColor
string
required
buttonTextColor
string
required
backgroundColor
string
required
iconUrl
string
required
Response

201
Enrollment customization was created

Response body
object
id
string
length ≥ 1
href
string

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v2/enrollment-customizations \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "siteId": "2",
  "displayName": "Example",
  "description": "Example description"
}
'
{
  "id": "1",
  "href": "https://yourJamfProUrl.jamf/api/v1/resource/1"
}
-----
Upload an image
post
https://yourServer.jamfcloud.com/api/v2/enrollment-customizations/images

Uploads an image

Body Params
file
file
required
The file to upload

No file chosen
Response

201
Image successfully uploaded

Response body
object
url
string
-----
curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v2/enrollment-customizations/images \
     --header 'accept: application/json' \
     --header 'content-type: multipart/form-data'

{
  "url": "https://jamfpro.jamf/image?1"
}
-----
Download an enrollment customization image
get
https://yourServer.jamfcloud.com/api/v2/enrollment-customizations/images/{id}

Download an enrollment customization image

Path Params
id
string
required
id of the enrollment customization image

Response

200
Successful response

Response body
file

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v2/enrollment-customizations/images/ \
     --header 'accept: image/*'
-----
Retrieve an Enrollment Customization with the supplied id
get
https://yourServer.jamfcloud.com/api/v2/enrollment-customizations/{id}

Retrieves an Enrollment Customization with the supplied id

Path Params
id
string
required
Enrollment Customization identifier

Responses

200
Success

Response body
object
id
string
siteId
string
required
displayName
string
required
description
string
required
enrollmentCustomizationBrandingSettings
object
required
textColor
string
required
buttonColor
string
required
buttonTextColor
string
required
backgroundColor
string
required
iconUrl
string
required

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v2/enrollment-customizations/ \
     --header 'accept: application/json'

{
  "id": "1",
  "siteId": "2",
  "displayName": "Example",
  "description": "Example description",
  "enrollmentCustomizationBrandingSettings": {
    "textColor": "0000FF",
    "buttonColor": "0000FF",
    "buttonTextColor": "0000FF",
    "backgroundColor": "0000FF",
    "iconUrl": "https://jamfUrl/api/v2/enrollment-customizations/images/1"
  }
}
-----
Update an Enrollment Customization
put
https://yourServer.jamfcloud.com/api/v2/enrollment-customizations/{id}

Updates an Enrollment Customization

Path Params
id
string
required
Enrollment Customization identifier

5
Body Params
Enrollment Customization to update

siteId
string
required
2
displayName
string
required
Example
description
string
required
Example description
enrollmentCustomizationBrandingSettings
object
required

enrollmentCustomizationBrandingSettings object
textColor
string
required
0000FF
buttonColor
string
required
0000FF
buttonTextColor
string
required
0000FF
backgroundColor
string
required
0000FF
iconUrl
string
required
https://jamfUrl/api/v2/enrollment-customizations/images/1
Responses

200
Success

Response body
object
id
string
siteId
string
required
displayName
string
required
description
string
required
enrollmentCustomizationBrandingSettings
object
required
textColor
string
required
buttonColor
string
required
buttonTextColor
string
required
backgroundColor
string
required
iconUrl
string
required

curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v2/enrollment-customizations/5 \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "enrollmentCustomizationBrandingSettings": {
    "iconUrl": "https://jamfUrl/api/v2/enrollment-customizations/images/1",
    "backgroundColor": "0000FF",
    "buttonTextColor": "0000FF",
    "buttonColor": "0000FF",
    "textColor": "0000FF"
  },
  "siteId": "2",
  "displayName": "Example",
  "description": "Example description"
}
'

{
  "id": "1",
  "siteId": "2",
  "displayName": "Example",
  "description": "Example description",
  "enrollmentCustomizationBrandingSettings": {
    "textColor": "0000FF",
    "buttonColor": "0000FF",
    "buttonTextColor": "0000FF",
    "backgroundColor": "0000FF",
    "iconUrl": "https://jamfUrl/api/v2/enrollment-customizations/images/1"
  }
}
-----
Delete an Enrollment Customization with the supplied id
delete
https://yourServer.jamfcloud.com/api/v2/enrollment-customizations/{id}

Deletes an Enrollment Customization with the supplied id

Path Params
id
string
required
Enrollment Customization identifier

Response
204
Success

curl --request DELETE \
     --url https://yourserver.jamfcloud.com/api/v2/enrollment-customizations/

-----
Get sorted and paged Enrollment Customization history objects
get
https://yourServer.jamfcloud.com/api/v2/enrollment-customizations/{id}/history

Gets sorted and paged enrollment customization history objects

Path Params
id
string
required
Enrollment Customization identifier

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
Sorting criteria in the format: property,asc/desc. Default sort order is descending. Multiple sort criteria are supported and must be entered on separate lines in Swagger UI. In the URI the 'sort' query param is duplicated for each sort criterion, e.g., ...&sort=name%2Casc&sort=date%2Cdesc


string

date:desc

string


ADD string
Response

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v2/enrollment-customizations//history?page=0&page-size=100&sort=date%3Adesc' \
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
Add Enrollment Customization history object notes
post
https://yourServer.jamfcloud.com/api/v2/enrollment-customizations/{id}/history

Adds enrollment customization history object notes

Path Params
id
string
required
Enrollment Customization identifier

Body Params
History notes to create

note
string
required
Responses

201
Notes of enrollment customization history were added

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
     --url https://yourserver.jamfcloud.com/api/v2/enrollment-customizations//history \
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
Retrieve the list of Prestages using this Enrollment Customization
get
https://yourServer.jamfcloud.com/api/v2/enrollment-customizations/{id}/prestages

Retrieves the list of Prestages using this Enrollment Customization

Path Params
id
string
required
Enrollment Customization identifier

Responses

200
Success

Response body
object
dependencies
array of objects
object
name
string
humanReadableName
string
hyperlink
string

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v2/enrollment-customizations//prestages \
     --header 'accept: application/json'

{
  "dependencies": [
    {
      "name": "Name",
      "humanReadableName": "Computer PreStage",
      "hyperlink": "/mobile-prestage/id"
    }
  ]
}
-----
