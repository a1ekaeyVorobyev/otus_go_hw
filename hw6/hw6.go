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
	size int
}

func (l *List) Len() int {
	return l.size
}

func (l *List) Remove(remoteItem *Item) error {
	if remoteItem.list  == nil {
		return fmt.Errorf("данный элемент  уже удален")
	}
	if l != remoteItem.list {
		return fmt.Errorf("данный элемент не принадлежит этому списку")
	}
	remoteItem.list = nil
	prev := remoteItem.prev
	next := remoteItem.next
	if prev != nil {
		prev.next = next
	}else{
		l.head = next
	}
	if next != nil {
		next.prev = prev
	}else{
		l.tail = prev
	}
	l.size--
	return nil
}

func (l *List) PushBack(data interface{}) {
	temp := l.Last()
	item := new(Item)
	item.data = data
	item.prev = temp
	l.tail = item
	if temp != nil {
		temp.next = item
	}
	if l.First() == nil {
		l.head = item
	}
	item.list = l
	l.size++
}

func (l *List) PushFont(data interface{}) {
	temp := l.First()
	item := new(Item) //:= new(Item)
	item.data = data
	item.prev = nil
	item.next = temp
	if temp != nil {
		temp.prev = item
	}
	l.head = item
	if l.Last() == nil {
		l.tail = item
	}
	item.list = l
	l.size++
}

func (l *List) First() *Item {
	return l.head
}

func (l *List) Last() *Item {
	return l.tail
}

func List_New() *List {
	l := new(List)
	l.size = 0
	return l
}
