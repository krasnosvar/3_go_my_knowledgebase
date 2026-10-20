// Упражнение 11: последовательность задержек exponential backoff (Решение).
package main

import "fmt"

func backoffDelay(attempt, base, maxDelay int) int {
	d := base
	for i := 0; i < attempt; i++ {
		d *= 2
		if d >= maxDelay {
			return maxDelay
		}
	}
	if d > maxDelay {
		return maxDelay
	}
	return d
}

func main() {
	fmt.Println("--- Проверка 11_backoff_sequence (Решение) ---")

	base, maxDelay := 100, 1000
	want := []int{100, 200, 400, 800, 1000, 1000}

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
