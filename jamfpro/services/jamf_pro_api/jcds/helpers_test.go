package jcds

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_readJCDSPackageTypes_Success(t *testing.T) {
	tmp := t.TempDir()
	pkgPath := filepath.Join(tmp, "test.pkg")
	require.NoError(t, os.WriteFile(pkgPath, []byte("test content"), 0644))

	reader, size, err := readJCDSPackageTypes(pkgPath)
	require.NoError(t, err)
	require.NotNil(t, reader)
	assert.Equal(t, int64(12), size)
}

func Test_readJCDSPackageTypes_InvalidExtension(t *testing.T) {
	tmp := t.TempDir()
	txtPath := filepath.Join(tmp, "test.txt")
	require.NoError(t, os.WriteFile(txtPath, []byte("test"), 0644))

	_, _, err := readJCDSPackageTypes(txtPath)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "not allowed")
}

func Test_readJCDSPackageTypes_FileNotFound(t *testing.T) {
	_, _, err := readJCDSPackageTypes("/nonexistent/path/test.pkg")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to read")
}

func Test_readJCDSPackageTypes_DMG(t *testing.T) {
	tmp := t.TempDir()
	dmgPath := filepath.Join(tmp, "test.dmg")
	content := []byte("dmg content")
	require.NoError(t, os.WriteFile(dmgPath, content, 0644))

	reader, size, err := readJCDSPackageTypes(dmgPath)
	require.NoError(t, err)
	require.NotNil(t, reader)
	assert.Equal(t, int64(len(content)), size)
}

func Test_readJCDSPackageTypes_ZIP(t *testing.T) {
	tmp := t.TempDir()
	zipPath := filepath.Join(tmp, "test.zip")
	content := []byte("zip content")
	require.NoError(t, os.WriteFile(zipPath, content, 0644))

	reader, size, err := readJCDSPackageTypes(zipPath)
	require.NoError(t, err)
	require.NotNil(t, reader)
	assert.Equal(t, int64(len(content)), size)
}

func Test_safeReadJCDSPackageFile_InvalidExtension(t *testing.T) {
	tmp := t.TempDir()
	txtPath := filepath.Join(tmp, "test.txt")
	require.NoError(t, os.WriteFile(txtPath, []byte("x"), 0644))

	_, err := safeReadJCDSPackageFile(txtPath, []string{".pkg", ".dmg", ".zip"})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "not allowed")
}

func Test_safeReadJCDSPackageFile_Success(t *testing.T) {
	tmp := t.TempDir()
	pkgPath := filepath.Join(tmp, "test.pkg")
	content := []byte("pkg data")
	require.NoError(t, os.WriteFile(pkgPath, content, 0644))

	data, err := safeReadJCDSPackageFile(pkgPath, []string{".pkg", ".dmg", ".zip"})
	require.NoError(t, err)
	assert.Equal(t, content, data)
}

func Test_safeReadJCDSPackageFile_ReadFileError(t *testing.T) {
	tmp := t.TempDir()
	dirPath := filepath.Join(tmp, "test.pkg")
	require.NoError(t, os.Mkdir(dirPath, 0755))

	_, err := safeReadJCDSPackageFile(dirPath, []string{".pkg", ".dmg", ".zip"})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to read Jamf Pro package")
}

func Test_resolveSymlinks_Success(t *testing.T) {
	tmp := t.TempDir()
	realPath := filepath.Join(tmp, "real.pkg")
	require.NoError(t, os.WriteFile(realPath, []byte("x"), 0644))
	linkPath := filepath.Join(tmp, "link.pkg")
	require.NoError(t, os.Symlink(realPath, linkPath))

	resolved, err := resolveSymlinks(linkPath)
	require.NoError(t, err)
	assert.NotEmpty(t, resolved)
}

func Test_resolveSymlinks_InvalidPath(t *testing.T) {
	_, err := resolveSymlinks("/nonexistent/symlink")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unable to resolve")
}

func Test_cleanPath(t *testing.T) {
	assert.Equal(t, "a/b/c", cleanPath("a/b/c"))
	assert.Equal(t, "a/b", cleanPath("a/./b"))
	assert.Equal(t, "b", cleanPath("a/../b"))
}

func Test_isValidExtension(t *testing.T) {
	allowed := []string{".pkg", ".dmg", ".zip"}
	assert.True(t, isValidExtension("/path/to/file.pkg", allowed))
	assert.True(t, isValidExtension("/path/to/file.PKG", allowed))
	assert.True(t, isValidExtension("/path/to/file.dmg", allowed))
	assert.True(t, isValidExtension("/path/to/file.zip", allowed))
	assert.False(t, isValidExtension("/path/to/file.txt", allowed))
	assert.False(t, isValidExtension("/path/to/file", allowed))
}
