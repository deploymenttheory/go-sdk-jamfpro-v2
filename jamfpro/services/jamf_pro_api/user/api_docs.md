Changes the user account password.
post
https://yourServer.jamfcloud.com/api/v1/user/change-password

Changes the account password for a currently authenticated user.

Body Params
Current account password and new password.

currentPassword
password
required
•••••••••••
newPassword
password
required
•••••••••••••••
Responses
202
Password for user successfully changed.

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/user/change-password \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "currentPassword": "oldPassword",
  "newPassword": "updatedPassword"
}
'

