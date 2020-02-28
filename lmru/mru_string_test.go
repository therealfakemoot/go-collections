package main

import (
	"testing"
	"time"
)

func dummyMRU() *MRUString {
	m := NewMRUString(time.Second * 5)
	m.SetWithTimestamp("demo1", "testing1", time.Unix(0, 0))
	m.SetWithTimestamp("demo2", "testing2", time.Unix(1, 0))
	m.SetWithTimestamp("demo3", "testing3", time.Unix(2, 0))
	m.SetWithTimestamp("demo4", "testing4", time.Unix(3, 0))
	m.SetWithTimestamp("demo5", "testing5", time.Unix(4, 0))
	m.SetWithTimestamp("demo6", "testing6", time.Unix(5, 0))
	return m
}

func TestMRU(t *testing.T) {
	t.Run("short test", func(t *testing.T) {
		m := dummyMRU()
		c := make(chan time.Time)
		go m.Start(c)
		c <- time.Unix(3, 0)

		_, ok := m.Get("demo1")
		if !ok {
			t.Logf("key demo1 evicted prematurely")
			t.Fail()
		}

		c <- time.Unix(5, 0)
		// c <- time.Unix(6, 0)
		_, ok = m.Get("demo1")
		if ok {
			t.Logf("key demo1 not evicted in time")
			t.Fail()
		}
	})

}
