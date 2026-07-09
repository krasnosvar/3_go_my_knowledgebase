package main

import (
	"fmt"
	"slices"
)

// sortAndFind сортирует слайс и ищет элемент
func sortAndFind(nums []int, target int) ([]int, int, bool) {
	slices.Sort(nums)
	idx, found := slices.BinarySearch(nums, target)
	return nums, idx, found
}

func main() {
	fmt.Println("--- Проверка 05_sort_search (Решение) ---")

	s := []int{50, 10, 30, 20, 40}
	sorted, idx, found := sortAndFind(s, 30)

	t1 := found && idx == 2
	t2 := false
	if len(sorted) == 5 {
		t2 = sorted[0] == 10 && sorted[4] == 50
	}

	if t1 && t2 {
		fmt.Println("Тест sortAndFind: PASS")
	} else {
		fmt.Printf("Тест sortAndFind: FAIL (found=%t, idx=%d, sorted=%v)\n", found, idx, sorted)
	}
}
