package jcds

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"resty.dev/v3"
)

// ServiceInterface defines the interface for Jamf Cloud Distribution Service (JCDS) operations.
//
// JCDS is Jamf's cloud-based content delivery service using AWS S3 for storing
// and distributing packages to managed devices. This service provides operations
// for uploading, downloading, and managing packages in JCDS.
//
// Reference: https://grahamrpugh.com/2023/08/21/introducing-jcds2.html
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-jcds-files
type ServiceInterface interface {
	// GetPackagesV1 retrieves a list of all files stored in JCDS.
	//
	// Returns metadata for all packages including filename, size, MD5, SHA3, and region.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-jcds-files
	GetPackagesV1(ctx context.Context) ([]ResourceJCDSFile, *resty.Response, error)

	// GetPackageURIByNameV1 retrieves the S3 URI for a specific package by name.
	//
	// Returns the S3 URI that can be used to access the package directly.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-jcds-files-filename
	GetPackageURIByNameV1(ctx context.Context, packageName string) (*ResponseJCDSFile, *resty.Response, error)

	// RenewCredentialsV1 obtains fresh AWS credentials for JCDS operations.
	//
	// Returns temporary AWS credentials (access key, secret key, session token)
	// that can be used to perform S3 operations on the JCDS bucket.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-jcds-renew-credentials
	RenewCredentialsV1(ctx context.Context) (*ResourceJCDSUploadCredentials, *resty.Response, error)

	// CreatePackageV1 uploads a package file to JCDS using AWS S3.
	//
	// This is a multi-step operation:
	// 1. Obtains temporary AWS credentials from Jamf Pro
	// 2. Configures AWS SDK with those credentials
	// 3. Uploads the file to S3 with progress tracking
	// 4. Returns the S3 URI of the uploaded file
	//
	// The filePath must be an absolute path to a .pkg, .dmg, or .zip file.
	// Progress is reported to stdout during upload.
	//
	// Note: This operation uploads the file but does not create package metadata
	// in Jamf Pro. Use the Packages service to register the package after upload.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-jcds-files
	CreatePackageV1(ctx context.Context, filePath string) (*ResponseJCDSFile, *resty.Response, error)

	// DeletePackageV1 deletes a package file from JCDS using AWS S3.
	//
	// This is a multi-step operation:
	// 1. Obtains temporary AWS credentials from Jamf Pro
	// 2. Configures AWS SDK with those credentials
	// 3. Deletes the file from S3
	//
	// The filePath should be the same path used during upload. Only the filename
	// is used for deletion.
	//
	// Note: This operation deletes the file from S3 but does not remove package
	// metadata from Jamf Pro. Use the Packages service to manage package records.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-jcds-files (for credentials)
	DeletePackageV1(ctx context.Context, filePath string) (*resty.Response, error)

	// RefreshInventoryV1 triggers Jamf Pro to refresh its inventory of JCDS packages.
	//
	// This endpoint should be called after uploading or deleting files from JCDS
	// to ensure Jamf Pro's package inventory is synchronized with S3.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-jcds-refresh-inventory
	RefreshInventoryV1(ctx context.Context) (*resty.Response, error)
}

type (
	// Service handles communication with the JCDS-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-jcds-files
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ ServiceInterface = (*Service)(nil)

// NewService creates a new JCDS service.
func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// GetPackagesV1 retrieves a list of all files stored in JCDS.
// URL: GET /api/v1/jcds/files
// https://developer.jamf.com/jamf-pro/reference/get_v1-jcds-files
func (s *Service) GetPackagesV1(ctx context.Context) ([]ResourceJCDSFile, *resty.Response, error) {
	endpoint := EndpointJCDSV1 + "/files"

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	var result []ResourceJCDSFile
	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get JCDS packages: %w", err)
	}

	return result, resp, nil
}

// GetPackageURIByNameV1 retrieves the S3 URI for a specific package by name.
// URL: GET /api/v1/jcds/files/{filename}
// https://developer.jamf.com/jamf-pro/reference/get_v1-jcds-files-filename
func (s *Service) GetPackageURIByNameV1(ctx context.Context, packageName string) (*ResponseJCDSFile, *resty.Response, error) {
	if packageName == "" {
		return nil, nil, fmt.Errorf("package name is required")
	}

	endpoint := fmt.Sprintf("%s/files/%s", EndpointJCDSV1, packageName)

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	var result ResponseJCDSFile
	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get JCDS package URI for '%s': %w", packageName, err)
	}

	return &result, resp, nil
}

// RenewCredentialsV1 obtains fresh AWS credentials for JCDS operations.
// URL: POST /api/v1/jcds/renew-credentials
// https://developer.jamf.com/jamf-pro/reference/post_v1-jcds-renew-credentials
func (s *Service) RenewCredentialsV1(ctx context.Context) (*ResourceJCDSUploadCredentials, *resty.Response, error) {
	endpoint := EndpointJCDSV1 + "/renew-credentials"

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	var result ResourceJCDSUploadCredentials
	resp, err := s.client.Post(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to renew JCDS credentials: %w", err)
	}

	return &result, resp, nil
}

// CreatePackageV1 uploads a package file to JCDS using AWS S3.
// URL: POST /api/v1/jcds/files (for credentials) + AWS S3 upload
// https://developer.jamf.com/jamf-pro/reference/post_v1-jcds-files
func (s *Service) CreatePackageV1(ctx context.Context, filePath string) (*ResponseJCDSFile, *resty.Response, error) {
	if filePath == "" {
		return nil, nil, fmt.Errorf("file path is required")
	}

	// Step 1: Obtain AWS credentials for the package upload endpoint
	var uploadCredentials ResourceJCDSUploadCredentials
	endpoint := EndpointJCDSV1 + "/files"

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, nil, headers, &uploadCredentials)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to obtain upload credentials: %w", err)
	}

	// Validate if we received necessary details
	if uploadCredentials.Region == "" || uploadCredentials.BucketName == "" || uploadCredentials.Path == "" {
		return nil, resp, fmt.Errorf("incomplete upload credentials received")
	}

	// Step 2: Use the obtained credentials to configure AWS SDK v2
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(uploadCredentials.Region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			uploadCredentials.AccessKeyID,
			uploadCredentials.SecretAccessKey,
			uploadCredentials.SessionToken,
		)),
	)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to create AWS config: %w", err)
	}

	// Create S3 service client
	s3Client := s3.NewFromConfig(cfg)

	// Step 3: Create an Uploader with the configuration and default options
	uploader := manager.NewUploader(s3Client)

	// Step 4: Use the secure file reading helper
	fileReader, fileSize, err := readJCDSPackageTypes(filePath)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to read package file securely: %w", err)
	}

	// Create a progress reader
	progressReader := &ProgressReader{
		reader:     fileReader,
		totalBytes: fileSize,
		progressFn: func(read, total int64, unit string) {
			fmt.Printf("\rUploaded %d / %d %s (%.2f%%)", read, total, unit, float64(read)/float64(total)*100)
		},
	}

	// Create the upload input
	uploadInput := &s3.PutObjectInput{
		Bucket: aws.String(uploadCredentials.BucketName),
		Key:    aws.String(uploadCredentials.Path + filepath.Base(filePath)),
		Body:   progressReader,
	}

	// Step 5: Perform the upload
	_, err = uploader.Upload(ctx, uploadInput)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to upload file: %w", err)
	}

	fmt.Println("\nUpload completed successfully")

	// Construct the final file upload response
	finalResponse := &ResponseJCDSFile{
		URI: fmt.Sprintf("s3://%s/%s%s", uploadCredentials.BucketName, uploadCredentials.Path, filepath.Base(filePath)),
	}

	return finalResponse, resp, nil
}

// DeletePackageV1 deletes a package file from JCDS using AWS S3.
// URL: POST /api/v1/jcds/files (for credentials) + AWS S3 delete
// https://developer.jamf.com/jamf-pro/reference/post_v1-jcds-files
func (s *Service) DeletePackageV1(ctx context.Context, filePath string) (*resty.Response, error) {
	if filePath == "" {
		return nil, fmt.Errorf("file path is required")
	}

	// Step 1: Obtain AWS credentials for the package deletion endpoint
	var uploadCredentials ResourceJCDSUploadCredentials
	endpoint := EndpointJCDSV1 + "/files"

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, nil, headers, &uploadCredentials)
	if err != nil {
		return resp, fmt.Errorf("failed to obtain deletion credentials: %w", err)
	}

	// Validate if we received necessary details
	if uploadCredentials.Region == "" || uploadCredentials.BucketName == "" || uploadCredentials.Path == "" {
		return resp, fmt.Errorf("incomplete deletion credentials received")
	}

	// Step 2: Use the obtained credentials to configure AWS SDK v2
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(uploadCredentials.Region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			uploadCredentials.AccessKeyID,
			uploadCredentials.SecretAccessKey,
			uploadCredentials.SessionToken,
		)),
	)
	if err != nil {
		return resp, fmt.Errorf("failed to create AWS config: %w", err)
	}

	// Create S3 service client
	s3Client := s3.NewFromConfig(cfg)

	// Step 3: Define the object to delete
	objectToDelete := &s3.DeleteObjectInput{
		Bucket: aws.String(uploadCredentials.BucketName),
		Key:    aws.String(uploadCredentials.Path + filepath.Base(filePath)),
	}

	// Step 4: Perform the deletion
	_, err = s3Client.DeleteObject(ctx, objectToDelete)
	if err != nil {
		return resp, fmt.Errorf("failed to delete file: %w", err)
	}

	fmt.Printf("File '%s' successfully deleted from JCDS\n", filepath.Base(filePath))
	return resp, nil
}

// RefreshInventoryV1 triggers Jamf Pro to refresh its inventory of JCDS packages.
// URL: POST /api/v1/jcds/refresh-inventory
// https://developer.jamf.com/jamf-pro/reference/post_v1-jcds-refresh-inventory
func (s *Service) RefreshInventoryV1(ctx context.Context) (*resty.Response, error) {
	endpoint := EndpointJCDSV1 + "/refresh-inventory"

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, fmt.Errorf("failed to refresh JCDS inventory: %w", err)
	}

	return resp, nil
}
