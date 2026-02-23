1. ensure that all api function have an example
2. ensure that all functions have correct comment styling and valid
doc url.
3. ensure that all unit tests are named TestUnit_<ServiceName>_<FunctionName>
4. ensure that all acceptance tests are named TestAcceptance_<ServiceName>_<FunctionName>
5. ensure that all List functions that support pagination use the GetPagination transport method
6. ensure that jamf pro api functions that support rsqlQuery have implementation
7. Ensure that all unit tests for List and Get use externised json for mock
responses (jamf pro api) and xml (classic api) and are maximal in field definition to validate the full data model
8. ensure that all services are registered in new.go
9. ensure that all classic api endpoint constants are named EndpointClassic<Service_Name>
10. Ensure that all headers are accurate and relfect the documentation for each function.
11. Ensure that all examples for all responses unmarshall into either json or xml to show the full response body.
12. Ensure that all unit tests pass
13. Ensure that all acceptance tests pass or, have handlers in place when the service is not enabled in the jamf pro tenant. with comments.
14. ensure that all accpetance tests follow one of the pre-defined acc test
strategies.
15. Ensure that all jamf pro api fuction describe the http method verb with
the api version. e.g GetByIDV1 Or CreateV2. Classic api doesnt have versions so just a verb is okay. when there's multipe functions that don't fix this patten. use concise naming aligned with the api endpoint.