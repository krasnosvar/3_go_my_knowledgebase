// Упражнение 04: переменная цикла в горутине (Решение).
package main

import (
	"fmt"
	"sort"
)

func collectIDs(ids []int) []int {
	out := make(chan int, len(ids))
	for _, id := range ids {
		go func(id int) {
			out <- id
		}(id)
	}
	res := make([]int, 0, len(ids))
	for range ids {
		res = append(res, <-out)
	}
	sort.Ints(res)
	return res
}

func main() {
	fmt.Println("--- Проверка 04_loop_capture (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	got := collectIDs([]int{5, 3, 1, 4, 2})
	want := []int{1, 2, 3, 4, 5}
	check(fmt.Sprint(got) == fmt.Sprint(want), fmt.Sprintf("collectIDs = %v, ожидалось %v", got, want))

	got = collectIDs(nil)
	check(len(got) == 0, "для пустого входа результат должен быть пустым")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
