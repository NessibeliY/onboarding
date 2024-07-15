package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	cond := sync.NewCond(&mu)

	wg.Add(2)

	go func() {
		fmt.Println("goroutine 1 is started")
		defer wg.Done()

		cond.L.Lock()
		defer cond.L.Unlock()

		fmt.Println("goroutine 1 is waiting for condition")
		cond.Wait()
		fmt.Println("goroutine 1 met the condition")

		fmt.Println("goroutine 1 is done")
	}()

	go func() {
		fmt.Println("goroutine 2 is started")
		defer wg.Done()

		time.Sleep(1 * time.Second) // simulating some work

		cond.L.Lock()
		defer cond.L.Unlock()

		fmt.Println("goroutine 2 is signaling for condition")
		cond.Signal()
		fmt.Println("goroutine 2 completed signaling")

		fmt.Println("goroutine 2 is done")
	}()

	wg.Wait()
}
