// Упражнение 08: проверка дедлайна на детерминированном времени.
//
// Задача: реализуйте remaining(now, deadline), которая возвращает:
//   - (deadline - now, false), если дедлайн ещё не наступил;
//   - (0, true), если now >= deadline (дедлайн истёк).
//
// Время передаётся параметром, а не берётся из time.Now() внутри функции —
// так логику легко тестировать без реального ожидания.
// Используйте deadline.Sub(now) и/или now.Before(deadline).
//
// Запуск: go run ./03_exercises/08_deadline_check/main.go
package main

import (
	"fmt"
	"time"
)

// remaining возвращает остаток времени до дедлайна и флаг «истёк».
func remaining(now, deadline time.Time) (time.Duration, bool) {
	// TODO: если now.Before(deadline) — вернуть deadline.Sub(now), false.
	return 0, false
}

func main() {
	fmt.Println("--- Проверка 08_deadline_check ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	base := time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC)
	deadline := base.Add(90 * time.Second)

	d, expired := remaining(base, deadline)
	check(!expired && d == 90*time.Second, fmt.Sprintf("за 90s до дедлайна: получено %v, expired=%t", d, expired))

	d, expired = remaining(deadline, deadline)
	check(expired && d == 0, fmt.Sprintf("ровно в дедлайн: получено %v, expired=%t (ожидалось 0, true)", d, expired))

	d, expired = remaining(deadline.Add(time.Minute), deadline)
	check(expired && d == 0, fmt.Sprintf("после дедлайна: получено %v, expired=%t", d, expired))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
