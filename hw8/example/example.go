package main

import (
	"fmt"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw8"
)

func withError() error {

	fmt.Println("Hello, error")
	return fmt.Errorf("Error")
}

func withOutError() error {

	fmt.Println("Hello, fun")
	return nil
}

func main() {

	cntExecuteFunction := 2
	cntError := 3
	f := withOutError
	l := withError
	task := []func() error{f, l, f, l, l, l, l}
	errors := hw8.Run(task, cntExecuteFunction, cntError)
	if errors != nil {
		//text := strings.Split(errors.Error(), " ")
		//fmt.Println("Кол-во ошибок", text[2])
		fmt.Println("Превышенно число ошибок", errors)
	} else {
		fmt.Println("Finish")
	}
}
