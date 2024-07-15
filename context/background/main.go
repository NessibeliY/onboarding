package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	process(ctx)
}

func process(ctx context.Context) {
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("process completed")
	case <-ctx.Done():
		fmt.Println("process cancelled")
	}
}
