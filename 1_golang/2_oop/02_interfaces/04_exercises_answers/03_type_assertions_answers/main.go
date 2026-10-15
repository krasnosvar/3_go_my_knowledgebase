// Упражнение 03: безопасный type assertion (Решение).
package main

import (
	"fmt"
)

func toInt(v any) (int, bool) {
	n, ok := v.(int)
	return n, ok
}

func main() {
	fmt.Println("--- Проверка 03_type_assertions (Решение) ---")

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
