package queue

import "container/list"

type Queue struct {
	l *list.List
}

func NewQueue() Queue {
	return Queue{l: list.New()}
}

func (q *Queue) PutFront(x interface{}) {
	q.l.PushFront(x)
}

func (q *Queue) PutBack(x interface{}) {
	q.l.PushBack(x)
}

func (q *Queue) Front() interface{} {
	front := q.l.Front()
	if front != nil {
		return front.Value
	}
	return nil
}

func (q *Queue) Back() interface{} {
	back := q.l.Back()
	if back != nil {
		return back.Value
	}
	return nil
}

func (q *Queue) PollFront() interface{} {
	front := q.l.Front()
	if front != nil {
		return q.l.Remove(front)
	}
	return nil
}
func (q *Queue) PollBack() interface{} {
	back := q.l.Back()
	if back != nil {
		return q.l.Remove(back)
	}
	return nil
}

func (q *Queue) Empty() bool {
	return q.l.Len() == 0
}

func (q *Queue) Size() int {
	return q.l.Len()
}


