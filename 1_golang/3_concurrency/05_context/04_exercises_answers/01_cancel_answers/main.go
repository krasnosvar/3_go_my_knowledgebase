// Упражнение 01: отмена через context.WithCancel (Решение).
package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func waitUntilCancelled(ctx context.Context) error {
	<-ctx.Done()
	return ctx.Err()
}

func startAndCancel() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	return ctx
}

func main() {
	fmt.Println("--- Проверка 01_cancel (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	ctx := startAndCancel()
	if ctx == nil {
		check(false, "startAndCancel вернула nil")
	} else {
		check(errors.Is(ctx.Err(), context.Canceled), fmt.Sprintf("ctx.Err() = %v, ожидалось context.Canceled", ctx.Err()))
	}

	ctx2, cancel := context.WithCancel(context.Background())
	res := make(chan error, 1)
	go func() { res <- waitUntilCancelled(ctx2) }()
	time.Sleep(10 * time.Millisecond)
	select {
	case err := <-res:
		check(false, fmt.Sprintf("waitUntilCancelled вернулась (%v) до отмены", err))
	default:
	}
	cancel()
	select {
	case err := <-res:
		check(errors.Is(err, context.Canceled), fmt.Sprintf("после cancel получено %v, ожидалось context.Canceled", err))
	case <-time.After(300 * time.Millisecond):
		check(false, "waitUntilCancelled не вернулась после cancel")
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
