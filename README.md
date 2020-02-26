[![GoDoc](https://godoc.org/github.com/therealfakemoot/go-counter?status.svg)](https://godoc.org/github.com/therealfakemoot/go-counter)
[![Go Report Card](https://goreportcard.com/badge/github.com/therealfakemoot/go-counter)](https://goreportcard.com/report/github.com/therealfakemoot/go-counter)

Counter is a general purpose package. It's a rough port of Python's [collections.Counter](https://docs.python.org/3.7/library/collections.html#collections.Counter).

# Examples
```go
import (
    "github.com/therealfakemoot/go-counter"
    "fmt"
)

func main() {
    c := New()
    c.Add("foo")
    c.Add("foo")
    c.Add("bar")

    fmt.Println(c.Get("foo")) // 2
    fmt.Println(c.Get("bar")) // 1
}
```

# Generate
Generating a Counter for a type of your choosing is straightforward. From the project root, you man run the following command:

`go run gen/main.go -type typeName > filename.go`

Note: Use caution as not all types may be appropriate for use as keys. Be mindful, this package is released without warranty or guarantee.

