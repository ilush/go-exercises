package hw04_lru_cache //nolint:golint,stylecheck

type List interface {
	Len() int                          // длина списка
	Front() *listItem                  // первый listItem
	Back() *listItem                   // последний listItem
	PushFront(v interface{}) *listItem // добавить значение в начало
	PushBack(v interface{}) *listItem  // добавить значение в конец
	Remove(i *listItem)                // удалить элемент
	MoveToFront(i *listItem)           // переместить элемент в начало
}

type listItem struct {
	Next  *listItem
	Prev  *listItem
	Value interface{}
	list  *list
}

type list struct {
	front *listItem
	back  *listItem
	len   int
}

func (l *list) Init() *list {
	l.front = nil
	l.back = nil
	l.len = 0
	return l
}

func NewList() List {
	new(list).Init()
	return &list{}
}

func (l *list) Len() int { return l.len }

func (l *list) Front() *listItem {
	if l.len == 0 {
		return nil
	}
	return l.front
}

func (l *list) Back() *listItem {
	if l.len == 0 {
		return nil
	}
	return l.back
}

func (l *list) PushFront(v interface{}) *listItem { //nolint:dupl
	i := &listItem{Value: v}
	f := l.front
	b := l.back
	switch {
	case f == nil: // push to an empty list
		l.front = i
		l.back = i
		l.len++
		i.list = l
	case f.Value != nil && f == b: // push to front of one-item list
		i.list = l
		l.front = i
		l.front.Prev = f
		l.front.Next = nil
		l.back = f
		l.back.Next = i
		l.back.Prev = nil
		l.len++
	default:
		i.list = l
		l.front = i
		l.front.Prev = f
		l.front.Next = nil
		f.Next = i
		l.len++
	}
	return i
}

func (l *list) PushBack(v interface{}) *listItem { //nolint:dupl
	i := &listItem{Value: v}
	f := l.front
	b := l.back
	switch {
	case f == nil: // push to an empty list
		l.front = i
		l.back = i
		l.len++
		i.list = l
	case b.Value != nil && f == b: // push to back of one-item list
		i.list = l
		l.front = f
		l.front.Prev = i
		l.front.Next = nil
		l.back = i
		l.back.Next = f
		l.back.Prev = nil
		l.len++
	default:
		i.list = l
		l.back = i
		l.back.Prev = nil
		l.back.Next = b
		b.Prev = i
		l.len++
	}

	return i
}

func (l *list) MoveToFront(i *listItem) {
	if i == l.front {
		return
	}
	if i == l.back {
		l.back = i.Next
	}

	if i.Prev != nil {
		i.Prev.Next = i.Next
	}

	i.Next.Prev = i.Prev

	n := l.front
	l.front = i
	i.Prev = n
	i.Next = nil
	n.Next = i
}

func (l *list) Remove(i *listItem) {
	if i == l.front && i == l.back { // removing last listItem from list
		l.front = nil
		l.back = nil
		l.len = 0
	}
	if i.Prev != nil { // l.Back()
		i.Prev.Next = i.Next
	} else {
		l.back = i.Next
	}
	if i.Next != nil { // l.Front()
		i.Next.Prev = i.Prev
	} else {
		l.front = i.Prev
	}

	i.Next = nil
	i.Prev = nil
	i.list = nil
	l.len--
}
