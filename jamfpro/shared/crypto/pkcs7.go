package crypto

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
)

var (
	ErrAlreadyUnsigned          = errors.New("data is already unsigned")
	ErrSignatureStrippingFailed = errors.New("signature stripping failed")
)

// StripSignature removes CMS/PKCS#7 signatures from mobileconfig data.
func StripSignature(data []byte) ([]byte, error) {
	if !bytes.HasPrefix(data, []byte("-----BEGIN PKCS7-----")) &&
		!bytes.HasPrefix(data, []byte{0x30, 0x80}) &&
		!bytes.HasPrefix(data, []byte{0x30, 0x82}) {
		return nil, ErrAlreadyUnsigned
	}

	cmd := exec.Command("openssl", "smime", "-verify", "-noverify", "-inform", "DER")
	cmd.Stdin = bytes.NewReader(data)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		cmd = exec.Command("openssl", "smime", "-verify", "-noverify", "-inform", "PEM")
		cmd.Stdin = bytes.NewReader(data)
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr

		if err := cmd.Run(); err != nil {
			return nil, fmt.Errorf("%w: %v", ErrSignatureStrippingFailed, stderr.String())
		}
	}

	return stdout.Bytes(), nil
}
