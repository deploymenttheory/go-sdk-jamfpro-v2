package managed_software_updates

// ResponseAvailableUpdates represents the response structure for available updates.
type ResponseAvailableUpdates struct {
	AvailableUpdates ResourceAvailableUpdates `json:"availableUpdates"`
}

// ResourceAvailableUpdates represents the available updates for macOS and iOS.
type ResourceAvailableUpdates struct {
	MacOS []string `json:"macOS"`
	IOS   []string `json:"iOS"`
}

// ResponsePlanList represents the paginated response for managed software update plans.
type ResponsePlanList struct {
	TotalCount int              `json:"totalCount"`
	Results    []ResourcePlan   `json:"results"`
}

// ResourcePlan represents a managed software update plan.
type ResourcePlan struct {
	PlanUuid                  string               `json:"planUuid,omitempty"`
	Device                    PlanDevice           `json:"device,omitempty"`
	UpdateAction              string               `json:"updateAction,omitempty"`
	VersionType               string               `json:"versionType,omitempty"`
	SpecificVersion           string               `json:"specificVersion,omitempty"`
	BuildVersion              string               `json:"buildVersion,omitempty"`
	MaxDeferrals              int                  `json:"maxDeferrals,omitempty"`
	ForceInstallLocalDateTime string               `json:"forceInstallLocalDateTime,omitempty"`
	RecipeId                  string               `json:"recipeId,omitempty"`
	Status                    PlanStatus           `json:"status,omitempty"`
}

// PlanDevice represents the device information in a managed software update plan.
type PlanDevice struct {
	DeviceId   string `json:"deviceId,omitempty"`
	ObjectType string `json:"objectType,omitempty"`
	Href       string `json:"href,omitempty"`
}

// PlanStatus represents the status of a managed software update plan.
type PlanStatus struct {
	State        string   `json:"state,omitempty"`
	ErrorReasons []string `json:"errorReasons"`
}

// ResponseDeclarationsList represents the response structure for the list of declarations.
type ResponseDeclarationsList struct {
	Declarations []ResourceDeclaration `json:"declarations"`
}

// ResourceDeclaration represents a declaration associated with a managed software update plan.
type ResourceDeclaration struct {
	UUID        string `json:"uuid"`
	PayloadJson string `json:"payloadJson"`
	Type        string `json:"type"`
	Group       string `json:"group"`
}

// ResponsePlanCreate represents the response structure when creating a managed software update plan.
type ResponsePlanCreate struct {
	Plans []PlanCreateItem `json:"plans"`
}

// PlanCreateItem represents a single plan item in the create response.
type PlanCreateItem struct {
	Device PlanCreateDevice `json:"device"`
	PlanID string           `json:"planId"`
	Href   string           `json:"href"`
}

// PlanCreateDevice represents the device information in a plan create response.
type PlanCreateDevice struct {
	DeviceID   string `json:"deviceId"`
	ObjectType string `json:"objectType"`
	Href       string `json:"href"`
}

// ResponseFeatureToggle represents the response structure for the feature toggle.
type ResponseFeatureToggle struct {
	Toggle                       bool `json:"toggle"`
	ForceInstallLocalDateEnabled bool `json:"forceInstallLocalDateEnabled"`
	DssEnabled                   bool `json:"dssEnabled"`
	RecipeEnabled                bool `json:"recipeEnabled"`
}

// ResponseFeatureToggleStatus represents the response structure for the feature toggle status.
type ResponseFeatureToggleStatus struct {
	ToggleOn  *FeatureToggleStatusDetail `json:"toggleOn"`
	ToggleOff *FeatureToggleStatusDetail `json:"toggleOff"`
}

// FeatureToggleStatusDetail represents the detailed status of the feature toggle (on/off).
type FeatureToggleStatusDetail struct {
	StartTime                string  `json:"startTime"`
	EndTime                  string  `json:"endTime"`
	ElapsedTime              int     `json:"elapsedTime"`
	State                    string  `json:"state"`
	TotalRecords             int64   `json:"totalRecords"`
	ProcessedRecords         int64   `json:"processedRecords"`
	PercentComplete          float64 `json:"percentComplete"`
	FormattedPercentComplete string  `json:"formattedPercentComplete"`
	ExitState                string  `json:"exitState"`
	ExitMessage              string  `json:"exitMessage"`
}

// RequestPlanCreate represents the payload structure for creating a managed software update plan.
type RequestPlanCreate struct {
	Devices []PlanObject `json:"devices,omitempty"`
	Group   PlanObject   `json:"group,omitempty"`
	Config  PlanConfig   `json:"config,omitempty"`
}

// PlanObject defines the structure for device or group objects in the managed software update plan.
type PlanObject struct {
	ObjectType string `json:"objectType"`
	DeviceId   string `json:"deviceId,omitempty"`
	GroupId    string `json:"groupId,omitempty"`
}

// PlanConfig defines the configuration for a managed software update plan.
type PlanConfig struct {
	UpdateAction              string `json:"updateAction"`
	VersionType               string `json:"versionType"`
	SpecificVersion           string `json:"specificVersion,omitempty"`
	BuildVersion              string `json:"buildVersion,omitempty"`
	MaxDeferrals              int    `json:"maxDeferrals,omitempty"`
	ForceInstallLocalDateTime string `json:"forceInstallLocalDateTime,omitempty"`
}

// RequestFeatureToggle represents the payload for updating the feature toggle.
type RequestFeatureToggle struct {
	Toggle bool `json:"toggle"`
}

// ResponseError represents a standard error response.
type ResponseError struct {
	HTTPStatus int               `json:"httpStatus"`
	Errors     []ErrorDetail     `json:"errors"`
}

// ErrorDetail represents a single error detail.
type ErrorDetail struct {
	Code        string `json:"code"`
	Field       string `json:"field"`
	Description string `json:"description"`
	ID          string `json:"id"`
}
