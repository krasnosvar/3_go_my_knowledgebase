// Упражнение 08: повторы с таймаутом на попытку.
//
// Задача: retry(attempt, attempts, perTry) вызывает attempt() до attempts
// раз. Каждый вызов возвращает канал, в который ПОЗЖЕ (или никогда) придёт
// true при успехе. На каждую попытку ждём не дольше perTry (select с
// time.After). Вернуть номер успешной попытки (с 1) и true, либо
// (attempts, false), если все попытки не удались.
//
// Запуск: go run ./03_exercises/08_retry_loop/main.go
package main

import (
	"fmt"
	"time"
)

// retry повторяет attempt, пока не придёт успех или не кончатся попытки.
func retry(attempt func() <-chan bool, attempts int, perTry time.Duration) (int, bool) {
	// TODO: цикл попыток, внутри select по каналу попытки и time.After(perTry).
	return 0, false
}

func main() {
	fmt.Println("--- Проверка 08_retry_loop ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	calls := 0
	flaky := func() <-chan bool {
		calls++
		ch := make(chan bool, 1)
		if calls == 3 {
			ch <- true // третья попытка отвечает сразу
		}
		return ch // первые две «зависают»
	}
	n, ok := retry(flaky, 5, 20*time.Millisecond)
	check(ok && n == 3, fmt.Sprintf("retry = (%d, %t), ожидалось (3, true)", n, ok))

	never := func() <-chan bool { return make(chan bool) }
	start := time.Now()
	n, ok = retry(never, 3, 10*time.Millisecond)
	check(!ok && n == 3, fmt.Sprintf("все попытки неудачны: (%d, %t), ожидалось (3, false)", n, ok))
	check(time.Since(start) >= 25*time.Millisecond, "таймаут каждой попытки должен выдерживаться")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
