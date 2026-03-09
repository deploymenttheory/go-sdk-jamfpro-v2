package mobile_device_enrollment_profiles

import (
	"encoding/xml"
)

// ListResponse is the response for List (GET /JSSResource/mobiledeviceenrollmentprofiles).
type ListResponse struct {
	XMLName xml.Name `xml:"mobile_device_enrollment_profiles"`
	Size    int      `xml:"size"`
	Results []ListItem `xml:"mobile_device_enrollment_profile"`
}

// ListItem represents a single mobile device enrollment profile item in the list.
type ListItem struct {
	ID         int     `xml:"id"`
	Name       string  `xml:"name"`
	Invitation float64 `xml:"invitation"`
}

// Resource represents the full mobile device enrollment profile structure.
type Resource struct {
	XMLName     xml.Name            `xml:"mobile_device_enrollment_profile"`
	General     SubsetGeneral       `xml:"general"`
	Location    *SubsetLocation     `xml:"location,omitempty"`
	Purchasing  *SubsetPurchasing   `xml:"purchasing,omitempty"`
	Attachments []SubsetAttachments `xml:"attachments>attachment,omitempty"`
}

// SubsetGeneral represents the general subset of a mobile device enrollment profile.
type SubsetGeneral struct {
	ID          int    `xml:"id"`
	Name        string `xml:"name"`
	Invitation  string `xml:"invitation,omitempty"`
	UDID        string `xml:"udid,omitempty"`
	Description string `xml:"description,omitempty"`
}

// SubsetLocation represents the location subset of a mobile device enrollment profile.
type SubsetLocation struct {
	Username     string `xml:"username,omitempty"`
	Realname     string `xml:"realname,omitempty"`
	RealName     string `xml:"real_name,omitempty"`
	EmailAddress string `xml:"email_address,omitempty"`
	Position     string `xml:"position,omitempty"`
	Phone        string `xml:"phone,omitempty"`
	PhoneNumber  string `xml:"phone_number,omitempty"`
	Department   string `xml:"department,omitempty"`
	Building     string `xml:"building,omitempty"`
	Room         int    `xml:"room,omitempty"`
}

// SubsetPurchasing represents the purchasing subset of a mobile device enrollment profile.
type SubsetPurchasing struct {
	IsPurchased          bool   `xml:"is_purchased"`
	IsLeased             bool   `xml:"is_leased"`
	PONumber             string `xml:"po_number,omitempty"`
	Vendor               string `xml:"vendor,omitempty"`
	ApplecareID          string `xml:"applecare_id,omitempty"`
	PurchasePrice        string `xml:"purchase_price,omitempty"`
	PurchasingAccount    string `xml:"purchasing_account,omitempty"`
	PODate               string `xml:"po_date,omitempty"`
	PODateEpoch          int64  `xml:"po_date_epoch,omitempty"`
	PODateUTC            string `xml:"po_date_utc,omitempty"`
	WarrantyExpires      string `xml:"warranty_expires,omitempty"`
	WarrantyExpiresEpoch int64  `xml:"warranty_expires_epoch,omitempty"`
	WarrantyExpiresUTC   string `xml:"warranty_expires_utc,omitempty"`
	LeaseExpires         string `xml:"lease_expires,omitempty"`
	LeaseExpiresEpoch    int64  `xml:"lease_expires_epoch,omitempty"`
	LeaseExpiresUTC      string `xml:"lease_expires_utc,omitempty"`
	LifeExpectancy       int    `xml:"life_expectancy,omitempty"`
	PurchasingContact    string `xml:"purchasing_contact,omitempty"`
}

// SubsetAttachments represents an attachment in a mobile device enrollment profile.
type SubsetAttachments struct {
	ID       int    `xml:"id"`
	Filename string `xml:"filename"`
	URI      string `xml:"uri"`
}
