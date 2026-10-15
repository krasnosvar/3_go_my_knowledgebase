// Упражнение 03: сортировка чисел по убыванию без изменения входа.
//
// Задача: реализуйте sortDesc(nums), которая возвращает НОВЫЙ срез,
// отсортированный по убыванию. Исходный срез менять нельзя.
//   - сначала скопируйте данные (make + copy или append([]int(nil), nums...));
//   - затем sort.Sort(sort.Reverse(sort.IntSlice(cp))) или sort.Slice.
//
// Запуск: go run ./03_exercises/03_sort_numbers/main.go
package main

import (
	"fmt"
	"sort"
)

var _ = sort.Ints // держим импорт живым в заготовке

// sortDesc возвращает отсортированную по убыванию копию nums.
func sortDesc(nums []int) []int {
	// TODO: скопируйте nums и отсортируйте копию по убыванию.
	return nums
}

func main() {
	fmt.Println("--- Проверка 03_sort_numbers ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	in := []int{3, 1, 4, 1, 5, 9, 2}
	got := sortDesc(in)
	check(fmt.Sprint(got) == "[9 5 4 3 2 1 1]", fmt.Sprintf("ожидалось [9 5 4 3 2 1 1], получено %v", got))
	check(fmt.Sprint(in) == "[3 1 4 1 5 9 2]", fmt.Sprintf("исходный срез изменён: %v", in))
	check(len(sortDesc(nil)) == 0, "sortDesc(nil) должен вернуть пустой срез")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
