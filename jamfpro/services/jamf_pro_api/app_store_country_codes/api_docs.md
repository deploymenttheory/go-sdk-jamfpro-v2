Return a list of Countries and the associated Codes
get
https://yourServer.jamfcloud.com/api/v1/app-store-country-codes


Returns a list of countries and the associated codes that can be use for the App Store locale

Response

200
Successful response

Response body
object
countryCodes
array of objects
object
code
string
name
string

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/app-store-country-codes \
     --header 'accept: application/json'

{
  "countryCodes": [
    {
      "code": "US",
      "name": "United States"
    }
  ]
}