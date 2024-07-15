package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.TODO()
	process(ctx)
}

func process(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("process cancelled")
	default:
		fmt.Println("process running")
	}
}
