package shared

import (
	"fmt"
	"io"
	"sync"
	"time"
)

type ProgressReader struct {
	reader       io.Reader
	totalSize    int64
	bytesRead    int64
	lastReported int64
	fileName     string
	mu           sync.Mutex
	startTime    time.Time
}

func NewProgressReader(reader io.Reader, totalSize int64, fileName string) *ProgressReader {
	return &ProgressReader{
		reader:    reader,
		totalSize: totalSize,
		fileName:  fileName,
		startTime: time.Now(),
	}
}

func (pr *ProgressReader) Read(p []byte) (int, error) {
	n, err := pr.reader.Read(p)
	
	pr.mu.Lock()
	pr.bytesRead += int64(n)
	currentBytes := pr.bytesRead
	pr.mu.Unlock()

	pr.reportProgress(currentBytes)

	return n, err
}

func (pr *ProgressReader) reportProgress(currentBytes int64) {
	const reportIntervalMB = 1024 * 1024

	pr.mu.Lock()
	defer pr.mu.Unlock()

	if currentBytes-pr.lastReported >= reportIntervalMB || currentBytes == pr.totalSize {
		pr.lastReported = currentBytes
		
		currentMB := float64(currentBytes) / (1024 * 1024)
		totalMB := float64(pr.totalSize) / (1024 * 1024)
		percent := float64(currentBytes) / float64(pr.totalSize) * 100
		elapsed := time.Since(pr.startTime).Seconds()
		
		var speed float64
		if elapsed > 0 {
			speed = currentMB / elapsed
		}

		if currentBytes == pr.totalSize {
			fmt.Printf("\r✓ Upload complete: %s (%.1f MB in %.1fs, avg %.1f MB/s)\n", 
				pr.fileName, totalMB, elapsed, speed)
		} else {
			fmt.Printf("\rUploading %s: %.1f/%.1f MB (%.1f%%, %.1f MB/s)", 
				pr.fileName, currentMB, totalMB, percent, speed)
		}
	}
}
