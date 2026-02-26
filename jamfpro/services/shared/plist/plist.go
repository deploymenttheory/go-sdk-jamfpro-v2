package plist

import (
	"fmt"
	"strings"

	"howett.net/plist"
)

// -----------------------------------------------------------------------------
// UUID Preservation Helper Functions
// -----------------------------------------------------------------------------

// PreservePlistUUIDs extracts UUIDs from an existing plist string and injects them into a new plist string.
//
// This function handles nested plist structures and ensures that PayloadUUID and PayloadIdentifier
// values are preserved from the existing plist to the new one. This prevents Jamf Pro from treating
// the update as a brand new plist.
//
// Jamf Pro modifies only the top-level PayloadUUID and PayloadIdentifier upon profile creation.
// All nested payload UUIDs/identifiers remain unchanged. This function:
// 1. Copies top-level PayloadUUID and PayloadIdentifier from existing to new
// 2. Builds UUID maps by PayloadDisplayName for nested payloads
// 3. Updates nested PayloadUUID and PayloadIdentifier values in new plist
func PreservePlistUUIDs(existingPlist, newPlist string) (string, error) {
	// Decode existing plist
	var existingData map[string]any
	decoder := plist.NewDecoder(strings.NewReader(existingPlist))
	if err := decoder.Decode(&existingData); err != nil {
		return "", fmt.Errorf("failed to decode existing plist: %w", err)
	}

	// Decode new plist
	var newData map[string]any
	decoder = plist.NewDecoder(strings.NewReader(newPlist))
	if err := decoder.Decode(&newData); err != nil {
		return "", fmt.Errorf("failed to decode new plist: %w", err)
	}

	// Step 1: Preserve top-level PayloadUUID and PayloadIdentifier
	// Jamf Pro modifies these values upon profile creation
	if uuid, ok := existingData["PayloadUUID"].(string); ok {
		newData["PayloadUUID"] = uuid
	}
	if identifier, ok := existingData["PayloadIdentifier"].(string); ok {
		newData["PayloadIdentifier"] = identifier
	}

	// Step 2: Build UUID maps from existing plist for nested payloads
	// Map payloads by PayloadDisplayName to ensure correct UUID matching
	uuidMap := make(map[string]string)
	identifierMap := make(map[string]string)
	extractUUIDsFromPlist(existingData, uuidMap, identifierMap)

	// Step 3: Update UUIDs in new plist using the maps
	updateUUIDsInPlist(newData, uuidMap, identifierMap)

	// Step 4: Re-encode the updated plist
	var buf strings.Builder
	encoder := plist.NewEncoder(&buf)
	encoder.Indent("\t")
	if err := encoder.Encode(newData); err != nil {
		return "", fmt.Errorf("failed to encode updated plist: %w", err)
	}

	return buf.String(), nil
}

// extractUUIDsFromPlist recursively extracts PayloadUUID and PayloadIdentifier values from a plist.
//
// It builds maps keyed by PayloadDisplayName (or PayloadType as fallback) for later matching.
// This ensures that UUIDs are correctly matched between existing and new plists even if
// payload order changes.
func extractUUIDsFromPlist(data map[string]any, uuidMap, identifierMap map[string]string) {
	// Determine the key for this payload
	displayName, _ := data["PayloadDisplayName"].(string)
	if displayName == "" {
		displayName, _ = data["PayloadType"].(string)
	}

	// Extract UUID and identifier if display name exists
	if displayName != "" {
		if uuid, ok := data["PayloadUUID"].(string); ok {
			uuidMap[displayName] = uuid
		}
		if identifier, ok := data["PayloadIdentifier"].(string); ok {
			identifierMap[displayName] = identifier
		}
	}

	// Recursively process PayloadContent array (nested payloads)
	if payloadContent, ok := data["PayloadContent"].([]any); ok {
		for _, item := range payloadContent {
			if itemMap, ok := item.(map[string]any); ok {
				extractUUIDsFromPlist(itemMap, uuidMap, identifierMap)
			}
		}
	}
}

// updateUUIDsInPlist recursively updates PayloadUUID and PayloadIdentifier values in a plist
// using the provided maps.
//
// This function matches payloads by PayloadDisplayName (or PayloadType as fallback) and
// injects the corresponding UUIDs from the existing plist.
func updateUUIDsInPlist(data map[string]any, uuidMap, identifierMap map[string]string) {
	// Determine the key for this payload
	displayName, _ := data["PayloadDisplayName"].(string)
	if displayName == "" {
		displayName, _ = data["PayloadType"].(string)
	}

	// Update UUID and identifier if found in maps
	if displayName != "" {
		if uuid, found := uuidMap[displayName]; found {
			data["PayloadUUID"] = uuid
		}
		if identifier, found := identifierMap[displayName]; found {
			data["PayloadIdentifier"] = identifier
		}
	}

	// Recursively process PayloadContent array (nested payloads)
	if payloadContent, ok := data["PayloadContent"].([]any); ok {
		for _, item := range payloadContent {
			if itemMap, ok := item.(map[string]any); ok {
				updateUUIDsInPlist(itemMap, uuidMap, identifierMap)
			}
		}
	}
}
