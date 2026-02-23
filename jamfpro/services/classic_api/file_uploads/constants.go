package file_uploads

// EndpointFileUploads is the Classic API endpoint for file uploads.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/fileuploads
const EndpointFileUploads = "/JSSResource/fileuploads"

// ValidFileUploadResources contains the list of valid resources for file uploads.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/fileuploads
var ValidFileUploadResources = []string{
	"computers",
	"mobiledevices",
	"enrollmentprofiles",
	"printers",
	"peripherals",
	"policies",
	"ebooks",
	"mobiledeviceapplications",
	"icon",
	"mobiledeviceapplicationsipa",
	"diskencryptionconfigurations",
}
