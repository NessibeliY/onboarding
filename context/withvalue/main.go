package main

import (
	"context"
	"fmt"
)

type key string

func main() {

	ctx := context.WithValue(context.Background(), key("username"), "johndoe")

	processRequest(ctx)
}

func processRequest(ctx context.Context) {
	if username, ok := ctx.Value(key("username")).(string); ok {
		fmt.Println("Username:", username)
	} else {
		fmt.Println("Username not found")
	}
}
