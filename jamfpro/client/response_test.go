package client

import (
	"net/http"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/stretchr/testify/assert"
)

func TestIsResponseSuccess(t *testing.T) {
	assert.False(t, IsResponseSuccess(nil))
	assert.True(t, IsResponseSuccess(&interfaces.Response{StatusCode: 200}))
	assert.True(t, IsResponseSuccess(&interfaces.Response{StatusCode: 299}))
	assert.False(t, IsResponseSuccess(&interfaces.Response{StatusCode: 400}))
}

func TestIsResponseError(t *testing.T) {
	assert.False(t, IsResponseError(nil))
	assert.False(t, IsResponseError(&interfaces.Response{StatusCode: 200}))
	assert.True(t, IsResponseError(&interfaces.Response{StatusCode: 400}))
	assert.True(t, IsResponseError(&interfaces.Response{StatusCode: 500}))
}

func TestGetResponseHeader(t *testing.T) {
	assert.Empty(t, GetResponseHeader(nil, "X-Foo"))
	assert.Empty(t, GetResponseHeader(&interfaces.Response{}, "X-Foo"))
	h := make(http.Header)
	h.Set("X-Foo", "bar")
	assert.Equal(t, "bar", GetResponseHeader(&interfaces.Response{Headers: h}, "X-Foo"))
}

func TestGetResponseHeaders(t *testing.T) {
	assert.NotNil(t, GetResponseHeaders(nil))
	assert.Empty(t, GetResponseHeaders(nil).Get("X-Foo"))
	h := make(http.Header)
	h.Set("X-Foo", "bar")
	out := GetResponseHeaders(&interfaces.Response{Headers: h})
	assert.Equal(t, "bar", out.Get("X-Foo"))
}
