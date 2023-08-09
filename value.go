package pointer

import "time"

// IntValue is like *p but it returns 0 if p is nil.
func IntValue(p *int) int {
	if p == nil {
		return 0
	}
	return *p
}

// Int8Value is like *p but it returns 0 if p is nil.
func Int8Value(p *int8) int8 {
	if p == nil {
		return 0
	}
	return *p
}

// Int16Value is like *p but it returns 0 if p is nil.
func Int16Value(p *int16) int16 {
	if p == nil {
		return 0
	}
	return *p
}

// Int32Value is like *p but it returns 0 if p is nil.
func Int32Value(p *int32) int32 {
	if p == nil {
		return 0
	}
	return *p
}

// Int64Value is like *p but it returns 0 if p is nil.
func Int64Value(p *int64) int64 {
	if p == nil {
		return 0
	}
	return *p
}

// UintValue is like *p but it returns 0 if p is nil.
func UintValue(p *uint) uint {
	if p == nil {
		return 0
	}
	return *p
}

// Uint8Value is like *p but it returns 0 if p is nil.
func Uint8Value(p *uint8) uint8 {
	if p == nil {
		return 0
	}
	return *p
}

// Uint16Value is like *p but it returns 0 if p is nil.
func Uint16Value(p *uint16) uint16 {
	if p == nil {
		return 0
	}
	return *p
}

// Uint32Value is like *p but it returns 0 if p is nil.
func Uint32Value(p *uint32) uint32 {
	if p == nil {
		return 0
	}
	return *p
}

// Uint64Value is like *p but it returns 0 if p is nil.
func Uint64Value(p *uint64) uint64 {
	if p == nil {
		return 0
	}
	return *p
}

// StringValue is like *p but it returns "" if p is nil.
func StringValue(p *string) string {
	if p == nil {
		return ""
	}
	return *p
}

// Float32Value is like *p but it returns 0 if p is nil.
func Float32Value(p *float32) float32 {
	if p == nil {
		return 0
	}
	return *p
}

// Float64Value is like *p but it returns 0 if p is nil.
func Float64Value(p *float64) float64 {
	if p == nil {
		return 0
	}
	return *p
}

// BoolValue is like *p but it returns false if p is nil.
func BoolValue(p *bool) bool {
	if p == nil {
		return false
	}
	return *p
}

// ByteValue is like *p but it returns 0 if p is nil.
func ByteValue(p *byte) byte {
	if p == nil {
		return 0
	}
	return *p
}

// RuneValue is like *p but it returns 0 if p is nil.
func RuneValue(p *rune) rune {
	if p == nil {
		return 0
	}
	return *p
}

// TimeValue is like *p but it returns time.Time{} if p is nil.
func TimeValue(p *time.Time) time.Time {
	if p == nil {
		return time.Time{}
	}
	return *p
}

// Complex64Value is like *p but it returns 0 if p is nil.
func Complex64Value(p *complex64) complex64 {
	if p == nil {
		return 0
	}
	return *p
}

// Complex128Value is like *p but it returns 0 if p is nil.
func Complex128Value(p *complex128) complex128 {
	if p == nil {
		return 0
	}
	return *p
}
