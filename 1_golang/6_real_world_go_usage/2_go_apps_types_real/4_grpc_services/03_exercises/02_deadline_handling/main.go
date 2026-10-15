// Упражнение 02: бюджет времени из deadline клиента.
//
// gRPC передаёт deadline клиента серверу в context. Если оставшегося
// времени заведомо не хватит на операцию, лучше сразу отказать, чем
// нагружать БД работой, результат которой никто не дождётся.
//
// Задача: реализуйте handle(ctx, now time.Time, need time.Duration, work func() string) (string, error):
//   - нет deadline (ctx.Deadline() ok == false) → выполнить work;
//   - оставшееся время = deadline - now; если оно < need →
//     вернуть context.DeadlineExceeded, work НЕ вызывать;
//   - иначе выполнить work.
//
// now передаётся параметром, чтобы тест был детерминированным.
//
// Запуск: go run ./03_exercises/02_deadline_handling
package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func handle(ctx context.Context, now time.Time, need time.Duration, work func() string) (string, error) {
	// TODO: ctx.Deadline(), сравнение остатка с need.
	return "", nil
}

func main() {
	fmt.Println("--- Проверка 02_deadline_handling ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	now := time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC)
	calls := 0
	work := func() string { calls++; return "done" }

	ctx, cancel := context.WithDeadline(context.Background(), now.Add(time.Hour))
	defer cancel()
	res, err := handle(ctx, now, 100*time.Millisecond, work)
	check(err == nil && res == "done", "времени хватает → work выполнен")

	ctx2, cancel2 := context.WithDeadline(context.Background(), now.Add(50*time.Millisecond))
	defer cancel2()
	_, err = handle(ctx2, now, 100*time.Millisecond, work)
	check(errors.Is(err, context.DeadlineExceeded), "остаток 50ms < 100ms → DeadlineExceeded")
	check(calls == 1, "при нехватке времени work не вызывается")

	res, err = handle(context.Background(), now, time.Hour, work)
	check(err == nil && res == "done", "без deadline → выполняем")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
