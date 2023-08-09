package pointer

import (
	"testing"
	"time"
)

func TestIntValue(t *testing.T) {
	var v int = 42
	var got int

	got = IntValue(&v)
	if got != 42 {
		t.Errorf("want %d, got %d", 42, got)
	}

	got = IntValue(nil)
	if got != 0 {
		t.Errorf("want %d, got %d", 0, got)
	}
}

func TestInt8Value(t *testing.T) {
	var v int8 = 42
	var got int8

	got = Int8Value(&v)
	if got != 42 {
		t.Errorf("want %d, got %d", 42, got)
	}

	got = Int8Value(nil)
	if got != 0 {
		t.Errorf("want %d, got %d", 0, got)
	}
}

func TestInt16Value(t *testing.T) {
	var v int16 = 42
	var got int16

	got = Int16Value(&v)
	if got != 42 {
		t.Errorf("want %d, got %d", 42, got)
	}

	got = Int16Value(nil)
	if got != 0 {
		t.Errorf("want %d, got %d", 0, got)
	}
}

func TestInt32Value(t *testing.T) {
	var v int32 = 42
	var got int32

	got = Int32Value(&v)
	if got != 42 {
		t.Errorf("want %d, got %d", 42, got)
	}

	got = Int32Value(nil)
	if got != 0 {
		t.Errorf("want %d, got %d", 0, got)
	}
}

func TestInt64Value(t *testing.T) {
	var v int64 = 42
	var got int64

	got = Int64Value(&v)
	if got != 42 {
		t.Errorf("want %d, got %d", 42, got)
	}

	got = Int64Value(nil)
	if got != 0 {
		t.Errorf("want %d, got %d", 0, got)
	}
}

func TestUintValue(t *testing.T) {
	var v uint = 42
	var got uint

	got = UintValue(&v)
	if got != 42 {
		t.Errorf("want %d, got %d", 42, got)
	}

	got = UintValue(nil)
	if got != 0 {
		t.Errorf("want %d, got %d", 0, got)
	}
}

func TestUint8Value(t *testing.T) {
	var v uint8 = 42
	var got uint8

	got = Uint8Value(&v)
	if got != 42 {
		t.Errorf("want %d, got %d", 42, got)
	}

	got = Uint8Value(nil)
	if got != 0 {
		t.Errorf("want %d, got %d", 0, got)
	}
}

func TestUint16Value(t *testing.T) {
	var v uint16 = 42
	var got uint16

	got = Uint16Value(&v)
	if got != 42 {
		t.Errorf("want %d, got %d", 42, got)
	}

	got = Uint16Value(nil)
	if got != 0 {
		t.Errorf("want %d, got %d", 0, got)
	}
}

func TestUint32Value(t *testing.T) {
	var v uint32 = 42
	var got uint32

	got = Uint32Value(&v)
	if got != 42 {
		t.Errorf("want %d, got %d", 42, got)
	}

	got = Uint32Value(nil)
	if got != 0 {
		t.Errorf("want %d, got %d", 0, got)
	}
}

func TestUint64Value(t *testing.T) {
	var v uint64 = 42
	var got uint64

	got = Uint64Value(&v)
	if got != 42 {
		t.Errorf("want %d, got %d", 42, got)
	}

	got = Uint64Value(nil)
	if got != 0 {
		t.Errorf("want %d, got %d", 0, got)
	}
}

func TestStringValue(t *testing.T) {
	var v string = "foo"
	var got string

	got = StringValue(&v)
	if got != "foo" {
		t.Errorf("want %q, got %q", "foo", got)
	}

	got = StringValue(nil)
	if got != "" {
		t.Errorf("want %q, got %q", "", got)
	}
}

func TestFloat32Value(t *testing.T) {
	var v float32 = 42
	var got float32

	got = Float32Value(&v)
	if got != 42.0 {
		t.Errorf("want %f, got %f", 42.0, got)
	}

	got = Float32Value(nil)
	if got != 0.0 {
		t.Errorf("want %f, got %f", 0.0, got)
	}
}

func TestFloat64Value(t *testing.T) {
	var v float64 = 42
	var got float64

	got = Float64Value(&v)
	if got != 42.0 {
		t.Errorf("want %f, got %f", 42.0, got)
	}

	got = Float64Value(nil)
	if got != 0.0 {
		t.Errorf("want %f, got %f", 0.0, got)
	}
}

func TestByteValue(t *testing.T) {
	var v byte = 42
	var got byte

	got = ByteValue(&v)
	if got != 42 {
		t.Errorf("want %d, got %d", 42, got)
	}

	got = ByteValue(nil)
	if got != 0 {
		t.Errorf("want %d, got %d", 0, got)
	}
}

func TestRuneValue(t *testing.T) {
	var v rune = 42
	var got rune

	got = RuneValue(&v)
	if got != 42 {
		t.Errorf("want %d, got %d", 42, got)
	}

	got = RuneValue(nil)
	if got != 0 {
		t.Errorf("want %d, got %d", 0, got)
	}
}

func TestBoolValue(t *testing.T) {
	var v bool = true
	var got bool

	got = BoolValue(&v)
	if got != true {
		t.Errorf("want %t, got %t", true, got)
	}

	got = BoolValue(nil)
	if got != false {
		t.Errorf("want %t, got %t", false, got)
	}
}

func TestTimeValue(t *testing.T) {
	v := time.Date(2020, time.January, 17, 2, 19, 0, 0, time.UTC)
	var got time.Time

	got = TimeValue(&v)
	if !got.Equal(v) {
		t.Errorf("want %s, got %s", v, got)
	}

	got = TimeValue(nil)
	if !got.IsZero() {
		t.Errorf("want %s, got %s", time.Time{}, got)
	}
}

func TestComplex64Value(t *testing.T) {
	var v complex64 = 4 + 2i
	var got complex64

	got = Complex64Value(&v)
	if got != 4+2i {
		t.Errorf("want %f, got %f", 4+2i, got)
	}

	got = Complex64Value(nil)
	if got != 0 {
		t.Errorf("want %f, got %f", 4+2i, got)
	}
}

func TestComplex128Value(t *testing.T) {
	var v complex128 = 4 + 2i
	var got complex128

	got = Complex128Value(&v)
	if got != 4+2i {
		t.Errorf("want %f, got %f", 4+2i, got)
	}

	got = Complex128Value(nil)
	if got != 0 {
		t.Errorf("want %f, got %f", 4+2i, got)
	}
}
