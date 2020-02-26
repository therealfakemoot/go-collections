[![GoDoc](https://godoc.org/github.com/therealfakemoot/go-collections?status.svg)](https://godoc.org/github.com/therealfakemoot/go-collections)
[![Go Report Card](https://goreportcard.com/badge/github.com/therealfakemoot/go-collections)](https://goreportcard.com/report/github.com/therealfakemoot/go-collections)

Collections is a general purpose package. It's a rough port of Python's [collections.](https://docs.python.org/3.7/library/collections.html).

# Counter

## Examples
```go
import (
    "github.com/therealfakemoot/go-collections"
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

## Generate
Generating a Counter for a type of your choosing is straightforward. From the project root, you man run the following command:

`go run gen/main.go -type typeName > filename.go`

Note: Use caution as not all types may be appropriate for use as keys.

# LMRU
The {L,M}RU is a cache structure that evicts members based on how long they have been in cache. A Least Recently Used cache evicts the oldest items and the Most Recently Used cache evicts members newer than the provided duration.
