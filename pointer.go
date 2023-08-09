package pointer

import "time"

// Int returns a pointer to an int whose value is v.
func Int(v int) *int {
	return &v
}

// Int8 returns a pointer to an int8 whose value is v.
func Int8(v int8) *int8 {
	return &v
}

// Int16 returns a pointer to an int16 whose value is v.
func Int16(v int16) *int16 {
	return &v
}

// Int32 returns a pointer to an int32 whose value is v.
func Int32(v int32) *int32 {
	return &v
}

// Int64 returns a pointer to an int64 whose value is v.
func Int64(v int64) *int64 {
	return &v
}

// Uint returns a pointer to an uint whose value is v.
func Uint(v int) *int {
	return &v
}

// Uint8 returns a pointer to an uint8 whose value is v.
func Uint8(v uint8) *uint8 {
	return &v
}

// Uint16 returns a pointer to an uint16 whose value is v.
func Uint16(v uint16) *uint16 {
	return &v
}

// Uint32 returns a pointer to an uint32 whose value is v.
func Uint32(v uint32) *uint32 {
	return &v
}

// Uint64 returns a pointer to an uint64 whose value is v.
func Uint64(v uint64) *uint64 {
	return &v
}

// String returns a pointer to a string whose value is v.
func String(v string) *string {
	return &v
}

// Float32 returns a pointer to a string whose value is v.
func Float32(v float32) *float32 {
	return &v
}

// Float64 returns a pointer to a string whose value is v.
func Float64(v float64) *float64 {
	return &v
}

// Bool returns a pointer to a string whose value is v.
func Bool(v bool) *bool {
	return &v
}

// Byte returns a pointer to a byte whose value is v.
func Byte(v byte) *byte {
	return &v
}

// Rune returns a pointer to a rune whose value is v.
func Rune(v rune) *rune {
	return &v
}

// Time returns a pointer to a time.Time whose value is v.
func Time(v time.Time) *time.Time {
	return &v
}

// Complex64 returns a pointer to a complex64 whose value is v.
func Complex64(v complex64) *complex64 {
	return &v
}

// Complex128 returns a pointer to a complex128 whose value is v.
func Complex128(v complex128) *complex128 {
	return &v
}
