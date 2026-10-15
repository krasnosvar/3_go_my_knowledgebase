// Упражнение 03: базовый type switch (Решение).
package main

import (
	"fmt"
)

func kind(v any) string {
	switch v.(type) {
	case int:
		return "int"
	case float64:
		return "float64"
	case string:
		return "string"
	case bool:
		return "bool"
	case nil:
		return "nil"
	default:
		return "unknown"
	}
}

func main() {
	fmt.Println("--- Проверка 03_type_switch_basic (Решение) ---")

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
