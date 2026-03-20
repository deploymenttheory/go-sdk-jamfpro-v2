package crypto

import (
	"bytes"
	"testing"
)

func TestStripSignature_AlreadyUnsigned(t *testing.T) {
	unsignedData := []byte(`<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>Test</key>
	<string>Data</string>
</dict>
</plist>`)

	_, err := StripSignature(unsignedData)
	if err == nil {
		t.Fatal("Expected error for already unsigned data")
	}

	if err != ErrAlreadyUnsigned {
		t.Errorf("Expected ErrAlreadyUnsigned, got %v", err)
	}
}

func TestStripSignature_PEMFormat(t *testing.T) {
	pemData := []byte(`-----BEGIN PKCS7-----
MIIBogYJKoZIhvcNAQcCoIIBkzCCAY8CAQExADALBgkqhkiG9w0BBwGgggFvMIIB
azCCARGgAwIBAgIBADAKBggqhkjOPQQDAjAaMRgwFgYDVQQDDA9UZXN0IENlcnRp
ZmljYXRlMB4XDTI2MDMxMDAwMDAwMFoXDTI3MDMxMDAwMDAwMFowGjEYMBYGA1UE
AwwPVGVzdCBDZXJ0aWZpY2F0ZTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABPQR
-----END PKCS7-----`)

	_, err := StripSignature(pemData)

	if err == ErrAlreadyUnsigned {
		t.Error("Should not return ErrAlreadyUnsigned for PEM data")
	}
}

func TestStripSignature_DERFormat(t *testing.T) {
	derData := []byte{0x30, 0x82, 0x01, 0x00}

	_, err := StripSignature(derData)

	if err == ErrAlreadyUnsigned {
		t.Error("Should not return ErrAlreadyUnsigned for DER data")
	}
}

func TestStripSignature_DetectsSignatureMarkers(t *testing.T) {
	tests := []struct {
		name        string
		data        []byte
		shouldError bool
		expectedErr error
	}{
		{
			name:        "unsigned_xml",
			data:        []byte(`<?xml version="1.0"?><plist></plist>`),
			shouldError: true,
			expectedErr: ErrAlreadyUnsigned,
		},
		{
			name:        "pem_marker",
			data:        []byte("-----BEGIN PKCS7-----\ndata\n-----END PKCS7-----"),
			shouldError: false,
		},
		{
			name:        "der_indefinite",
			data:        []byte{0x30, 0x80, 0x00, 0x00},
			shouldError: false,
		},
		{
			name:        "der_definite",
			data:        []byte{0x30, 0x82, 0x00, 0x00},
			shouldError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := StripSignature(tt.data)

			if tt.shouldError && err != tt.expectedErr {
				t.Errorf("Expected error %v, got %v", tt.expectedErr, err)
			}

			if !tt.shouldError && err == ErrAlreadyUnsigned {
				t.Errorf("Should not return ErrAlreadyUnsigned for signed data")
			}
		})
	}
}

func TestStripSignature_PreservesContent(t *testing.T) {
	originalData := []byte(`<?xml version="1.0" encoding="UTF-8"?>
<plist version="1.0">
<dict>
	<key>PayloadDisplayName</key>
	<string>Test</string>
</dict>
</plist>`)

	_, err := StripSignature(originalData)
	if err != ErrAlreadyUnsigned {
		t.Errorf("Expected ErrAlreadyUnsigned, got %v", err)
	}

	if !bytes.Contains(originalData, []byte("PayloadDisplayName")) {
		t.Error("Original data was modified")
	}
}
