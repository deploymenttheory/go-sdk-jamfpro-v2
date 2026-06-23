package groups

import "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared/smartgroupvalidation"

// validateGroupUpdate enforces the Jamf Pro 11.28 V2-criteria andOr rule
// (each criterion's andOr must be "and"/"or") used by the v2 update endpoint.
func validateGroupUpdate(req *RequestUpdateGroup) error {
	for _, c := range req.Criteria {
		if err := smartgroupvalidation.ValidateAndOr(c.AndOr); err != nil {
			return err
		}
	}
	return nil
}
