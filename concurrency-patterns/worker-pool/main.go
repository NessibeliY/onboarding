package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("worker %d started job %d\n", id, job)
		time.Sleep(time.Second)
		fmt.Printf("worker %d finished job %d\n", id, job)
		results <- job * 2
	}
}

func main() {
	numJobs := 10
	numWorkers := 3

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	var wg sync.WaitGroup

	// start worker goroutines
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			worker(workerID, jobs, results)
		}(i)
	}

	// enqueue jobs
	for i := 1; i <= numJobs; i++ {
		jobs <- i
	}
	close(jobs)

	// wait for all workers to finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// collect results
	for result := range results {
		fmt.Printf("result: %d\n", result)
	}
}
