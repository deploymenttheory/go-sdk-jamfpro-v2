package crypto

import (
	"crypto/md5"
	"crypto/sha3"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

// CalculateSHA3_512 calculates the SHA3-512 hash of the file at filePath.
func CalculateSHA3_512(filePath string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("open file for SHA3-512: %w", err)
	}
	defer f.Close()

	h := sha3.New512()
	if _, err := io.Copy(h, f); err != nil {
		return "", fmt.Errorf("calculate SHA3-512: %w", err)
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

// CalculateMD5 calculates the MD5 hash of the file at filePath.
func CalculateMD5(filePath string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("open file for MD5: %w", err)
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", fmt.Errorf("calculate MD5: %w", err)
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
