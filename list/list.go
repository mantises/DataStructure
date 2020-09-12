package list

type Element struct {
	next, prev *Element
	list       *List
	Value      interface{}
}

func (e *Element) Next() *Element {
	if p := e.next; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}
func (e *Element) Prev() *Element {
	if p := e.prev; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

type List struct {
	root Element
	len  int
}

func (l *List) Init() *List {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}

func New() *List { return new(List).Init() }

func (l *List) Len() int {
	return l.len
}

func (l *List) Front() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

func (l *List) Back() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}

func (l *List) LazyInit() {
	if l.root.next == nil {
		l.Init()
	}
}

func (l *List) insertAt(e, at *Element) *Element {
	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e

	e.list = l
	l.len++
	return e
}

func (l *List) InsertFront(v interface{}) *Element {
	l.LazyInit()
	return l.insertAt(&Element{Value: v}, &l.root)
}

func (l *List) InsertBack(v interface{}) *Element {
	l.LazyInit()
	return l.insertAt(&Element{Value: v}, l.root.prev)
}

func (l *List) Remove(e *Element) interface{} {
	if e.list == l {
		e.next.prev = e.prev
		e.prev.next = e.next
		e.next = nil
		e.prev = nil
		e.list = nil

		l.len--
	}
	return e.Value
}
