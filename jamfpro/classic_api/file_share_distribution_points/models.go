package file_share_distribution_points

import "encoding/xml"

// ResourceFileShareDistributionPoint represents a Jamf Pro Classic API file share distribution point resource.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/distributionpoints
type ResourceFileShareDistributionPoint struct {
	XMLName                    xml.Name `xml:"distribution_point"`
	ID                         int      `xml:"id,omitempty"`
	Name                       string   `xml:"name"`
	IPAddress                  string   `xml:"ipAddress,omitempty"`
	IPAddressAlt                string   `xml:"ip_address,omitempty"`
	IsMaster                   bool     `xml:"is_master"`
	FailoverPoint              string   `xml:"failover_point,omitempty"`
	FailoverPointURL           string   `xml:"failover_point_url,omitempty"`
	EnableLoadBalancing        bool     `xml:"enable_load_balancing"`
	LocalPath                  string   `xml:"local_path,omitempty"`
	SSHUsername                string   `xml:"ssh_username,omitempty"`
	Password                   string   `xml:"password,omitempty"`
	ConnectionType             string   `xml:"connection_type,omitempty"`
	ShareName                  string   `xml:"share_name,omitempty"`
	WorkgroupOrDomain          string   `xml:"workgroup_or_domain,omitempty"`
	SharePort                  int      `xml:"share_port,omitempty"`
	ReadOnlyUsername           string   `xml:"read_only_username,omitempty"`
	ReadOnlyPassword           string   `xml:"read_only_password,omitempty"`
	ReadWriteUsername          string   `xml:"read_write_username,omitempty"`
	ReadWritePassword          string   `xml:"read_write_password,omitempty"`
	HTTPDownloadsEnabled       bool     `xml:"http_downloads_enabled"`
	HTTPURL                    string   `xml:"http_url,omitempty"`
	Context                    string   `xml:"context,omitempty"`
	Protocol                   string   `xml:"protocol,omitempty"`
	Port                       int      `xml:"port,omitempty"`
	NoAuthenticationRequired  bool     `xml:"no_authentication_required"`
	UsernamePasswordRequired   bool     `xml:"username_password_required"`
	HTTPUsername               string   `xml:"http_username,omitempty"`
	HTTPPassword               string   `xml:"http_password,omitempty"`
}

// ListResponse is the response for List (GET /JSSResource/distributionpoints).
type ListResponse struct {
	XMLName xml.Name                   `xml:"distribution_points"`
	Size    int                        `xml:"size"`
	Results []DistributionPointListItem `xml:"distribution_point"`
}

// DistributionPointListItem represents a single distribution point item in the list.
type DistributionPointListItem struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// RequestFileShareDistributionPoint is the body for creating or updating a file share distribution point.
// The ID field is not included; the target is specified via the URL path.
type RequestFileShareDistributionPoint struct {
	XMLName                   xml.Name `xml:"distribution_point"`
	Name                      string   `xml:"name"`
	IPAddress                 string   `xml:"ipAddress,omitempty"`
	IPAddressAlt              string   `xml:"ip_address,omitempty"`
	IsMaster                  bool     `xml:"is_master"`
	FailoverPoint             string   `xml:"failover_point,omitempty"`
	FailoverPointURL          string   `xml:"failover_point_url,omitempty"`
	EnableLoadBalancing       bool     `xml:"enable_load_balancing"`
	LocalPath                 string   `xml:"local_path,omitempty"`
	SSHUsername               string   `xml:"ssh_username,omitempty"`
	Password                  string   `xml:"password,omitempty"`
	ConnectionType            string   `xml:"connection_type,omitempty"`
	ShareName                 string   `xml:"share_name,omitempty"`
	WorkgroupOrDomain         string   `xml:"workgroup_or_domain,omitempty"`
	SharePort                 int      `xml:"share_port,omitempty"`
	ReadOnlyUsername          string   `xml:"read_only_username,omitempty"`
	ReadOnlyPassword          string   `xml:"read_only_password,omitempty"`
	ReadWriteUsername         string   `xml:"read_write_username,omitempty"`
	ReadWritePassword         string   `xml:"read_write_password,omitempty"`
	HTTPDownloadsEnabled     bool     `xml:"http_downloads_enabled"`
	HTTPURL                   string   `xml:"http_url,omitempty"`
	Context                   string   `xml:"context,omitempty"`
	Protocol                  string   `xml:"protocol,omitempty"`
	Port                      int      `xml:"port,omitempty"`
	NoAuthenticationRequired bool     `xml:"no_authentication_required"`
	UsernamePasswordRequired  bool     `xml:"username_password_required"`
	HTTPUsername              string   `xml:"http_username,omitempty"`
	HTTPPassword              string   `xml:"http_password,omitempty"`
}

// CreateUpdateResponse represents the response from Create/Update operations.
// The Classic API returns only the ID for these operations.
// Note: The API returns <file_share_distribution_point> for Create/Update operations.
type CreateUpdateResponse struct {
	XMLName xml.Name `xml:"file_share_distribution_point"`
	ID      int      `xml:"id"`
}
