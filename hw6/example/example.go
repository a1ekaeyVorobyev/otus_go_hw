package main

import (
	"fmt"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw6"
)

func main() {
	fmt.Println("Создаем список ")
	s := hw6.List_New()

	s.PushBack("dd")
	s.PushBack("4")
	s.PushBack(5)

	fmt.Println("Размер = ", s.Len())

	elemLast := s.Last()

	fmt.Println("Последний = ", elemLast.Value())

	elemFirst := s.First()
	fmt.Println("Первый =", elemFirst.Value())

	fmt.Print("\nВыводим список с конца :	")
	for e := s.Last(); e != nil; e = e.Next() {
		fmt.Print(e.Value(), ";")
	}

	fmt.Println("\nДобавляем в начала sds")
	s.PushFont("sds")

	fmt.Print("Выводим список с конца :	")
	for e := s.Last(); e != nil; e = e.Next() {
		fmt.Print(e.Value(), ";")
	}

	fmt.Println("\nДобавляем в конец 15")
	s.PushBack(15)

	fmt.Print("Выводим список с конца :	")
	for e := s.Last(); e != nil; e = e.Next() {
		fmt.Print(e.Value(), ";")
	}
	fmt.Print("\nВыводим список с начала :	")

	for e := s.First(); e != nil; e = e.Prev() {
		fmt.Print(e.Value(), ";")
	}

	fmt.Println("\nУдаляем 2 элементы", elemLast.Value(), ";", elemFirst.Value())
	if err :=s.Remove(elemLast);err!=nil{
		fmt.Println(err.Error())
	}
	if err :=s.Remove(elemFirst);err!=nil{
		fmt.Println(err.Error())
	}
	fmt.Println("Удаляем элементы", elemFirst.Value())
	if err := s.Remove(elemFirst);err!=nil{
		fmt.Println(err.Error())
	}
	fmt.Print("Выводим список с конца : 	")

	for e := s.Last(); e != nil; e = e.Next() {
		fmt.Print(e.Value(), ";")
	}
	fmt.Print("\nВыводим список с начала : 	")

	for e := s.First(); e != nil; e = e.Prev() {
		fmt.Print(e.Value(), ";")
	}
	fmt.Print("\nУдаляем элемент из другого списка : 	")
	f:= hw6.List_New()
	f.PushBack("dd")
	if err := s.Remove(f.First()); err!=nil{
		fmt.Println(err.Error())
	}
}
