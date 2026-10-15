// Упражнение 01: параллельная сумма через WaitGroup (Решение).
package main

import (
	"fmt"
	"sync"
)

func parallelSum(parts [][]int) int {
	sums := make([]int, len(parts))
	var wg sync.WaitGroup
	for i, p := range parts {
		wg.Add(1)
		go func(i int, p []int) {
			defer wg.Done()
			for _, v := range p {
				sums[i] += v
			}
		}(i, p)
	}
	wg.Wait()
	total := 0
	for _, s := range sums {
		total += s
	}
	return total
}

func main() {
	fmt.Println("--- Проверка 01_waitgroup (Решение) ---")

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
