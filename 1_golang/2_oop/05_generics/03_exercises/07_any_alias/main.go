// Упражнение 07: any как алиас interface{}.
//
// Задача: any — это просто псевдоним interface{} (type any = interface{}).
// Реализуйте:
//   - Join(vals ...any) string — fmt.Sprint каждого значения через ", ";
//   - Count[T any](xs []T, pred func(T) bool) int — сколько элементов
//     удовлетворяет pred (здесь any — constraint «любой тип»).
//
// Обратите внимание: []any и []interface{} — один и тот же тип.
//
// Запуск: go run ./03_exercises/07_any_alias/main.go
package main

import (
	"fmt"
	"strings"
)

// Join склеивает значения любых типов.
func Join(vals ...any) string {
	// TODO: fmt.Sprint(v) для каждого + strings.Join(..., ", ").
	_ = strings.Join
	return ""
}

// Count считает элементы, удовлетворяющие pred.
func Count[T any](xs []T, pred func(T) bool) int {
	// TODO
	return 0
}

func main() {
	fmt.Println("--- Проверка 07_any_alias ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	got := Join(1, "go", true, 2.5)
	check(got == "1, go, true, 2.5", fmt.Sprintf("Join: получено %q", got))

	var legacy []interface{} = []any{"a", "b"} // один и тот же тип
	got = Join(legacy...)
	check(got == "a, b", fmt.Sprintf("Join([]interface{}...): получено %q", got))

	n := Count([]string{"go", "rust", "c"}, func(s string) bool { return len(s) <= 2 })
	check(n == 2, fmt.Sprintf("Count коротких строк: ожидалось 2, получено %d", n))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
