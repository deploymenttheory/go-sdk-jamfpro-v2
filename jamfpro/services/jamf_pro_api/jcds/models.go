package jcds

import "io"

// ResourceJCDSFile represents a file stored in Jamf Cloud Distribution Service.
type ResourceJCDSFile struct {
	FileName string `json:"fileName"` // The name of the file
	Length   int64  `json:"length"`   // The size of the file in bytes
	MD5      string `json:"md5"`      // The MD5 hash of the file
	Region   string `json:"region"`   // The AWS region where the file is stored
	SHA3     string `json:"sha3"`     // The SHA3 hash of the file
}

// ResourceJCDSUploadCredentials contains AWS credentials for uploading to JCDS.
type ResourceJCDSUploadCredentials struct {
	AccessKeyID     string `json:"accessKeyID"`     // AWS access key ID
	SecretAccessKey string `json:"secretAccessKey"` // AWS secret access key
	SessionToken    string `json:"sessionToken"`    // AWS session token
	Region          string `json:"region"`          // AWS region
	BucketName      string `json:"bucketName"`      // S3 bucket name
	Path            string `json:"path"`            // S3 path prefix
	UUID            string `json:"uuid"`            // Unique identifier
}

// ResponseJCDSFile contains the URI of a file in JCDS.
type ResponseJCDSFile struct {
	URI string `json:"uri"` // S3 URI of the file
}

// ProgressReader wraps an io.Reader to report progress on read operations.
type ProgressReader struct {
	reader     io.Reader
	totalBytes int64
	readBytes  int64
	progressFn func(readBytes, totalBytes int64, unit string)
}

// Read implements the io.Reader interface for progress tracking.
func (r *ProgressReader) Read(p []byte) (int, error) {
	n, err := r.reader.Read(p)
	r.readBytes += int64(n)

	const kb = 1024
	const mb = 1024 * kb

	if r.totalBytes > mb { // report in MB if file is larger than 1MB
		r.progressFn(r.readBytes/mb, r.totalBytes/mb, "MB")
	} else { // otherwise, report in KB
		r.progressFn(r.readBytes/kb, r.totalBytes/kb, "KB")
	}

	return n, err
}
