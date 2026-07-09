package main

import (
	"fmt"
	"slices"
)

var _ = slices.Sort[[]int] // Держим импорт активным для компиляции заготовки

// sortAndFind сортирует переданный слайс nums по возрастанию,
// затем ищет в нем target с помощью slices.BinarySearch.
// Возвращает отсортированный слайс, найденный индекс и флаг нахождения (bool).
func sortAndFind(nums []int, target int) ([]int, int, bool) {
	// TODO: Реализуйте сортировку и бинарный поиск.
	return nil, -1, false
}

func main() {
	fmt.Println("--- Проверка 05_sort_search ---")

	s := []int{50, 10, 30, 20, 40}
	sorted, idx, found := sortAndFind(s, 30)

	t1 := found && idx == 2 // [10, 20, 30, 40, 50] -> 30 на индексе 2
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
