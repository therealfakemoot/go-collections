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
	t.Run("single eviction", func(t *testing.T) {
		m := dummyMRU()
		c := make(chan time.Time)
		e := make(chan Eviction)
		go m.Start(c, e)

		c <- time.Unix(5, 0)
		select {
		case ev, _ := <-e:
			t.Logf("received unexpected eviction: %+v", ev)
			t.Fail()
		default:
			break // break the select and continue to the rest of the test body???
		}

		t.Logf("T: +6s")
		c <- time.Unix(6, 0)
		select {
		case member := <-e:
			if member.Key != "demo1" {
				t.Logf("expected demo1 to be evicted, received %s", member.Key)
				t.Fail()
			}
			return
		case <-time.After(time.Second * 10):
			t.Logf("timeout exceeded")
			t.Fail()
		}
	})

	t.Run("butts", func(t *testing.T) {
		m := dummyMRU()

	})
}
