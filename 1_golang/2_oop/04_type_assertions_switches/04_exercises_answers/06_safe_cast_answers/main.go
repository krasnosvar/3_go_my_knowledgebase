// Упражнение 06: безопасное приведение к int (Решение).
package main

import (
	"fmt"
	"strconv"
)

func toIntOr(v any, def int) int {
	switch x := v.(type) {
	case int:
		return x
	case int64:
		return int(x)
	case float64:
		if x == float64(int(x)) {
			return int(x)
		}
		return def
	case string:
		n, err := strconv.Atoi(x)
		if err != nil {
			return def
		}
		return n
	default:
		return def
	}
}

func main() {
	fmt.Println("--- Проверка 06_safe_cast (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	cases := []struct {
		in   any
		want int
	}{
		{5, 5}, {int64(8), 8}, {3.0, 3}, {3.5, -1},
		{"42", 42}, {"abc", -1}, {true, -1}, {nil, -1},
	}
	for _, c := range cases {
		got := toIntOr(c.in, -1)
		check(got == c.want, fmt.Sprintf("toIntOr(%#v, -1): ожидалось %d, получено %d", c.in, c.want, got))
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
