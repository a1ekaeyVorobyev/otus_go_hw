package hw8

import (
	"fmt"
	"testing"
)

func withError() error {

	return fmt.Errorf("Error")
}

func withOutError() error {

	return nil
}

//если ошибок маньше количства допустимых
func Test_Without_Errors(t *testing.T) {

	cntExecuteFunction := 2
	cntError := 13
	f := withOutError
	l := withError
	task := []func() error{f, l, f, l, l, l, l}
	errors := Run(task, cntExecuteFunction, cntError)
	if errors != nil {
		t.Error("Expected : ", errors.Error())
	}
}

func Test_With_Errors(t *testing.T) {

	cntExecuteFunction := 2
	cntError := 2
	f := withOutError
	l := withError
	task := []func() error{f, l, f, l, l, l, l}
	errors := Run(task, cntExecuteFunction, cntError)
	if errors == nil {
		t.Error("Expected : ", errors.Error())
	}
}
