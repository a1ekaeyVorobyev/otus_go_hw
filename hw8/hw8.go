package hw8

import (
	"fmt"
	"sync"
	"sync/atomic"
)

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
	var cntExecute int32
	var cntExecuteGoroutine int32

	for isExecute {

		for _, v := range args {
			goroutines <- struct{}{}
			if cntErrorCurrent >= int32(cntError) {
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
				<-goroutines
				atomic.AddInt32(&cntExecuteGoroutine, 1)
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
