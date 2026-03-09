Reinstall App Config for Managed iOS Apps
post
https://yourServer.jamfcloud.com/api/v1/mobile-device-apps/reinstall-app-config

Redeploys the managed app configuration for a specific app on a specific device using the $APP_CONFIG_REINSTALL_CODE generated during deployment. This endpoint does not require authorization, only the re-install code. The code does not contain any user authentication information. For example usage, see the following Teacher app documentation: Teacher App Manged App Configuration

Body Params
The $APP_CONFIG_REINSTALL_CODE variable for the specific device and app supplied by the managed iOS app's current App Config.

reinstallCode
string
975767FE-074E-4F42-BB8B-925B1627CA6F
Responses
204
Install App Config command successfully issued

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/mobile-device-apps/reinstall-app-config \
     --header 'content-type: application/json' \
     --data '
{
  "reinstallCode": "975767FE-074E-4F42-BB8B-925B1627CA6F"
}
'