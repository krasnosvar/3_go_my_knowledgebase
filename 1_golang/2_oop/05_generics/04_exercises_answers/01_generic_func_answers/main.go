// Упражнение 01: обобщённые функции Map и Filter (Решение).
package main

import (
	"fmt"
)

func Map[T, U any](xs []T, f func(T) U) []U {
	out := make([]U, 0, len(xs))
	for _, x := range xs {
		out = append(out, f(x))
	}
	return out
}

func Filter[T any](xs []T, keep func(T) bool) []T {
	out := make([]T, 0, len(xs))
	for _, x := range xs {
		if keep(x) {
			out = append(out, x)
		}
	}
	return out
}

func main() {
	fmt.Println("--- Проверка 01_generic_func (Решение) ---")

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
