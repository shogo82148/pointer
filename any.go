//go:build go1.18
// +build go1.18

package pointer

// Ptr returns a pointer whose value is v.
func Ptr[T any](v T) *T {
	return &v
}

// PtrOrNil is like Any but returns nil if v is the zero value.
func PtrOrNil[T comparable](v T) *T {
	var zero T
	if v == zero {
		return nil
	}
	return &v
}

// ValueWithDefault is like *p but it returns def if p is nil.
func ValueWithDefault[T any](p *T, def T) T {
	if p == nil {
		return def
	}
	return *p
}

// Value is like *p but it returns the zero value if p is nil.
func Value[T any](p *T) T {
	if p == nil {
		var zero T
		return zero
	}
	return *p
}
