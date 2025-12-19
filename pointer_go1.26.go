//go:build go1.26
// +build go1.26

package pointer

import "time"

// Int returns a pointer to an int whose value is v.
//
// Deprecated: Use new(v) instead.
//
//go:fix inline
func Int(v int) *int {
	return new(v)
}

// Int8 returns a pointer to an int8 whose value is v.
//
// Deprecated: Use new(v) instead.
//
//go:fix inline
func Int8(v int8) *int8 {
	return new(v)
}

// Int16 returns a pointer to an int16 whose value is v.
//
// Deprecated: Use new(v) instead.
//
//go:fix inline
func Int16(v int16) *int16 {
	return new(v)
}

// Int32 returns a pointer to an int32 whose value is v.
//
// Deprecated: Use new(v) instead.
//
//go:fix inline
func Int32(v int32) *int32 {
	return new(v)
}

// Int64 returns a pointer to an int64 whose value is v.
//
// Deprecated: Use new(v) instead.
//
//go:fix inline
func Int64(v int64) *int64 {
	return new(v)
}

// Uint returns a pointer to an uint whose value is v.
//
// Deprecated: Use new(v) instead.
//
//go:fix inline
func Uint(v uint) *uint {
	return new(v)
}

// Uint8 returns a pointer to an uint8 whose value is v.
//
// Deprecated: Use new(v) instead.
//
//go:fix inline
func Uint8(v uint8) *uint8 {
	return new(v)
}

// Uint16 returns a pointer to an uint16 whose value is v.
//
// Deprecated: Use new(v) instead.
//
//go:fix inline
func Uint16(v uint16) *uint16 {
	return new(v)
}

// Uint32 returns a pointer to an uint32 whose value is v.
//
// Deprecated: Use new(v) instead.
//
//go:fix inline
func Uint32(v uint32) *uint32 {
	return new(v)
}

// Uint64 returns a pointer to an uint64 whose value is v.
//
// Deprecated: Use new(v) instead.
//
//go:fix inline
func Uint64(v uint64) *uint64 {
	return new(v)
}

// String returns a pointer to a string whose value is v.
//
// Deprecated: Use new(v) instead.
//
//go:fix inline
func String(v string) *string {
	return new(v)
}

// Float32 returns a pointer to a string whose value is v.
//
// Deprecated: Use new(v) instead.
//
//go:fix inline
func Float32(v float32) *float32 {
	return new(v)
}

// Float64 returns a pointer to a string whose value is v.
//
// Deprecated: Use new(v) instead.
//
//go:fix inline
func Float64(v float64) *float64 {
	return new(v)
}

// Bool returns a pointer to a string whose value is v.
//
// Deprecated: Use new(v) instead.
//
//go:fix inline
func Bool(v bool) *bool {
	return new(v)
}

// Byte returns a pointer to a byte whose value is v.
//
// Deprecated: Use new(v) instead.
//
//go:fix inline
func Byte(v byte) *byte {
	return new(v)
}

// Rune returns a pointer to a rune whose value is v.
//
// Deprecated: Use new(v) instead.
//
//go:fix inline
func Rune(v rune) *rune {
	return new(v)
}

// Time returns a pointer to a time.Time whose value is v.
//
// Deprecated: Use new(v) instead.
//
//go:fix inline
func Time(v time.Time) *time.Time {
	return new(v)
}

// Complex64 returns a pointer to a complex64 whose value is v.
//
// Deprecated: Use new(v) instead.
//
//go:fix inline
func Complex64(v complex64) *complex64 {
	return new(v)
}

// Complex128 returns a pointer to a complex128 whose value is v.
//
// Deprecated: Use new(v) instead.
//
//go:fix inline
func Complex128(v complex128) *complex128 {
	return new(v)
}
