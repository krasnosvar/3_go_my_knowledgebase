// Упражнение 03: базовый type switch.
//
// Задача: реализуйте kind через type switch:
//
//	int -> "int", float64 -> "float64", string -> "string", bool -> "bool",
//	nil -> "nil", любое другое -> "unknown".
//
// Запуск: go run ./03_exercises/03_type_switch_basic/main.go
package main

import (
	"fmt"
)

// kind возвращает название типа.
func kind(v any) string {
	// TODO: switch v.(type) { ... }
	return ""
}

func main() {
	fmt.Println("--- Проверка 03_type_switch_basic ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	cases := []struct {
		in   any
		want string
	}{
		{1, "int"}, {1.5, "float64"}, {"x", "string"}, {false, "bool"},
		{nil, "nil"}, {int64(1), "unknown"}, {[]int{1}, "unknown"},
	}
	for _, c := range cases {
		got := kind(c.in)
		check(got == c.want, fmt.Sprintf("kind(%#v): ожидалось %q, получено %q", c.in, c.want, got))
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
