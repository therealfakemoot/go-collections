package main

var lmruTemplate = `package {{ .Package }}

import (
	"sync"
	"time"
)

// New{{ .Name}} creates a {{ .Name }} and prepares it for use.
// Note: The {{ .Name }} will NOT start cleanup until you run Start().
func New{{ .Name}}(t time.Duration) *{{ .Name }} {
	m := {{ .Name }}{
		timeout: t,
		m: make(map[string]{{ .Type }}),
		c: make(map[string]time.Time),
	}

	return &m
}

{{ if .L }}// {{ .Name }} is a map that evicts items based on their last time of use.{{ else }}// {{ .Name }} is a map that evicts items based on their last time of use.{{ end }}
type {{ .Name }} struct {
	sync.Mutex
	timeout time.Duration
	m       map[string]{{ .Type }}      // this is the payload data
	c       map[string]time.Time // maps each key to its insertion time
}

// Get returns the value associated with k
func (m *{{ .Name }}) Get(k string) (string, bool) {
       m.Lock()
       defer m.Unlock()
       v, ok := m.m[k]
       return v, ok
}

// Set assigns the value to k and stores the creation time of this entry.
func (m *{{ .Name }}) Set(k string, v {{ .Type }}) {
	m.SetWithTimestamp(k,v, time.Now())
}

// SetWithTimestamp assigns the value to k and stores the creation time of this entry.
func (m *{{ .Name }}) SetWithTimestamp(k string, v {{ .Type }}, t time.Time) {
	m.Lock()
	defer m.Unlock()
	m.m[k] = v
	m.c[k] = t
}

// Start begins the goroutine that cleans up entries from the LMRU.
func (m *{{ .Name }}) Start(c chan time.Time) {
	go func() {
		for t := range c {
			m.Lock()
			for k := range m.m {
				if t.Sub(m.c[k]){{if .L}} < {{ else }} > {{ end }}m.timeout {
					delete(m.m, k)
				}

			}
			m.Unlock()
		}
	}()
}
`
