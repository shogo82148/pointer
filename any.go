//go:build go1.18
// +build go1.18

package pointer

// Int returns a pointer whose value is v.
func Any[T any](v T) *T {
	return &v
}

// AnyOrNil is like Any but returns nil if v is the zero value.
func AnyOrNil[T any](v T) *T {
	var zero T
	if v == zero {
		return nil
	}
	return &v
}

// IntValueWithDefault is like *p but it returns def if p is nil.
func AnyValueWithDefault[T any](p *T, def T) T {
	if p == nil {
		return def
	}
	return *p
}

// AnyValue is like *p but it returns the zero value if p is nil.
func AnyValue[T any](p *T) T {
	if p == nil {
		var zero T
		return zero
	}
	return *p
}
