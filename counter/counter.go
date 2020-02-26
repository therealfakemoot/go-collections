package main

var counterTemplate = `
package counter

import (
	"sort"
)

// New initializes a new Counter and its backing map.
func New{{ .Name }}Counter() {{ .Name }}Counter {
	var c {{ .Name }}Counter
	c.members = make(map[{{ .Type }}]int)

	return c
}

// Counter uses a map[<type>]int to count how many instances of a value have been observed.
// Inspired by Python's collections.Counter type.
type {{ .Name }}Counter struct {
	members map[{{ .Type }}]int
}

// Add accepts a value and increments its observation count by 1.
func (c *{{ .Name }}Counter) Add(m {{ .Type }}) {
	c.members[m]++
}

// Reset resets all observations, replacing the backing map with a new one.
func (c *{{ .Name }}Counter) Reset() {
	c.members = make(map[{{ .Type }}]int)
}

func (c *{{ .Name }}Counter) sorted() []{{ .Type }} {
	var s []{{ .Type }}
	for k := range c.members {
		s = append(s, k)
	}

	sort.Slice(s, func(i, j int) bool {
		return c.Get(s[i]) > c.Get(s[j])
	})

	return s
}

// Get returns the number of times a value has been observed.
func (c *{{ .Name }}Counter) Get(m {{ .Type }}) int {
	n, ok := c.members[m]
	if !ok {
		return 0
	}

	return n
}

// Set sets the observation count for a given value to the provided int.
func (c *{{ .Name }}Counter) Set(m {{ .Type }}, n int) {
	c.members[m] = n
}

// Most returns the n most frequently observed values, as a Counter.
// This is to facilitate further filtering of the provided data.
func (c *{{ .Name }}Counter) Most(n int) {{ .Name }}Counter {
	sorted := c.sorted()
	r := New{{ .Name }}Counter()

	for i := 0; i < len(sorted)-1; i++ {
		r.Set(sorted[i], c.Get(sorted[i]))
	}

	return r
}

// Least returns the n least frequently observed values, as a Counter.
// This is to facilitate further filtering of the provided data.
func (c *{{ .Name }}Counter) Least(n int) {{ .Name }}Counter {
	sorted := c.sorted()
	r := New{{ .Name }}Counter()

	for i := len(sorted) - 1; i >= 0; i-- {
		r.Set(sorted[i], c.Get(sorted[i]))
	}

	return r
}
`
