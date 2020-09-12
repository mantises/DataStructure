package list

import (
	"testing"
)

func TestList(t *testing.T) {
	l := New()
	for i := 0; i < 100; i++ {
		l.InsertBack(i)
	}
	for e := l.root.next; e != &l.root; {
		t.Log(e.Value)
		e = e.next
	}
}
