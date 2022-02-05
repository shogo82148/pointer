//go:build go1.18
// +build go1.18

package pointer_test

import (
	"fmt"

	"github.com/shogo82148/pointer"
)

func ExamplePtr() {
	v := pointer.Ptr(42)
	fmt.Printf("%d\n", *v)

	//Output:
	// 42
}

func ExamplePtrOrNil() {
	v := pointer.PtrOrNil(42)
	fmt.Printf("%d\n", *v)

	v = pointer.PtrOrNil(0)
	fmt.Printf("%t\n", v == nil)

	//Output:
	// 42
	// true
}

func ExampleValue() {
	v := pointer.Ptr(42)
	fmt.Printf("%d\n", pointer.Value(v))

	v = nil
	fmt.Printf("%d\n", pointer.Value(v))

	//Output:
	// 42
	// 0
}

func ExampleEqual() {
	// a and b have the same look.
	a := pointer.Ptr(42)
	b := pointer.Ptr(42)

	// But it reports false because a and b have differrent addresses.
	fmt.Printf("a == b: %t\n", a == b)

	// Equal returns true.
	fmt.Printf("pointer.Equal(a, b): %t\n", pointer.Equal(a, b))

	//Output:
	// a == b: false
	// pointer.Equal(a, b): true
}
