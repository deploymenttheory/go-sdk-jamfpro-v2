package computer_groups

import "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared/smartgroupvalidation"

// validateSmartGroupV3 enforces the Jamf Pro 11.28 V2-criteria rules: the group
// name must be present and within 255 characters, and each criterion's andOr
// must be "and"/"or".
func validateSmartGroupV3(req *RequestSmartGroupV3) error {
	if err := smartgroupvalidation.ValidateGroupName(req.Name); err != nil {
		return err
	}
	for _, c := range req.Criteria {
		if err := smartgroupvalidation.ValidateAndOr(c.AndOr); err != nil {
			return err
		}
	}
	return nil
}

// validateStaticGroupV3 enforces the 255-character name cap for static groups.
func validateStaticGroupV3(req *RequestStaticGroupV3) error {
	return smartgroupvalidation.ValidateGroupName(req.Name)
}
