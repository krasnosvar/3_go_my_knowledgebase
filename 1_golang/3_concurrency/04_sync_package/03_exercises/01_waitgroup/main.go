// Упражнение 01: параллельная сумма через WaitGroup.
//
// Задача: parallelSum(parts) считает сумму каждой части в отдельной
// горутине, кладёт частичные суммы в срез sums[i] (у каждой горутины
// свой индекс — гонки нет), ждёт всех через sync.WaitGroup и возвращает
// общую сумму.
//
// Запуск: go run ./03_exercises/01_waitgroup/main.go
package main

import "fmt"

// parallelSum суммирует parts параллельно.
func parallelSum(parts [][]int) int {
	// TODO: sums := make([]int, len(parts)); горутина на часть; wg.Wait(); сложить sums.
	return 0
}

func main() {
	fmt.Println("--- Проверка 01_waitgroup ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	got := parallelSum([][]int{{1, 2, 3}, {10, 20}, {}, {100}})
	check(got == 136, fmt.Sprintf("parallelSum = %d, ожидалось 136", got))
	check(parallelSum(nil) == 0, "сумма пустого входа должна быть 0")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
