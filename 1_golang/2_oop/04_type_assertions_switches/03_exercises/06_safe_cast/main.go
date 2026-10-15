// Упражнение 06: безопасное приведение к int.
//
// Задача: реализуйте toIntOr(v, def) — привести «что угодно» к int или
// вернуть def:
//
//	int     -> как есть
//	int64   -> int(v)
//	float64 -> int(v), только если дробной части нет (3.0 -> 3, 3.5 -> def)
//	string  -> strconv.Atoi; при ошибке -> def
//	иначе   -> def
//
// Запуск: go run ./03_exercises/06_safe_cast/main.go
package main

import (
	"fmt"
	"strconv"
)

// toIntOr приводит значение к int или возвращает def.
func toIntOr(v any, def int) int {
	// TODO: type switch по int, int64, float64, string.
	_ = strconv.Atoi
	return def
}

func main() {
	fmt.Println("--- Проверка 06_safe_cast ---")

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
