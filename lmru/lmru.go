package main

var lmruTemplate = `package {{ .Package }}

import (
	"sync"
	"time"
)

// Eviction contains information about keys evicted from the cache
type Eviction struct {
	Key string
	Value {{ .Type }}
	Lifetime time.Duration
}

// New{{ .Name}} creates a {{ .Name }} and prepares it for use.
// Note: The {{ .Name }} will NOT start cleanup until you run Start().
func New{{ .Name}}(t time.Duration) *{{ .Name }} {
	m := {{ .Name }}{
		timeout: t,
		cache: make(map[string]{{ .Type }}),
		inserts: make(map[string]time.Time),
	}

	return &m
}

{{ if .L }}// {{ .Name }} is a map that evicts items based on their last time of use.{{ else }}// {{ .Name }} is a map that evicts items based on their last time of use.{{ end }}
type {{ .Name }} struct {
	sync.Mutex
	timeout time.Duration
	cache       map[string]{{ .Type }}      // this is the payload data
	inserts       map[string]time.Time // maps each key to its insertion time
}

// Get returns the value associated with k
func (m *{{ .Name }}) Get(k string) (string, bool) {
       m.Lock()
       defer m.Unlock()
       v, ok := m.cache[k]
       if ok {
	       m.inserts[k] = time.Now()
       }
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
	m.cache[k] = v
	m.inserts[k] = t
}

// Start begins the goroutine that cleans up entries from the LMRU.
func (m *{{ .Name }}) Start(c chan time.Time, e chan Eviction) {
		for t := range c {
			m.Lock()
			for k := range m.cache {
				lifetime := t.Sub(m.inserts[k])
				if lifetime{{if .L}} < {{ else }} > {{ end }}m.timeout {
				var eviction Eviction
				if e != nil {
					eviction.Key = k
					eviction.Value = m.cache[k]
					eviction.Lifetime = lifetime
				}
					delete(m.cache, k)
					delete(m.inserts, k)
					e <- eviction
				}

			}
			m.Unlock()
		}
}
`
