Finds all accounts
get
https://yourServer.jamfcloud.com/JSSResource/accounts

This resource corresponds with the Jamf Pro User Accounts and Groups feature. This endpoint is commonly mistaken with the users endpoint which allows for interaction with users associated with managed devices.

Privileges required to interact with this endpoint: Read - User accounts and groups

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
Response

200
OK

Response body

application/xml
object
users
object
user
array of objects
object
id
integer
name
string
required
Name of the account

groups
object
group
array of objects
object
id
integer
name
string
required
Name of the group

curl --request GET \
     --url https://yourserver.jamfcloud.com/JSSResource/accounts \
     --header 'accept: application/xml'

{
  "users": {
    "user": [
      {
        "id": 1,
        "name": "Steve Jobs"
      }
    ]
  },
  "groups": {
    "group": [
      {
        "id": 1,
        "name": "Information Technology"
      }
    ]
  }
}
-----
Finds groups by ID
get
https://yourServer.jamfcloud.com/JSSResource/accounts/groupid/{id}


Privileges required to interact with this endpoint: Read - User accounts and groups

Path Params
id
integer
required
ID value to filter by

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
Response

200
OK

Response body

application/xml
object
id
integer
name
string
required
Group name

access_level
string
enum
Full Access Site Access Group Access

privilege_set
string
enum
Administrator Auditor Enrollment Only Custom

site
object
id
integer
name
string
required
Name of the site

privileges
object
jss_objects
array of objects
object
privilege
string
jss_settings
array of objects
object
privilege
string
jss_actions
array of objects
object
privilege
string
recon
array of objects
object
privilege
string
casper_admin
array of objects
object
privilege
string
casper_remote
array of objects
object
privilege
string
casper_imaging
array of objects
object
privilege
string
members
array of objects
object
user
object

user object
id
integer
name
string

curl --request GET \
     --url https://yourserver.jamfcloud.com/JSSResource/accounts/groupid/id \
     --header 'accept: application/xml'

{
  "id": 1,
  "name": "Administrators",
  "access_level": "Full Access",
  "privilege_set": "Administrator",
  "site": {
    "id": 1,
    "name": "Minneapolis"
  },
  "privileges": {
    "jss_objects": [
      {
        "privilege": "string"
      }
    ],
    "jss_settings": [
      {
        "privilege": "string"
      }
    ],
    "jss_actions": [
      {
        "privilege": "string"
      }
    ],
    "recon": [
      {
        "privilege": "string"
      }
    ],
    "casper_admin": [
      {
        "privilege": "string"
      }
    ],
    "casper_remote": [
      {
        "privilege": "string"
      }
    ],
    "casper_imaging": [
      {
        "privilege": "string"
      }
    ]
  },
  "members": [
    {
      "user": {
        "id": 1,
        "name": "string"
      }
    }
  ]
}
-----
Updates an existing group by ID
put
https://yourServer.jamfcloud.com/JSSResource/accounts/groupid/{id}


Groups updated via this endpoint are updated only within Jamf Pro and will not affect the LDAP server.

Privileges required to interact with this endpoint: Update - User accounts and groups

Sample Request Body

<group>
	<id>1</id>
	<name>Administrators</name>
	<access_level>Full Access</access_level>
	<privilege_set>Administrator</privilege_set>
	<site>
		<id>-1</id>
		<name>None</name>
	</site>
	<privileges>
		<jss_objects>
			<privilege>string</privilege>
		</jss_objects>
		<jss_settings>
			<privilege>string</privilege>
		</jss_settings>
		<jss_actions>
			<privilege>string</privilege>
		</jss_actions>
		<recon>
			<privilege>string</privilege>
		</recon>
		<casper_admin>
			<privilege>string</privilege>
		</casper_admin>
		<casper_remote>
			<privilege>string</privilege>
		</casper_remote>
		<casper_imaging>
			<privilege>string</privilege>
		</casper_imaging>
	</privileges>
	<members>
		<user>
			<id>1</id>
			<name>string</name>
		</user>
	</members>
</group>
Path Params
id
integer
required
ID value to filter by

1
Response
201
Created

curl --request PUT \
     --url https://yourserver.jamfcloud.com/JSSResource/accounts/groupid/1
-----
Creates a new group by ID
post
https://yourServer.jamfcloud.com/JSSResource/accounts/groupid/{id}


Groups created via this endpoint are created only within Jamf Pro and will not affect the LDAP server.

Privileges required to interact with this endpoint: Create - User accounts and groups

Sample Request Body

<group>
	<id>1</id>
	<name>Administrators</name>
	<access_level>Full Access</access_level>
	<privilege_set>Administrator</privilege_set>
	<site>
		<id>-1</id>
		<name>None</name>
	</site>
	<privileges>
		<jss_objects>
			<privilege>string</privilege>
		</jss_objects>
		<jss_settings>
			<privilege>string</privilege>
		</jss_settings>
		<jss_actions>
			<privilege>string</privilege>
		</jss_actions>
		<recon>
			<privilege>string</privilege>
		</recon>
		<casper_admin>
			<privilege>string</privilege>
		</casper_admin>
		<casper_remote>
			<privilege>string</privilege>
		</casper_remote>
		<casper_imaging>
			<privilege>string</privilege>
		</casper_imaging>
	</privileges>
	<members>
		<user>
			<id>1</id>
			<name>string</name>
		</user>
	</members>
</group>
Path Params
id
integer
required
Defaults to 0
ID value to filter by

0
Response
201
Created

curl --request POST \
     --url https://yourserver.jamfcloud.com/JSSResource/accounts/groupid/0
-----
Deletes a group by ID
delete
https://yourServer.jamfcloud.com/JSSResource/accounts/groupid/{id}


Privileges required to interact with this endpoint: Delete - User accounts and groups

Path Params
id
integer
required
ID value to filter by

Response
200
OK

curl --request DELETE \
     --url https://yourserver.jamfcloud.com/JSSResource/accounts/groupid/id
-----
Finds groups by name
get
https://yourServer.jamfcloud.com/JSSResource/accounts/groupname/{name}


Privileges required to interact with this endpoint: Read - User accounts and groups

Path Params
name
string
required
Name to filter by

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
Response

curl --request GET \
     --url https://yourserver.jamfcloud.com/JSSResource/accounts/groupname/name \
     --header 'accept: application/xml'
-----
Updates an existing group by name
put
https://yourServer.jamfcloud.com/JSSResource/accounts/groupname/{name}


Groups updated via this endpoint are updated only within Jamf Pro and will not affect the LDAP server.

Privileges required to interact with this endpoint: Update - User accounts and groups

Sample Request Body

<group>
	<id>1</id>
	<name>Administrators</name>
	<access_level>Full Access</access_level>
	<privilege_set>Administrator</privilege_set>
	<site>
		<id>-1</id>
		<name>None</name>
	</site>
	<privileges>
		<jss_objects>
			<privilege>string</privilege>
		</jss_objects>
		<jss_settings>
			<privilege>string</privilege>
		</jss_settings>
		<jss_actions>
			<privilege>string</privilege>
		</jss_actions>
		<recon>
			<privilege>string</privilege>
		</recon>
		<casper_admin>
			<privilege>string</privilege>
		</casper_admin>
		<casper_remote>
			<privilege>string</privilege>
		</casper_remote>
		<casper_imaging>
			<privilege>string</privilege>
		</casper_imaging>
	</privileges>
	<members>
		<user>
			<id>1</id>
			<name>string</name>
		</user>
	</members>
</group>
Path Params
name
integer
required
Name value to filter by

curl --request PUT \
     --url https://yourserver.jamfcloud.com/JSSResource/accounts/groupname/name
-----

Deletes a group by name
delete
https://yourServer.jamfcloud.com/JSSResource/accounts/groupname/{name}


Privileges required to interact with this endpoint: Delete - User accounts and groups

Path Params
name
integer
required
Name value to filter by

Response
200
OK

curl --request DELETE \
     --url https://yourserver.jamfcloud.com/JSSResource/accounts/groupname/name