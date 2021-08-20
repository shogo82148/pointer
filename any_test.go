//go:build go1.18
// +build go1.18

package pointer

import (
	"testing"
	"time"
)

func TestAny(t *testing.T) {
	if v := Any(42); *v != 42 {
		t.Errorf("want %d, got %d", 42, *v)
	}

	if v := Any("foo"); *v != "foo" {
		t.Errorf("want %q, got %q", "foo", *v)
	}

	if v := Any(42.0); *v != 42.0 {
		t.Errorf("want %f, got %f", 42.0, *v)
	}

	if v := Bool(true); *v != true {
		t.Errorf("want %t, got %t", true, *v)
	}

	if v := Time(time.Date(2020, time.January, 17, 2, 19, 0, 0, time.UTC)); !v.Equal(time.Date(2020, time.January, 17, 2, 19, 0, 0, time.UTC)) {
		t.Errorf("want %s, got %s", time.Date(2020, time.January, 17, 2, 19, 0, 0, time.UTC), v)
	}
}
func TestAnyOrNil(t *testing.T) {
	if v := AnyOrNil(42); *v != 42 {
		t.Errorf("want %d, got %d", 42, *v)
	}

	if v := AnyOrNil("foo"); *v != "foo" {
		t.Errorf("want %q, got %q", "foo", *v)
	}

	if v := Any32OrNil(42.0); *v != 42.0 {
		t.Errorf("want %f, got %f", 42.0, *v)
	}

	if v := AnyOrNil(true); *v != true {
		t.Errorf("want %t, got %t", true, *v)
	}

	if v := AnyOrNil(time.Date(2020, time.January, 17, 2, 19, 0, 0, time.UTC)); !v.Equal(time.Date(2020, time.January, 17, 2, 19, 0, 0, time.UTC)) {
		t.Errorf("want %s, got %s", time.Date(2020, time.January, 17, 2, 19, 0, 0, time.UTC), v)
	}

	if v := AnyOrNil(0); v != nil {
		t.Errorf("want nil, got %p", v)
	}

	if v := AnyOrNil(""); v != nil {
		t.Errorf("want nil, got %p", v)
	}

	if v := AnyOrNil(0.0); v != nil {
		t.Errorf("want nil, got %p", v)
	}

	if v := AnyOrNil(false); v != nil {
		t.Errorf("want nil, got %p", v)
	}

	if v := TimeOrNil(time.Time{}); v != nil {
		t.Errorf("want nil, got %p", v)
	}
}

func TestAnyValueWithDefault(t *testing.T) {
	var v int = 42
	var got int

	got = AnyValueWithDefault(&v, 0)
	if got != 42 {
		t.Errorf("want %d, got %d", 42, got)
	}

	got = AnyValueWithDefault(nil, 42)
	if got != 42 {
		t.Errorf("want %d, got %d", 42, got)
	}
}

func TestAnyValue(t *testing.T) {
	var v int = 42
	var got int

	got = IntValue(&v)
	if got != 42 {
		t.Errorf("want %d, got %d", 42, got)
	}

	got = AnyValue((*int)(nil))
	if got != 0 {
		t.Errorf("want %d, got %d", 0, got)
	}
}
