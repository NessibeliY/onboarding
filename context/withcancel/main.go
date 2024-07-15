package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		time.Sleep(2 * time.Second)
		cancel()
	}()

	select {
	case <-ctx.Done():
		fmt.Println("Context cancelled")
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout")
	}
}
