Get Cache Settings
get
https://yourServer.jamfcloud.com/api/v1/cache-settings

gets cache settings

Response

200
cache settings

Response body
object
id
string
length ≥ 1
Defaults to 0
name
string
Defaults to cache configuration
cacheType
string
required
timeToLiveSeconds
int32
required
timeToIdleSeconds
int32
directoryTimeToLiveSeconds
int32
ehcacheMaxBytesLocalHeap
string
Defaults to null
cacheUniqueId
string
required
The default is for Jamf Pro to generate a UUID, so we can only give an example instead.

elasticache
boolean
Defaults to false
memcachedEndpoints
array of objects
required
object
id
string
length ≥ 1
name
string
hostName
string
port
integer
enabled
boolean
jssCacheConfigurationId
integer

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/cache-settings \
     --header 'accept: application/json'

{
  "id": "0",
  "name": "cache configuration",
  "cacheType": "ehcache",
  "timeToLiveSeconds": 120,
  "timeToIdleSeconds": 120,
  "directoryTimeToLiveSeconds": 120,
  "ehcacheMaxBytesLocalHeap": "null",
  "cacheUniqueId": "24864549-94ea-4cc1-bb80-d7fb392c6556",
  "elasticache": false,
  "memcachedEndpoints": []
}

-----

Update Cache Settings
put
https://yourServer.jamfcloud.com/api/v1/cache-settings

updates cache settings

Body Params
name
string
Defaults to cache configuration
cache configuration
cacheType
string
required
timeToLiveSeconds
int32
required
timeToIdleSeconds
int32
directoryTimeToLiveSeconds
int32
ehcacheMaxBytesLocalHeap
string
Defaults to null
null
cacheUniqueId
string
required
The default is for Jamf Pro to generate a UUID, so we can only give an example instead.

elasticache
boolean
Defaults to false

false
memcachedEndpoints
array of objects
required

ADD object
Responses

200
cache has been updated

Response body
object
id
string
length ≥ 1
Defaults to 0
name
string
Defaults to cache configuration
cacheType
string
required
timeToLiveSeconds
int32
required
timeToIdleSeconds
int32
directoryTimeToLiveSeconds
int32
ehcacheMaxBytesLocalHeap
string
Defaults to null
cacheUniqueId
string
required
The default is for Jamf Pro to generate a UUID, so we can only give an example instead.

elasticache
boolean
Defaults to false
memcachedEndpoints
array of objects
required
object
id
string
length ≥ 1
name
string
hostName
string
port
integer
enabled
boolean
jssCacheConfigurationId
integer

curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v1/cache-settings \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "name": "cache configuration",
  "ehcacheMaxBytesLocalHeap": "null",
  "elasticache": false
}
'

{
  "id": "0",
  "name": "cache configuration",
  "cacheType": "ehcache",
  "timeToLiveSeconds": 120,
  "timeToIdleSeconds": 120,
  "directoryTimeToLiveSeconds": 120,
  "ehcacheMaxBytesLocalHeap": "null",
  "cacheUniqueId": "24864549-94ea-4cc1-bb80-d7fb392c6556",
  "elasticache": false,
  "memcachedEndpoints": []
}