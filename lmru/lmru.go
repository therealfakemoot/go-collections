package main

var lmruTemplate = `
import (
	"fmt"
	"sync"
	"time"
)

{{ if .L }}// LRU is a map that evicts items based on their last time of use.{{ else }}// MRU is a map that evicts items based on their last time of use.{{ end }}
type {{if .L}}L{{else}}M{{end}}{{.Name}} struct {
	sync.Mutex
	timeout time.Duration
	m       map[string]{{ .Type }}      // this is the payload data
	c       map[string]time.Time // maps each key to its insertion time
}

// Get returns the value associated with k
func (m *MRU{{.Name}}) Get(k string) {{ .Type }} {
	m.Lock()
	defer m.Unlock()
	return m.m[k]
}

// Set assigns the value to k and stores the creation time of this entry.
func (m *MRU{{.Name}}) Set(k string, v {{ .Type }}) {
	m.Lock()
	defer m.Unlock()
	m.m[k] = v
	m.c[k] = time.Now()
}

// Start begins the goroutine that cleans up entries from the LMRU.
func (m *MRU{{.Name}}) Start(d time.Duration) {
	go func() {
		for range time.NewTicker(d).C {
			m.Lock()
			defer m.Unlock()
			for k := range m.m {
				if time.Now().Sub(m.c[k]){{if .L}} > {{ else }} < {{ end }} m.timeout {
					delete(m.m, k)
				}

			}
		}
	}()
}
`
