//go:build go1.18
// +build go1.18

package pointer

import (
	"testing"
	"time"
)

func TestAnyPtr(t *testing.T) {
	if v := Ptr(42); *v != 42 {
		t.Errorf("want %d, got %d", 42, *v)
	}

	if v := Ptr("foo"); *v != "foo" {
		t.Errorf("want %q, got %q", "foo", *v)
	}

	if v := Ptr(42.0); *v != 42.0 {
		t.Errorf("want %f, got %f", 42.0, *v)
	}

	if v := Ptr(true); *v != true {
		t.Errorf("want %t, got %t", true, *v)
	}

	if v := Ptr(time.Date(2020, time.January, 17, 2, 19, 0, 0, time.UTC)); !v.Equal(time.Date(2020, time.January, 17, 2, 19, 0, 0, time.UTC)) {
		t.Errorf("want %s, got %s", time.Date(2020, time.January, 17, 2, 19, 0, 0, time.UTC), v)
	}
}
func TestAnyPtrOrNil(t *testing.T) {
	if v := PtrOrNil(42); *v != 42 {
		t.Errorf("want %d, got %d", 42, *v)
	}

	if v := PtrOrNil("foo"); *v != "foo" {
		t.Errorf("want %q, got %q", "foo", *v)
	}

	if v := PtrOrNil(42.0); *v != 42.0 {
		t.Errorf("want %f, got %f", 42.0, *v)
	}

	if v := PtrOrNil(true); *v != true {
		t.Errorf("want %t, got %t", true, *v)
	}

	if v := PtrOrNil(time.Date(2020, time.January, 17, 2, 19, 0, 0, time.UTC)); !v.Equal(time.Date(2020, time.January, 17, 2, 19, 0, 0, time.UTC)) {
		t.Errorf("want %s, got %s", time.Date(2020, time.January, 17, 2, 19, 0, 0, time.UTC), v)
	}

	if v := PtrOrNil(0); v != nil {
		t.Errorf("want nil, got %p", v)
	}

	if v := PtrOrNil(""); v != nil {
		t.Errorf("want nil, got %p", v)
	}

	if v := PtrOrNil(0.0); v != nil {
		t.Errorf("want nil, got %p", v)
	}

	if v := PtrOrNil(false); v != nil {
		t.Errorf("want nil, got %p", v)
	}

	if v := PtrOrNil(time.Time{}); v != nil {
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

func TestEqual(t *testing.T) {
	var v int

	// a == b
	if !Equal(&v, &v) {
		t.Error("Equal(&v, &v) should be true, got false")
	}

	// a != b but *a == *b
	if !Equal(&v, Ptr(0)) {
		t.Error("should be true, got false")
	}

	a := [...]int{1, 2, 3}
	b := [...]int{1, 2, 3}
	if !Equal(&a, &b) {
		t.Error("should be true, got false")
	}

	if Equal(nil, &v) {
		t.Error("Equal(nil, &v) should be false, got true")
	}

	if Equal(&v, nil) {
		t.Error("Equal(&v, nil) should be false, got true")
	}

	if !Equal((*int)(nil), (*int)(nil)) {
		t.Error("Equal((*int)(nil), (*int)(nil)) should be true, got false")
	}
}
