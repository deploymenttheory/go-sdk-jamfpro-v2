package version_locking

import "reflect"

// versionLockFieldName is the Go struct field name that carries a Jamf Pro
// optimistic-locking token. Jamf marshals it as "versionLock".
const versionLockFieldName = "VersionLock"

// SyncAll copies every versionLock value found in current onto the matching
// field in request, walking the entire struct tree.
//
// Jamf Pro resources carry version locks at more than one level: the resource
// itself plus each nested subset (location information, purchasing information,
// account settings, ...). Every one of them must echo the value returned by the
// most recent GET, or the write is rejected with OPTIMISTIC_LOCK_FAILED.
//
// Enumerating those fields by hand at each call site is what lets a newly added
// subset silently ship with an unsynchronised lock, so SyncAll discovers them
// reflectively instead. Fields are matched by name and traversed in parallel;
// a field present in one struct but not the other is skipped.
//
// current and request must be pointers to structs of the same type. Anything
// else is a no-op.
func SyncAll(current, request any) {
	c, r := reflect.ValueOf(current), reflect.ValueOf(request)
	if !c.IsValid() || !r.IsValid() || c.Type() != r.Type() {
		return
	}
	syncValue(c, r)
}

// ZeroAll sets every versionLock in the struct tree to NewResourceVersionLock.
//
// Jamf requires a lock of 0 on POST: a create carries no prior state to echo,
// and a non-zero value is rejected.
func ZeroAll(request any) {
	v := reflect.ValueOf(request)
	if !v.IsValid() {
		return
	}
	zeroValue(v)
}

// TopLock reports the resource-level versionLock, and whether one was found.
// Used to detect whether a write landed when the API reports an error: a lock
// that has advanced past the value we submitted means the server accepted it.
func TopLock(v any) (int, bool) {
	rv := reflect.ValueOf(v)
	for rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface {
		if rv.IsNil() {
			return 0, false
		}
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
		return 0, false
	}
	f := rv.FieldByName(versionLockFieldName)
	if !f.IsValid() || !isIntKind(f.Kind()) {
		return 0, false
	}
	return int(f.Int()), true
}

// syncValue walks current and request in lockstep, copying version locks.
func syncValue(current, request reflect.Value) {
	for current.Kind() == reflect.Ptr || current.Kind() == reflect.Interface {
		if current.IsNil() || request.IsNil() {
			return
		}
		current, request = current.Elem(), request.Elem()
	}

	switch current.Kind() {
	case reflect.Struct:
		t := current.Type()
		for i := range t.NumField() {
			if t.Field(i).PkgPath != "" { // unexported
				continue
			}
			cf, rf := current.Field(i), request.Field(i)
			if t.Field(i).Name == versionLockFieldName && isIntKind(cf.Kind()) {
				if rf.CanSet() {
					rf.SetInt(cf.Int())
				}
				continue
			}
			syncValue(cf, rf)
		}
	case reflect.Slice, reflect.Array:
		// Positional match is the only correlation available; anything the
		// request adds beyond what the server returned keeps its own value.
		n := min(current.Len(), request.Len())
		for i := range n {
			syncValue(current.Index(i), request.Index(i))
		}
	}
}

// zeroValue sets every version lock in the tree to NewResourceVersionLock.
func zeroValue(v reflect.Value) {
	for v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		if v.IsNil() {
			return
		}
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Struct:
		t := v.Type()
		for i := range t.NumField() {
			if t.Field(i).PkgPath != "" {
				continue
			}
			f := v.Field(i)
			if t.Field(i).Name == versionLockFieldName && isIntKind(f.Kind()) {
				if f.CanSet() {
					f.SetInt(NewResourceVersionLock)
				}
				continue
			}
			zeroValue(f)
		}
	case reflect.Slice, reflect.Array:
		for i := range v.Len() {
			zeroValue(v.Index(i))
		}
	}
}

func isIntKind(k reflect.Kind) bool {
	switch k {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return true
	}
	return false
}
