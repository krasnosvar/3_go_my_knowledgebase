// Упражнение 06: распределение работы между P (основа work stealing).
//
// Задача: реализуйте parallelSum(nums, workers): разбейте nums на workers
// примерно равных кусков, посчитайте сумму каждого куска в отдельной
// горутине и сложите результаты. Планировщик раскидает горутины по P, а
// простаивающие P «украдут» работу из чужих очередей (work stealing).
//
// Учтите крайние случаи: workers <= 0 трактуйте как 1; workers > len(nums)
// — пустые куски не должны ломать результат; пустой nums -> 0.
//
// Запуск: go run ./03_exercises/06_work_stealing/main.go
package main

import (
	"fmt"
	"sync"
)

// parallelSum считает сумму nums, распределяя работу по workers горутинам.
func parallelSum(nums []int, workers int) int {
	// TODO: нарежьте nums на куски, посчитайте каждый в своей горутине
	// (sync.WaitGroup + слайс частичных сумм), сложите частичные суммы.
	_ = sync.WaitGroup{}
	return 0
}

func main() {
	fmt.Println("--- Проверка 06_work_stealing ---")

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
