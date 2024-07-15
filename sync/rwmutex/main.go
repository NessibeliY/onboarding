package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	counter int
	rwMu    sync.RWMutex
)

func read(wg *sync.WaitGroup) {
	defer wg.Done()
	rwMu.RLock()
	defer rwMu.RUnlock()
	fmt.Println("reading counter:", counter)
	time.Sleep(time.Millisecond)
}

func write(wg *sync.WaitGroup) {
	defer wg.Done()
	rwMu.Lock()
	counter++
	fmt.Println("incremented counter to:", counter)
	rwMu.Unlock()
	time.Sleep(time.Millisecond)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go write(&wg)
	}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go read(&wg)
	}

	wg.Wait()
}
