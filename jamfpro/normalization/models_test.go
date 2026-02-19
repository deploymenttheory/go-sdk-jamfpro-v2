package normalization

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIDAsString_UnmarshalJSON_FromString(t *testing.T) {
	var id IDAsString
	err := json.Unmarshal([]byte(`"123"`), &id)
	require.NoError(t, err)
	assert.Equal(t, IDAsString("123"), id)
	assert.Equal(t, "123", string(id))
}

func TestIDAsString_UnmarshalJSON_FromNumber(t *testing.T) {
	var id IDAsString
	err := json.Unmarshal([]byte(`37013`), &id)
	require.NoError(t, err)
	assert.Equal(t, IDAsString("37013"), id)
	assert.Equal(t, "37013", string(id))
}

func TestIDAsString_UnmarshalJSON_FromNumberZero(t *testing.T) {
	var id IDAsString
	err := json.Unmarshal([]byte(`0`), &id)
	require.NoError(t, err)
	assert.Equal(t, IDAsString("0"), id)
}

func TestIDAsString_UnmarshalJSON_FromBool_ReturnsError(t *testing.T) {
	var id IDAsString
	err := json.Unmarshal([]byte(`true`), &id)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unexpected type")
}

func TestIDAsString_UnmarshalJSON_FromObject_ReturnsError(t *testing.T) {
	var id IDAsString
	err := json.Unmarshal([]byte(`{"key":"value"}`), &id)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unexpected type")
}

func TestIDAsString_UnmarshalJSON_FromArray_ReturnsError(t *testing.T) {
	var id IDAsString
	err := json.Unmarshal([]byte(`["a"]`), &id)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unexpected type")
}

func TestIDAsString_UnmarshalJSON_InStruct(t *testing.T) {
	var v struct {
		ID IDAsString `json:"id"`
	}
	err := json.Unmarshal([]byte(`{"id":999}`), &v)
	require.NoError(t, err)
	assert.Equal(t, IDAsString("999"), v.ID)
}

func TestIDAsString_UnmarshalJSON_InStruct_StringID(t *testing.T) {
	var v struct {
		ID IDAsString `json:"id"`
	}
	err := json.Unmarshal([]byte(`{"id":"abc-123"}`), &v)
	require.NoError(t, err)
	assert.Equal(t, IDAsString("abc-123"), v.ID)
}
