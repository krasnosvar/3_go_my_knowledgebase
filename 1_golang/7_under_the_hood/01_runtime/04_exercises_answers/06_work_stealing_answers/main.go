// Упражнение 06: распределение работы между P (основа work stealing) (Решение).
package main

import (
	"fmt"
	"sync"
)

// parallelSum считает сумму nums, распределяя работу по workers горутинам.
func parallelSum(nums []int, workers int) int {
	if workers <= 0 {
		workers = 1
	}
	partial := make([]int, workers) // у каждой горутины своя ячейка — без гонок
	chunk := (len(nums) + workers - 1) / workers
	var wg sync.WaitGroup
	for w := 0; w < workers; w++ {
		lo := w * chunk
		if lo >= len(nums) {
			break
		}
		hi := min(lo+chunk, len(nums))
		wg.Add(1)
		go func(w int, part []int) {
			defer wg.Done()
			for _, v := range part {
				partial[w] += v
			}
		}(w, nums[lo:hi])
	}
	wg.Wait()

	total := 0
	for _, p := range partial {
		total += p
	}
	return total
}

func main() {
	fmt.Println("--- Проверка 06_work_stealing (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	nums := make([]int, 10001)
	want := 0
	for i := range nums {
		nums[i] = i
		want += i
	}

	for _, w := range []int{1, 2, 3, 8, 64} {
		got := parallelSum(nums, w)
		check(got == want, fmt.Sprintf("workers=%d: ожидалось %d, получено %d", w, want, got))
	}
	check(parallelSum([]int{1, 2, 3}, 10) == 6, "workers > len(nums): ожидалось 6")
	check(parallelSum(nil, 4) == 0, "пустой слайс: ожидалось 0")
	check(parallelSum([]int{5, 5}, 0) == 10, "workers=0 трактуется как 1: ожидалось 10")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
