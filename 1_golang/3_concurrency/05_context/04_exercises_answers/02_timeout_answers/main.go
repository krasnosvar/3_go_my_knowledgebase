// Упражнение 02: ограничение операции по времени (Решение).
package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func slowOp(ctx context.Context, d time.Duration) error {
	select {
	case <-time.After(d):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func runWithTimeout(timeout, work time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return slowOp(ctx, work)
}

func main() {
	fmt.Println("--- Проверка 02_timeout (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	start := time.Now()
	err := runWithTimeout(20*time.Millisecond, 500*time.Millisecond)
	check(errors.Is(err, context.DeadlineExceeded), fmt.Sprintf("долгая работа: %v, ожидалось context.DeadlineExceeded", err))
	check(time.Since(start) < 300*time.Millisecond, "операция не прервалась по таймауту")

	err = runWithTimeout(200*time.Millisecond, 5*time.Millisecond)
	check(err == nil, fmt.Sprintf("быстрая работа: %v, ожидалось nil", err))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
