endpoints that support rqql are as follows:

Get user accounts
get
https://yourServer.jamfcloud.com/api/v1/accounts

Get all user accounts with pagination, sorting, and filtering support.

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
Defaults to username:asc
Sorting criteria in the format: property:asc/desc. Default sort is username:desc. Multiple sort criteria are supported and must be separated with a comma. Accepts fields: id, lastPasswordChange, failedLoginAttempts, username, realname, email, phone, ldapServerId, distinguishedName, siteId, privilegeLevel, changePasswordOnNextLogin, accountStatus. If any other field is passed it will be ignored in sorting operation and/or create unpredictable results.


string

username:asc

ADD string
filter
string
Query in the RSQL format to filter user accounts collection. An empty query returns all results for the requested page. Supported fields: id, lastPasswordChange, failedLoginAttempts, username, realname, email, phone, ldapServerId, distinguishedName, siteId, privilegeLevel, changePasswordOnNextLogin, accountStatus. Multiple conditions can be combined using logical operators. This parameter can be used with paging and sorting parameters. Example: username=="admin" and accountStatus==Enabled and failedLoginAttempts==0

Response

-----
Get Activation Code history object
get
https://yourServer.jamfcloud.com/api/v1/activation-code/history


Get Activation Code history object

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
Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Fields allowed in the query: id, username, date, note, details Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,note:asc


string

date:desc

ADD string
filter
string
Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: id, username, date, note, details. This param can be combined with paging and sorting. Example: filter=username!=admin and details==disabled and date<2019-12-15

Responses

--------
Get specified AD CS Settings history object
get
https://yourServer.jamfcloud.com/api/v1/pki/adcs-settings/{id}/history


Get specified AD CS Settings history object.

Path Params
id
string
required
ID of the AD CS Settings configuration.

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

---------

Get the current API Integrations
get
https://yourServer.jamfcloud.com/api/v1/api-integrations

Get Jamf|Pro API Integrations with Search Criteria

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
Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. Fields allowed in the query: id, displayName. Example: sort=displayName:desc


string

id:asc

ADD string
filter
string
Query in the RSQL format, allowing to filter app titles collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: id, displayName. Example: displayName=="IntegrationName"

---------

Get the current Jamf API Roles
get
https://yourServer.jamfcloud.com/api/v1/api-roles

Get roles with Search Criteria

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
Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. Fields allowed in the query: id, displayName. Example: sort=displayName:desc


string

id:asc

ADD string
filter
string
Query in the RSQL format, allowing to filter app titles collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: id, displayName. Example: displayName=="myRole"

---------

Search for clients with push notifications disabled
get
https://yourServer.jamfcloud.com/api/v1/apns-client-push-status

Retrieve a paginated, sortable, and filterable list of MDM clients that have push notifications disabled. The endpoint queries the mdm_client table and returns information about when push was disabled and links to the device records.

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
Defaults to pushDisabledTime:asc
Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. Sortable fields: pushDisabledTime, deviceType, managementId


string

pushDisabledTime:asc

ADD string
filter
string
Query in the RSQL format, allowing to filter results. Fields allowed in the query: deviceType, disabledAt, managementId. This param can be combined with paging and sorting.

Example: filter=deviceType=="MOBILE_DEVICE" Example: filter=disabledAt>2024-11-01T00:00:00Z Example: filter=deviceType=="COMPUTER";disabledAt>2024-01-01T00:00:00Z

--------

Search for sorted and paged Buildings
get
https://yourServer.jamfcloud.com/api/v1/buildings

Search for sorted and paged buildings

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
filter
string
Query in the RSQL format, allowing to filter buildings collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: name, streetAddress1, streetAddress2, city, stateProvince, zipPostalCode, country. This param can be combined with paging and sorting. Example: filter=city=="Chicago" and name=="build"


-----------

Get specified Building History object
get
https://yourServer.jamfcloud.com/api/v1/buildings/{id}/history

Gets specified Building history object

Path Params
id
string
required
instance id of building history record

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

-------

Get Category objects
get
https://yourServer.jamfcloud.com/api/v1/categories

Gets Category objects.

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
filter
string
Query in the RSQL format, allowing to filter categories collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: name, priority. This param can be combined with paging and sorting. Example: filter=name=="Apps*" and priority>=5

-----

Get specified Category history object
get
https://yourServer.jamfcloud.com/api/v1/categories/{id}/history


Gets specified Category history object

Path Params
id
string
required
instance id of category history record

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
Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc


string

date:desc

ADD string
filter
string
Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: username, date, note, details. This param can be combined with paging and sorting. Example: filter=username!=admin and details==disabled and date<2019-12-15

-----

Get Client Check-In history object
get
https://yourServer.jamfcloud.com/api/v3/check-in/history

Gets Client Check-In history object

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
Sorting criteria in the format: property:asc/desc. Default sort is name:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,username:asc


string

date:desc

ADD string
filter
string
Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: username, date, note, details. This param can be combined with paging and sorting. Example: filter=username!=admin and details==disabled and date<2019-12-15

-----

Get cloud distribution point history details
get
https://yourServer.jamfcloud.com/api/v1/cloud-distribution-point/history


Get cloud distribution point history details

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

-----

Get the cloud distribution point Inventory files details
get
https://yourServer.jamfcloud.com/api/v1/cloud-distribution-point/files


Retrieves the details of the inventory files associated with a cloud distribution point.This includes information about the files used for content distribution, such as their type, status, and categorization.The response provides a comprehensive list of inventory files, which may include packages, ebooks, or mobile device apps, allowing users to view the current state and metadata for each file in the distribution system.

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
Defaults to id.asc
Sorts results by one or more criteria, following the format property:asc/desc.
Default sort is id:asc.
If using multiple criteria, separate with commas. Allows sort for id, fileName, inventoryId and type etc.


string

id.asc

ADD string
filter
string
Filters results. Use RSQL format for query. Allows for many fields, including fileName and type
Can be combined with paging and sorting.
Fields allowed in the query: fileName, inventoryId and type
Default filter is an empty query and returns all results from the requested page.

-----

Get Cloud Identity Provider history
get
https://yourServer.jamfcloud.com/api/v1/cloud-idp/{id}/history

Gets specified Cloud Identity Provider object history

Path Params
id
string
required
Cloud Identity Provider identifier

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


-----

Retrieve Computer Extension Attributes.
get
https://yourServer.jamfcloud.com/api/v1/computer-extension-attributes


Retrieves All Computer Extension Attributes Configuration.

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
-----

Retrieve Computer Extension Attributes.
get
https://yourServer.jamfcloud.com/api/v1/computer-extension-attributes


Retrieves All Computer Extension Attributes Configuration.

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

-----

Search for Smart Computer Groups
get
https://yourServer.jamfcloud.com/api/v2/computer-groups/smart-groups


Search for Smart Computer Groups

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
Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=name:asc


string

id:asc

ADD string
filter
string
Query in the RSQL format, allowing to filter smart computer group collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: id, name, siteId. The siteId field can only be filtered by admins with full access. Any sited admin will have siteId filtered automatically. Example: name=="group"

-----

Search for Static Computer Groups
get
https://yourServer.jamfcloud.com/api/v2/computer-groups/static-groups


Search for Static Computer Groups

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
Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=name:asc


string

id:asc

ADD string
filter
string
Query in the RSQL format, allowing to filter static computer group collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: id, name, siteId. The siteId field can only be filtered by admins with full access. Any sited admin will have siteId filtered automatically. Example: name=="group"
-----

Search for Departments
get
https://yourServer.jamfcloud.com/api/v1/departments

Search for Departments

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
filter
string
Query in the RSQL format, allowing to filter department collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: id, name. Example: name=="department"

-----

Get specified Department history object
get
https://yourServer.jamfcloud.com/api/v1/departments/{id}/history


Gets specified Department history object

Path Params
id
string
required
instance id of department history record

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

-----

Get Device Communication settings history
get
https://yourServer.jamfcloud.com/api/v1/device-communication-settings/history


Gets Device Communication settings history

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

-----

Get sorted and paged Device Enrollment history objects
get
https://yourServer.jamfcloud.com/api/v1/device-enrollments/{id}/history


Gets sorted and paged device enrollment history objects

Path Params
id
string
required
Device Enrollment Instance identifier

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
filter
string
Query in the RSQL format, allowing to filter history notes collection. Default search is empty query - returning all results for the requested page. Fields allowed in the query: username, date, note, details. This param can be combined with paging and sorting. Example: search=username!=admin and details==disabled and date<2019-12-15

-----

Finds all Distribution Points
get
https://yourServer.jamfcloud.com/api/v1/distribution-points

Finds all Distribution Points

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
Sorts results by one or more criteria, following the format property:asc/desc. Default sort is id:asc. If using multiple criteria, separate with commas. Allows fields such as - name, serverName


string

id:asc

ADD string
filter
string
Filters results. Use RSQL format for query. Allows fields such as - name, serverName, principal, fileSharingConnectionType, and httpsEnabled Can be combined with paging and sorting. Default filter is an empty query and returns all results from the requested page.

-----

Returns group information for all Mobile Device and Computer groups
get
https://yourServer.jamfcloud.com/api/v1/groups

Returns group information for all Mobile Device and Computer groups. The type of groups returned will be dependent upon the corresponding group type READ privileges. Results can be sorted by name, description, group type, or isSmart. Default sorting is by group name in ascending order.

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
Defaults to groupName:asc
Sorting criteria in the format: property:asc/desc. Default sort is groupName:asc. Multiple sort criteria are supported and must be separated with a comma. Fields allowed in sorting: groupName, groupDescription, groupType, isSmart. Example: sort=groupName:asc,groupType:desc


string

groupName:asc

ADD string
filter
string
Query in the RSQL format, allowing to filter group collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: groupName, groupDescription, groupType, isSmart. This param can be combined with paging and sorting. When using groupType in the filter, the value must be either "MOBILE" or "COMPUTER" but not both. When using groupType in the filter, the value is case sensitive. When using groupType in the filter, it will exclude groups of the other type regardless of or/and conditionals. Example: filter=groupName=="Managed" and isSmart=="true" Example: filter=groupType=="COMPUTER" and groupDescription=="Admin"

-----

Get specified GSX Connection History object
get
https://yourServer.jamfcloud.com/api/v1/gsx-connection/history

Gets specified GSX Connection history object

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

------

Search for config profiles linked to Jamf Connect
get
https://yourServer.jamfcloud.com/api/v1/jamf-connect/config-profiles


Search for config profiles linked to Jamf Connect

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
Defaults to
Sorting criteria in the format: property:asc/desc. Default sort order is descending. Multiple sort criteria are supported and must be entered on separate lines in Swagger UI. In the URI the 'sort' query param is not duplicated for each sort criterion, e.g., ...&sort=name:asc,date:desc. Fields that can be sorted: status, updated


ADD string
filter
string
Query in the RSQL format, allowing to filter results. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: status, updated, version This param can be combined with paging and sorting. Example: filter=username!=admin and details==disabled and date<2019-12-15

-----

Search for deployment tasks for a config profile linked to Jamf Connect
get
https://yourServer.jamfcloud.com/api/v1/jamf-connect/deployments/{id}/tasks


Search for config profiles linked to Jamf Connect

Path Params
id
string
required
the UUID of the Jamf Connect deployment

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
Defaults to
Sorting criteria in the format: property:asc/desc. Default sort order is descending. Multiple sort criteria are supported and must be entered on separate lines in Swagger UI. In the URI the 'sort' query param is not duplicated for each sort criterion, e.g., ...&sort=name:asc,date:desc. Fields that can be sorted: status, updated


ADD string
filter
string
Query in the RSQL format, allowing to filter results. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: status, updated, version This param can be combined with paging and sorting. Example: filter=username!=admin and details==disabled and date<2019-12-15

-----

Get Jamf Connect history
get
https://yourServer.jamfcloud.com/api/v1/jamf-connect/history


Get Jamf Connect history

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
Defaults to
Sorting criteria in the format: property:asc/desc. Default sort order is descending. Multiple sort criteria are supported and must be entered on separate lines in Swagger UI. In the URI the 'sort' query param is not duplicated for each sort criterion, e.g., ...&sort=name:asc,date:desc. Fields that can be sorted: status, updated


ADD string
filter
string
Query in the RSQL format, allowing to filter results. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: status, updated, version This param can be combined with paging and sorting. Example: filter=username!=admin and details==disabled and date<2019-12-15

-----

Search for deployment tasks for a config profile linked to Jamf Protect
get
https://yourServer.jamfcloud.com/api/v1/jamf-protect/deployments/{id}/tasks


Search for config profiles linked to Jamf Protect

Path Params
id
string
required
the UUID of the Jamf Protect deployment

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
Defaults to
Sorting criteria in the format: property:asc/desc. Default sort order is descending. Multiple sort criteria are supported and must be entered on separate lines in Swagger UI. In the URI the 'sort' query param is not duplicated for each sort criterion, e.g., ...&sort=name:asc,date:desc. Fields that can be sorted: status, updated


ADD string
filter
string
Query in the RSQL format, allowing to filter results. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: status, updated, version This param can be combined with paging and sorting. Example: filter=username!=admin and details==disabled and date<2019-12-15

-----

Get Jamf Protect history
get
https://yourServer.jamfcloud.com/api/v1/jamf-protect/history


Get Jamf Protect history

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
Defaults to
Sorting criteria in the format: property:asc/desc. Default sort order is descending. Multiple sort criteria are supported and must be entered on separate lines in Swagger UI. In the URI the 'sort' query param is not duplicated for each sort criterion, e.g., ...&sort=name:asc,date:desc. Fields that can be sorted: status, updated


ADD string
filter
string
Query in the RSQL format, allowing to filter results. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: status, updated, version This param can be combined with paging and sorting. Example: filter=username!=admin and details==disabled and date<2019-12-15

-----

Get all of the previously synced Jamf Protect Plans with information about their associated configuration profile
get
https://yourServer.jamfcloud.com/api/v1/jamf-protect/plans

Get all of the previously synced Jamf Protect Plans with information about their associated configuration profile

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
Defaults to
Sorting criteria in the format: property:asc/desc. Default sort order is descending. Multiple sort criteria are supported and must be entered on separate lines in Swagger UI. In the URI the 'sort' query param is not duplicated for each sort criterion, e.g., ...&sort=name:asc,date:desc. Fields that can be sorted: status, updated


ADD string
filter
string
Query in the RSQL format, allowing to filter results. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: status, updated, version This param can be combined with paging and sorting. Example: filter=username!=admin and details==disabled and date<2019-12-15

-----

Gets session history items.
get
https://yourServer.jamfcloud.com/api/v2/jamf-remote-assist/session


Returns tenants sessions history.

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
Defaults to sessionId:desc
Sorting criteria in the format: property:asc/desc. Default sort is sessionId:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=sessionId:desc,deviceId:asc


string

sessionId:desc

ADD string
filter
string
Query in the RSQL format, allowing to filter session history items collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: sessionId, deviceId, sessionAdminId. This param can be combined with paging and sorting. Example: sessionAdminId=="Andrzej"

-----

Retrieve Managed Software Update Plans
get
https://yourServer.jamfcloud.com/api/v1/managed-software-updates/plans


Retrieve Managed Software Update Plans

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
Defaults to planUuid:asc
Sorting criteria in the format: property:asc/desc. Default sort is planUuid:asc. Multiple sort criteria are supported and must be separated with a comma.


string

planUuid:asc

ADD string
filter
string
Query in the RSQL format, allowing to filter Managed Software Updates collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: planUuid, device.deviceId, device.objectType, updateAction, versionType, specificVersion, maxDeferrals, recipeId, forceInstallLocalDateTime, state.

-----

Get information about mdm commands made by Jamf Pro.
get
https://yourServer.jamfcloud.com/api/v2/mdm/commands

Get information about mdm commands made by Jamf Pro.

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
Defaults to dateSent:asc
Default sort is dateSent:asc. Multiple sort criteria are supported and must be separated with a comma.


string

dateSent:asc

ADD string
filter
string
Query in the RSQL format, allowing to filter, for a list of commands. All url must contain minimum one filter field. Fields allowed in the query: uuid, clientManagementId, command, status, clientType, dateSent, validAfter, dateCompleted, profileId, profileIdentifier, and active. This param can be combined with paging. Please note that any date filters must be used with gt, lt, ge, le Example: clientManagementId==fb511aae-c557-474f-a9c1-5dc845b90d0f;status==Pending;command==INSTALL_PROFILE;uuid==9e18f849-e689-4f2d-b616-a99d3da7db42;clientType==COMPUTER_USER;profileId==1;profileIdentifier==18cc61c2-01fc-11ed-b939-0242ac120002;dateCompleted=ge=2021-08-04T14:25:18.26Z;dateCompleted=le=2021-08-04T14:25:18.26Z;validAfter=ge=2021-08-05T14:25:18.26Z;active==true

-----

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

-----

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

-----

Get Static Groups
get
https://yourServer.jamfcloud.com/api/v1/mobile-device-groups/static-groups


Get Static Groups

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
Defaults to groupId:asc
Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Available criteria to sort on: groupId, groupName, siteId.


string

groupId:asc

ADD string
filter
string
Query in the RSQL format, allowing to filter department collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: groupId, groupName, siteId. The siteId field can only be filtered by admins with full access. Any sited admin will have siteId filtered automatically. This param can be combined with paging and sorting. Example: groupName=="staticGroup1"

-----

Get Smart Groups
get
https://yourServer.jamfcloud.com/api/v1/mobile-device-groups/smart-groups


Get Smart Groups

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
Defaults to groupId:asc
Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Available criteria to sort on: groupId, groupName, siteId.


string

groupId:asc

ADD string
filter
string
Query in the RSQL format, allowing to filter smart group collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: groupId, groupName, siteId. The siteId field can only be filtered by admins with full access. Any sited admin will have siteId filtered automatically. This param can be combined with paging and sorting. Example: groupName=="smartGroup1"

-----

Return paginated Computer Inventory records
get
https://yourServer.jamfcloud.com/api/v3/computers-inventory

Return paginated Computer Inventory records

Query Params
section
array of strings
Defaults to GENERAL
section of computer details, if not specified, General section data is returned. Multiple section parameters are supported, e.g. section=GENERAL&section=HARDWARE


string


GENERAL

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
Defaults to general.name:asc
Sorting criteria in the format: property:asc/desc. Default sort is general.name:asc. Multiple sort criteria are supported and must be separated with a comma.

Fields allowed in the sort: general.name, udid, id, general.assetTag, general.jamfBinaryVersion, general.lastContactTime, general.lastEnrolledDate, general.lastCloudBackupDate, general.reportDate, general.mdmCertificateExpiration, general.platform, general.lastLoggedInUsernameSelfService, general.lastLoggedInUsernameSelfServiceTimestamp, general.mdmCertificateExpiration, general.platform, general.lastLoggedInUsernameBinary, general.lastLoggedInUsernameBinaryTimestamp hardware.make, hardware.model, operatingSystem.build, operatingSystem.supplementalBuildVersion, operatingSystem.rapidSecurityResponse, operatingSystem.name, operatingSystem.version, userAndLocation.realname, purchasing.lifeExpectancy, purchasing.warrantyDate

Example: sort=udid:desc,general.name:asc.


string

general.name:asc

string


ADD string
filter
string
Query in the RSQL format, allowing to filter computer inventory collection. Default filter is empty query - returning all results for the requested page.

Fields allowed in the query: general.name, udid, id, general.assetTag, general.barcode1, general.barcode2, general.enrolledViaAutomatedDeviceEnrollment, general.lastIpAddress, general.itunesStoreAccountActive, general.jamfBinaryVersion, general.lastContactTime, general.lastEnrolledDate, general.lastCloudBackupDate, general.reportDate, general.lastReportedIp, general.lastReportedIpV4, general.lastReportedIpV6, general.managementId, general.remoteManagement.managed, general.mdmCapable.capable, general.mdmCertificateExpiration, general.platform, general.supervised, general.userApprovedMdm, general.declarativeDeviceManagementEnabled, general.lastLoggedInUsernameSelfService, general.lastLoggedInUsernameSelfServiceTimestamp, general.mdmCapable.capable, general.mdmCertificateExpiration, general.platform, general.supervised, general.userApprovedMdm, general.declarativeDeviceManagementEnabled, general.lastLoggedInUsernameBinary, general.lastLoggedInUsernameBinaryTimestamp, hardware.bleCapable, hardware.macAddress, hardware.make, hardware.model, hardware.modelIdentifier, hardware.serialNumber, hardware.supportsIosAppInstalls,hardware.appleSilicon, operatingSystem.activeDirectoryStatus, operatingSystem.fileVault2Status, operatingSystem.build, operatingSystem.supplementalBuildVersion, operatingSystem.rapidSecurityResponse, operatingSystem.name, operatingSystem.version, security.activationLockEnabled, security.recoveryLockEnabled,security.firewallEnabled,userAndLocation.buildingId, userAndLocation.departmentId, userAndLocation.email, userAndLocation.realname, userAndLocation.phone, userAndLocation.position,userAndLocation.room, userAndLocation.username, diskEncryption.fileVault2Enabled, purchasing.appleCareId, purchasing.lifeExpectancy, purchasing.purchased, purchasing.leased, purchasing.vendor, purchasing.warrantyDate,

This param can be combined with paging and sorting. Example: filter=general.name=="Orchard"

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

-----

Retrieve Managed Software Update Statuses
get
https://yourServer.jamfcloud.com/api/v1/managed-software-updates/update-statuses


Retrieve Managed Software Update Statuses

Query Params
filter
string
Query in the RSQL format, allowing to filter Managed Software Updates collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: osUpdatesStatusId, device.deviceId, device.objectType, downloaded, downloadPercentComplete, productKey, status, deferralsRemaining, maxDeferrals, nextScheduledInstall, created and updated.

-----


Get Static Group Membership by Id
get
https://yourServer.jamfcloud.com/api/v1/mobile-device-groups/static-group-membership/{id}


Get Static Group Membership by Id

Path Params
id
string
required
instance id of static-group

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
Defaults to displayName:asc
Sorting criteria in the format: property:asc/desc. Default sort is mobileDeviceId:asc. Multiple sort criteria are supported and must be separated with a comma.

Fields allowed in the sort: airPlayPassword, appAnalyticsEnabled, assetTag, availableSpaceMb, batteryLevel, batteryHealth, bluetoothLowEnergyCapable, bluetoothMacAddress, capacityMb, lostModeEnabledDate, declarativeDeviceManagementEnabled, deviceId, deviceLocatorServiceEnabled, devicePhoneNumber, diagnosticAndUsageReportingEnabled, displayName, doNotDisturbEnabled, enrollmentSessionTokenValid, exchangeDeviceId, cloudBackupEnabled, osBuild, osRapidSecurityResponse, osSupplementalBuildVersion, osVersion, ipAddress, itunesStoreAccountActive, mobileDeviceId, managementId, languages, lastBackupDate, lastEnrolledDate, lastCloudBackupDate, lastInventoryUpdateDate, locales, locationServicesForSelfServiceMobileEnabled, lostModeEnabled, managed, mdmProfileExpirationDate, model, modelIdentifier, modelNumber, modemFirmwareVersion, preferredVoiceNumber, quotaSize, residentUsers, serialNumber, sharedIpad, supervised, tethered, timeZone, udid, usedSpacePercentage, wifiMacAddress, deviceOwnershipType, building, department, emailAddress, fullName, userPhoneNumber, position, room, username, appleCareId, leaseExpirationDate,lifeExpectancyYears, poDate, poNumber, purchasePrice, purchasedOrLeased, purchasingAccount, purchasingContact, vendor, warrantyExpirationDate, activationLockEnabled, blockEncryptionCapable, dataProtection, fileEncryptionCapable, hardwareEncryptionSupported, jailbreakStatus, passcodeCompliant, passcodeCompliantWithProfile, passcodeLockGracePeriodEnforcedSeconds, passcodePresent, carrierSettingsVersion, cellularTechnology, currentCarrierNetwork, currentMobileCountryCode, currentMobileNetworkCode, dataRoamingEnabled, eid, network, homeMobileCountryCode, homeMobileNetworkCode, iccid, imei, imei2, meid, personalHotspotEnabled, voiceRoamingEnabled, roaming, lastLoggedInUsernameSelfService, lastLoggedInUsernameSelfServiceTimestamp

Extension attributes can be sorted by using the format EA+ID where ID is the ID of the extension attribute, for example EA+1!=null

Example: sort=displayName:desc,username:asc


string

displayName:asc

ADD string
filter
string
Query in the RSQL format, allowing to filter mobile device collection. Default filter is empty query - returning all results for the requested page.

Fields allowed in the query: airPlayPassword, appAnalyticsEnabled, assetTag, availableSpaceMb, batteryLevel, bluetoothLowEnergyCapable, bluetoothMacAddress, capacityMb, declarativeDeviceManagementEnabled, deviceId, deviceLocatorServiceEnabled, devicePhoneNumber, diagnosticAndUsageReportingEnabled, displayName, doNotDisturbEnabled, exchangeDeviceId, cloudBackupEnabled, osBuild, osSupplementalBuildVersion, osVersion, osRapidSecurityResponse, ipAddress, itunesStoreAccountActive, mobileDeviceId, managementId, languages, lastInventoryUpdateDate, locales, locationServicesForSelfServiceMobileEnabled, lostModeEnabled, managed, model, modelIdentifier, modelNumber, modemFirmwareVersion, preferredVoiceNumber, quotaSize, residentUsers, serialNumber, sharedIpad, supervised, tethered, timeZone, udid, usedSpacePercentage, wifiMacAddress, building, department, emailAddress, fullName, userPhoneNumber, position, room, username, appleCareId, lifeExpectancyYears, poNumber, purchasePrice, purchasedOrLeased, purchasingAccount, purchasingContact, vendor, activationLockEnabled, blockEncryptionCapable, dataProtection, fileEncryptionCapable, passcodeCompliant, passcodeCompliantWithProfile, passcodeLockGracePeriodEnforcedSeconds, passcodePresent, carrierSettingsVersion, currentCarrierNetwork, currentMobileCountryCode, currentMobileNetworkCode, dataRoamingEnabled, eid, network, homeMobileCountryCode, homeMobileNetworkCode, iccid, imei, imei2, meid, personalHotspotEnabled, roaming, lastLoggedInUsernameSelfService, lastLoggedInUsernameSelfServiceTimestamp

Extension attributes can be filtered by using the format EA+ID where ID is the ID of the extension attribute, for example EA+1!=null

This param can be combined with paging and sorting. Example: filter=displayName=="iPad"

-----



Get Smart Group Membership by Id
get
https://yourServer.jamfcloud.com/api/v1/mobile-device-groups/smart-group-membership/{id}


Get Smart Group Membership by Id

Path Params
id
string
required
instance id of smart-group

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
Defaults to displayName:asc
Sorting criteria in the format: property:asc/desc. Default sort is mobileDeviceId:asc. Multiple sort criteria are supported and must be separated with a comma.

Fields allowed in the sort: airPlayPassword, appAnalyticsEnabled, assetTag, availableSpaceMb, batteryLevel, batteryHealth, bluetoothLowEnergyCapable, bluetoothMacAddress, capacityMb, lostModeEnabledDate, declarativeDeviceManagementEnabled, deviceId, deviceLocatorServiceEnabled, devicePhoneNumber, diagnosticAndUsageReportingEnabled, displayName, doNotDisturbEnabled, enrollmentSessionTokenValid, exchangeDeviceId, cloudBackupEnabled, osBuild, osRapidSecurityResponse, osSupplementalBuildVersion, osVersion, ipAddress, itunesStoreAccountActive, mobileDeviceId, managementId, languages, lastBackupDate, lastEnrolledDate, lastCloudBackupDate, lastInventoryUpdateDate, locales, locationServicesForSelfServiceMobileEnabled, lostModeEnabled, managed, mdmProfileExpirationDate, model, modelIdentifier, modelNumber, modemFirmwareVersion, preferredVoiceNumber, quotaSize, residentUsers, serialNumber, sharedIpad, supervised, tethered, timeZone, udid, usedSpacePercentage, wifiMacAddress, deviceOwnershipType, building, department, emailAddress, fullName, userPhoneNumber, position, room, username, appleCareId, leaseExpirationDate,lifeExpectancyYears, poDate, poNumber, purchasePrice, purchasedOrLeased, purchasingAccount, purchasingContact, vendor, warrantyExpirationDate, activationLockEnabled, blockEncryptionCapable, dataProtection, fileEncryptionCapable, hardwareEncryptionSupported, jailbreakStatus, passcodeCompliant, passcodeCompliantWithProfile, passcodeLockGracePeriodEnforcedSeconds, passcodePresent, carrierSettingsVersion, cellularTechnology, currentCarrierNetwork, currentMobileCountryCode, currentMobileNetworkCode, dataRoamingEnabled, eid, network, homeMobileCountryCode, homeMobileNetworkCode, iccid, imei, imei2, meid, personalHotspotEnabled, voiceRoamingEnabled, roaming, lastLoggedInUsernameSelfService, lastLoggedInUsernameSelfServiceTimestamp

Extension attributes can be sorted by using the format EA+ID where ID is the ID of the extension attribute, for example EA+1!=null

Example: sort=displayName:desc,username:asc


string

displayName:asc

ADD string
filter
string
Query in the RSQL format, allowing to filter mobile device collection. Default filter is empty query - returning all results for the requested page.

Fields allowed in the query: airPlayPassword, appAnalyticsEnabled, assetTag, availableSpaceMb, batteryLevel, bluetoothLowEnergyCapable, bluetoothMacAddress, capacityMb, declarativeDeviceManagementEnabled, deviceId, deviceLocatorServiceEnabled, devicePhoneNumber, diagnosticAndUsageReportingEnabled, displayName, doNotDisturbEnabled, exchangeDeviceId, cloudBackupEnabled, osBuild, osSupplementalBuildVersion, osVersion, osRapidSecurityResponse, ipAddress, itunesStoreAccountActive, mobileDeviceId, managementId, languages, lastInventoryUpdateDate, locales, locationServicesForSelfServiceMobileEnabled, lostModeEnabled, managed, model, modelIdentifier, modelNumber, modemFirmwareVersion, preferredVoiceNumber, quotaSize, residentUsers, serialNumber, sharedIpad, supervised, tethered, timeZone, udid, usedSpacePercentage, wifiMacAddress, building, department, emailAddress, fullName, userPhoneNumber, position, room, username, appleCareId, lifeExpectancyYears, poNumber, purchasePrice, purchasedOrLeased, purchasingAccount, purchasingContact, vendor, activationLockEnabled, blockEncryptionCapable, dataProtection, fileEncryptionCapable, passcodeCompliant, passcodeCompliantWithProfile, passcodeLockGracePeriodEnforcedSeconds, passcodePresent, carrierSettingsVersion, currentCarrierNetwork, currentMobileCountryCode, currentMobileNetworkCode, dataRoamingEnabled, eid, network, homeMobileCountryCode, homeMobileNetworkCode, iccid, imei, imei2, meid, personalHotspotEnabled, roaming, lastLoggedInUsernameSelfService, lastLoggedInUsernameSelfServiceTimestamp

Extension attributes can be filtered by using the format EA+ID where ID is the ID of the extension attribute, for example EA+1!=null

This param can be combined with paging and sorting. Example: filter=displayName=="iPad"

-----


Return paginated Mobile Device Inventory records
get
https://yourServer.jamfcloud.com/api/v2/mobile-devices/detail

Return paginated Mobile Device Inventory records

Query Params
section
array of strings
Defaults to GENERAL
section of mobile device details, if not specified, General section data is returned. Multiple section parameters are supported, e.g. section=GENERAL&section=HARDWARE


string


GENERAL

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
Defaults to displayName:asc
Sorting criteria in the format: property:asc/desc. Default sort is displayName:asc. Multiple sort criteria are supported and must be separated with a comma.

Fields allowed in the sort: airPlayPassword, appAnalyticsEnabled, assetTag, availableSpaceMb, batteryLevel, batteryHealth, bluetoothLowEnergyCapable, bluetoothMacAddress, capacityMb, lostModeEnabledDate, declarativeDeviceManagementEnabled, deviceId, deviceLocatorServiceEnabled, devicePhoneNumber, diagnosticAndUsageReportingEnabled, displayName, doNotDisturbEnabled, enrollmentSessionTokenValid, exchangeDeviceId, cloudBackupEnabled, osBuild, osSupplementalBuildVersion, osVersion, osRapidSecurityResponse, ipAddress, itunesStoreAccountActive, mobileDeviceId, managementId, languages, lastBackupDate, lastEnrolledDate, lastCloudBackupDate, lastInventoryUpdateDate, locales, locationServicesForSelfServiceMobileEnabled, lostModeEnabled, managed, mdmProfileExpirationDate, model, modelIdentifier, modelNumber, modemFirmwareVersion, preferredVoiceNumber, quotaSize, residentUsers, serialNumber, sharedIpad, supervised, tethered, timeZone, udid, usedSpacePercentage, wifiMacAddress, deviceOwnershipType, building, department, emailAddress, fullName, userPhoneNumber, position, room, username, appleCareId, leaseExpirationDate,lifeExpectancyYears, poDate, poNumber, purchasePrice, purchasedOrLeased, purchasingAccount, purchasingContact, vendor, warrantyExpirationDate, activationLockEnabled, blockEncryptionCapable, dataProtection, fileEncryptionCapable, hardwareEncryptionSupported, jailbreakStatus, passcodeCompliant, passcodeCompliantWithProfile, passcodeLockGracePeriodEnforcedSeconds, passcodePresent, carrierSettingsVersion, cellularTechnology, currentCarrierNetwork, currentMobileCountryCode, currentMobileNetworkCode, dataRoamingEnabled, eid, network, homeMobileCountryCode, homeMobileNetworkCode, iccid, imei, imei2, meid, personalHotspotEnabled, voiceRoamingEnabled, roaming, lastLoggedInUsernameSelfService, lastLoggedInUsernameSelfServiceTimestamp

Example: sort=displayName:desc,username:asc


string

displayName:asc

ADD string
filter
string
Query in the RSQL format, allowing to filter mobile device collection. Default filter is empty query - returning all results for the requested page.

Fields allowed in the query: airPlayPassword, appAnalyticsEnabled, assetTag, availableSpaceMb, batteryLevel, bluetoothLowEnergyCapable, bluetoothMacAddress, capacityMb, declarativeDeviceManagementEnabled, deviceId, deviceLocatorServiceEnabled, devicePhoneNumber, diagnosticAndUsageReportingEnabled, displayName, doNotDisturbEnabled, exchangeDeviceId, cloudBackupEnabled, osBuild, osSupplementalBuildVersion, osVersion, osRapidSecurityResponse, ipAddress, itunesStoreAccountActive, mobileDeviceId, managementId, languages, lastInventoryUpdateDate, locales, locationServicesForSelfServiceMobileEnabled, lostModeEnabled, managed, model, modelIdentifier, modelNumber, modemFirmwareVersion, preferredVoiceNumber, quotaSize, residentUsers, serialNumber, sharedIpad, supervised, tethered, timeZone, udid, usedSpacePercentage, wifiMacAddress, building, department, emailAddress, fullName, userPhoneNumber, position, room, username, appleCareId, lifeExpectancyYears, poNumber, purchasePrice, purchasedOrLeased, purchasingAccount, purchasingContact, vendor, activationLockEnabled, blockEncryptionCapable, dataProtection, fileEncryptionCapable, passcodeCompliant, passcodeCompliantWithProfile, passcodeLockGracePeriodEnforcedSeconds, passcodePresent, carrierSettingsVersion, currentCarrierNetwork, currentMobileCountryCode, currentMobileNetworkCode, dataRoamingEnabled, eid, network, homeMobileCountryCode, homeMobileNetworkCode, iccid, imei, imei2, meid, personalHotspotEnabled, roaming, lastLoggedInUsernameSelfService, lastLoggedInUsernameSelfServiceTimestamp, groupId, groupName

This param can be combined with paging and sorting. Example: filter=displayName=="iPad"

-----



Return paginated Mobile Device Inventory records of all paired devices for the device
get
https://yourServer.jamfcloud.com/api/v2/mobile-devices/{id}/paired-devices


Return paginated Mobile Device Inventory records of all paired devices for the device

Path Params
id
string
required
instance id of mobile device record

Query Params
section
array of strings
Defaults to GENERAL
section of mobile device details, if not specified, Paired Devices section data is returned. Multiple section parameters are supported, e.g. section=GENERAL&section=HARDWARE


string


GENERAL

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
Defaults to displayName:asc
Sorting criteria in the format: property:asc/desc. Default sort is displayName:asc. Multiple sort criteria are supported and must be separated with a comma.

Fields allowed in the sort: airPlayPassword, appAnalyticsEnabled, assetTag, availableSpaceMb, batteryLevel, bluetoothLowEnergyCapable, bluetoothMacAddress, capacityMb, lostModeEnabledDate, declarativeDeviceManagementEnabled, deviceId, deviceLocatorServiceEnabled, devicePhoneNumber, diagnosticAndUsageReportingEnabled, displayName, doNotDisturbEnabled, enrollmentSessionTokenValid, osBuild, osSupplementalBuildVersion, osVersion, osRapidSecurityResponse, ipAddress, itunesStoreAccountActive, mobileDeviceId, languages, lastEnrolledDate, lastCloudBackupDate, lastInventoryUpdateDate, locales, lostModeEnabled, managed, mdmProfileExpirationDate, model, modelIdentifier, modelNumber, modemFirmwareVersion, preferredVoiceNumber, serialNumber, supervised, timeZone, udid, usedSpacePercentage, wifiMacAddress, deviceOwnershipType, building, department, emailAddress, fullName, userPhoneNumber, position, room, username, appleCareId, leaseExpirationDate,lifeExpectancyYears, poDate, poNumber, purchasePrice, purchasedOrLeased, purchasingAccount, purchasingContact, vendor, warrantyExpirationDate, activationLockEnabled, blockEncryptionCapable, dataProtection, fileEncryptionCapable, hardwareEncryptionSupported, jailbreakStatus, passcodeCompliant, passcodeCompliantWithProfile, passcodeLockGracePeriodEnforcedSeconds, passcodePresent, carrierSettingsVersion, cellularTechnology, currentCarrierNetwork, currentMobileCountryCode, currentMobileNetworkCode, dataRoamingEnabled, eid, network, homeMobileCountryCode, homeMobileNetworkCode, iccid, imei, imei2, meid, personalHotspotEnabled, voiceRoamingEnabled, roaming, lastLoggedInUsernameSelfService, lastLoggedInUsernameSelfServiceTimestamp

Example: sort=displayName:desc,username:asc


string

displayName:asc

ADD string
filter
string
Query in the RSQL format, allowing to filter mobile device collection. Default filter is empty query - returning all results for the requested page.

Fields allowed in the query: airPlayPassword, appAnalyticsEnabled, assetTag, availableSpaceMb, batteryLevel, bluetoothLowEnergyCapable, bluetoothMacAddress, capacityMb, declarativeDeviceManagementEnabled, deviceId, deviceLocatorServiceEnabled, devicePhoneNumber, diagnosticAndUsageReportingEnabled, displayName, doNotDisturbEnabled, osBuild, osSupplementalBuildVersion, osVersion, osRapidSecurityResponse, ipAddress, itunesStoreAccountActive, mobileDeviceId, languages, lastInventoryUpdateDate, locales, lostModeEnabled, managed, model, modelIdentifier, modelNumber, modemFirmwareVersion, preferredVoiceNumber, serialNumber, supervised, timeZone, udid, usedSpacePercentage, wifiMacAddress, building, department, emailAddress, fullName, userPhoneNumber, position, room, username, appleCareId, lifeExpectancyYears, poNumber, purchasePrice, purchasedOrLeased, purchasingAccount, purchasingContact, vendor, activationLockEnabled, blockEncryptionCapable, dataProtection, fileEncryptionCapable, passcodeCompliant, passcodeCompliantWithProfile, passcodeLockGracePeriodEnforcedSeconds, passcodePresent, carrierSettingsVersion, currentCarrierNetwork, currentMobileCountryCode, currentMobileNetworkCode, dataRoamingEnabled, eid, network, homeMobileCountryCode, homeMobileNetworkCode, iccid, imei, imei2, meid, personalHotspotEnabled, roaming, lastLoggedInUsernameSelfService, lastLoggedInUsernameSelfServiceTimestamp, groupId, groupName

This param can be combined with paging and sorting. Example: filter=displayName=="iPad"

-----

Get Onboarding history object
get
https://yourServer.jamfcloud.com/api/v1/onboarding/history

Gets Onboarding history object

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
Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: username, date, note, details. This param can be combined with paging and sorting. Example: filter=username!=admin and date<2019-12-15

-----

Retrieve Packages
get
https://yourServer.jamfcloud.com/api/v1/packages

Retrieves packages

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
Filters results. Use RSQL format for query. Allows for many fields, including ID, name, etc. Can be combined with paging and sorting. Fields allowed in the query: id, fileName, packageName, categoryId, info, notes, manifestFileName, cloudTransferStatus. Default filter is an empty query and returns all results from the requested page.

Response

-----

Get Jamf Parent app settings history
get
https://yourServer.jamfcloud.com/api/v1/parent-app/history

Gets Jamf Parent app settings history

Query Params
page
integer
Defaults to 0
0
page-size
integer
Defaults to 100
100
filter
string
Query in the RSQL format, allowing to filter results. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: status, updated, version This param can be combined with paging and sorting. Example: filter=username!=admin and details==disabled and date<2019-12-15

sort
string
Defaults to date:desc
Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc

date:desc

-----

Retrieve Patch Policies
get
https://yourServer.jamfcloud.com/api/v2/patch-policies

Retrieves a list of patch policies.

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
Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma.


string

id:asc

ADD string
filter
string
Query in the RSQL format, allowing to filter Patch Policy collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: id, policyName, policyEnabled, policyTargetVersion, policyDeploymentMethod, softwareTitle, softwareTitleConfigurationId, pending, completed, deferred, and failed. This param can be combined with paging and sorting.

-----

Retrieve Patch Policies
get
https://yourServer.jamfcloud.com/api/v2/patch-policies/policy-details


Retrieves a list of patch policies.

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
Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma.


string

id:asc

ADD string
filter
string
Query in the RSQL format, allowing to filter Patch Policy collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: id, name, enabled, targetPatchVersion, deploymentMethod, softwareTitleId, softwareTitleConfigurationId, killAppsDelayMinutes, killAppsMessage, isDowngrade, isPatchUnknownVersion, notificationHeader, selfServiceEnforceDeadline, selfServiceDeadline, installButtonText, selfServiceDescription, iconId, reminderFrequency, reminderEnabled. This param can be combined with paging and sorting.

-----

Retrieve Patch Policy Logs
get
https://yourServer.jamfcloud.com/api/v2/patch-policies/{id}/logs


Retrieves Patch Policy Logs

Path Params
id
string
required
patch policy id

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
Defaults to deviceName:asc
Sorting criteria in the format: property:asc/desc. Default sort is deviceName:asc. Multiple sort criteria are supported and must be separated with a comma.


string

deviceName:asc

ADD string
filter
string
Query in the RSQL format, allowing to filter Patch Policy Logs collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: deviceId, deviceName, statusCode, statusDate, attemptNumber, ignoredForPatchPolicyId. This param can be combined with paging and sorting.

-----

Retrieve Patch Software Title Definitions with the supplied id
get
https://yourServer.jamfcloud.com/api/v2/patch-software-title-configurations/{id}/definitions


Retrieves patch software title definitions with the supplied id

Path Params
id
string
required
Patch Software Title identifier

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
Defaults to absoluteOrderId:asc
Sorting criteria in the format: property:asc/desc. Default sort is absoluteOrderId:asc. Multiple sort criteria are supported and must be separated with a comma.


string

absoluteOrderId:asc

ADD string
filter
string
Query in the RSQL format, allowing to filter Patch Software Title Definition collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: id, version, minimumOperatingSystem, releaseDate, reboot, standalone and absoluteOrderId. This param can be combined with paging and sorting.

-----

Export Patch Reporting Data
get
https://yourServer.jamfcloud.com/api/v2/patch-software-title-configurations/{id}/export-report


Export Patch Reporting Data

Path Params
id
string
required
Patch Software Title Configurations identifier

Query Params
filter
string
Query in the RSQL format, allowing to filter Patch Report collection on version equality only. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: version. Comparators allowed in the query: ==, != This param can be combined with paging and sorting.

columns-to-export
array of strings
required
Defaults to computerName,deviceId,username,operatingSystemVersion,lastContactTime,buildingName,departmentName,siteName,version
List of column names to export


string

computerName

string

deviceId

string

username

string

operatingSystemVersion

string

lastContactTime

string

buildingName

string

departmentName

string

siteName

string

version

ADD string
Headers
accept
string
File

-----

Get specified Patch Software Title Configuration history object
get
https://yourServer.jamfcloud.com/api/v2/patch-software-title-configurations/{id}/history


Gets specified Patch Software Title Configuration history object

Path Params
id
string
required
Patch Software Title Configuration Id

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


string

date:desc

ADD string
filter
string
Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: username, date, note, details. This param can be combined with paging and sorting. Example: filter=username!=admin and details==disabled and date<2019-12-15

-----

Retrieve Patch Software Title Configuration Patch Report
get
https://yourServer.jamfcloud.com/api/v2/patch-software-title-configurations/{id}/patch-report


Retrieve Patch Software Title Configuration Patch Report

Path Params
id
string
required
Patch Software Title Configurations identifier

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
Defaults to computerName:asc
Sorting criteria in the format: property:asc/desc. Default sort is computerName:asc. Multiple sort criteria are supported and must be separated with a comma. Supported fields: computerName, deviceId, username, operatingSystemVersion, lastContactTime, buildingName, departmentName, siteName, version


string

computerName:asc

ADD string
filter
string
Query in the RSQL format, allowing to filter Patch Report collection on version equality only. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: version. Comparators allowed in the query: ==, != This param can be combined with paging and sorting.

-----

Retrieve all triggers for a Jamf Pro Scheduler job
get
https://yourServer.jamfcloud.com/api/v1/scheduler/jobs/{jobKey}/triggers


Retrieves all triggers for a Jamf Pro Scheduler job

Path Params
jobKey
string
required
Jamf Pro Scheduler Job Key

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
Defaults to nextFireTime:asc
Sorts results by one or more criteria, following the format property:asc/desc. Default sort is nextFireTime:asc. If using multiple criteria, separate with commas.


string

nextFireTime:asc

ADD string
filter
string
Query in the RSQL format, allowing to filter the Jamf Pro Scheduler triggers collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: triggerKey, previousFireTime, nextFireTime.

-----

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

-----

Get a page of Self Service settings history
get
https://yourServer.jamfcloud.com/api/v1/self-service/settings/history


Get a page of Self Service settings history

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
Defaults to
Sorting criteria in the format: property:asc/desc. Default sort order is descending. Multiple sort criteria are supported and must be entered on separate lines in Swagger UI. In the URI the 'sort' query param is not duplicated for each sort criterion, e.g., ...&sort=name:asc,date:desc. Fields that can be sorted: status, updated


ADD string
filter
string
Query in the RSQL format, allowing to filter results. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: status, updated, version This param can be combined with paging and sorting. Example: filter=username!=admin and details==disabled and date<2019-12-15

Response

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

-----

Get specified SMTP Server history object
get
https://yourServer.jamfcloud.com/api/v1/smtp-server/history

Get specified SMTP Server history object

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
Sorting criteria in the format: property:asc/desc. Default sort is name:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,username:asc


string

date:desc

ADD string
filter
string
Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: username, date, note, details. This param can be combined with paging and sorting. Example: filter=username!=admin and details==disabled and date<2019-12-15

-----

Get SSO history object
get
https://yourServer.jamfcloud.com/api/v3/sso/history

Gets SSO history object

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
Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc


string

id:desc

ADD string
filter
string
Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: username, date, note, details. This param can be combined with paging and sorting. Example: filter=username!=admin and details==disabled and date<2019-12-15

-----


Get Jamf Teacher app settings history
get
https://yourServer.jamfcloud.com/api/v1/teacher-app/history

Gets Jamf Teacher app settings history

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
Defaults to
Sorting criteria in the format: property:asc/desc. Default sort order is descending. Multiple sort criteria are supported and must be entered on separate lines in Swagger UI. In the URI the 'sort' query param is not duplicated for each sort criterion, e.g., ...&sort=name:asc,date:desc. Fields that can be sorted: status, updated


ADD string
filter
string
Query in the RSQL format, allowing to filter results. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: status, updated, version This param can be combined with paging and sorting. Example: filter=username!=admin and details==disabled and date<2019-12-15


-----

Get a paginated list of sessions
get
https://yourServer.jamfcloud.com/api/preview/remote-administration-configurations/team-viewer/{configurationId}/sessions


Returns a paginated list of sessions for a given configuration ID

Path Params
configurationId
string
required
ID of the Team Viewer connection configuration

Query Params
page
integer
Defaults to 0
0
page-size
integer
Defaults to 100
100
filter
string
Query in the RSQL format, allowing to filter sessions collection. Default filter is empty query - returning all results for the requested page.

Fields allowed in the query: deviceId, deviceType, state

This param can be combined with paging.

Responses

-----

Get specified Venafi CA history object
get
https://yourServer.jamfcloud.com/api/v1/pki/venafi/{id}/history


Get specified Venafi CA history object

Path Params
id
string
required
ID of the Venafi configuration

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

-----

Retrieve Volume Purchasing Locations
get
https://yourServer.jamfcloud.com/api/v1/volume-purchasing-locations


Retrieves Volume Purchasing Locations

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
Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma.


string

id:asc

ADD string
filter
string
Query in the RSQL format, allowing to filter Volume Purchasing Location collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: id, name, appleId, email, organizationName, tokenExpiration, countryCode, locationName, automaticallyPopulatePurchasedContent, sendNotificationWhenNoLongerAssigned, siteId and siteName. This param can be combined with paging and sorting.

-----

Retrieve the Volume Purchasing Content for the Volume Purchasing Location with the supplied id
get
https://yourServer.jamfcloud.com/api/v1/volume-purchasing-locations/{id}/content


Retrieves the Volume Purchasing Content for the Volume Purchasing Location with the supplied id

Path Params
id
string
required
Volume Purchasing Location identifier

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
Sorting criteria in the format: property:asc/desc. Default sort is name:asc. Multiple sort criteria are supported and must be separated with a comma.


string

id:asc

ADD string
filter
string
Query in the RSQL format, allowing to filter Volume Purchasing Content collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: name, licenseCountTotal, licenseCountInUse, licenseCountReported, contentType, and pricingParam. This param can be combined with paging and sorting.

-----

Get specified Volume Purchasing Location history object
get
https://yourServer.jamfcloud.com/api/v1/volume-purchasing-locations/{id}/history


Gets specified Volume Purchasing Location history object

Path Params
id
string
required
instance id of Volume Purchasing Location history record

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


string

date:desc

ADD string
filter
string
Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: username, date, note, details. This param can be combined with paging and sorting. Example: filter=username!=admin and details==disabled and date<2019-12-15

