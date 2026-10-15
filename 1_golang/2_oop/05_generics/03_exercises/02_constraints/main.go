// Упражнение 02: ограничения типов (constraints).
//
// Задача: реализуйте:
//   - Sum[T Number](xs []T) T — сумма элементов (Number — свой union-constraint);
//   - Max[T cmp.Ordered](xs []T) (T, bool) — максимум; для пустого слайса
//     (нулевое значение, false). cmp.Ordered — стандартный constraint (Go 1.21+)
//     для всех типов с операторами < > (числа и строки).
//
// Запуск: go run ./03_exercises/02_constraints/main.go
package main

import (
	"cmp"
	"fmt"
)

// Number — числовые типы (с ~, чтобы подходили и пользовательские типы).
type Number interface {
	~int | ~int64 | ~float64
}

// Celsius — пользовательский тип на базе float64.
type Celsius float64

// Sum складывает элементы.
func Sum[T Number](xs []T) T {
	// TODO: var s T; цикл; return s.
	var zero T
	return zero
}

// Max возвращает максимальный элемент.
func Max[T cmp.Ordered](xs []T) (T, bool) {
	// TODO: пустой -> (zero, false); иначе пройти и найти максимум.
	var zero T
	return zero, false
}

func main() {
	fmt.Println("--- Проверка 02_constraints ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	check(Sum([]int{1, 2, 3}) == 6, fmt.Sprintf("Sum int: ожидалось 6, получено %d", Sum([]int{1, 2, 3})))
	check(Sum([]float64{0.5, 0.25}) == 0.75, "Sum float64: ожидалось 0.75")
	check(Sum([]Celsius{1.5, 2.5}) == 4, "Sum []Celsius: ожидалось 4 (работает благодаря ~float64)")

	m, ok := Max([]int{3, 9, 2})
	check(ok && m == 9, fmt.Sprintf("Max int: ожидалось (9, true), получено (%d, %t)", m, ok))
	s, ok := Max([]string{"b", "c", "a"})
	check(ok && s == "c", fmt.Sprintf("Max string: ожидалось (\"c\", true), получено (%q, %t)", s, ok))
	_, ok = Max([]int{})
	check(!ok, "Max пустого слайса: ожидалось ok=false")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
