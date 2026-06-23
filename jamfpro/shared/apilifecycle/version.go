// Package apilifecycle provides Jamf Pro SDK API lifecycle primitives:
// semantic version parsing/comparison, runtime deprecation warnings, and a
// removal guard that errors when a capability has been removed from the
// connected Jamf Pro server version.
//
// The package deliberately does not import jamfpro/client; the dependency
// direction is client -> apilifecycle only, avoiding an import cycle. The
// removal guard depends on the narrow ServerVersionProvider interface, which
// the client.Client (and the test mock) satisfy.
package apilifecycle

import (
	"fmt"
	"strconv"
	"strings"
)

// Version is a parsed Jamf Pro semantic version (major.minor.patch). Any
// trailing build or qualifier segment (e.g. the "-t1234" in "11.28.0-t1234")
// is ignored.
type Version struct {
	Major int
	Minor int
	Patch int
}

// Parse parses "11.28.0", "11.28", or "11" into a Version. Missing segments
// default to 0. Each segment is read up to its first non-digit, so a build
// qualifier such as "0-t1776264729651" parses to 0. An empty string is the
// only error case.
func Parse(s string) (Version, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return Version{}, fmt.Errorf("apilifecycle: empty version string")
	}

	parts := strings.Split(s, ".")
	segment := func(i int) int {
		if i >= len(parts) {
			return 0
		}
		seg := parts[i]
		end := 0
		for end < len(seg) && seg[end] >= '0' && seg[end] <= '9' {
			end++
		}
		if end == 0 {
			return 0
		}
		n, _ := strconv.Atoi(seg[:end])
		return n
	}

	return Version{Major: segment(0), Minor: segment(1), Patch: segment(2)}, nil
}

// MustParse is Parse that panics on error. Use it only for compile-time
// constant versions defined in SDK source (removal/deprecation markers), never
// for runtime server input.
func MustParse(s string) Version {
	v, err := Parse(s)
	if err != nil {
		panic(err)
	}
	return v
}

// Compare returns -1 if v < other, 0 if equal, +1 if v > other.
func (v Version) Compare(other Version) int {
	switch {
	case v.Major != other.Major:
		return sign(v.Major - other.Major)
	case v.Minor != other.Minor:
		return sign(v.Minor - other.Minor)
	case v.Patch != other.Patch:
		return sign(v.Patch - other.Patch)
	default:
		return 0
	}
}

// AtLeast reports whether v >= other.
func (v Version) AtLeast(other Version) bool { return v.Compare(other) >= 0 }

// String renders the version as "major.minor.patch".
func (v Version) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}

func sign(n int) int {
	if n < 0 {
		return -1
	}
	return 1
}
