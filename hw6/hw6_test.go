package hw6

import (
	"fmt"
	"testing"
)

func Test_CheckCreate(t *testing.T) {

	checkValue := []string{"dd", "4", "5", "true"}
	s := List_New()
	s.PushFont("dd")
	s.PushFont("4")
	s.PushFont(5)
	s.PushFont(true)
	if s.Len() != 4 {
		t.Error("Не верный размер ", s.Len())
	}
	count := 0

	for e := s.Last(); e != nil; e = e.Prev() {
		st := fmt.Sprintf("%v", e.Value())
		if st != checkValue[count] {

			t.Error(e.Value(), " Не cooтветствует ", checkValue[count])
		}
		count++
	}
	if s.First().Value() != true {
		t.Error(s.First().Value(), " не cooтветствует true")
	}
	if s.Last().Value() != "dd" {
		t.Error(s.Last().Value(), " не cooтветствует dd")
	}
}

func Test_Add(t *testing.T) {

	checkValue := []string{"dd", "4", "5", "true"}
	s := List_New()
	s.PushFont("dd")
	s.PushFont("4")
	s.PushFont(5)
	s.PushFont(true)
	count := 0

	for e := s.Last(); e != nil; e = e.Prev() {
		st := fmt.Sprintf("%v", e.Value())
		if st != checkValue[count] {

			t.Error(e.Value(), " не cooтветствует ", checkValue[count])
		}
		count++
	}
	checkValue = []string{"sds", "dd", "4", "5", "true", "15"}
	s.PushBack("sds")
	s.PushFont(15)
	count = 0

	for e := s.Last(); e != nil; e = e.Prev() {
		st := fmt.Sprintf("%v", e.Value())
		if st != checkValue[count] {

			t.Error(e.Value(), " не cooтветствует ", checkValue[count])
		}
		count++
	}
}
