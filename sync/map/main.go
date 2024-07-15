package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map

	// store values
	m.Store("foo", 1)
	m.Store("bar", 2)

	// load values
	if value, ok := m.Load("foo"); ok {
		fmt.Println("foo:", value)
	}

	// delete values
	m.Delete("bar")

	// iterate over values
	m.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
}
