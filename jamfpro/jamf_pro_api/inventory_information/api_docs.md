Get statistics about managed/unmanaged devices and computers in the inventory
get
https://yourServer.jamfcloud.com/api/v1/inventory-information


Gets statistics about managed/unmanaged devices and computers in the inventory.

Response

200
Successful response

Response body
object
managedComputers
integer
≥ 0
Number of managed computers in inventory.

unmanagedComputers
integer
≥ 0
Number of unmanaged computers in inventory.

managedDevices
integer
≥ 0
Number of managed devices in inventory.

unmanagedDevices
integer
≥ 0
Number of unmanaged devices in inventory.

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/inventory-information \
     --header 'accept: application/json'

{
  "managedComputers": 1200,
  "unmanagedComputers": 1100,
  "managedDevices": 1200,
  "unmanagedDevices": 1100
}