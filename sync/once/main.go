package main

import (
	"fmt"
	"sync"
	"time"
)

var once sync.Once

func initialize() {
	fmt.Println("initializing...")
}

func main() {
	for i := 0; i < 10; i++ {
		go once.Do(initialize)
	}

	// wait to ensure all goroutines have run
	time.Sleep(time.Second)
}
