// Package smartgroupvalidation provides shared client-side validation for the
// Jamf Pro V2/V3 smart- and static-group endpoints introduced in Jamf Pro
// 11.28: stricter andOr criteria validation, the 255-character group-name cap,
// and set semantics (uniqueItems) for static-group assignments.
package smartgroupvalidation

import (
	"fmt"
	"strings"
)

// MaxGroupNameLength is the maximum length of a static/smart group name enforced
// by the Jamf Pro V2/V3 group endpoints (maxLength: 255).
const MaxGroupNameLength = 255

// ValidateAndOr reports an error if any supplied criterion andOr value is not
// "and" or "or" (case-insensitive). Empty values are allowed (the API treats a
// missing andOr as the default). The V2 criteria schemas validate andOr against
// this enum; the legacy V1 schema accepted any string.
func ValidateAndOr(values ...string) error {
	for _, v := range values {
		if v == "" {
			continue
		}
		switch strings.ToLower(v) {
		case "and", "or":
		default:
			return fmt.Errorf("invalid andOr value %q: must be \"and\" or \"or\"", v)
		}
	}
	return nil
}

// ValidateGroupName reports an error if the name is empty or exceeds
// MaxGroupNameLength characters.
func ValidateGroupName(name string) error {
	if name == "" {
		return fmt.Errorf("group name is required")
	}
	if len(name) > MaxGroupNameLength {
		return fmt.Errorf("group name exceeds the %d-character maximum (got %d)", MaxGroupNameLength, len(name))
	}
	return nil
}

// DedupeStrings returns a new slice with duplicate values removed while
// preserving first-seen order. Static-group assignments are a set
// (uniqueItems), so duplicates are collapsed before sending. A nil input
// returns nil.
func DedupeStrings(in []string) []string {
	if in == nil {
		return nil
	}
	seen := make(map[string]struct{}, len(in))
	out := make([]string, 0, len(in))
	for _, v := range in {
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		out = append(out, v)
	}
	return out
}
