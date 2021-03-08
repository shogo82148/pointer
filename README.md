# pointer

[![Go Reference](https://pkg.go.dev/badge/github.com/shogo82148/pointer.svg)](https://pkg.go.dev/github.com/shogo82148/pointer)
![test](https://github.com/shogo82148/pointer/workflows/test/badge.svg)

Pointer Utility Functions for Golang.
The Go's specification says that the operand of the address operation `&x` must be **addressable**
(ref. https://golang.org/ref/spec#Address_operators ).
It means that we can't get the addresses of [constants](https://golang.org/ref/spec#Constants),
literals([Integer literals](https://golang.org/ref/spec#Integer_literals), [Floating-point literals](https://golang.org/ref/spec#Floating-point_literals),
[String literals](https://golang.org/ref/spec#String_literals), etc.), and the return values of a function or method.
The `pointer` package makes them addressable, and returns their pointers.

## SYNOPSIS

```go
import "github.com/shogo82148/pointer"

type Stuff struct {
    Name    *string
    Comment *string
    Value   *int64
    Time    *time.Time
}

func main() {
    const defaultName = "some name"

    // convert to pointers
    stuff := &Stuff{
        // it is same as s := defaultName; &s
        // can't be &defaultName
        Name: pointer.String(defaultName),

        // can't be &"not yet"
        Comment: pointer.String("not yet"),

        // can't be &42 or &int64(42)
        Value: pointer.Int64(42),

        // can't be &time.Date(2014, 6, 25, 12, 24, 40, 0, time.UTC)
        Time: pointer.Time(time.Date(2014, 6, 25, 12, 24, 40, 0, time.UTC)),
    }

    // Int64OrNil(v) is like Int64, but it returns nil if v is the zero value.
    stuff.Value = pointer.Int64OrNil(0)

    // StringValue(p) is like *p but it returns "" if p is nil
    // It never panic.
    fmt.Printf("%s\n", pointer.StringValue(stuff.Name))

    // StringValueWithDefault returns the default value if p is nil
    fmt.Printf("%s\n", pointer.StringValueWithDefault(stuff.Name, "John"))
}
```

## RELATED WORKS

- https://github.com/AlekSi/pointer
- https://github.com/agrea/ptr
