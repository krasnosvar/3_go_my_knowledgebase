// Упражнение 04: type switch.
//
// Задача: реализуйте describe через type switch:
//   - int     -> "int"
//   - string  -> "string"
//   - bool    -> "bool"
//   - nil     -> "nil"
//   - error   -> "error: <текст ошибки>"
//   - иначе   -> "other"
//
// Запуск: go run ./03_exercises/04_type_switch/main.go
package main

import (
	"errors"
	"fmt"
)

// describe возвращает описание типа значения.
func describe(v any) string {
	// TODO: switch x := v.(type) { case int: ... case nil: ... case error: ... }
	return ""
}

func main() {
	fmt.Println("--- Проверка 04_type_switch ---")

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
