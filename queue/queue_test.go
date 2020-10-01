package queue

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T)  {
	q := NewQueue()
	fmt.Println(q.Front())
	fmt.Println(q.Back())
	for i := 0; i < 10; i++ {
		q.PutBack(i)
		q.PutFront(-i)
	}
	fmt.Println("size:", q.Size())
	for !q.Empty() {
		fmt.Println(q.PollBack())
	}
}
