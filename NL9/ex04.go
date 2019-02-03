package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("CPUs used:", runtime.NumCPU())
	fmt.Println("GoRoutines used:", runtime.NumGoroutine())

	counter := 0
	const gs = 100
	wg.Add(gs)

	mu := sync.Mutex{}

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			v++
			counter = v
			fmt.Println("count: ", counter)
			mu.Unlock()
			wg.Done()

		}()
		fmt.Println("GoRoutines used: ", runtime.NumGoroutine())

	}

	wg.Wait()
	fmt.Println("GoRoutines used:", runtime.NumGoroutine())

}
