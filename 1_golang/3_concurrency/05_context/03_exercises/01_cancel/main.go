// Упражнение 01: отмена через context.WithCancel.
//
// Задача: waitUntilCancelled(ctx) блокируется, пока ctx не отменён, и
// возвращает ctx.Err(). startAndCancel() создаёт контекст через
// context.WithCancel(context.Background()), вызывает cancel() и
// возвращает уже отменённый контекст.
//
// После отмены ctx.Done() закрыт, а ctx.Err() == context.Canceled.
//
// Запуск: go run ./03_exercises/01_cancel/main.go
package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// waitUntilCancelled ждёт отмены ctx и возвращает причину.
func waitUntilCancelled(ctx context.Context) error {
	// TODO: <-ctx.Done(); return ctx.Err()
	return nil
}

// startAndCancel возвращает уже отменённый контекст.
func startAndCancel() context.Context {
	// TODO: ctx, cancel := context.WithCancel(...); cancel(); return ctx
	return nil
}

func main() {
	fmt.Println("--- Проверка 01_cancel ---")

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
