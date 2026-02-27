package static_user_groups

import "fmt"

// ValidateRequest validates a static user group request.
func ValidateRequest(req *RequestStaticUserGroup) error {
	if req == nil {
		return fmt.Errorf("request is required")
	}
	if req.Name == "" {
		return fmt.Errorf("user group name is required")
	}
	if req.IsSmart {
		return fmt.Errorf("static user group must have IsSmart=false")
	}
	return nil
}
