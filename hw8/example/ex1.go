package main

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	Run([]func() error{
		func() error {
			fmt.Println("first")
			time.Sleep(time.Second)
			return errors.New("")
		},
		func() error {
			fmt.Println("second")
			time.Sleep(time.Second)
			return errors.New("")
		},
		func() error {
			fmt.Println("third")
			time.Sleep(time.Second)
			return errors.New("")
		},
	}, 1, 1)
}

//Run(task []func()error, cntExecuteFunction int, cntError int) error
func Run(task []func() error, cntExecuteFunction int, cntError int) error {
	error := Run1(cntExecuteFunction, cntError, task...)
	return error
}

func Run1(cntExecuteFunction int, cntError int, args ...func() error) error {

	var wg sync.WaitGroup
	goroutines := make(chan struct{}, cntExecuteFunction)
	var cntErrorCurrent int32
	isExecute := true
	var cntExecute int32 = 0
	var cntExecuteGoroutine int32 = 0
	for isExecute {
		for _, v := range args {
			goroutines <- struct{}{}
			if atomic.LoadInt32(&cntErrorCurrent) >= int32(cntError) {
				isExecute = false
				close(goroutines)
				return fmt.Errorf("колво ошибок %v из %v. Было запушенно горутин %v, выполненно %v", cntErrorCurrent, cntError, cntExecute, cntExecuteGoroutine)
			}
			wg.Add(1)
			cntExecute++
			go func(goroutines <-chan struct{}, wg *sync.WaitGroup, v func() error) {
				error := v()
				if error != nil {
					atomic.AddInt32(&cntErrorCurrent, 1)
				}
				atomic.AddInt32(&cntExecuteGoroutine, 1)
				<-goroutines
				wg.Done()
			}(goroutines, &wg, v)
		}

		isExecute = false
	}
	wg.Wait()
	close(goroutines)
	fmt.Printf("колво ошибок %v из %v. Было запушенно горутин %v, выполненно %v\n", cntErrorCurrent, cntError, cntExecute, cntExecuteGoroutine)
	return nil
}
