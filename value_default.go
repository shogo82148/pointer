package pointer

import "time"

// IntValueWithDefault is like *p but it returns def if p is nil.
func IntValueWithDefault(p *int, def int) int {
	if p == nil {
		return def
	}
	return *p
}

// Int8ValueWithDefault is like *p but it returns def if p is nil.
func Int8ValueWithDefault(p *int8, def int8) int8 {
	if p == nil {
		return def
	}
	return *p
}

// Int16ValueWithDefault is like *p but it returns def if p is nil.
func Int16ValueWithDefault(p *int16, def int16) int16 {
	if p == nil {
		return def
	}
	return *p
}

// Int32ValueWithDefault is like *p but it returns def if p is nil.
func Int32ValueWithDefault(p *int32, def int32) int32 {
	if p == nil {
		return def
	}
	return *p
}

// Int64ValueWithDefault is like *p but it returns def if p is nil.
func Int64ValueWithDefault(p *int64, def int64) int64 {
	if p == nil {
		return def
	}
	return *p
}

// UintValueWithDefault is like *p but it returns def if p is nil.
func UintValueWithDefault(p *uint, def uint) uint {
	if p == nil {
		return def
	}
	return *p
}

// Uint8ValueWithDefault is like *p but it returns def if p is nil.
func Uint8ValueWithDefault(p *uint8, def uint8) uint8 {
	if p == nil {
		return def
	}
	return *p
}

// Uint16ValueWithDefault is like *p but it returns def if p is nil.
func Uint16ValueWithDefault(p *uint16, def uint16) uint16 {
	if p == nil {
		return def
	}
	return *p
}

// Uint32ValueWithDefault is like *p but it returns def if p is nil.
func Uint32ValueWithDefault(p *uint32, def uint32) uint32 {
	if p == nil {
		return def
	}
	return *p
}

// Uint64ValueWithDefault is like *p but it returns def if p is nil.
func Uint64ValueWithDefault(p *uint64, def uint64) uint64 {
	if p == nil {
		return def
	}
	return *p
}

// StringValueWithDefault is like *p but it returns def if p is nil.
func StringValueWithDefault(p *string, def string) string {
	if p == nil {
		return def
	}
	return *p
}

// Float32ValueWithDefault is like *p but it returns def if p is nil.
func Float32ValueWithDefault(p *float32, def float32) float32 {
	if p == nil {
		return def
	}
	return *p
}

// Float64ValueWithDefault is like *p but it returns def if p is nil.
func Float64ValueWithDefault(p *float64, def float64) float64 {
	if p == nil {
		return def
	}
	return *p
}

// BoolValueWithDefault is like *p but it returns def if p is nil.
func BoolValueWithDefault(p *bool, def bool) bool {
	if p == nil {
		return def
	}
	return *p
}

// ByteValueWithDefault is like *p but it returns def if p is nil.
func ByteValueWithDefault(p *byte, def byte) byte {
	if p == nil {
		return def
	}
	return *p
}

// RuneValueWithDefault is like *p but it returns def if p is nil.
func RuneValueWithDefault(p *rune, def rune) rune {
	if p == nil {
		return def
	}
	return *p
}

// TimeValueWithDefault is like *p but it returns def if p is nil.
func TimeValueWithDefault(p *time.Time, def time.Time) time.Time {
	if p == nil {
		return def
	}
	return *p
}
