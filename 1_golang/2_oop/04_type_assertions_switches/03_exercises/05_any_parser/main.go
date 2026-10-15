// Упражнение 05: разбор значений any.
//
// Задача: реализуйте format(v any) string — короткое описание значения:
//
//	int              -> "int:<n>"
//	string           -> "str:<s>"
//	[]int            -> "ints:<len>"
//	map[string]any   -> "obj:<кол-во ключей>"
//	nil              -> "null"
//	иначе            -> "other"
//
// Запуск: go run ./03_exercises/05_any_parser/main.go
package main

import (
	"fmt"
)

// format описывает значение.
func format(v any) string {
	// TODO: type switch с переменной: switch x := v.(type) { ... }
	return ""
}

func main() {
	fmt.Println("--- Проверка 05_any_parser ---")

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
		{7, "int:7"},
		{"go", "str:go"},
		{[]int{1, 2, 3}, "ints:3"},
		{map[string]any{"a": 1, "b": "x"}, "obj:2"},
		{nil, "null"},
		{3.14, "other"},
	}
	for _, c := range cases {
		got := format(c.in)
		check(got == c.want, fmt.Sprintf("format(%#v): ожидалось %q, получено %q", c.in, c.want, got))
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
