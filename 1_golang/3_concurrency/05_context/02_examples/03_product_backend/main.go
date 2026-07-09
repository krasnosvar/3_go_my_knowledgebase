package main

import (
	"context"
	"fmt"
)

type ctxKey string

const requestIDKey ctxKey = "request-id"

func handle(ctx context.Context) {
	rid, _ := ctx.Value(requestIDKey).(string)
	fmt.Println("request-id:", rid)
}

func main() {
	ctx := context.WithValue(context.Background(), requestIDKey, "req-123")
	handle(ctx)
}
