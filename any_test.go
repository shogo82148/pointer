//go:build go1.18
// +build go1.18

package pointer

import (
	"testing"
	"time"
)

func TestAnyPointer(t *testing.T) {
	if v := Pointer(42); *v != 42 {
		t.Errorf("want %d, got %d", 42, *v)
	}

	if v := Pointer("foo"); *v != "foo" {
		t.Errorf("want %q, got %q", "foo", *v)
	}

	if v := Pointer(42.0); *v != 42.0 {
		t.Errorf("want %f, got %f", 42.0, *v)
	}

	if v := Pointer(true); *v != true {
		t.Errorf("want %t, got %t", true, *v)
	}

	if v := Pointer(time.Date(2020, time.January, 17, 2, 19, 0, 0, time.UTC)); !v.Equal(time.Date(2020, time.January, 17, 2, 19, 0, 0, time.UTC)) {
		t.Errorf("want %s, got %s", time.Date(2020, time.January, 17, 2, 19, 0, 0, time.UTC), v)
	}
}
func TestAnyPointerOrNil(t *testing.T) {
	if v := PointerOrNil(42); *v != 42 {
		t.Errorf("want %d, got %d", 42, *v)
	}

	if v := PointerOrNil("foo"); *v != "foo" {
		t.Errorf("want %q, got %q", "foo", *v)
	}

	if v := PointerOrNil(42.0); *v != 42.0 {
		t.Errorf("want %f, got %f", 42.0, *v)
	}

	if v := PointerOrNil(true); *v != true {
		t.Errorf("want %t, got %t", true, *v)
	}

	if v := PointerOrNil(time.Date(2020, time.January, 17, 2, 19, 0, 0, time.UTC)); !v.Equal(time.Date(2020, time.January, 17, 2, 19, 0, 0, time.UTC)) {
		t.Errorf("want %s, got %s", time.Date(2020, time.January, 17, 2, 19, 0, 0, time.UTC), v)
	}

	if v := PointerOrNil(0); v != nil {
		t.Errorf("want nil, got %p", v)
	}

	if v := PointerOrNil(""); v != nil {
		t.Errorf("want nil, got %p", v)
	}

	if v := PointerOrNil(0.0); v != nil {
		t.Errorf("want nil, got %p", v)
	}

	if v := PointerOrNil(false); v != nil {
		t.Errorf("want nil, got %p", v)
	}

	if v := PointerOrNil(time.Time{}); v != nil {
		t.Errorf("want nil, got %p", v)
	}
}

func TestValueWithDefault(t *testing.T) {
	var v int = 42
	var got int

	got = ValueWithDefault(&v, 0)
	if got != 42 {
		t.Errorf("want %d, got %d", 42, got)
	}

	got = ValueWithDefault(nil, 42)
	if got != 42 {
		t.Errorf("want %d, got %d", 42, got)
	}
}

func TestValue(t *testing.T) {
	var v int = 42
	var got int

	got = Value(&v)
	if got != 42 {
		t.Errorf("want %d, got %d", 42, got)
	}

	got = Value((*int)(nil))
	if got != 0 {
		t.Errorf("want %d, got %d", 0, got)
	}
}
