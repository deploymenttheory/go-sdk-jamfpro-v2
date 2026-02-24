

-----

Return General section of a Computer
get
https://yourServer.jamfcloud.com/api/v3/computers-inventory/{id}


Return General section of a Computer

Path Params
id
string
required
instance id of computer record

Query Params
section
array of strings
Defaults to GENERAL
section of computer details, if not specified, General section data is returned. Multiple section parameters are supported, e.g. section=general&section=hardware


string


GENERAL

ADD string
-----

Get sorted and paged Computer Prestages
get
https://yourServer.jamfcloud.com/api/v3/computer-prestages

Gets sorted and paged computer prestages

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
Defaults to id:desc
Sorting criteria in the format: property:asc/desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc


string

id:desc

string


ADD string
-----

Return a list of Computers
get
https://yourServer.jamfcloud.com/api/preview/computers

Returns a list of computers.

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
string
Defaults to name:asc
Sorting criteria in the format: property:asc/desc. Default sort is name:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc


Read all sorted and paged Device Enrollment instances
get
https://yourServer.jamfcloud.com/api/v1/device-enrollments

Search for sorted and paged device enrollment instances

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
-----

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
-----

Get sorted and paged Enrollment history object
get
https://yourServer.jamfcloud.com/api/v2/enrollment/history

Gets sorted and paged Enrollment history object

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
Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc.


string

date:desc

ADD string

-----

Retrieve the configured LDAP groups configured for User-Initiated Enrollment.
get
https://yourServer.jamfcloud.com/api/v3/enrollment/access-groups


Retrieves the configured LDAP groups configured for User-Initiated Enrollment.

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
Sorting criteria in the format: property:asc/desc. Default sort is name:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc.


string

name:asc

ADD string
all-users-option-first
boolean
Defaults to false
Return "All LDAP Users" option on the first position if it is present in the current page


false
-----

Get an array of the language codes that have Enrollment messaging
get
https://yourServer.jamfcloud.com/api/v3/enrollment/languages


Returns an array of the language codes that have enrollment messaging currently configured.

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
Defaults to languageCode:asc
Sorting criteria in the format: property:asc/desc. Default sort is languageCode:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc.


string

languageCode:asc

ADD string

-----

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

ADD string

-----



-----



-----

Get Jamf Pro Server URL settings history
get
https://yourServer.jamfcloud.com/api/v1/jamf-pro-server-url/history


Gets Jamf Pro Server URL settings history

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
string
Defaults to date:desc
Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc

-----

Retrieve the configured access groups that contain the text in the search param
get
https://yourServer.jamfcloud.com/api/ldap/groups

Retrieves the configured access groups that contain the text in the searchParam.

Query Params
q
string
Defaults to null
Will perform a "contains" search on the names of access groups

null

-----

Retrieve the configured access groups that contain the text in the search param
get
https://yourServer.jamfcloud.com/api/v1/ldap/groups

Retrieves the configured access groups that contain the text in the searchParam.

Query Params
q
string
Defaults to null
Will perform a "contains" search on the names of access groups

null

-----

Retrieve Managed Software Update Plans for a Group
get
https://yourServer.jamfcloud.com/api/v1/managed-software-updates/plans/group/{id}


Retrieves Managed Software Update Plans for a Group

Path Params
id
string
required
Managed Software Update Group Id

Query Params
group-type
string
enum
required
Managed Software Update Group Type, Available options are "COMPUTER_GROUP" or "MOBILE_DEVICE_GROUP"


COMPUTER_GROUP
Allowed:

COMPUTER_GROUP

MOBILE_DEVICE_GROUP

-----


-----

Get Mobile Device Extension Attribute values placed in select paramter
get
https://yourServer.jamfcloud.com/api/devices/extensionAttributes


Gets Mobile Device Extension Attribute values placed in select parameter.

Query Params
select
string
Defaults to name
Acceptable values currently include:

name
name
Response

-----


Get sorted and paged Mobile Device Prestages
get
https://yourServer.jamfcloud.com/api/v3/mobile-device-prestages


Gets sorted and paged mobile device prestages

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
Defaults to id:desc
Sorting criteria in the format: property:asc/desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc


string

id:desc

ADD string

-----

Get sorted and paged Mobile Device Prestage history objects
get
https://yourServer.jamfcloud.com/api/v3/mobile-device-prestages/{id}/history


Gets sorted and paged mobile device prestage history objects

Path Params
id
string
required
Mobile Device Prestage identifier

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

ADD string

-----

Get Mobile Device objects
get
https://yourServer.jamfcloud.com/api/v2/mobile-devices

Gets Mobile Device objects.

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
Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc


string

id:asc

ADD string
-----

Retrieves a list of applications that are eligible to be used in an onboarding configuration
get
https://yourServer.jamfcloud.com/api/v1/onboarding/eligible-apps


Retrieves a list of applications that are eligible to be used in an onboarding configuration

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
Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc


string

id:asc

ADD string

-----

Retrieves a list of configuration profiles that are eligible to be used in an onboarding configuration
get
https://yourServer.jamfcloud.com/api/v1/onboarding/eligible-configuration-profiles


Retrieves a list of configuration profiles that are eligible to be used in an onboarding configuration

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
Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc


string

id:asc

ADD string
Response

-----

Retrieves a list of policies that are eligible to be used in an onboarding configuration
get
https://yourServer.jamfcloud.com/api/v1/onboarding/eligible-policies


Retrieves a list of policies that are eligible to be used in an onboarding configuration

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
Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc


string

id:asc

ADD string
Response

-----

Get Re-enrollment history object
get
https://yourServer.jamfcloud.com/api/v1/reenrollment/history

Gets Re-enrollment history object

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
string
Defaults to date:desc
Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc

-----

Get information about all remote administration configurations.
get
https://yourServer.jamfcloud.com/api/preview/remote-administration-configurations


Remote administration feature creates a secure screen-sharing experience between Jamf Pro administrators and their end-users.

Query Params
page
integer
Defaults to 0
0
page-size
integer
Defaults to 100
100

-----

Search for sorted and paged iOS branding configurations
get
https://yourServer.jamfcloud.com/api/v1/self-service/branding/ios


Search for sorted and paged iOS branding configurations

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
Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=id:desc,brandingName:asc


string

id:asc

ADD string

-----

Search for sorted and paged macOS branding configurations
get
https://yourServer.jamfcloud.com/api/v1/self-service/branding/macos


Search for sorted and paged macOS branding configurations

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
Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=id:desc,brandingName:asc


string

id:asc

ADD string

-----

Search for sorted and paged Supervision Identities
get
https://yourServer.jamfcloud.com/api/v1/supervision-identities


Search for sorted and paged supervision identities

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
string
Defaults to id:asc
Sorting criteria in the format: property:asc/desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc

id:asc

-----

Retrieve Volume Purchasing Subscriptions
get
https://yourServer.jamfcloud.com/api/v1/volume-purchasing-subscriptions


Retrieves Volume Purchasing Subscriptions

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
Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. Allowable properties are id, name, and enabled.


string

id:asc

ADD string
