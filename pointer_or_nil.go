package pointer

import "time"

// IntOrNil is like Int but returns nil if v is 0.
func IntOrNil(v int) *int {
	if v == 0 {
		return nil
	}
	return &v
}

// Int8OrNil is like Int8 but returns nil if v is 0.
func Int8OrNil(v int8) *int8 {
	if v == 0 {
		return nil
	}
	return &v
}

// Int16OrNil is like Int but returns nil if v is 0.
func Int16OrNil(v int16) *int16 {
	if v == 0 {
		return nil
	}
	return &v
}

// Int32OrNil is like Int32 but returns nil if v is 0.
func Int32OrNil(v int32) *int32 {
	if v == 0 {
		return nil
	}
	return &v
}

// Int64OrNil is like Int64 but returns nil if v is 0.
func Int64OrNil(v int64) *int64 {
	if v == 0 {
		return nil
	}
	return &v
}

// UintOrNil is like Uint but returns nil if v is 0.
func UintOrNil(v uint) *uint {
	if v == 0 {
		return nil
	}
	return &v
}

// Uint8OrNil is like Uint8 but returns nil if v is 0.
func Uint8OrNil(v uint8) *uint8 {
	if v == 0 {
		return nil
	}
	return &v
}

// Uint16OrNil is like Uint16 but returns nil if v is 0.
func Uint16OrNil(v uint16) *uint16 {
	if v == 0 {
		return nil
	}
	return &v
}

// Uint32OrNil is like Uint32 but returns nil if v is 0.
func Uint32OrNil(v uint32) *uint32 {
	if v == 0 {
		return nil
	}
	return &v
}

// Uint64OrNil is like Uint64 but returns nil if v is 0.
func Uint64OrNil(v uint64) *uint64 {
	if v == 0 {
		return nil
	}
	return &v
}

// StringOrNil is like String but returns nil if v is "".
func StringOrNil(v string) *string {
	if v == "" {
		return nil
	}
	return &v
}

// Float32OrNil is like Float32 but returns nil if v is 0.
func Float32OrNil(v float32) *float32 {
	if v == 0 {
		return nil
	}
	return &v
}

// Float64OrNil is like Float64 but returns nil if v is 0.
func Float64OrNil(v float64) *float64 {
	if v == 0 {
		return nil
	}
	return &v
}

// BoolOrNil is like Int but returns nil if v is 0.
func BoolOrNil(v bool) *bool {
	if !v {
		return nil
	}
	return &v
}

// ByteOrNil is like Byte but returns nil if v is 0.
func ByteOrNil(v byte) *byte {
	if v == 0 {
		return nil
	}
	return &v
}

// RuneOrNil is like Byte but returns nil if v is 0.
func RuneOrNil(v rune) *rune {
	if v == 0 {
		return nil
	}
	return &v
}

// TimeOrNil is like Time but returns nil if v is the zero value of time.Time.
func TimeOrNil(v time.Time) *time.Time {
	if v.IsZero() {
		return nil
	}
	return &v
}
