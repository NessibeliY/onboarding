package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go doWork(ctx)

	select {
	case <-ctx.Done():
		fmt.Println("main context done", ctx.Err())
	}
}

func doWork(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("main context done", ctx.Err())
			return
		default:
			fmt.Println("working...")
			time.Sleep(1 * time.Second)
		}
	}
}
