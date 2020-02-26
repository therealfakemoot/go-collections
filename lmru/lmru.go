package main

var lmruTemplate = `package {{ .Package }}

import (
	"fmt"
	"sync"
	"time"
)

{{ if .L }}// {{ .Name }} is a map that evicts items based on their last time of use.{{ else }}// {{ .Name }} is a map that evicts items based on their last time of use.{{ end }}
type {{ .Name }} struct {
	sync.Mutex
	timeout time.Duration
	m       map[string]{{ .Type }}      // this is the payload data
	c       map[string]time.Time // maps each key to its insertion time
}

// Get returns the value associated with k
func (m *{{ .Name }}) Get(k string) {{ .Type }} {
	m.Lock()
	defer m.Unlock()
	return m.m[k]
}

// Set assigns the value to k and stores the creation time of this entry.
func (m *{{ .Name }}) Set(k string, v {{ .Type }}) {
	m.Lock()
	defer m.Unlock()
	m.m[k] = v
	m.c[k] = time.Now()
}

// Start begins the goroutine that cleans up entries from the LMRU.
func (m *{{ .Name }}) Start(c chan time.Time) {
	go func() {
		for t := range c {
			m.Lock()
			for k := range m.m {
				if t.Sub(m.c[k]){{if .L}} > {{ else }} < {{ end }}m.timeout {
					delete(m.m, k)
				}

			}
			m.Unlock()
		}
	}()
}
`
