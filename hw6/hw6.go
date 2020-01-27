package hw6

import "fmt"

type Item struct {
	data interface{}
	next *Item
	prev *Item
	list *List
}

func (i *Item) Prev() *Item {
	return i.prev
}

func (i *Item) Next() *Item {
	return i.next
}

func (i *Item) GetList() *List {
	return i.list
}

func (i *Item) Value() interface{} {
	return i.data
}

type List struct {
	head *Item
	tail *Item
	Size int
}

func (l *List) Len() int {
	return l.Size
}

func (l *List) Remove(remoteItem *Item) error {
	if (l != remoteItem.list){
		return fmt.Errorf("данный элемент не пренадлежит этому списку")
	}
	prev := remoteItem.prev
	next := remoteItem.next
	if prev != nil {
		if prev.next != remoteItem.data {
			return fmt.Errorf("нет данного элемента")
		}
		prev.next = next
	}
	if next != nil {
		if next.prev != remoteItem.data {
			return fmt.Errorf("нет данного элемента")
		}
		next.prev = prev
	}
	l.Size--
	return nil
}

func (l *List) PushFont(data interface{}) {
	item := new(Item)
	item.data = data
	temp := l.First()
	item.prev = temp
	l.head = item
	if temp != nil {
		temp.next = l.head
	}
	if l.Last() == nil {
		l.tail = item
	}
	item.list=l
	l.Size++
}

func (l *List) PushBack (data interface{}) {
	temp := l.Last()
	item := new(Item) //:= new(Item)
	item.data = data
	item.prev = nil
	item.next = temp
	temp.prev = item
	l.tail = item
	if l.First() == nil {
		l.head = item
	}
	item.list=l
	l.Size++
}

func (l *List) First() *Item {
	return l.head
}

func (l *List) Last() *Item {
	return l.tail
}

func List_New() *List {
	l := new(List)
	l.Size = 0
	return l
}
