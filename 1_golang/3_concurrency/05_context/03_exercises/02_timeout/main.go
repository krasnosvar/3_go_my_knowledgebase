// Упражнение 02: ограничение операции по времени.
//
// Задача:
//   - slowOp(ctx, d) «работает» d (time.After(d)) и возвращает nil, но если
//     ctx завершится раньше — сразу возвращает ctx.Err();
//   - runWithTimeout(timeout, work) создаёт context.WithTimeout от
//     Background, ОБЯЗАТЕЛЬНО вызывает cancel (defer cancel()) и запускает
//     slowOp(ctx, work).
//
// Запуск: go run ./03_exercises/02_timeout/main.go
package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// slowOp ждёт d или завершения ctx.
func slowOp(ctx context.Context, d time.Duration) error {
	// TODO: select { case <-time.After(d): ...; case <-ctx.Done(): ... }
	return nil
}

// runWithTimeout выполняет slowOp с таймаутом.
func runWithTimeout(timeout, work time.Duration) error {
	// TODO: ctx, cancel := context.WithTimeout(...); defer cancel(); ...
	return nil
}

func main() {
	fmt.Println("--- Проверка 02_timeout ---")

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
