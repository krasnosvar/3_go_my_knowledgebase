package main

import (
	"context"
	"fmt"
	"time"
)

func run(ctx context.Context) error {
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("done")
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	fmt.Println("result:", run(ctx))
}
