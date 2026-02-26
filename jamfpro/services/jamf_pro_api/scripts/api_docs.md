Search for sorted and paged Scripts
get
https://yourServer.jamfcloud.com/api/v1/scripts

Search for sorted and paged scripts

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
Sorting criteria in the format: property:asc/desc. Default sort is name:asc. Multiple sort criteria are supported and must be separated with a comma. Fields allowed in the query: id, name, info, notes, priority, categoryId, categoryName, parameter4 up to parameter11, osRequirements, scriptContents. Example: sort=date:desc,name:asc


string

name:asc

ADD string
filter
string
Query in the RSQL format, allowing to filter scripts collection. Default search is empty query - returning all results for the requested page. Fields allowed in the query: id, name, info, notes, priority, categoryId, categoryName, parameter4 up to parameter11, osRequirements, scriptContents. This param can be combined with paging and sorting. Example: filter=categoryName=="Category" and name=="script name"

Response

200
Found scripts matching search params.

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
name
string
required
info
string
notes
string
priority
string
enum
BEFORE AFTER AT_REBOOT

categoryId
string
categoryName
string
parameter4
string
parameter5
string
parameter6
string
parameter7
string
parameter8
string
parameter9
string
parameter10
string
parameter11
string
osRequirements
string
scriptContents
string

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v1/scripts?page=0&page-size=100&sort=name%3Aasc' \
     --header 'accept: application/json'

{
  "totalCount": 3,
  "results": [
    {
      "id": "1",
      "name": "Install Developer Utils Script",
      "info": "Installs utilities for developers",
      "notes": "Should be able to be re-run without problem.",
      "priority": "AFTER",
      "categoryId": "1",
      "categoryName": "Developer Tools",
      "parameter4": "1",
      "parameter5": "2",
      "parameter6": "3",
      "parameter7": "4",
      "parameter8": "5",
      "parameter9": "6",
      "parameter10": "7",
      "parameter11": "8",
      "osRequirements": "10.10.x",
      "scriptContents": "echo \"Trivial script.\""
    }
  ]
}
-----
Create a Script
post
https://yourServer.jamfcloud.com/api/v1/scripts

Creates a script

Body Params
new script to create. ids defined in this body will be ignored

name
string
required
Install Developer Utils Script
info
string
Installs utilities for developers
notes
string
Should be able to be re-run without problem.
priority
string
enum

BEFORE
Allowed:

BEFORE

AFTER

AT_REBOOT
categoryId
string
1
categoryName
string
Developer Tools
parameter4
string
1
parameter5
string
2
parameter6
string
3
parameter7
string
4
parameter8
string
5
parameter9
string
6
parameter10
string
7
parameter11
string
8
osRequirements
string
10.10.x
scriptContents
string
echo "Trivial script."
Response

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/scripts \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "name": "Install Developer Utils Script",
  "info": "Installs utilities for developers",
  "notes": "Should be able to be re-run without problem.",
  "priority": "BEFORE",
  "categoryId": "1",
  "categoryName": "Developer Tools",
  "parameter4": "1",
  "parameter5": "2",
  "parameter6": "3",
  "parameter7": "4",
  "parameter8": "5",
  "parameter9": "6",
  "parameter10": "7",
  "parameter11": "8",
  "osRequirements": "10.10.x",
  "scriptContents": "echo \"Trivial script.\""
}
'

{
  "id": "1",
  "href": "https://yourJamfProUrl.jamf/api/v1/resource/1"
}
-----
Retrieve a full script object
get
https://yourServer.jamfcloud.com/api/v1/scripts/{id}

Retrieves a full script object

Path Params
id
string
required
Script object identifier

Responses

200
Success

Response body
object
id
string
length ≥ 1
name
string
required
info
string
notes
string
priority
string
enum
BEFORE AFTER AT_REBOOT

categoryId
string
categoryName
string
parameter4
string
parameter5
string
parameter6
string
parameter7
string
parameter8
string
parameter9
string
parameter10
string
parameter11
string
osRequirements
string
scriptContents
string

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/scripts/ \
     --header 'accept: application/json'

{
  "id": "1",
  "name": "Install Developer Utils Script",
  "info": "Installs utilities for developers",
  "notes": "Should be able to be re-run without problem.",
  "priority": "AFTER",
  "categoryId": "1",
  "categoryName": "Developer Tools",
  "parameter4": "1",
  "parameter5": "2",
  "parameter6": "3",
  "parameter7": "4",
  "parameter8": "5",
  "parameter9": "6",
  "parameter10": "7",
  "parameter11": "8",
  "osRequirements": "10.10.x",
  "scriptContents": "echo \"Trivial script.\""
}
-----
Replace the script at the id with the supplied information
put
https://yourServer.jamfcloud.com/api/v1/scripts/{id}

Replaces the script at the id with the supplied information

Path Params
id
string
required
Script object identifier

1
Body Params
new script to upload to existing id. ids defined in this body will be ignored

name
string
required
Install Developer Utils Script
info
string
Installs utilities for developers
notes
string
Should be able to be re-run without problem.
priority
string
enum

BEFORE
Allowed:

BEFORE

AFTER

AT_REBOOT
categoryId
string
categoryName
string
parameter4
string
parameter5
string
parameter6
string
parameter7
string
parameter8
string
parameter9
string
parameter10
string
parameter11
string
osRequirements
string
scriptContents
string
Responses

200
Script at id was updated

Response body
object
id
string
length ≥ 1
name
string
required
info
string
notes
string
priority
string
enum
BEFORE AFTER AT_REBOOT

categoryId
string
categoryName
string
parameter4
string
parameter5
string
parameter6
string
parameter7
string
parameter8
string
parameter9
string
parameter10
string
parameter11
string
osRequirements
string
scriptContents
string

curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v1/scripts/1 \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "name": "Install Developer Utils Script",
  "info": "Installs utilities for developers",
  "notes": "Should be able to be re-run without problem.",
  "priority": "BEFORE"
}
'

{
  "id": "1",
  "name": "Install Developer Utils Script",
  "info": "Installs utilities for developers",
  "notes": "Should be able to be re-run without problem.",
  "priority": "AFTER",
  "categoryId": "1",
  "categoryName": "Developer Tools",
  "parameter4": "1",
  "parameter5": "2",
  "parameter6": "3",
  "parameter7": "4",
  "parameter8": "5",
  "parameter9": "6",
  "parameter10": "7",
  "parameter11": "8",
  "osRequirements": "10.10.x",
  "scriptContents": "echo \"Trivial script.\""
}
-----
Delete a Script at the specified id
delete
https://yourServer.jamfcloud.com/api/v1/scripts/{id}

Deletes a script at the specified id

Path Params
id
string
required
Script object identifier

Response
204
Successful response

curl --request DELETE \
     --url https://yourserver.jamfcloud.com/api/v1/scripts/
-----
Download a text file of the Script contents
get
https://yourServer.jamfcloud.com/api/v1/scripts/{id}/download

Download a text file of the script contents

Path Params
id
string
required
id of the script to be downloaded

Headers
accept
string
enum
Defaults to application/json
Generated from available response content types


text/plain
Allowed:

application/json

text/plain
Responses

200
The script of the specified id

Response body
file

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/scripts//download \
     --header 'accept: text/plain'

-----
Get specified Script history object
get
https://yourServer.jamfcloud.com/api/v1/scripts/{id}/history

Gets specified Script history object

Path Params
id
string
required
id of script history record

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
Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc


string

date:desc

ADD string
filter
string
Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: username, date, note, details. This param can be combined with paging and sorting. Example: filter=username!=admin and details==disabled and date<2019-12-15

Responses

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
     --url 'https://yourserver.jamfcloud.com/api/v1/scripts//history?page=0&page-size=100&sort=date%3Adesc' \
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
Add specified Script history object notes
post
https://yourServer.jamfcloud.com/api/v1/scripts/{id}/history

Adds specified Script history object notes

Path Params
id
string
required
instance id of script history record

Body Params
history notes to create

note
string
required
Responses

201
Notes of script history were added

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
     --url https://yourserver.jamfcloud.com/api/v1/scripts//history \
     --header 'accept: application/json' \
     --header 'content-type: application/json'

{
  "id": 1,
  "username": "admin",
  "date": "2019-02-04T21:09:31.661Z",
  "note": "Sso settings update",
  "details": "Is SSO Enabled false\\nSelected SSO Provider"
}