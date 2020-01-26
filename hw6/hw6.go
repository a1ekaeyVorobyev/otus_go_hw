package hw6

import "fmt"

type Item struct {
	data interface{}
	next *Item
	prev *Item
}

func (i *Item) Prev() *Item {

	return i.prev
}

func (i *Item) Next() *Item {

	return i.next
}

func (i *Item) Value() interface{} {

	return i.data
}

type List struct {
	head *Item
	Size int
}

func (l *List) Len() int {

	return l.Size
}

func (l *List) Item() *Item {

	return l.head
}

func (l *List) Remove(remoteItem *Item) error {

	prev := remoteItem.prev
	if prev.next != remoteItem.data {
		return fmt.Errorf("нет данного элемента")
	}
	next := remoteItem.next
	if prev != nil {
		prev.next = next
	}
	if next != nil {
		next.prev = prev
	}
	l.Size--
	return nil
}

func (l *List) PushFont(data interface{}) {

	Item := new(Item)
	Item.data = data
	temp := l.First()
	Item.prev = temp
	l.head = Item
	l.Size++
	if temp != nil {
		temp.next = l.head
	}
}

func (l *List) PushBack(data interface{}) {

	temp := l.Last()
	Item := new(Item) //:= new(Item)
	Item.data = data
	Item.prev = nil
	Item.next = temp
	l.Size++
	temp.prev = Item
}

func (l *List) First() *Item {

	return l.head
}

func (l *List) Last() *Item {

	var i *Item
	for e := l.Item(); e != nil; e = e.Prev() {
		i = e
	}
	return i
}

func List_New() *List {

	l := new(List)
	l.Size = 0
	return l
}
