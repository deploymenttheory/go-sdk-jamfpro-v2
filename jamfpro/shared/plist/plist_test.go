package plist

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"howett.net/plist"
)

// TestUnit_MacOSConfigurationProfiles_PreservePlistUUIDs_TopLevelOnly tests UUID preservation
// for a simple plist with only top-level UUIDs.
func TestUnit_MacOSConfigurationProfiles_PreservePlistUUIDs_TopLevelOnly(t *testing.T) {
	existingPlist := `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>PayloadUUID</key>
	<string>EXISTING-UUID-12345</string>
	<key>PayloadIdentifier</key>
	<string>com.example.existing</string>
	<key>PayloadDisplayName</key>
	<string>Test Profile</string>
	<key>PayloadType</key>
	<string>Configuration</string>
</dict>
</plist>`

	newPlist := `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>PayloadUUID</key>
	<string>NEW-UUID-67890</string>
	<key>PayloadIdentifier</key>
	<string>com.example.new</string>
	<key>PayloadDisplayName</key>
	<string>Test Profile Updated</string>
	<key>PayloadType</key>
	<string>Configuration</string>
</dict>
</plist>`

	result, err := PreservePlistUUIDs(existingPlist, newPlist)
	require.NoError(t, err)
	require.NotEmpty(t, result)

	// Decode result to verify UUIDs were preserved
	var resultData map[string]any
	decoder := plist.NewDecoder(strings.NewReader(result))
	err = decoder.Decode(&resultData)
	require.NoError(t, err)

	// Verify top-level UUID was preserved from existing
	assert.Equal(t, "EXISTING-UUID-12345", resultData["PayloadUUID"])
	assert.Equal(t, "com.example.existing", resultData["PayloadIdentifier"])
	// Verify other fields came from new plist
	assert.Equal(t, "Test Profile Updated", resultData["PayloadDisplayName"])
}

// TestUnit_MacOSConfigurationProfiles_PreservePlistUUIDs_NestedPayloads tests UUID preservation
// for a plist with nested PayloadContent arrays.
func TestUnit_MacOSConfigurationProfiles_PreservePlistUUIDs_NestedPayloads(t *testing.T) {
	existingPlist := `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>PayloadUUID</key>
	<string>TOP-LEVEL-UUID-EXISTING</string>
	<key>PayloadIdentifier</key>
	<string>com.example.top.existing</string>
	<key>PayloadContent</key>
	<array>
		<dict>
			<key>PayloadDisplayName</key>
			<string>WiFi Settings</string>
			<key>PayloadType</key>
			<string>com.apple.wifi.managed</string>
			<key>PayloadUUID</key>
			<string>WIFI-UUID-EXISTING</string>
			<key>PayloadIdentifier</key>
			<string>com.example.wifi.existing</string>
		</dict>
		<dict>
			<key>PayloadDisplayName</key>
			<string>VPN Settings</string>
			<key>PayloadType</key>
			<string>com.apple.vpn.managed</string>
			<key>PayloadUUID</key>
			<string>VPN-UUID-EXISTING</string>
			<key>PayloadIdentifier</key>
			<string>com.example.vpn.existing</string>
		</dict>
	</array>
</dict>
</plist>`

	newPlist := `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>PayloadUUID</key>
	<string>TOP-LEVEL-UUID-NEW</string>
	<key>PayloadIdentifier</key>
	<string>com.example.top.new</string>
	<key>PayloadContent</key>
	<array>
		<dict>
			<key>PayloadDisplayName</key>
			<string>WiFi Settings</string>
			<key>PayloadType</key>
			<string>com.apple.wifi.managed</string>
			<key>PayloadUUID</key>
			<string>WIFI-UUID-NEW</string>
			<key>PayloadIdentifier</key>
			<string>com.example.wifi.new</string>
		</dict>
		<dict>
			<key>PayloadDisplayName</key>
			<string>VPN Settings</string>
			<key>PayloadType</key>
			<string>com.apple.vpn.managed</string>
			<key>PayloadUUID</key>
			<string>VPN-UUID-NEW</string>
			<key>PayloadIdentifier</key>
			<string>com.example.vpn.new</string>
		</dict>
	</array>
</dict>
</plist>`

	result, err := PreservePlistUUIDs(existingPlist, newPlist)
	require.NoError(t, err)
	require.NotEmpty(t, result)

	// Decode result to verify UUIDs were preserved
	var resultData map[string]any
	decoder := plist.NewDecoder(strings.NewReader(result))
	err = decoder.Decode(&resultData)
	require.NoError(t, err)

	// Verify top-level UUIDs were preserved
	assert.Equal(t, "TOP-LEVEL-UUID-EXISTING", resultData["PayloadUUID"])
	assert.Equal(t, "com.example.top.existing", resultData["PayloadIdentifier"])

	// Verify nested payload UUIDs were preserved
	payloadContent, ok := resultData["PayloadContent"].([]any)
	require.True(t, ok)
	require.Len(t, payloadContent, 2)

	// Check WiFi payload
	wifiPayload, ok := payloadContent[0].(map[string]any)
	require.True(t, ok)
	assert.Equal(t, "WIFI-UUID-EXISTING", wifiPayload["PayloadUUID"])
	assert.Equal(t, "com.example.wifi.existing", wifiPayload["PayloadIdentifier"])

	// Check VPN payload
	vpnPayload, ok := payloadContent[1].(map[string]any)
	require.True(t, ok)
	assert.Equal(t, "VPN-UUID-EXISTING", vpnPayload["PayloadUUID"])
	assert.Equal(t, "com.example.vpn.existing", vpnPayload["PayloadIdentifier"])
}

// TestUnit_MacOSConfigurationProfiles_PreservePlistUUIDs_PayloadTypeFallback tests UUID preservation
// when PayloadDisplayName is missing and PayloadType is used as fallback.
func TestUnit_MacOSConfigurationProfiles_PreservePlistUUIDs_PayloadTypeFallback(t *testing.T) {
	existingPlist := `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>PayloadUUID</key>
	<string>TOP-UUID-EXISTING</string>
	<key>PayloadContent</key>
	<array>
		<dict>
			<key>PayloadType</key>
			<string>com.apple.security.firewall</string>
			<key>PayloadUUID</key>
			<string>FIREWALL-UUID-EXISTING</string>
			<key>PayloadIdentifier</key>
			<string>com.example.firewall.existing</string>
		</dict>
	</array>
</dict>
</plist>`

	newPlist := `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>PayloadUUID</key>
	<string>TOP-UUID-NEW</string>
	<key>PayloadContent</key>
	<array>
		<dict>
			<key>PayloadType</key>
			<string>com.apple.security.firewall</string>
			<key>PayloadUUID</key>
			<string>FIREWALL-UUID-NEW</string>
			<key>PayloadIdentifier</key>
			<string>com.example.firewall.new</string>
		</dict>
	</array>
</dict>
</plist>`

	result, err := PreservePlistUUIDs(existingPlist, newPlist)
	require.NoError(t, err)
	require.NotEmpty(t, result)

	// Decode result to verify UUIDs were preserved using PayloadType as key
	var resultData map[string]any
	decoder := plist.NewDecoder(strings.NewReader(result))
	err = decoder.Decode(&resultData)
	require.NoError(t, err)

	payloadContent, ok := resultData["PayloadContent"].([]any)
	require.True(t, ok)
	require.Len(t, payloadContent, 1)

	firewallPayload, ok := payloadContent[0].(map[string]any)
	require.True(t, ok)
	assert.Equal(t, "FIREWALL-UUID-EXISTING", firewallPayload["PayloadUUID"])
	assert.Equal(t, "com.example.firewall.existing", firewallPayload["PayloadIdentifier"])
}

// TestUnit_MacOSConfigurationProfiles_PreservePlistUUIDs_DeeplyNestedPayloads tests UUID preservation
// for deeply nested PayloadContent structures.
func TestUnit_MacOSConfigurationProfiles_PreservePlistUUIDs_DeeplyNestedPayloads(t *testing.T) {
	existingPlist := `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>PayloadUUID</key>
	<string>ROOT-UUID-EXISTING</string>
	<key>PayloadContent</key>
	<array>
		<dict>
			<key>PayloadDisplayName</key>
			<string>Parent Payload</string>
			<key>PayloadUUID</key>
			<string>PARENT-UUID-EXISTING</string>
			<key>PayloadContent</key>
			<array>
				<dict>
					<key>PayloadDisplayName</key>
					<string>Child Payload</string>
					<key>PayloadUUID</key>
					<string>CHILD-UUID-EXISTING</string>
					<key>PayloadIdentifier</key>
					<string>com.example.child.existing</string>
				</dict>
			</array>
		</dict>
	</array>
</dict>
</plist>`

	newPlist := `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>PayloadUUID</key>
	<string>ROOT-UUID-NEW</string>
	<key>PayloadContent</key>
	<array>
		<dict>
			<key>PayloadDisplayName</key>
			<string>Parent Payload</string>
			<key>PayloadUUID</key>
			<string>PARENT-UUID-NEW</string>
			<key>PayloadContent</key>
			<array>
				<dict>
					<key>PayloadDisplayName</key>
					<string>Child Payload</string>
					<key>PayloadUUID</key>
					<string>CHILD-UUID-NEW</string>
					<key>PayloadIdentifier</key>
					<string>com.example.child.new</string>
				</dict>
			</array>
		</dict>
	</array>
</dict>
</plist>`

	result, err := PreservePlistUUIDs(existingPlist, newPlist)
	require.NoError(t, err)
	require.NotEmpty(t, result)

	// Decode result to verify deeply nested UUIDs were preserved
	var resultData map[string]any
	decoder := plist.NewDecoder(strings.NewReader(result))
	err = decoder.Decode(&resultData)
	require.NoError(t, err)

	// Check root UUID
	assert.Equal(t, "ROOT-UUID-EXISTING", resultData["PayloadUUID"])

	// Navigate to parent payload
	payloadContent, ok := resultData["PayloadContent"].([]any)
	require.True(t, ok)
	require.Len(t, payloadContent, 1)

	parentPayload, ok := payloadContent[0].(map[string]any)
	require.True(t, ok)
	assert.Equal(t, "PARENT-UUID-EXISTING", parentPayload["PayloadUUID"])

	// Navigate to child payload
	childContent, ok := parentPayload["PayloadContent"].([]any)
	require.True(t, ok)
	require.Len(t, childContent, 1)

	childPayload, ok := childContent[0].(map[string]any)
	require.True(t, ok)
	assert.Equal(t, "CHILD-UUID-EXISTING", childPayload["PayloadUUID"])
	assert.Equal(t, "com.example.child.existing", childPayload["PayloadIdentifier"])
}

// TestUnit_MacOSConfigurationProfiles_PreservePlistUUIDs_InvalidExistingPlist tests error handling
// when the existing plist is invalid.
func TestUnit_MacOSConfigurationProfiles_PreservePlistUUIDs_InvalidExistingPlist(t *testing.T) {
	invalidPlist := `This is not a valid plist`
	validPlist := `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>PayloadUUID</key>
	<string>TEST-UUID</string>
</dict>
</plist>`

	result, err := PreservePlistUUIDs(invalidPlist, validPlist)
	assert.Error(t, err)
	assert.Empty(t, result)
	assert.Contains(t, err.Error(), "failed to decode existing plist")
}

// TestUnit_MacOSConfigurationProfiles_PreservePlistUUIDs_InvalidNewPlist tests error handling
// when the new plist is invalid.
func TestUnit_MacOSConfigurationProfiles_PreservePlistUUIDs_InvalidNewPlist(t *testing.T) {
	validPlist := `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>PayloadUUID</key>
	<string>TEST-UUID</string>
</dict>
</plist>`
	invalidPlist := `This is not a valid plist`

	result, err := PreservePlistUUIDs(validPlist, invalidPlist)
	assert.Error(t, err)
	assert.Empty(t, result)
	assert.Contains(t, err.Error(), "failed to decode new plist")
}

// TestUnit_MacOSConfigurationProfiles_PreservePlistUUIDs_MissingUUIDs tests behavior
// when existing plist has no UUIDs.
func TestUnit_MacOSConfigurationProfiles_PreservePlistUUIDs_MissingUUIDs(t *testing.T) {
	existingPlist := `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>PayloadDisplayName</key>
	<string>Test Profile</string>
</dict>
</plist>`

	newPlist := `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>PayloadUUID</key>
	<string>NEW-UUID</string>
	<key>PayloadDisplayName</key>
	<string>Test Profile</string>
</dict>
</plist>`

	result, err := PreservePlistUUIDs(existingPlist, newPlist)
	require.NoError(t, err)
	require.NotEmpty(t, result)

	// Decode result - should keep new UUID since existing had none
	var resultData map[string]any
	decoder := plist.NewDecoder(strings.NewReader(result))
	err = decoder.Decode(&resultData)
	require.NoError(t, err)

	// New UUID should remain since existing had no UUID
	assert.Equal(t, "NEW-UUID", resultData["PayloadUUID"])
}

// TestUnit_MacOSConfigurationProfiles_ExtractUUIDsFromPlist tests the extraction helper function.
func TestUnit_MacOSConfigurationProfiles_ExtractUUIDsFromPlist(t *testing.T) {
	data := map[string]any{
		"PayloadDisplayName": "Test Payload",
		"PayloadUUID":        "TEST-UUID-123",
		"PayloadIdentifier":  "com.example.test",
		"PayloadContent": []any{
			map[string]any{
				"PayloadDisplayName": "Nested Payload",
				"PayloadUUID":        "NESTED-UUID-456",
				"PayloadIdentifier":  "com.example.nested",
			},
		},
	}

	uuidMap := make(map[string]string)
	identifierMap := make(map[string]string)

	extractUUIDsFromPlist(data, uuidMap, identifierMap)

	// Verify top-level extraction
	assert.Equal(t, "TEST-UUID-123", uuidMap["Test Payload"])
	assert.Equal(t, "com.example.test", identifierMap["Test Payload"])

	// Verify nested extraction
	assert.Equal(t, "NESTED-UUID-456", uuidMap["Nested Payload"])
	assert.Equal(t, "com.example.nested", identifierMap["Nested Payload"])
}

// TestUnit_MacOSConfigurationProfiles_ExtractUUIDsFromPlist_PayloadTypeFallback tests extraction
// using PayloadType when PayloadDisplayName is missing.
func TestUnit_MacOSConfigurationProfiles_ExtractUUIDsFromPlist_PayloadTypeFallback(t *testing.T) {
	data := map[string]any{
		"PayloadType":       "com.apple.security",
		"PayloadUUID":       "SECURITY-UUID",
		"PayloadIdentifier": "com.example.security",
	}

	uuidMap := make(map[string]string)
	identifierMap := make(map[string]string)

	extractUUIDsFromPlist(data, uuidMap, identifierMap)

	// Verify extraction using PayloadType as key
	assert.Equal(t, "SECURITY-UUID", uuidMap["com.apple.security"])
	assert.Equal(t, "com.example.security", identifierMap["com.apple.security"])
}

// TestUnit_MacOSConfigurationProfiles_ExtractUUIDsFromPlist_NoDisplayNameOrType tests extraction
// when neither PayloadDisplayName nor PayloadType exist.
func TestUnit_MacOSConfigurationProfiles_ExtractUUIDsFromPlist_NoDisplayNameOrType(t *testing.T) {
	data := map[string]any{
		"PayloadUUID":       "ORPHAN-UUID",
		"PayloadIdentifier": "com.example.orphan",
	}

	uuidMap := make(map[string]string)
	identifierMap := make(map[string]string)

	extractUUIDsFromPlist(data, uuidMap, identifierMap)

	// Maps should be empty since there's no key to use
	assert.Empty(t, uuidMap)
	assert.Empty(t, identifierMap)
}

// TestUnit_MacOSConfigurationProfiles_UpdateUUIDsInPlist tests the update helper function.
func TestUnit_MacOSConfigurationProfiles_UpdateUUIDsInPlist(t *testing.T) {
	data := map[string]any{
		"PayloadDisplayName": "Test Payload",
		"PayloadUUID":        "OLD-UUID",
		"PayloadIdentifier":  "com.example.old",
		"PayloadContent": []any{
			map[string]any{
				"PayloadDisplayName": "Nested Payload",
				"PayloadUUID":        "OLD-NESTED-UUID",
				"PayloadIdentifier":  "com.example.old.nested",
			},
		},
	}

	uuidMap := map[string]string{
		"Test Payload":   "NEW-UUID-123",
		"Nested Payload": "NEW-NESTED-UUID-456",
	}
	identifierMap := map[string]string{
		"Test Payload":   "com.example.new",
		"Nested Payload": "com.example.new.nested",
	}

	updateUUIDsInPlist(data, uuidMap, identifierMap)

	// Verify top-level update
	assert.Equal(t, "NEW-UUID-123", data["PayloadUUID"])
	assert.Equal(t, "com.example.new", data["PayloadIdentifier"])

	// Verify nested update
	payloadContent := data["PayloadContent"].([]any)
	nestedPayload := payloadContent[0].(map[string]any)
	assert.Equal(t, "NEW-NESTED-UUID-456", nestedPayload["PayloadUUID"])
	assert.Equal(t, "com.example.new.nested", nestedPayload["PayloadIdentifier"])
}

// TestUnit_MacOSConfigurationProfiles_UpdateUUIDsInPlist_NoMatchingKey tests update behavior
// when the payload key is not in the maps.
func TestUnit_MacOSConfigurationProfiles_UpdateUUIDsInPlist_NoMatchingKey(t *testing.T) {
	data := map[string]any{
		"PayloadDisplayName": "Unmatched Payload",
		"PayloadUUID":        "ORIGINAL-UUID",
		"PayloadIdentifier":  "com.example.original",
	}

	uuidMap := map[string]string{
		"Different Payload": "NEW-UUID",
	}
	identifierMap := map[string]string{
		"Different Payload": "com.example.new",
	}

	updateUUIDsInPlist(data, uuidMap, identifierMap)

	// Original values should remain unchanged
	assert.Equal(t, "ORIGINAL-UUID", data["PayloadUUID"])
	assert.Equal(t, "com.example.original", data["PayloadIdentifier"])
}

// TestUnit_MacOSConfigurationProfiles_PreservePlistUUIDs_EmptyPayloadContent tests behavior
// with empty PayloadContent array.
func TestUnit_MacOSConfigurationProfiles_PreservePlistUUIDs_EmptyPayloadContent(t *testing.T) {
	existingPlist := `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>PayloadUUID</key>
	<string>EXISTING-UUID</string>
	<key>PayloadContent</key>
	<array/>
</dict>
</plist>`

	newPlist := `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>PayloadUUID</key>
	<string>NEW-UUID</string>
	<key>PayloadContent</key>
	<array/>
</dict>
</plist>`

	result, err := PreservePlistUUIDs(existingPlist, newPlist)
	require.NoError(t, err)
	require.NotEmpty(t, result)

	var resultData map[string]any
	decoder := plist.NewDecoder(strings.NewReader(result))
	err = decoder.Decode(&resultData)
	require.NoError(t, err)

	// Top-level UUID should be preserved
	assert.Equal(t, "EXISTING-UUID", resultData["PayloadUUID"])

	// PayloadContent should still be empty array
	payloadContent, ok := resultData["PayloadContent"].([]any)
	require.True(t, ok)
	assert.Empty(t, payloadContent)
}
