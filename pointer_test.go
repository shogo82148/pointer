package pointer

import (
	"testing"
	"time"
)

func TestPointer(t *testing.T) {
	if v := Int(42); *v != 42 {
		t.Errorf("want %d, got %d", 42, *v)
	}

	if v := Int8(42); *v != 42 {
		t.Errorf("want %d, got %d", 42, *v)
	}

	if v := Int16(42); *v != 42 {
		t.Errorf("want %d, got %d", 42, *v)
	}

	if v := Int32(42); *v != 42 {
		t.Errorf("want %d, got %d", 42, *v)
	}

	if v := Int64(42); *v != 42 {
		t.Errorf("want %d, got %d", 42, *v)
	}

	if v := Uint(42); *v != 42 {
		t.Errorf("want %d, got %d", 42, *v)
	}

	if v := Uint8(42); *v != 42 {
		t.Errorf("want %d, got %d", 42, *v)
	}

	if v := Uint16(42); *v != 42 {
		t.Errorf("want %d, got %d", 42, *v)
	}

	if v := Uint32(42); *v != 42 {
		t.Errorf("want %d, got %d", 42, *v)
	}

	if v := Uint64(42); *v != 42 {
		t.Errorf("want %d, got %d", 42, *v)
	}

	if v := String("foo"); *v != "foo" {
		t.Errorf("want %q, got %q", "foo", *v)
	}

	if v := Float32(42); *v != 42.0 {
		t.Errorf("want %f, got %f", 42.0, *v)
	}

	if v := Float64(42); *v != 42.0 {
		t.Errorf("want %f, got %f", 42.0, *v)
	}

	if v := Bool(true); *v != true {
		t.Errorf("want %t, got %t", true, *v)
	}

	if v := Byte(42); *v != 42 {
		t.Errorf("want %d, got %d", 42, *v)
	}

	if v := Rune('a'); *v != 'a' {
		t.Errorf("want %c, got %c", 'a', *v)
	}

	if v := Time(time.Date(2020, time.January, 17, 2, 19, 0, 0, time.UTC)); !v.Equal(time.Date(2020, time.January, 17, 2, 19, 0, 0, time.UTC)) {
		t.Errorf("want %s, got %s", time.Date(2020, time.January, 17, 2, 19, 0, 0, time.UTC), v)
	}

	if v := Complex64(1 + 2i); *v != 1+2i {
		t.Errorf("want %f, got %f", 1+2i, *v)
	}

	if v := Complex128(1 + 2i); *v != 1+2i {
		t.Errorf("want %f, got %f", 1+2i, *v)
	}
}
