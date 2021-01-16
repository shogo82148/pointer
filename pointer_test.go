package pointer

import "testing"

func TestPointer(t *testing.T) {
	if v := Int(42); *v != 42 {
		t.Errorf("want %d, got %d", 42, *v)
	}
}
