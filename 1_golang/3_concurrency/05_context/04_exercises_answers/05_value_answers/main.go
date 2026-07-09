package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.WithValue(context.Background(), "request-id", "req-42")
	fmt.Println(ctx.Value("request-id"))
}
