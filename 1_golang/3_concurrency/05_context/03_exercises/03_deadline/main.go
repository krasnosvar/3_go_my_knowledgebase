// Упражнение 03: абсолютный дедлайн через WithDeadline.
//
// Задача: deadlineIn(parent, d) возвращает контекст с АБСОЛЮТНЫМ дедлайном
// time.Now().Add(d) (context.WithDeadline) и его cancel-функцию.
//
// Разница с WithTimeout: WithTimeout(ctx, d) — это сокращение для
// WithDeadline(ctx, time.Now().Add(d)). Дедлайн удобен, когда момент
// окончания уже известен (например, пришёл в запросе).
//
// Запуск: go run ./03_exercises/03_deadline/main.go
package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// deadlineIn возвращает контекст, истекающий через d.
func deadlineIn(parent context.Context, d time.Duration) (context.Context, context.CancelFunc) {
	// TODO: return context.WithDeadline(parent, time.Now().Add(d))
	return parent, func() {}
}

func main() {
	fmt.Println("--- Проверка 03_deadline ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	before := time.Now()
	ctx, cancel := deadlineIn(context.Background(), 30*time.Millisecond)
	defer cancel()

	dl, ok := ctx.Deadline()
	check(ok, "у контекста нет дедлайна")
	if ok {
		left := dl.Sub(before)
		check(left > 20*time.Millisecond && left < 100*time.Millisecond, fmt.Sprintf("дедлайн через %v, ожидалось ~30ms", left))
	}

	select {
	case <-ctx.Done():
		check(errors.Is(ctx.Err(), context.DeadlineExceeded), fmt.Sprintf("Err = %v, ожидалось context.DeadlineExceeded", ctx.Err()))
	case <-time.After(300 * time.Millisecond):
		check(false, "контекст не завершился по дедлайну")
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
