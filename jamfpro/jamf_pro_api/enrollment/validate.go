package enrollment

import (
	"fmt"
	"slices"
	"strings"
)

// ValidateLanguageCode validates a language code against the list of available codes.
func ValidateLanguageCode(languageCode string, validCodes []ResourceLanguageCode) error {
	if languageCode == "" {
		return fmt.Errorf("language code is required")
	}

	validValues := make([]string, len(validCodes))
	for i, code := range validCodes {
		validValues[i] = code.Value
	}

	if !slices.Contains(validValues, languageCode) {
		return fmt.Errorf("invalid language code '%s': not found in available language codes (%s)",
			languageCode, strings.Join(validValues, ", "))
	}

	return nil
}
