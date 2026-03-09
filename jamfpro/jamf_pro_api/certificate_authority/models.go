package certificate_authority

// ResourceActiveCertificateAuthorityV1 represents the active certificate authority response.
type ResourceActiveCertificateAuthorityV1 struct {
	SubjectX500Principal string                                 `json:"subjectX500Principal"`
	IssuerX500Principal  string                                 `json:"issuerX500Principal"`
	SerialNumber         string                                 `json:"serialNumber"`
	Version              int                                    `json:"version"`
	NotAfter             int64                                  `json:"notAfter"`
	NotBefore            int64                                  `json:"notBefore"`
	Signature            ActiveCertificateAuthoritySignatureV1  `json:"signature"`
	KeyUsage             []string                               `json:"keyUsage"`
	KeyUsageExtended     []string                               `json:"keyUsageExtended"`
	SHA1Fingerprint      string                                 `json:"sha1Fingerprint"`
	SHA256Fingerprint    string                                 `json:"sha256Fingerprint"`
}

// ActiveCertificateAuthoritySignatureV1 represents the signature subset of the CA resource.
type ActiveCertificateAuthoritySignatureV1 struct {
	Algorithm    string `json:"algorithm"`
	AlgorithmOID string `json:"algorithmOid"`
	Value        string `json:"value"`
}
