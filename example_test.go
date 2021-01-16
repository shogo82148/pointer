package pointer_test

import (
	"fmt"

	"github.com/shogo82148/pointer"
)

func ExampleInt() {
	v := pointer.Int(42)
	fmt.Printf("%d\n", *v)

	//Output:
	// 42
}

func ExampleIntOrNil() {
	v := pointer.IntOrNil(42)
	fmt.Printf("%d\n", *v)

	v = pointer.IntOrNil(0)
	fmt.Printf("%t\n", v == nil)

	//Output:
	// 42
	// true
}

func ExampleIntValue() {
	v := pointer.Int(42)
	fmt.Printf("%d\n", pointer.IntValue(v))
	fmt.Printf("%d\n", pointer.IntValue(nil))

	//Output:
	// 42
	// 0
}
