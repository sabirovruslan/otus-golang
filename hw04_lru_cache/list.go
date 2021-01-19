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
	front *listItem
	back  *listItem
}

func (l list) Len() int {
	return 0
}

func (l list) Front() *listItem {
	return l.front
}

func (l list) Back() *listItem {
	return l.back
}

func (l list) PushFront(v interface{}) *listItem {

}

func (l list) PushBack(v interface{}) *listItem {

}

func (l list) Remove(i *listItem) {

}

func (l list) MoveToFront(i *listItem) {

}

func NewList() List {
	return &list{}
}
