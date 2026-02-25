Create a DockItem
post
https://yourServer.jamfcloud.com/api/v1/dock-items

Creates a DockItem

Body Params
new DockItem to create. ids defined in this body will be ignored

name
string
required
DockItem Name
type
string
enum
required

APP
Allowed:

APP

FILE

FOLDER
path
string
required
file://localhost/Applications/iTunes.app
Responses

201
DockItem created successfully

Response body
object
id
string
length ≥ 1
href
string

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/dock-items \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "type": "APP",
  "path": "file://localhost/Applications/iTunes.app",
  "name": "DockItem Name"
}
'

{
  "id": "1",
  "href": "https://yourJamfProUrl.jamf/api/v1/resource/1"
}
-----
Retrieve a full dockItem object
get
https://yourServer.jamfcloud.com/api/v1/dock-items/{id}

Retrieves a full dockItem object

Path Params
id
string
required
DockItem object identifier

1
Responses

200
Success

Response body
object
id
string
≥ 1
name
string
required
type
string
enum
required
APP FILE FOLDER

path
string
required
contents
string

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/dock-items/1 \
     --header 'accept: application/json'

{
  "id": "1",
  "name": "DockItem Name",
  "type": "FILE",
  "path": "file://localhost/Applications/iTunes.app",
  "contents": "<dict><key>GUID</key><integer>-91117049</integer><key>tile-data</key><dict><key>file-data</key><dict><key>_CFURLString</key><string>file://localhost/Applications/iTunes.app</string><key>_CFURLStringType</key><integer>15</integer></dict><key>file-label</key><string>MyDockItem2</string></dict><key>tile-type</key><string>directory-tile</string></dict>"
}
-----
Replace the dockItem at the id with the supplied information
put
https://yourServer.jamfcloud.com/api/v1/dock-items/{id}

Replaces the dockItem at the id with the supplied information

Path Params
id
string
required
DockItem object identifier

1
Body Params
new dockItem to upload to existing id. ids defined in this body will be ignored

name
string
required
DockItem Name
type
string
enum
required

APP
Allowed:

APP

FILE

FOLDER
path
string
required
file://localhost/Applications/iTunes.app
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
DockItem at id was updated

Response body
string

curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v1/dock-items/1 \
     --header 'accept: text/plain' \
     --header 'content-type: application/json' \
     --data '
{
  "type": "APP",
  "name": "DockItem Name",
  "path": "file://localhost/Applications/iTunes.app"
}
'
-----
Delete a DockItem at the specified id
delete
https://yourServer.jamfcloud.com/api/v1/dock-items/{id}

Deletes a dockItem at the specified id

Path Params
id
string
required
DockItem object identifier

curl --request DELETE \
     --url https://yourserver.jamfcloud.com/api/v1/dock-items/ \
     --header 'accept: application/json'