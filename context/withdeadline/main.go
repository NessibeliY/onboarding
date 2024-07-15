package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	deadline := time.Now().Add(2 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	select {
	case <-ctx.Done():
		fmt.Println("context deadline reached")
	case <-time.After(3 * time.Second):
		fmt.Println("operation completed")
	}
}
