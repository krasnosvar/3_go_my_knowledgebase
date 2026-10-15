// Упражнение 01: обобщённые функции Map и Filter.
//
// Задача: реализуйте две обобщённые функции:
//   - Map[T, U any](xs []T, f func(T) U) []U — применяет f к каждому элементу;
//   - Filter[T any](xs []T, keep func(T) bool) []T — оставляет элементы,
//     для которых keep вернул true.
//
// Для пустого входа обе возвращают пустой (len 0) слайс.
//
// Запуск: go run ./03_exercises/01_generic_func/main.go
package main

import (
	"fmt"
)

// Map применяет f к каждому элементу.
func Map[T, U any](xs []T, f func(T) U) []U {
	// TODO: make([]U, 0, len(xs)) + цикл.
	return nil
}

// Filter оставляет элементы, удовлетворяющие keep.
func Filter[T any](xs []T, keep func(T) bool) []T {
	// TODO: цикл + append.
	return nil
}

func main() {
	fmt.Println("--- Проверка 01_generic_func ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	sq := Map([]int{1, 2, 3}, func(x int) int { return x * x })
	check(fmt.Sprint(sq) == "[1 4 9]", fmt.Sprintf("Map квадраты: получено %v", sq))

	strs := Map([]int{1, 2}, func(x int) string { return fmt.Sprintf("#%d", x) })
	check(fmt.Sprint(strs) == "[#1 #2]", fmt.Sprintf("Map int->string: получено %v", strs))

	even := Filter([]int{1, 2, 3, 4}, func(x int) bool { return x%2 == 0 })
	check(fmt.Sprint(even) == "[2 4]", fmt.Sprintf("Filter чётные: получено %v", even))

	empty := Filter([]string{}, func(string) bool { return true })
	check(empty != nil && len(empty) == 0, "Filter пустого входа: ожидался пустой не-nil слайс")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
