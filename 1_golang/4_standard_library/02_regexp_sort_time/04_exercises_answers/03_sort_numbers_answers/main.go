// Упражнение 03: сортировка чисел по убыванию без изменения входа (Решение).
package main

import (
	"fmt"
	"sort"
)

// sortDesc возвращает отсортированную по убыванию копию nums.
func sortDesc(nums []int) []int {
	cp := make([]int, len(nums))
	copy(cp, nums)
	sort.Sort(sort.Reverse(sort.IntSlice(cp)))
	return cp
}

func main() {
	fmt.Println("--- Проверка 03_sort_numbers (Решение) ---")

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
