package normalization

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// IDAsString unmarshals from JSON as either string or integer.
// Jamf Pro returns IDs in both formats depending on context;
// we always store them as strings so that we don't need to
// worry about what the API returns.
type IDAsString string

func (f *IDAsString) UnmarshalJSON(data []byte) error {
	var raw any
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	switch v := raw.(type) {
	case string:
		*f = IDAsString(v)
	case float64:
		*f = IDAsString(strconv.FormatInt(int64(v), 10))
	default:
		return fmt.Errorf("id: unexpected type %T", v)
	}
	return nil
}
