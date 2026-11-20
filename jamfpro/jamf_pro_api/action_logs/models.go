package action_logs

type ResourceActionLog struct {
	ID          string          `json:"id"`
	CreatedTime string          `json:"createdTime"`
	Requestor   SubsetRequestor `json:"requestor"`
	ActionType  string          `json:"actionType"`
	Target      SubsetTarget    `json:"target"`
	Status      string          `json:"status"`
}

type ResourceActionLogDetail struct {
	ID            string              `json:"id"`
	CreatedTime   string              `json:"createdTime"`
	Requestor     SubsetRequestor     `json:"requestor"`
	ActionType    string              `json:"actionType"`
	Target        SubsetTarget        `json:"target"`
	Status        string              `json:"status"`
	StatusSummary SubsetStatusSummary `json:"statusSummary"`
}

type SubsetRequestor struct {
	ID   *string `json:"id"`
	Type string  `json:"type"`
	Name string  `json:"name"`
}

type SubsetTarget struct {
	Type        string  `json:"type"`
	ID          *string `json:"id"`
	Name        *string `json:"name"`
	DeviceClass *string `json:"deviceClass"`
	DeviceCount *int    `json:"deviceCount"`
}

type SubsetStatusSummary struct {
	Completed int `json:"completed"`
	Failed    int `json:"failed"`
	Pending   int `json:"pending"`
	Canceled  int `json:"canceled"`
}

type ResourceActionLogDevice struct {
	DeviceName   string  `json:"deviceName"`
	DeviceClass  string  `json:"deviceClass"`
	DeviceID     *string `json:"deviceId"`
	DeviceType   *string `json:"deviceType"`
	DeviceModel  *string `json:"deviceModel"`
	SerialNumber *string `json:"serialNumber"`
	ManagementID *string `json:"managementId"`
	Outcome      string  `json:"outcome"`
}

type ListResponse struct {
	TotalCount int                 `json:"totalCount"`
	Results    []ResourceActionLog `json:"results"`
}

type DeviceListResponse struct {
	TotalCount int                       `json:"totalCount"`
	Results    []ResourceActionLogDevice `json:"results"`
}
