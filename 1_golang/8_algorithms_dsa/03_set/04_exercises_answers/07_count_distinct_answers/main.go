// Упражнение 07: подсчёт уникальных и первый повтор (Решение).
package main

import "fmt"

// countDistinct возвращает число различных значений.
func countDistinct(xs []int) int {
	seen := make(map[int]struct{}, len(xs))
	for _, x := range xs {
		seen[x] = struct{}{}
	}
	return len(seen)
}

// firstDuplicate возвращает первое повторно встретившееся значение.
func firstDuplicate(xs []int) (int, bool) {
	seen := make(map[int]struct{}, len(xs))
	for _, x := range xs {
		if _, ok := seen[x]; ok {
			return x, true
		}
		seen[x] = struct{}{}
	}
	return 0, false
}

func main() {
	fmt.Println("--- Проверка 07_count_distinct (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	xs := []int{3, 1, 4, 1, 5, 3}
	check(countDistinct(xs) == 4, fmt.Sprintf("countDistinct: ожидалось 4, получено %d", countDistinct(xs)))
	check(countDistinct(nil) == 0, "countDistinct(nil): ожидалось 0")

	v, ok := firstDuplicate(xs)
	check(ok && v == 1, fmt.Sprintf("firstDuplicate: ожидалось (1, true), получено (%d, %t)", v, ok))

	_, ok = firstDuplicate([]int{1, 2, 3})
	check(!ok, "firstDuplicate без повторов: ожидалось ok=false")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
