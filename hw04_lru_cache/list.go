package hw04_lru_cache //nolint:golint,stylecheck

type List interface {
	Len() int
	Front() *listItem
	Back() *listItem
	PushFront(v interface{}) *listItem
	PushBack(v interface{}) *listItem
	Remove(i *listItem)
	MoveToFront(i *listItem)
}

type listItem struct {
	Value interface{}
	Next  *listItem
	Prev  *listItem
}

type list struct {
	tail *listItem
	size int
}

func NewList() List {
	return &list{}
}

func (l *list) Len() int {
	return l.size
}

func (l *list) Front() *listItem {
	if l.tail == nil {
		return l.tail
	}

	front := l.tail
	for front.Prev != nil {
		front = front.Prev
	}
	return front
}

func (l *list) Back() *listItem {
	return l.tail
}

func (l *list) PushFront(v interface{}) *listItem {
	defer func() {
		l.size++
	}()

	item := &listItem{v, nil, nil}
	if l.tail == nil {
		l.tail = item
		return item
	}

	front := l.tail
	for front.Prev != nil {
		front = front.Prev
	}
	front.Prev, item.Next = item, front

	return item
}

func (l *list) PushBack(v interface{}) *listItem {
	defer func() {
		l.size++
	}()

	item := &listItem{v, nil, l.tail}
	if l.tail == nil {
		l.tail = item
		return l.tail
	}

	l.tail.Next, item.Prev = item, l.tail
	l.tail = item

	return l.tail
}

func (l *list) Remove(i *listItem) {
	defer func() {
		l.size--
	}()

	if i.Prev == nil && i.Next == nil {
		return
	}

	if i.Next == nil {
		i.Prev.Next = nil
		l.tail = i.Prev
		return
	}

	if i.Prev == nil {
		i.Next.Prev = nil
		l.tail = i.Next
		return
	}

	i.Prev.Next, i.Next.Prev = i.Next, i.Prev
}

func (l *list) MoveToFront(i *listItem) {
	if i.Prev == nil {
		return
	}

	v := i.Prev
	for v.Prev != nil {
		v = v.Prev
	}

	l.Remove(i)

	i.Next, v.Prev = v, i
	i.Prev = nil
	l.size++
}
