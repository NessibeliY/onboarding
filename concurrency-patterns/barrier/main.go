package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup, barrier *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("worker %d is doing some work\n", id)
	time.Sleep(time.Duration(id) * time.Second)
	fmt.Printf("worker %d reached the barrier\n", id)

	barrier.Done()
	barrier.Wait()

	fmt.Printf("worker %d is proceeding with further work\n", id)
}

func main() {
	var wg sync.WaitGroup
	var barrier sync.WaitGroup

	numWorkers := 5

	wg.Add(numWorkers)
	barrier.Add(numWorkers)

	for i := 1; i <= numWorkers; i++ {
		go worker(i, &wg, &barrier)
	}

	wg.Wait()
}
