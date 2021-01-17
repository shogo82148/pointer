package pointer

import (
	"testing"
	"time"
)

func TestIntValueWithDefault(t *testing.T) {
	var v int = 42
	var got int

	got = IntValueWithDefault(&v, 0)
	if got != 42 {
		t.Errorf("want %d, got %d", 42, got)
	}

	got = IntValueWithDefault(nil, 42)
	if got != 42 {
		t.Errorf("want %d, got %d", 42, got)
	}
}

func TestInt8ValueWithDefault(t *testing.T) {
	var v int8 = 42
	var got int8

	got = Int8ValueWithDefault(&v, 0)
	if got != 42 {
		t.Errorf("want %d, got %d", 42, got)
	}

	got = Int8ValueWithDefault(nil, 42)
	if got != 42 {
		t.Errorf("want %d, got %d", 42, got)
	}
}

func TestInt16ValueWithDefault(t *testing.T) {
	var v int16 = 42
	var got int16

	got = Int16ValueWithDefault(&v, 0)
	if got != 42 {
		t.Errorf("want %d, got %d", 42, got)
	}

	got = Int16ValueWithDefault(nil, 42)
	if got != 42 {
		t.Errorf("want %d, got %d", 42, got)
	}
}

func TestInt32ValueWithDefault(t *testing.T) {
	var v int32 = 42
	var got int32

	got = Int32ValueWithDefault(&v, 0)
	if got != 42 {
		t.Errorf("want %d, got %d", 42, got)
	}

	got = Int32ValueWithDefault(nil, 42)
	if got != 42 {
		t.Errorf("want %d, got %d", 42, got)
	}
}

func TestInt64ValueWithDefault(t *testing.T) {
	var v int64 = 42
	var got int64

	got = Int64ValueWithDefault(&v, 0)
	if got != 42 {
		t.Errorf("want %d, got %d", 42, got)
	}

	got = Int64ValueWithDefault(nil, 42)
	if got != 42 {
		t.Errorf("want %d, got %d", 42, got)
	}
}

func TestUintValueWithDefault(t *testing.T) {
	var v uint = 42
	var got uint

	got = UintValueWithDefault(&v, 0)
	if got != 42 {
		t.Errorf("want %d, got %d", 42, got)
	}

	got = UintValueWithDefault(nil, 42)
	if got != 42 {
		t.Errorf("want %d, got %d", 42, got)
	}
}

func TestUint8ValueWithDefault(t *testing.T) {
	var v uint8 = 42
	var got uint8

	got = Uint8ValueWithDefault(&v, 0)
	if got != 42 {
		t.Errorf("want %d, got %d", 42, got)
	}

	got = Uint8ValueWithDefault(nil, 42)
	if got != 42 {
		t.Errorf("want %d, got %d", 42, got)
	}
}

func TestUint16ValueWithDefault(t *testing.T) {
	var v uint16 = 42
	var got uint16

	got = Uint16ValueWithDefault(&v, 0)
	if got != 42 {
		t.Errorf("want %d, got %d", 42, got)
	}

	got = Uint16ValueWithDefault(nil, 42)
	if got != 42 {
		t.Errorf("want %d, got %d", 42, got)
	}
}

func TestUint32ValueWithDefault(t *testing.T) {
	var v uint32 = 42
	var got uint32

	got = Uint32ValueWithDefault(&v, 0)
	if got != 42 {
		t.Errorf("want %d, got %d", 42, got)
	}

	got = Uint32ValueWithDefault(nil, 42)
	if got != 42 {
		t.Errorf("want %d, got %d", 42, got)
	}
}

func TestUint64ValueWithDefault(t *testing.T) {
	var v uint64 = 42
	var got uint64

	got = Uint64ValueWithDefault(&v, 0)
	if got != 42 {
		t.Errorf("want %d, got %d", 42, got)
	}

	got = Uint64ValueWithDefault(nil, 42)
	if got != 42 {
		t.Errorf("want %d, got %d", 42, got)
	}
}

func TestStringValueWithDefault(t *testing.T) {
	var v string = "foo"
	var got string

	got = StringValueWithDefault(&v, "")
	if got != "foo" {
		t.Errorf("want %q, got %q", "foo", got)
	}

	got = StringValueWithDefault(nil, "bar")
	if got != "bar" {
		t.Errorf("want %q, got %q", "bar", got)
	}
}

func TestFloat32ValueWithDefault(t *testing.T) {
	var v float32 = 42
	var got float32

	got = Float32ValueWithDefault(&v, 0.0)
	if got != 42.0 {
		t.Errorf("want %f, got %f", 42.0, got)
	}

	got = Float32ValueWithDefault(nil, 42.0)
	if got != 42.0 {
		t.Errorf("want %f, got %f", 42.0, got)
	}
}

func TestFloat64ValueWithDefault(t *testing.T) {
	var v float64 = 42
	var got float64

	got = Float64ValueWithDefault(&v, 0)
	if got != 42.0 {
		t.Errorf("want %f, got %f", 42.0, got)
	}

	got = Float64ValueWithDefault(nil, 42.0)
	if got != 42.0 {
		t.Errorf("want %f, got %f", 42.0, got)
	}
}

func TestByteValueWithDefault(t *testing.T) {
	var v byte = 42
	var got byte

	got = ByteValueWithDefault(&v, 0)
	if got != 42 {
		t.Errorf("want %d, got %d", 42, got)
	}

	got = ByteValueWithDefault(nil, 42)
	if got != 42 {
		t.Errorf("want %d, got %d", 42, got)
	}
}

func TestRuneValueWithDefault(t *testing.T) {
	var v rune = 42
	var got rune

	got = RuneValueWithDefault(&v, 0)
	if got != 42 {
		t.Errorf("want %d, got %d", 42, got)
	}

	got = RuneValueWithDefault(nil, 'a')
	if got != 'a' {
		t.Errorf("want %d, got %d", 'a', got)
	}
}

func TestBoolValueWithDefault(t *testing.T) {
	var v bool = true
	var got bool

	got = BoolValueWithDefault(&v, false)
	if got != true {
		t.Errorf("want %t, got %t", true, got)
	}

	got = BoolValueWithDefault(nil, true)
	if got != true {
		t.Errorf("want %t, got %t", true, got)
	}
}

func TestTimeValueWithDefault(t *testing.T) {
	v := time.Date(2020, time.January, 17, 2, 19, 0, 0, time.UTC)
	var got time.Time

	got = TimeValueWithDefault(&v, time.Time{})
	if !got.Equal(v) {
		t.Errorf("want %s, got %s", v, got)
	}

	got = TimeValueWithDefault(nil, v)
	if !got.Equal(v) {
		t.Errorf("want %s, got %s", v, got)
	}
}
