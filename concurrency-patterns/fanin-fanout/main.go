package main

import (
	"fmt"
	"sync"
)

func main() {
	data := []int{1, 2, 3, 4, 5}
	input := make(chan int, len(data))
	for _, d := range data {
		input <- d
	}
	close(input)

	// fan-out: Launch multiple worker goroutines
	numWorkers := 3
	results := make(chan int, len(data))

	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for num := range input {
				// simulate some processing
				result := num * 2
				results <- result
			}
		}()
	}

	// fan-in: Aggregate results from workers
	go func() {
		wg.Wait()
		close(results)
	}()

	// process aggregated results
	for result := range results {
		fmt.Println(result)
	}
}
