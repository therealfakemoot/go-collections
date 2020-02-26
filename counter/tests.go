package main

var counterTestTemplate = `
import (
	"testing"
)

func prepareTestCounter() Counter {

	c := New()
	c.Add("foo")
	c.Add("foo")
	c.Add("bar")
	return c
}

func Test_CounterBasic(t *testing.T) {
	c := prepareTestCounter()

	t.Run("get observations", func(t *testing.T) {
		n := c.Get("foo")
		if n != 2 {
			t.Logf("expected 2 'foo', got %d", n)
			t.Fail()
		}

		n = c.Get("bar")
		if n != 1 {
			t.Logf("expected 1 'bar', got %d", n)
			t.Fail()
		}
	})
}

func Test_CounterFilter(t *testing.T) {
	c := prepareTestCounter()
	t.Run("most", func(t *testing.T) {
		m := c.Most(1)
		if m.Get("foo") != 2 {
			t.Logf("%#v", m)
			t.Fail()
		}
	})

	t.Run("least", func(t *testing.T) {
		l := c.Least(1)
		if l.Get("bar") != 1 {
			t.Logf("%#v", l)
			t.Fail()
		}

	})
}
`
