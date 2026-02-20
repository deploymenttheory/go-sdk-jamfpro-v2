package classes

import (
	"encoding/xml"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"
)

// ResourceClass represents a Jamf Pro Classic API class resource.
type ResourceClass struct {
	XMLName             xml.Name                     `xml:"class"`
	ID                  int                          `xml:"id"`
	Source              string                       `xml:"source,omitempty"`
	Name                string                       `xml:"name"`
	Description         string                       `xml:"description,omitempty"`
	Site                shared.SharedResourceSite    `xml:"site"`
	MobileDeviceGroup   MobileDeviceGroup            `xml:"mobile_device_group,omitempty"`
	Students            []Student                    `xml:"students>student"`
	Teachers            []Teacher                    `xml:"teachers>teacher,omitempty"`
	TeacherIDs          []TeacherID                  `xml:"teacher_ids>id,omitempty"`
	StudentGroupIDs     []StudentGroupID             `xml:"student_group_ids>id"`
	TeacherGroupIDs     []TeacherGroupID             `xml:"teacher_group_ids>id"`
	MobileDevices       []MobileDevice               `xml:"mobile_devices>mobile_device"`
	MobileDeviceGroupID []MobileDeviceGroupID        `xml:"mobile_device_group_id>id,omitempty"`
	MeetingTimes        MeetingTimesContainer        `xml:"meeting_times,omitempty"`
	AppleTVs            []AppleTV                    `xml:"apple_tvs>apple_tv,omitempty"`
}

// ListResponse is the response for ListClasses (GET /JSSResource/classes).
type ListResponse struct {
	XMLName xml.Name    `xml:"classes"`
	Size    int         `xml:"size"`
	Results []ClassItem `xml:"class"`
}

// ClassItem represents a single class item in the list.
type ClassItem struct {
	ID          int    `xml:"id"`
	Name        string `xml:"name"`
	Description string `xml:"description"`
}

// RequestClass is the body for creating or updating a class.
// The ID field is not included; the target is specified via the URL path.
type RequestClass struct {
	XMLName             xml.Name                  `xml:"class"`
	Name                string                    `xml:"name"`
	Description         string                    `xml:"description,omitempty"`
	Site                *shared.SharedResourceSite `xml:"site,omitempty"`
	MobileDeviceGroup   *MobileDeviceGroup        `xml:"mobile_device_group,omitempty"`
	Students            []Student                 `xml:"students>student,omitempty"`
	Teachers            []Teacher                 `xml:"teachers>teacher,omitempty"`
	TeacherIDs          []TeacherID               `xml:"teacher_ids>id,omitempty"`
	StudentGroupIDs     []StudentGroupID          `xml:"student_group_ids>id,omitempty"`
	TeacherGroupIDs     []TeacherGroupID          `xml:"teacher_group_ids>id,omitempty"`
	MobileDevices       []MobileDevice            `xml:"mobile_devices>mobile_device,omitempty"`
	MobileDeviceGroupID []MobileDeviceGroupID     `xml:"mobile_device_group_id>id,omitempty"`
	MeetingTimes        *MeetingTimesContainer    `xml:"meeting_times,omitempty"`
	AppleTVs            []AppleTV                 `xml:"apple_tvs>apple_tv,omitempty"`
}

// CreateUpdateResponse represents the response from Create/Update operations.
// The Classic API returns only the ID for these operations.
type CreateUpdateResponse struct {
	XMLName xml.Name `xml:"class"`
	ID      int      `xml:"id"`
}

// MobileDeviceGroup represents a mobile device group reference.
type MobileDeviceGroup struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

// MobileDevice represents a mobile device in a class.
type MobileDevice struct {
	Name           string `xml:"name,omitempty"`
	UDID           string `xml:"udid,omitempty"`
	WifiMacAddress string `xml:"wifi_mac_address,omitempty"`
}

// MobileDeviceGroupID represents a mobile device group ID reference.
type MobileDeviceGroupID struct {
	ID int `xml:"id,omitempty"`
}

// Student represents a student in a class.
type Student struct {
	Student string `xml:"student,omitempty"`
}

// Teacher represents a teacher in a class.
type Teacher struct {
	Teacher string `xml:"teacher,omitempty"`
}

// TeacherID represents a teacher ID reference.
type TeacherID struct {
	ID int `xml:"id,omitempty"`
}

// StudentGroupID represents a student group ID reference.
type StudentGroupID struct {
	ID int `xml:"id,omitempty"`
}

// TeacherGroupID represents a teacher group ID reference.
type TeacherGroupID struct {
	ID int `xml:"id,omitempty"`
}

// MeetingTimesContainer wraps meeting time information.
type MeetingTimesContainer struct {
	MeetingTime MeetingTime `xml:"meeting_time,omitempty"`
}

// MeetingTime represents the meeting schedule for a class.
type MeetingTime struct {
	Days      string `xml:"days,omitempty"`
	StartTime int    `xml:"start_time,omitempty"`
	EndTime   int    `xml:"end_time,omitempty"`
}

// AppleTV represents an Apple TV in a class.
type AppleTV struct {
	Name            string `xml:"name,omitempty"`
	UDID            string `xml:"udid,omitempty"`
	WifiMacAddress  string `xml:"wifi_mac_address,omitempty"`
	DeviceID        string `xml:"device_id,omitempty"`
	AirplayPassword string `xml:"airplay_password,omitempty"`
}
