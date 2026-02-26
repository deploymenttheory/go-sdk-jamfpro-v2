package jcds

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProgressReader_Read(t *testing.T) {
	data := []byte("hello world")
	reader := bytes.NewReader(data)
	var progressCalls int
	progressFn := func(readBytes, totalBytes int64, unit string) {
		progressCalls++
		_, _, _ = readBytes, totalBytes, unit
	}

	pr := &ProgressReader{
		reader:     reader,
		totalBytes: int64(len(data)),
		progressFn: progressFn,
	}

	out := make([]byte, 20)
	n, err := pr.Read(out)
	require.NoError(t, err)
	assert.Equal(t, len(data), n)
	assert.Equal(t, data, out[:n])
	assert.Greater(t, progressCalls, 0)
}

func TestProgressReader_ReadPartial(t *testing.T) {
	// Use 2KB so progress reports in KB (files under 1MB use KB)
	data := make([]byte, 2048)
	for i := range data {
		data[i] = byte(i % 256)
	}
	reader := bytes.NewReader(data)
	var lastRead, lastTotal int64
	var lastUnit string
	progressFn := func(readBytes, totalBytes int64, unit string) {
		lastRead = readBytes
		lastTotal = totalBytes
		lastUnit = unit
	}

	pr := &ProgressReader{
		reader:     reader,
		totalBytes: int64(len(data)),
		progressFn: progressFn,
	}

	out := make([]byte, 1024)
	n, err := pr.Read(out)
	require.NoError(t, err)
	assert.Equal(t, 1024, n)
	assert.Equal(t, int64(1), lastRead)
	assert.Equal(t, int64(2), lastTotal)
	assert.Equal(t, "KB", lastUnit)
}

func TestProgressReader_ReadLargeFileMB(t *testing.T) {
	// Use 2MB to hit the MB progress branch (files > 1MB report in MB)
	data := make([]byte, 2*1024*1024)
	for i := range data {
		data[i] = byte(i % 256)
	}
	reader := bytes.NewReader(data)
	var lastUnit string
	progressFn := func(readBytes, totalBytes int64, unit string) {
		lastUnit = unit
		_, _ = readBytes, totalBytes
	}

	pr := &ProgressReader{
		reader:     reader,
		totalBytes: int64(len(data)),
		progressFn: progressFn,
	}

	out := make([]byte, 1024*1024)
	n, err := pr.Read(out)
	require.NoError(t, err)
	assert.Equal(t, 1024*1024, n)
	assert.Equal(t, "MB", lastUnit)
}
