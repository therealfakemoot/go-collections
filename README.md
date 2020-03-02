[![GoDoc](https://godoc.org/github.com/therealfakemoot/go-collections?status.svg)](https://godoc.org/github.com/therealfakemoot/go-collections)
[![Go Report Card](https://goreportcard.com/badge/github.com/therealfakemoot/go-collections)](https://goreportcard.com/report/github.com/therealfakemoot/go-collections)

Collections is a general purpose package. It's a rough port of Python's [collections.](https://docs.python.org/3.7/library/collections.html).

# Counter

## Generate
Generating a Counter for a type of your choosing is straightforward. From the project root, you man run the following command:

`go run gen/main.go -type typeName > filename.go`

Note: Use caution as not all types may be appropriate for use as keys.

## Examples
This example assumes you have generated your Counter type and have placed it in the counterPackage directory.

```go
import (
    "example.com/user/repo/counterPackage"
    "fmt"
)

func main() {
    c := counterPackage.New()
    c.Add("foo")
    c.Add("foo")
    c.Add("bar")

    fmt.Println(c.Get("foo")) // 2
    fmt.Println(c.Get("bar")) // 1
}
```


# LMRU
The {L,M}RU is a cache structure that evicts members based on how long they have been in cache. A Least Recently Used cache evicts the oldest items and the Most Recently Used cache evicts members newer than the provided duration.

## Generate
Help output for lmru generation:

```
  -lru
        least recently used?
  -pacakge string
        package name (default "main")
  -type string
        value type to store in cache (default "bool")
```

Generating an LMRU for a type of your choosing is straightforward. From the project root, you man run the following command. The `-l` flag is optional; omitting it indicates an MRU, including it indicates LRU.

`go run lmru/*.go -value valueType -package mycache > filename.go`

Note: Currently, the structure only supports string keys. Changing this would be relatively simple, but I'm just lazy enough that the ROI isn't there for me. Feel free to PR if you want this feature.

## Usage

```
import (
    "example.com/user/repo/mycache"
    "fmt"
)

func main() {
    c := mycache.New(time.Second * 5) // how long entries are allowed to stay in the cache without use before eviction
    e := make(chan mycache.Eviction) // The eviction channel emits an event when an entry is evicted with its name, value, and how long it was in the cache
    go c.Start(time.Ticker(time.Second * 30).C, e) // we pass all the channels in to the eviction goroutine

    c.Set("192.168.13.15", someValue)

    // at this point, you have a cache that makes an eviction pass every 30 seconds, and evicts any key that hasn't been used in 5 seconds. the structure accepts any time value, so you can have timeouts of days and eviction loops every 2 seconds, or however you want to mix it up
}
```
