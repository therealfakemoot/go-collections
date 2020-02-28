package main

import (
	"sync"
	"time"
)

// NewMRUString creates a MRUString and prepares it for use.
// Note: The MRUString will NOT start cleanup until you run Start().
func NewMRUString(t time.Duration) *MRUString {
	m := MRUString{
		timeout: t,
		cache:   make(map[string]string),
		inserts: make(map[string]time.Time),
	}

	return &m
}

// MRUString is a map that evicts items based on their last time of use.
type MRUString struct {
	sync.Mutex
	timeout time.Duration
	cache   map[string]string    // this is the payload data
	inserts map[string]time.Time // maps each key to its insertion time
}

// Get returns the value associated with k
func (m *MRUString) Get(k string) (string, bool) {
	m.Lock()
	defer m.Unlock()
	v, ok := m.cache[k]
	return v, ok
}

// Set assigns the value to k and stores the creation time of this entry.
func (m *MRUString) Set(k string, v string) {
	m.SetWithTimestamp(k, v, time.Now())
}

// SetWithTimestamp assigns the value to k and stores the creation time of this entry.
func (m *MRUString) SetWithTimestamp(k string, v string, t time.Time) {
	m.Lock()
	defer m.Unlock()
	m.cache[k] = v
	m.inserts[k] = t
}

// Start begins the goroutine that cleans up entries from the LMRU.
func (m *MRUString) Start(c chan time.Time) {
	for t := range c {
		m.Lock()
		for k := range m.cache {
			if t.Sub(m.inserts[k]) > m.timeout {
				delete(m.cache, k)
				delete(m.inserts, k)
			}

		}
		m.Unlock()
	}
}
