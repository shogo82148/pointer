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
