package jcds

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// readJCDSPackageTypes returns a reader and size for a package file securely after applying multiple checks.
func readJCDSPackageTypes(filePath string) (io.Reader, int64, error) {
	allowedExtensions := []string{".pkg", ".dmg", ".zip"}

	data, err := safeReadJCDSPackageFile(filePath, allowedExtensions)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to read package file securely: %v", err)
	}

	size := int64(len(data))
	reader := bytes.NewReader(data)

	return reader, size, nil
}

// safeReadJCDSPackageFile reads a package file securely after applying multiple checks.
func safeReadJCDSPackageFile(filePath string, allowedExtensions []string) ([]byte, error) {
	cleanedPath := cleanPath(filePath)

	if !isValidExtension(cleanedPath, allowedExtensions) {
		return nil, fmt.Errorf("file extension '%s' is not allowed", filepath.Ext(cleanedPath))
	}

	resolvedPath, err := resolveSymlinks(cleanedPath)
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(resolvedPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read Jamf Pro package: %v", err)
	}
	return data, nil
}

// resolveSymlinks resolves symbolic links and returns the absolute path.
func resolveSymlinks(filePath string) (string, error) {
	cleanPath := filepath.Clean(filePath)
	absPath, err := filepath.EvalSymlinks(cleanPath)
	if err != nil {
		return "", fmt.Errorf("unable to resolve the absolute path: %s, error: %w", filePath, err)
	}
	return absPath, nil
}

// cleanPath sanitizes the file path to prevent directory traversal.
func cleanPath(filePath string) string {
	return filepath.Clean(filePath)
}

// isValidExtension checks if the file has one of the allowed extensions.
func isValidExtension(filePath string, allowedExtensions []string) bool {
	ext := strings.ToLower(filepath.Ext(filePath))
	for _, allowedExt := range allowedExtensions {
		if ext == strings.ToLower(allowedExt) {
			return true
		}
	}
	return false
}
