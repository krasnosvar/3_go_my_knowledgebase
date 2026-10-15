// Упражнение 07: подсчёт уникальных и первый повтор.
//
// Задача: реализуйте две функции на множестве:
//   - countDistinct(xs) — число различных значений;
//   - firstDuplicate(xs) — первое значение, которое встретилось ПОВТОРНО
//     (по позиции второго вхождения): (v, true) или (0, false).
//
// Пример: [3 1 4 1 5 3] -> distinct=4, firstDuplicate=1.
//
// Запуск: go run ./03_exercises/07_count_distinct/main.go
package main

import "fmt"

// countDistinct возвращает число различных значений.
func countDistinct(xs []int) int {
	// TODO: множество значений.
	return 0
}

// firstDuplicate возвращает первое повторно встретившееся значение.
func firstDuplicate(xs []int) (int, bool) {
	// TODO: при повторной встрече значения верните его.
	return 0, false
}

func main() {
	fmt.Println("--- Проверка 07_count_distinct ---")

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
