// Упражнение 07: any как алиас interface{} (Решение).
package main

import (
	"fmt"
	"strings"
)

func Join(vals ...any) string {
	parts := make([]string, 0, len(vals))
	for _, v := range vals {
		parts = append(parts, fmt.Sprint(v))
	}
	return strings.Join(parts, ", ")
}

func Count[T any](xs []T, pred func(T) bool) int {
	n := 0
	for _, x := range xs {
		if pred(x) {
			n++
		}
	}
	return n
}

func main() {
	fmt.Println("--- Проверка 07_any_alias (Решение) ---")

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
