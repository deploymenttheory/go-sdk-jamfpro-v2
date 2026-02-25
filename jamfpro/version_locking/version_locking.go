// Package version_locking provides helpers for Jamf Pro API optimistic locking.
//
// The Jamf Pro API uses a versionLock integer field to prevent concurrent
// modifications to sensitive resources (e.g. computer prestages, mobile device
// prestages).
//
// Workflow:
//  1. GET the resource – the response body contains the current versionLock.
//  2. Store the versionLock value.
//  3. PUT / POST / DELETE – echo the stored versionLock back in the request body.
//
// If another process modifies the resource between the GET and the write, the
// supplied versionLock becomes stale and the API returns a conflict error.  The
// caller must re-fetch and retry.
//
// New resources (POST) must carry versionLock: 0.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/docs/optimistic-locking
package version_locking

// NewResourceVersionLock is the versionLock value that must be supplied when
// creating a new resource via POST.
const NewResourceVersionLock = 0

// VersionLocker is implemented by any Jamf Pro resource that participates in
// optimistic locking via a top-level versionLock field.
type VersionLocker interface {
	// GetVersionLock returns the current versionLock held by this resource.
	GetVersionLock() int
	// SetVersionLock writes a versionLock value onto this resource.
	SetVersionLock(lock int)
}

// PropagateVersionLock copies the versionLock from current into request when
// request's versionLock is still zero (i.e. the caller has not explicitly
// provided one).
//
// Use this in flows where the service performs the GET internally and wants to
// forward the lock to the subsequent PUT only when the caller has not already
// set one.
func PropagateVersionLock(current, request VersionLocker) {
	if request.GetVersionLock() == 0 {
		request.SetVersionLock(current.GetVersionLock())
	}
}

// EnsureVersionLock always copies the versionLock from current into request,
// overwriting any value the caller may have set.
//
// Use this when the service owns the GET and must guarantee the request always
// carries the freshest lock – for example in "update by name" flows where a
// caller-supplied stale lock would cause a 409 conflict.
func EnsureVersionLock(current, request VersionLocker) {
	request.SetVersionLock(current.GetVersionLock())
}
