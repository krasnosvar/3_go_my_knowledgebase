// Упражнение 05: ограничение comparable (Решение).
package main

import (
	"fmt"
)

// Point — comparable-структура.
type Point struct{ X, Y int }

func Index[T comparable](xs []T, v T) int {
	for i, x := range xs {
		if x == v {
			return i
		}
	}
	return -1
}

func Unique[T comparable](xs []T) []T {
	seen := make(map[T]struct{}, len(xs))
	out := make([]T, 0, len(xs))
	for _, x := range xs {
		if _, ok := seen[x]; ok {
			continue
		}
		seen[x] = struct{}{}
		out = append(out, x)
	}
	return out
}

func main() {
	fmt.Println("--- Проверка 05_comparable (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	check(Index([]string{"a", "b", "c"}, "c") == 2, "Index строки: ожидалось 2")
	check(Index([]int{1, 2}, 9) == -1, "Index отсутствующего: ожидалось -1")
	check(Index([]Point{{1, 2}, {3, 4}}, Point{3, 4}) == 1, "Index структуры: ожидалось 1")

	u := Unique([]int{3, 1, 3, 2, 1})
	check(fmt.Sprint(u) == "[3 1 2]", fmt.Sprintf("Unique int: ожидалось [3 1 2], получено %v", u))
	us := Unique([]string{"go", "go"})
	check(fmt.Sprint(us) == "[go]", fmt.Sprintf("Unique string: получено %v", us))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
