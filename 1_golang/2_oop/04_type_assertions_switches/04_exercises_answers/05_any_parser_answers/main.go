// Упражнение 05: разбор значений any (Решение).
package main

import (
	"fmt"
)

func format(v any) string {
	switch x := v.(type) {
	case int:
		return fmt.Sprintf("int:%d", x)
	case string:
		return "str:" + x
	case []int:
		return fmt.Sprintf("ints:%d", len(x))
	case map[string]any:
		return fmt.Sprintf("obj:%d", len(x))
	case nil:
		return "null"
	default:
		return "other"
	}
}

func main() {
	fmt.Println("--- Проверка 05_any_parser (Решение) ---")

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
