package main

import (
	"fmt"
	"sync"
)

func main() {
	var pool = sync.Pool{
		New: func() interface{} {
			return "new"
		},
	}

	pool.Put("foo")
	pool.Put("bar")

	for i := 0; i < 3; i++ {
		fmt.Println(pool.Get())
	}
}
