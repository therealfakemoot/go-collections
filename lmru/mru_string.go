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
		m:       make(map[string]string),
		c:       make(map[string]time.Time),
	}

	return &m
}

// MRUString is a map that evicts items based on their last time of use.
type MRUString struct {
	sync.Mutex
	timeout time.Duration
	m       map[string]string    // this is the payload data
	c       map[string]time.Time // maps each key to its insertion time
}

// Get returns the value associated with k
func (m *MRUString) Get(k string) (string, bool) {
	m.Lock()
	defer m.Unlock()
	v, ok := m.m[k]
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
	m.m[k] = v
	m.c[k] = t
}

// Start begins the goroutine that cleans up entries from the LMRU.
func (m *MRUString) Start(c chan time.Time) {
	go func() {
		for t := range c {
			m.Lock()
			for k := range m.m {
				if t.Sub(m.c[k]) > m.timeout {
					delete(m.m, k)
				}

			}
			m.Unlock()
		}
	}()
}
