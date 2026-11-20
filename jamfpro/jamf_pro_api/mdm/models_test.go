package mdm

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUnit_Mdm_ReturnToServiceRetryEnrollment(t *testing.T) {
	value := false
	body, err := json.Marshal(ReturnToService{ShouldRetryEnrollment: &value})
	require.NoError(t, err)
	var fields map[string]any
	require.NoError(t, json.Unmarshal(body, &fields))
	require.Equal(t, false, fields["shouldRetryEnrollment"])
	body, err = json.Marshal(ReturnToService{})
	require.NoError(t, err)
	require.NotContains(t, string(body), "shouldRetryEnrollment")
}
