package heap

import (
	"fmt"
	"testing"
)

func TestNewHeap(t *testing.T) {
	f := func(x interface{}, y interface{}) bool {
		if x.(int) > y.(int) {
			return true
		}
		return false
	}
	h := NewPriorityQueue(f)
	h.Push(10)
	h.Push(99)
	h.Push(6)
	h.Push(27)
	h.Push(88)
	fmt.Println(h.val)
	fmt.Println(h.Pop())
	fmt.Println(h.val)
	h.Push(1000)
	h.Push(90)
	fmt.Println(h.val)
	fmt.Println(h.Pop())
	fmt.Println(h.val)
}
