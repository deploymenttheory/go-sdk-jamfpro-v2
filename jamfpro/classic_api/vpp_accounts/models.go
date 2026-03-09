package vpp_accounts

import "encoding/xml"

// ResourceVPPAccount represents a Jamf Pro Classic API VPP account resource.
type ResourceVPPAccount struct {
	XMLName                       xml.Name `xml:"vpp_account"`
	ID                            int      `xml:"id"`
	Name                          string   `xml:"name"`
	Contact                       string   `xml:"contact,omitempty"`
	ServiceToken                  string   `xml:"service_token,omitempty"`
	AccountName                   string   `xml:"account_name,omitempty"`
	ExpirationDate                string   `xml:"expiration_date,omitempty"`
	Country                       string   `xml:"country,omitempty"`
	AppleID                       string   `xml:"apple_id,omitempty"`
	Site                          *Site    `xml:"site,omitempty"`
	PopulateCatalogFromVPPContent bool     `xml:"populate_catalog_from_vpp_content"`
	NotifyDisassociation          bool     `xml:"notify_disassociation"`
	AutoRegisterManagedUsers      bool     `xml:"auto_register_managed_users"`
}

// Site is a minimal site reference embedded in a VPP account.
type Site struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

// ListItemVPPAccount is the slim representation returned in list responses.
type ListItemVPPAccount struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// ListResponse is the response for ListVPPAccounts (GET /JSSResource/vppaccounts).
type ListResponse struct {
	XMLName xml.Name             `xml:"vpp_accounts"`
	Size    int                  `xml:"size"`
	Results []ListItemVPPAccount `xml:"vpp_account"`
}

// RequestVPPAccount is the body for creating or updating a VPP account.
// The ID field is not included; the target is specified via the URL path.
type RequestVPPAccount struct {
	XMLName                       xml.Name `xml:"vpp_account"`
	Name                          string   `xml:"name"`
	Contact                       string   `xml:"contact,omitempty"`
	ServiceToken                  string   `xml:"service_token,omitempty"`
	AccountName                   string   `xml:"account_name,omitempty"`
	Country                       string   `xml:"country,omitempty"`
	AppleID                       string   `xml:"apple_id,omitempty"`
	Site                          *Site    `xml:"site,omitempty"`
	PopulateCatalogFromVPPContent bool     `xml:"populate_catalog_from_vpp_content"`
	NotifyDisassociation          bool     `xml:"notify_disassociation"`
	AutoRegisterManagedUsers      bool     `xml:"auto_register_managed_users"`
}
