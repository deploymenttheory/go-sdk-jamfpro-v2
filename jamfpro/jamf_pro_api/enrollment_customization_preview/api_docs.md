Parse the given string as markdown text and return Html output
post
https://yourServer.jamfcloud.com/api/v1/enrollment-customization/parse-markdown

Parse the given string as markdown text and return Html output

Body Params
Enrollment Customization Panel to create

markdown
string
Response

200
Success

Response body
object
markdown
string

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/enrollment-customization/parse-markdown \
     --header 'accept: application/json' \
     --header 'content-type: application/json'

{
  "markdown": "**markdown**"
}
-----
Get all Panels for single Enrollment Customization
get
https://yourServer.jamfcloud.com/api/v1/enrollment-customization/{id}/all

Get all panels for single enrollment customization

Path Params
id
integer
required
Enrollment Customization identifier

Responses

200
Success

Response body
object
panels
array of objects
object
displayName
string
required
rank
integer
required
id
integer
type
string

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/enrollment-customization//all \
     --header 'accept: application/json'

{
  "panels": [
    {
      "displayName": "A Panel",
      "rank": 0,
      "id": 2,
      "type": "text"
    }
  ]
}
-----
Get a single Panel for a single Enrollment Customization
get
https://yourServer.jamfcloud.com/api/v1/enrollment-customization/{id}/all/{panel-id}

Get a single panel for a single enrollment customization

Path Params
id
integer
required
Enrollment Customization identifier

panel-id
integer
required
Panel object identifier

Responses

200
Success

Response body
object
displayName
string
required
rank
integer
required
id
integer
type
string

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/enrollment-customization//all/ \
     --header 'accept: application/json'

{
  "displayName": "A Panel",
  "rank": 0,
  "id": 2,
  "type": "text"
}
-----
Delete a single Panel from an Enrollment Customization
delete
https://yourServer.jamfcloud.com/api/v1/enrollment-customization/{id}/all/{panel-id}

Delete a single panel from an Enrollment Customization

Path Params
id
integer
required
Enrollment Customization identifier

5
panel-id
integer
required
Panel object identifier

3

curl --request DELETE \
     --url https://yourserver.jamfcloud.com/api/v1/enrollment-customization/5/all/3
-----
Create an LDAP Panel for a single Enrollment Customization
post
https://yourServer.jamfcloud.com/api/v1/enrollment-customization/{id}/ldap

Create an LDAP panel for a single enrollment customization. If multiple LDAP access groups are defined with the same name and id, only one will be saved.

Path Params
id
integer
required
Enrollment Customization identifier

Body Params
Enrollment Customization Panel to create

displayName
string
required
A Panel
rank
integer
required
0
usernameLabel
string
required
Username
passwordLabel
string
required
Password
title
string
required
My Ldap Panel
backButtonText
string
required
Back
continueButtonText
string
required
Continue
ldapGroupAccess
array of objects

object

ldapServerId
integer
1
groupName
string
admins

ADD object
Responses

201
LDAP panel was created

Response body
object
displayName
string
required
rank
integer
required
usernameLabel
string
required
passwordLabel
string
required
title
string
required
backButtonText
string
required
continueButtonText
string
required
ldapGroupAccess
array of objects
object
ldapServerId
integer
groupName
string
id
integer
type
string

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/enrollment-customization//ldap \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "displayName": "A Panel",
  "rank": 0,
  "usernameLabel": "Username",
  "passwordLabel": "Password",
  "title": "My Ldap Panel",
  "backButtonText": "Back",
  "continueButtonText": "Continue",
  "ldapGroupAccess": [
    {
      "ldapServerId": 1,
      "groupName": "admins"
    }
  ]
}
'
{
  "displayName": "A Panel",
  "rank": 0,
  "usernameLabel": "Username",
  "passwordLabel": "Password",
  "title": "My Ldap Panel",
  "backButtonText": "Back",
  "continueButtonText": "Continue",
  "ldapGroupAccess": [
    {
      "ldapServerId": 1,
      "groupName": "admins"
    }
  ],
  "id": 2,
  "type": "ldap"
}
-----
Get a single LDAP panel for a single Enrollment Customization
get
https://yourServer.jamfcloud.com/api/v1/enrollment-customization/{id}/ldap/{panel-id}

Get a single LDAP panel for a single enrollment customization

Path Params
id
integer
required
Enrollment Customization identifier

panel-id
integer
required
Panel object identifier

Responses

200
Success

Response body
object
displayName
string
required
rank
integer
required
usernameLabel
string
required
passwordLabel
string
required
title
string
required
backButtonText
string
required
continueButtonText
string
required
ldapGroupAccess
array of objects
object
ldapServerId
integer
groupName
string
id
integer
type
string

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/enrollment-customization//ldap/ \
     --header 'accept: application/json'

{
  "displayName": "A Panel",
  "rank": 0,
  "usernameLabel": "Username",
  "passwordLabel": "Password",
  "title": "My Ldap Panel",
  "backButtonText": "Back",
  "continueButtonText": "Continue",
  "ldapGroupAccess": [
    {
      "ldapServerId": 1,
      "groupName": "admins"
    }
  ],
  "id": 2,
  "type": "ldap"
}
-----
Update a single LDAP Panel for a single Enrollment Customization
put
https://yourServer.jamfcloud.com/api/v1/enrollment-customization/{id}/ldap/{panel-id}

Update a single LDAP panel for a single enrollment customization. If multiple LDAP access groups are defined with the same name and id, only one will be saved.

Path Params
id
integer
required
Enrollment Customization identifier

panel-id
integer
required
Panel object identifier

Body Params
Enrollment Customization Panel to update

displayName
string
required
A Panel
rank
integer
required
0
usernameLabel
string
required
Username
passwordLabel
string
required
Password
title
string
required
My Ldap Panel
backButtonText
string
required
Back
continueButtonText
string
required
Continue
ldapGroupAccess
array of objects

object

ldapServerId
integer
1
groupName
string
admins

object

ldapServerId
integer
1
groupName
string
admins

ADD object
Responses

200
Success

Response body
object
displayName
string
required
rank
integer
required
usernameLabel
string
required
passwordLabel
string
required
title
string
required
backButtonText
string
required
continueButtonText
string
required
ldapGroupAccess
array of objects
object
ldapServerId
integer
groupName
string
id
integer
type
string

curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v1/enrollment-customization//ldap/ \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "displayName": "A Panel",
  "rank": 0,
  "usernameLabel": "Username",
  "passwordLabel": "Password",
  "title": "My Ldap Panel",
  "backButtonText": "Back",
  "continueButtonText": "Continue",
  "ldapGroupAccess": [
    {
      "ldapServerId": 1,
      "groupName": "admins"
    },
    {
      "ldapServerId": 1,
      "groupName": "admins"
    }
  ]
}
'
{
  "displayName": "A Panel",
  "rank": 0,
  "usernameLabel": "Username",
  "passwordLabel": "Password",
  "title": "My Ldap Panel",
  "backButtonText": "Back",
  "continueButtonText": "Continue",
  "ldapGroupAccess": [
    {
      "ldapServerId": 1,
      "groupName": "admins"
    }
  ],
  "id": 2,
  "type": "ldap"
}
-----
Delete an LDAP single panel from an Enrollment Customization
delete
https://yourServer.jamfcloud.com/api/v1/enrollment-customization/{id}/ldap/{panel-id}

Delete an LDAP single Panel from an Enrollment Customization

Path Params
id
integer
required
Enrollment Customization identifier

panel-id
integer
required
Panel object identifier

Response
204
Success

curl --request DELETE \
     --url https://yourserver.jamfcloud.com/api/v1/enrollment-customization//ldap/
-----
Create an SSO Panel for a single Enrollment Customization
post
https://yourServer.jamfcloud.com/api/v1/enrollment-customization/{id}/sso

Create an SSO panel for a single enrollment customization

Path Params
id
integer
required
Enrollment Customization identifier

Body Params
Enrollment Customization Panel to create

displayName
string
required
A Panel
rank
integer
required
0
isUseJamfConnect
boolean
required

true
longNameAttribute
string
required
long name
shortNameAttribute
string
required
name
isGroupEnrollmentAccessEnabled
boolean
required

true
groupEnrollmentAccessName
string
required
GroupNameA
Responses

201
Auth panel was created

Response body
object
displayName
string
required
rank
integer
required
isUseJamfConnect
boolean
required
longNameAttribute
string
required
shortNameAttribute
string
required
isGroupEnrollmentAccessEnabled
boolean
required
groupEnrollmentAccessName
string
required
id
integer
type
string

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/enrollment-customization//sso \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "isUseJamfConnect": true,
  "isGroupEnrollmentAccessEnabled": true,
  "displayName": "A Panel",
  "rank": 0,
  "longNameAttribute": "long name",
  "shortNameAttribute": "name",
  "groupEnrollmentAccessName": "GroupNameA"
}
'
{
  "displayName": "A Panel",
  "rank": 0,
  "isUseJamfConnect": false,
  "longNameAttribute": "long name",
  "shortNameAttribute": "name",
  "isGroupEnrollmentAccessEnabled": false,
  "groupEnrollmentAccessName": "GroupNameA",
  "id": 2,
  "type": "sso"
}
-----
Get a single SSO Panel for a single Enrollment Customization
get
https://yourServer.jamfcloud.com/api/v1/enrollment-customization/{id}/sso/{panel-id}

Get a single SSO panel for a single enrollment customization

Path Params
id
integer
required
Enrollment Customization identifier

panel-id
integer
required
Panel object identifier

Responses

200
Success

Response body
object
displayName
string
required
rank
integer
required
isUseJamfConnect
boolean
required
longNameAttribute
string
required
shortNameAttribute
string
required
isGroupEnrollmentAccessEnabled
boolean
required
groupEnrollmentAccessName
string
required
id
integer
type
string

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/enrollment-customization//sso/ \
     --header 'accept: application/json'

{
  "displayName": "A Panel",
  "rank": 0,
  "isUseJamfConnect": false,
  "longNameAttribute": "long name",
  "shortNameAttribute": "name",
  "isGroupEnrollmentAccessEnabled": false,
  "groupEnrollmentAccessName": "GroupNameA",
  "id": 2,
  "type": "sso"
}
-----
Update a single SSO Panel for a single Enrollment Customization
put
https://yourServer.jamfcloud.com/api/v1/enrollment-customization/{id}/sso/{panel-id}

Update a single SSO panel for a single enrollment customization

Path Params
id
integer
required
Enrollment Customization identifier

panel-id
integer
required
Panel object identifier

Body Params
Enrollment Customization Panel to update

displayName
string
required
rank
integer
required
isUseJamfConnect
boolean
required

true
longNameAttribute
string
required
shortNameAttribute
string
required
isGroupEnrollmentAccessEnabled
boolean
required

true
groupEnrollmentAccessName
string
required
Responses


curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v1/enrollment-customization//sso/ \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "isUseJamfConnect": true,
  "isGroupEnrollmentAccessEnabled": true
}
'

{
  "displayName": "A Panel",
  "rank": 0,
  "isUseJamfConnect": false,
  "longNameAttribute": "long name",
  "shortNameAttribute": "name",
  "isGroupEnrollmentAccessEnabled": false,
  "groupEnrollmentAccessName": "GroupNameA",
  "id": 2,
  "type": "sso"
}
-----
Delete a single SSO Panel from an Enrollment Customization
delete
https://yourServer.jamfcloud.com/api/v1/enrollment-customization/{id}/sso/{panel-id}

Delete a single SSO panel from an Enrollment Customization

Path Params
id
integer
required
Enrollment Customization identifier

panel-id
integer
required
Panel object identifier

curl --request DELETE \
     --url https://yourserver.jamfcloud.com/api/v1/enrollment-customization//sso/
-----
Create a Text Panel for a single Enrollment Customization
post
https://yourServer.jamfcloud.com/api/v1/enrollment-customization/{id}/text

Create a Text panel for a single enrollment customization

Path Params
id
integer
required
Enrollment Customization identifier

3
Body Params
Enrollment Customization Panel to create

displayName
string
required
A Panel
rank
integer
required
0
body
string
required
Welcome!
subtext
string
World!
title
string
required
My text panel
backButtonText
string
required
Back
continueButtonText
string
required
Continue
Responses

201
Text panel was created

Response body
object
displayName
string
required
rank
integer
required
body
string
required
subtext
string
title
string
required
backButtonText
string
required
continueButtonText
string
required
id
integer
type
string

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/enrollment-customization/3/text \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "displayName": "A Panel",
  "rank": 0,
  "body": "Welcome!",
  "subtext": "World!",
  "title": "My text panel",
  "backButtonText": "Back",
  "continueButtonText": "Continue"
}
'

{
  "displayName": "A Panel",
  "rank": 0,
  "body": "Welcome!",
  "subtext": "World!",
  "title": "My text panel",
  "backButtonText": "Back",
  "continueButtonText": "Continue",
  "id": 2,
  "type": "text"
}
-----
Get a single Text Panel for a single Enrollment Customization
get
https://yourServer.jamfcloud.com/api/v1/enrollment-customization/{id}/text/{panel-id}

Get a single Text panel for a single enrollment customization

Path Params
id
integer
required
Enrollment Customization identifier

1
panel-id
integer
required
Panel object identifier

4
Responses

200
Success

Response body
object
displayName
string
required
rank
integer
required
body
string
required
subtext
string
title
string
required
backButtonText
string
required
continueButtonText
string
required
id
integer
type
string

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/enrollment-customization/1/text/4 \
     --header 'accept: application/json'

{
  "displayName": "A Panel",
  "rank": 0,
  "body": "Welcome!",
  "subtext": "World!",
  "title": "My text panel",
  "backButtonText": "Back",
  "continueButtonText": "Continue",
  "id": 2,
  "type": "text"
}
-----
Update a single Text Panel for a single Enrollment Customization
put
https://yourServer.jamfcloud.com/api/v1/enrollment-customization/{id}/text/{panel-id}

Update a single Text panel for a single enrollment customization

Path Params
id
integer
required
Enrollment Customization identifier

2
panel-id
integer
required
Panel object identifier

4
Body Params
Enrollment Customization Panel to update

displayName
string
required
A Panel
rank
integer
required
0
body
string
required
Welcome!
subtext
string
World!
title
string
required
My text panel
backButtonText
string
required
Back
continueButtonText
string
required
Continue
Responses

curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v1/enrollment-customization/2/text/4 \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "displayName": "A Panel",
  "rank": 0,
  "body": "Welcome!",
  "subtext": "World!",
  "title": "My text panel",
  "backButtonText": "Back",
  "continueButtonText": "Continue"
}
'

{
  "displayName": "A Panel",
  "rank": 0,
  "body": "Welcome!",
  "subtext": "World!",
  "title": "My text panel",
  "backButtonText": "Back",
  "continueButtonText": "Continue",
  "id": 2,
  "type": "text"
}
-----
Delete a Text single Panel from an Enrollment Customization
delete
https://yourServer.jamfcloud.com/api/v1/enrollment-customization/{id}/text/{panel-id}

Delete a Text single panel from an Enrollment Customization

Path Params
id
integer
required
Enrollment Customization identifier

3
panel-id
integer
required
Panel object identifier

4

curl --request DELETE \
     --url https://yourserver.jamfcloud.com/api/v1/enrollment-customization/3/text/4
-----
Get the markdown output of a single Text Panel for a single Enrollment
get
https://yourServer.jamfcloud.com/api/v1/enrollment-customization/{id}/text/{panel-id}/markdown

Get the markdown output of a single Text panel for a single enrollment customization

Path Params
id
integer
required
Enrollment Customization identifier

1
panel-id
integer
required
Panel object identifier

3
Responses

200
Success

Response body
object
markdown
string

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/enrollment-customization/1/text/3/markdown \
     --header 'accept: application/json'

{
  "markdown": "**markdown**"
}
-----