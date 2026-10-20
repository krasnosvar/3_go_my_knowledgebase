// Упражнение 11: последовательность задержек exponential backoff.
//
// Задача: реализуйте backoffDelay(attempt, base, maxDelay int) int —
// экспоненциальную задержку БЕЗ jitter:
//
//	delay = base * 2^attempt, но не больше maxDelay.
//
// attempt отсчитывается с 0: attempt=0 -> base, 1 -> base*2, 2 -> base*4 ...
// Результат ограничен сверху значением maxDelay (потолок).
//
// Пример при base=100, maxDelay=1000:
//
//	attempt: 0    1    2    3    4
//	delay:   100  200  400  800  1000  (упёрлись в потолок)
//
// Запуск: go run ./03_exercises/11_backoff_sequence/main.go
package main

import "fmt"

// backoffDelay возвращает задержку для попытки attempt.
func backoffDelay(attempt, base, maxDelay int) int {
	// TODO: вычислите base * 2^attempt с потолком maxDelay.
	return 0
}

func main() {
	fmt.Println("--- Проверка 11_backoff_sequence ---")

	base, maxDelay := 100, 1000
	want := []int{100, 200, 400, 800, 1000, 1000} // attempt 0..5

	allOK := true
	for attempt, w := range want {
		got := backoffDelay(attempt, base, maxDelay)
		if got != w {
			allOK = false
			fmt.Printf("FAIL: backoffDelay(%d) = %d, ожидалось %d\n", attempt, got, w)
		}
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
