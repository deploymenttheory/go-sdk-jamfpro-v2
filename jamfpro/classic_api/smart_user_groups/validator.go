package smart_user_groups

import "fmt"

// ValidateRequest validates a smart user group request.
func ValidateRequest(req *RequestSmartUserGroup) error {
	if req == nil {
		return fmt.Errorf("request is required")
	}
	if req.Name == "" {
		return fmt.Errorf("user group name is required")
	}
	if !req.IsSmart {
		return fmt.Errorf("smart user group must have IsSmart=true")
	}
	if req.Criteria == nil || req.Criteria.Size == 0 {
		return fmt.Errorf("smart user group must have criteria defined")
	}
	return nil
}
