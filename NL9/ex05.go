package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("CPUs used:", runtime.NumCPU())
	fmt.Println("GoRoutines used:", runtime.NumGoroutine())

	var counter int32
	const gs = 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt32(&counter, 1)
			fmt.Println(atomic.LoadInt32(&counter))
			wg.Done()

		}()
		fmt.Println("GoRoutines used: ", runtime.NumGoroutine())

	}

	wg.Wait()
	fmt.Println("GoRoutines used:", runtime.NumGoroutine())

}
