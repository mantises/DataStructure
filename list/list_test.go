package list

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	l := New()
	for i := 0; i < 10; i++ {
		l.InsertBack(i)
	}
	s := 100
	l.InsertFront(s)
	for e := l.root.next; e != &l.root; {
		fmt.Println(e.Value)
		e = e.next
	}
}
