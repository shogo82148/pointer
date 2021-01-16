package pointer

import (
	"testing"
	"time"
)

func TestPointerOrNil(t *testing.T) {
	if v := IntOrNil(42); *v != 42 {
		t.Errorf("want %d, got %d", 42, *v)
	}

	if v := Int8OrNil(42); *v != 42 {
		t.Errorf("want %d, got %d", 42, *v)
	}

	if v := Int16OrNil(42); *v != 42 {
		t.Errorf("want %d, got %d", 42, *v)
	}

	if v := Int32OrNil(42); *v != 42 {
		t.Errorf("want %d, got %d", 42, *v)
	}

	if v := Int64OrNil(42); *v != 42 {
		t.Errorf("want %d, got %d", 42, *v)
	}

	if v := UintOrNil(42); *v != 42 {
		t.Errorf("want %d, got %d", 42, *v)
	}

	if v := Uint8OrNil(42); *v != 42 {
		t.Errorf("want %d, got %d", 42, *v)
	}

	if v := Uint16OrNil(42); *v != 42 {
		t.Errorf("want %d, got %d", 42, *v)
	}

	if v := Uint32OrNil(42); *v != 42 {
		t.Errorf("want %d, got %d", 42, *v)
	}

	if v := Uint64OrNil(42); *v != 42 {
		t.Errorf("want %d, got %d", 42, *v)
	}

	if v := StringOrNil("foo"); *v != "foo" {
		t.Errorf("want %q, got %q", "foo", *v)
	}

	if v := Float32OrNil(42); *v != 42.0 {
		t.Errorf("want %f, got %f", 42.0, *v)
	}

	if v := Float64OrNil(42); *v != 42.0 {
		t.Errorf("want %f, got %f", 42.0, *v)
	}

	if v := BoolOrNil(true); *v != true {
		t.Errorf("want %t, got %t", true, *v)
	}

	if v := ByteOrNil(42); *v != 42 {
		t.Errorf("want %d, got %d", 42, *v)
	}

	if v := RuneOrNil('a'); *v != 'a' {
		t.Errorf("want %c, got %c", 'a', *v)
	}

	if v := TimeOrNil(time.Date(2020, time.January, 17, 2, 19, 0, 0, time.UTC)); !v.Equal(time.Date(2020, time.January, 17, 2, 19, 0, 0, time.UTC)) {
		t.Errorf("want %s, got %s", time.Date(2020, time.January, 17, 2, 19, 0, 0, time.UTC), v)
	}

	if v := IntOrNil(0); v != nil {
		t.Errorf("want nil, got %p", v)
	}

	if v := Int8OrNil(0); v != nil {
		t.Errorf("want nil, got %p", v)
	}

	if v := Int16OrNil(0); v != nil {
		t.Errorf("want nil, got %p", v)
	}

	if v := Int32OrNil(0); v != nil {
		t.Errorf("want nil, got %p", v)
	}

	if v := Int64OrNil(0); v != nil {
		t.Errorf("want nil, got %p", v)
	}

	if v := UintOrNil(0); v != nil {
		t.Errorf("want nil, got %p", v)
	}

	if v := Uint8OrNil(0); v != nil {
		t.Errorf("want nil, got %p", v)
	}

	if v := Uint16OrNil(0); v != nil {
		t.Errorf("want nil, got %p", v)
	}

	if v := Uint32OrNil(0); v != nil {
		t.Errorf("want nil, got %p", v)
	}

	if v := Uint64OrNil(0); v != nil {
		t.Errorf("want nil, got %p", v)
	}

	if v := StringOrNil(""); v != nil {
		t.Errorf("want nil, got %p", v)
	}

	if v := Float32OrNil(0); v != nil {
		t.Errorf("want nil, got %p", v)
	}

	if v := Float64OrNil(0); v != nil {
		t.Errorf("want nil, got %p", v)
	}

	if v := BoolOrNil(false); v != nil {
		t.Errorf("want nil, got %p", v)
	}

	if v := ByteOrNil(0); v != nil {
		t.Errorf("want nil, got %p", v)
	}

	if v := RuneOrNil(0); v != nil {
		t.Errorf("want nil, got %p", v)
	}

	if v := TimeOrNil(time.Time{}); v != nil {
		t.Errorf("want nil, got %p", v)
	}
}
