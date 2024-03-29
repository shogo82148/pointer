//go:build go1.18
// +build go1.18

package pointer

// Ptr returns a pointer whose value is v.
func Ptr[T any](v T) *T {
	return &v
}

// PtrOrNil is like Any but returns nil if v is the zero value.
func PtrOrNil[T comparable](v T) *T {
	if v == Zero[T]() {
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
		return Zero[T]()
	}
	return *p
}

// Equal returns *a == *b but it also accepts nil pointers.
// Special cases are:
//
//	Equal(nil, nil) = true
//	Equal(nil, &v) = false
//	Equal(&v, nil) = false
func Equal[T comparable](a, b *T) bool {
	return a == b || (a != nil && b != nil && *a == *b)
}

// ShallowCopy returns a shallow copy of v.
// If v is nil, it returns a nil pointer.
func ShallowCopy[T any](v *T) *T {
	if v == nil {
		return nil
	}
	u := *v
	return &u
}

// Zero returns the zero value of type T.
func Zero[T any]() T {
	var zero T
	return zero
}
