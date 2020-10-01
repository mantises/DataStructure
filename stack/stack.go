package stack

import "container/list"

type Stack struct {
	l *list.List
}

/** initialize your data structure here. */
func Constructor() Stack {
	return Stack{l: list.New()}
}

func (s *Stack) Push(x interface{}) {
	s.l.PushBack(x)
}

func (s *Stack) Pop() {
	s.l.Remove(s.l.Back())
}

func (s*Stack) Empty() bool {
	return s.l.Len() == 0
}

func (s *Stack) Top() interface{} {
	back := s.l.Back()
	if back != nil {
		return back.Value
	}
	return nil
}

/**
 * Your Stack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 */
