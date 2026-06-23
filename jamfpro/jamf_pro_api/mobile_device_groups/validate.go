package mobile_device_groups

import "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared/smartgroupvalidation"

// validateSmartMobileDeviceGroupV2 enforces the Jamf Pro 11.28 V2 smart-group
// rules: the group name must be present and within 255 characters, and each
// criterion's andOr must be "and"/"or".
func validateSmartMobileDeviceGroupV2(req *RequestSmartMobileDeviceGroup) error {
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

// validateStaticMobileDeviceGroupV2 enforces the 255-character name cap for
// static mobile device groups.
func validateStaticMobileDeviceGroupV2(req *RequestStaticMobileDeviceGroup) error {
	return smartgroupvalidation.ValidateGroupName(req.Name)
}
