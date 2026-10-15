// Упражнение 08: повторы с таймаутом на попытку (Решение).
package main

import (
	"fmt"
	"time"
)

func retry(attempt func() <-chan bool, attempts int, perTry time.Duration) (int, bool) {
	for i := 1; i <= attempts; i++ {
		select {
		case ok := <-attempt():
			if ok {
				return i, true
			}
		case <-time.After(perTry):
		}
	}
	return attempts, false
}

func main() {
	fmt.Println("--- Проверка 08_retry_loop (Решение) ---")

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
