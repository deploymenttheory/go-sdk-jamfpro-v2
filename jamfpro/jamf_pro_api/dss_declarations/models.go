package dss_declarations

// ResponseDSSDeclaration represents the response containing DSS declarations.
type ResponseDSSDeclaration struct {
	Declarations []ResourceDSSDeclaration `json:"declarations"`
}

// ResourceDSSDeclaration represents a single DSS declaration.
type ResourceDSSDeclaration struct {
	UUID        string `json:"uuid"`
	PayloadJson string `json:"payloadJson"`
	Type        string `json:"type"`
	Group       string `json:"group"`
}
