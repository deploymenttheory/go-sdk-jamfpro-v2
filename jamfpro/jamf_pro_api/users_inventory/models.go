package users_inventory

// ResourceUser represents a user resource in inventory.
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-users-id
type ResourceUser struct {
	ID                   string `json:"id"`
	Username             string `json:"username,omitempty"`
	Realname             string `json:"realname,omitempty"`
	Email                string `json:"email,omitempty"`
	Phone                string `json:"phone,omitempty"`
	Position             string `json:"position,omitempty"`
	EnableCustomPhotoUrl bool   `json:"enableCustomPhotoUrl,omitempty"`
	CustomPhotoUrl       string `json:"customPhotoUrl,omitempty"`
	ManagedAppleId       string `json:"managedAppleId,omitempty"`
}

// RequestUserInventory is the request body for creating and updating users.
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-users
type RequestUserInventory struct {
	Username             string `json:"username,omitempty"`
	Realname             string `json:"realname,omitempty"`
	Email                string `json:"email,omitempty"`
	Phone                string `json:"phone,omitempty"`
	Position             string `json:"position,omitempty"`
	EnableCustomPhotoUrl bool   `json:"enableCustomPhotoUrl,omitempty"`
	CustomPhotoUrl       string `json:"customPhotoUrl,omitempty"`
	ManagedAppleId       string `json:"managedAppleId,omitempty"`
}

// ListUsersResponse is the paginated response for ListV1.
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-users
type ListUsersResponse struct {
	Page        int            `json:"page"`
	PageSize    int            `json:"pageSize"`
	TotalCount  int            `json:"totalCount"`
	TotalPages  int            `json:"totalPages"`
	HasNext     bool           `json:"hasNext"`
	HasPrevious bool           `json:"hasPrevious"`
	Results     []ResourceUser `json:"results"`
}

// CreateUserResponse is the response for CreateV1.
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-users
type CreateUserResponse struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}
