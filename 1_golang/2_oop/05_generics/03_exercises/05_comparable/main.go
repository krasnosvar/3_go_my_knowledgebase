// Упражнение 05: ограничение comparable.
//
// Задача: реализуйте функции с ограничением comparable (типы, которые
// можно сравнивать через ==, например int, string, структуры из них):
//   - Index(xs, v) int — индекс первого вхождения или -1;
//   - Unique(xs) []T — элементы без повторов, в порядке первого появления.
//
// Запуск: go run ./03_exercises/05_comparable/main.go
package main

import (
	"fmt"
)

// Point — comparable-структура.
type Point struct{ X, Y int }

// Index возвращает индекс v в xs или -1.
func Index[T comparable](xs []T, v T) int {
	// TODO: цикл со сравнением ==.
	return -1
}

// Unique убирает повторы, сохраняя порядок.
func Unique[T comparable](xs []T) []T {
	// TODO: map[T]struct{} для «уже видели» + результирующий слайс.
	return nil
}

func main() {
	fmt.Println("--- Проверка 05_comparable ---")

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
