// Упражнение 04: type switch (Решение).
package main

import (
	"errors"
	"fmt"
)

func describe(v any) string {
	switch x := v.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case nil:
		return "nil"
	case error:
		return "error: " + x.Error()
	default:
		return "other"
	}
}

func main() {
	fmt.Println("--- Проверка 04_type_switch (Решение) ---")

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
		{1, "int"},
		{"go", "string"},
		{true, "bool"},
		{nil, "nil"},
		{errors.New("boom"), "error: boom"},
		{3.14, "other"},
	}
	for _, c := range cases {
		got := describe(c.in)
		check(got == c.want, fmt.Sprintf("describe(%#v): ожидалось %q, получено %q", c.in, c.want, got))
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
