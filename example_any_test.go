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
