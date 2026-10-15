// Упражнение 03: безопасный type assertion.
//
// Задача: реализуйте toInt — безопасно извлеките int из any через
// type assertion в форме "comma ok" (v.(int) с двумя результатами).
// Для не-int значений (включая int64, float64, string, nil) верните (0, false).
// Паники быть не должно.
//
// Запуск: go run ./03_exercises/03_type_assertions/main.go
package main

import (
	"fmt"
)

// toInt извлекает int из any.
func toInt(v any) (int, bool) {
	// TODO: n, ok := v.(int)
	return 0, false
}

func main() {
	fmt.Println("--- Проверка 03_type_assertions ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	n, ok := toInt(7)
	check(ok && n == 7, fmt.Sprintf("toInt(7): ожидалось (7, true), получено (%d, %t)", n, ok))

	for _, v := range []any{"7", int64(7), 7.0, nil} {
		n, ok := toInt(v)
		check(!ok && n == 0, fmt.Sprintf("toInt(%#v): ожидалось (0, false), получено (%d, %t)", v, n, ok))
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
