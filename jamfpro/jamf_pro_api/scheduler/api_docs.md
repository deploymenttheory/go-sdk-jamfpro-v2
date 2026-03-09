Retrieve all Jamf Pro Scheduler jobs
get
https://yourServer.jamfcloud.com/api/v1/scheduler/jobs

Retrieves the names of all Jamf Pro Scheduler jobs

Response

200
Success

Response body
object
jobKeys
array of strings

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/scheduler/jobs \
     --header 'accept: application/json'

{
  "jobKeys": [
    "EXAMPLE_JOB"
  ]
}
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

Response

200
Success

Response body
object
totalCount
integer
results
array of objects
object
triggerKey
string
previousFireTime
date-time
nextFireTime
date-time

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v1/scheduler/jobs//triggers?page=0&page-size=100&sort=nextFireTime%3Aasc' \
     --header 'accept: application/json'

{
  "totalCount": 1,
  "results": [
    {
      "triggerKey": "EXAMPLE_TRIGGER",
      "previousFireTime": "2023-08-21T21:20:34.35Z",
      "nextFireTime": "2023-08-21T21:30:34.35Z"
    }
  ]
}
-----
Retrieve a summary of the Jamf Pro Scheduler
get
https://yourServer.jamfcloud.com/api/v1/scheduler/summary

Retrieves a summary of the Jamf Pro Scheduler

Response

200
Success

Response body
object
numberOfPendingJobs
integer
numberOfExecutingJobs
integer
numberOfExecutedJobs
integer
started
boolean

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/scheduler/summary \
     --header 'accept: application/json'

{
  "numberOfPendingJobs": 1,
  "numberOfExecutingJobs": 1,
  "numberOfExecutedJobs": 1,
  "started": true
}
-----