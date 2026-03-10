package crypto

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCalculateSHA3_512(t *testing.T) {
	tmp := t.TempDir()
	f := filepath.Join(tmp, "test.bin")
	require.NoError(t, os.WriteFile(f, []byte("hello"), 0644))

	hash, err := CalculateSHA3_512(f)
	require.NoError(t, err)
	assert.Len(t, hash, 128) // SHA3-512 hex = 128 chars
	assert.Regexp(t, `^[a-f0-9]+$`, hash)
}

func TestCalculateSHA3_512_NotFound(t *testing.T) {
	_, err := CalculateSHA3_512("/nonexistent/path")
	assert.Error(t, err)
}

func TestCalculateMD5(t *testing.T) {
	tmp := t.TempDir()
	f := filepath.Join(tmp, "test.bin")
	require.NoError(t, os.WriteFile(f, []byte("hello"), 0644))

	hash, err := CalculateMD5(f)
	require.NoError(t, err)
	assert.Len(t, hash, 32) // MD5 hex = 32 chars
	assert.Equal(t, "5d41402abc4b2a76b9719d911017c592", hash)
}

func TestCalculateMD5_NotFound(t *testing.T) {
	_, err := CalculateMD5("/nonexistent/path")
	assert.Error(t, err)
}
