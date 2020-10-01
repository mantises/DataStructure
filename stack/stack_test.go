package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T)  {
	st := NewStack()
	for i := 0 ; i < 10; i++ {
		st.Push(i)
	}
	fmt.Println(st.Top())
	st.Pop()
	fmt.Println(st.Top())
	st.Push(3)
	for !st.Empty() {
		fmt.Println(st.Top())
		st.Pop()
	}
}
