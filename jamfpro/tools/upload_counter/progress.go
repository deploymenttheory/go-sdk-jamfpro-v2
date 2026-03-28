// Package upload_counter provides a terminal progress bar for multipart file uploads.
package upload_counter

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

const barWidth = 30

// Bar writes an upload progress bar to w (defaults to os.Stderr when nil).
// Call NewBar once per upload, then pass Bar.Callback to SetMultipartFile.
//
// Example output (overwrites same line):
//
//	Uploading Firefox_147.0.pkg  [████████████░░░░░░░░░░░░░░░░]  45.2 MB / 98.7 MB  45.7%
type Bar struct {
	w    io.Writer
	mu   sync.Mutex
	done bool
}

// New creates a new Bar that writes to w. Pass nil to write to os.Stderr.
func New(w io.Writer) *Bar {
	if w == nil {
		w = os.Stderr
	}
	return &Bar{w: w}
}

// Callback satisfies client.MultipartProgressCallback.
// Pass this to SetMultipartFile as the progress callback.
func (b *Bar) Callback(fieldName, fileName string, written, total int64) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if total <= 0 {
		return
	}

	pct := float64(written) / float64(total)
	filled := int(pct * barWidth)
	if filled > barWidth {
		filled = barWidth
	}
	empty := barWidth - filled

	bar := strings.Repeat("█", filled) + strings.Repeat("░", empty)
	writtenMB := float64(written) / (1024 * 1024)
	totalMB := float64(total) / (1024 * 1024)

	fmt.Fprintf(b.w, "\rUploading %-30s [%s]  %6.1f MB / %.1f MB  %5.1f%%",
		fileName, bar, writtenMB, totalMB, pct*100)

	if written >= total && !b.done {
		b.done = true
		fmt.Fprintln(b.w)
	}
}
