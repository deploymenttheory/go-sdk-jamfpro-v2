package command_flush

import "encoding/xml"

// RequestCommandFlush represents the XML request body for batch command flush operations.
type RequestCommandFlush struct {
	XMLName       xml.Name      `xml:"commandflush"`
	Status        string        `xml:"status"`
	MobileDevices *MobileDevices `xml:"mobile_devices,omitempty"`
	Computers     *Computers    `xml:"computers,omitempty"`
}

// MobileDevices represents a list of mobile devices in the command flush request.
type MobileDevices struct {
	MobileDevice []DeviceID `xml:"mobile_device"`
}

// Computers represents a list of computers in the command flush request.
type Computers struct {
	Computer []DeviceID `xml:"computer"`
}

// DeviceID represents a device ID in the command flush request.
type DeviceID struct {
	ID int `xml:"id"`
}
