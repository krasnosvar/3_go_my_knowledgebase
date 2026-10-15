// Упражнение 01: type assertion с проверкой ok.
//
// Задача: реализуйте asString — извлеките string из any формой
// `s, ok := v.(string)`. Для не-string значений верните ("", false).
//
// Запуск: go run ./03_exercises/01_assertion_ok/main.go
package main

import (
	"fmt"
)

// asString пытается получить строку.
func asString(v any) (string, bool) {
	// TODO: s, ok := v.(string)
	return "", false
}

func main() {
	fmt.Println("--- Проверка 01_assertion_ok ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	s, ok := asString("go")
	check(ok && s == "go", fmt.Sprintf("asString(\"go\"): получено (%q, %t)", s, ok))

	s, ok = asString("")
	check(ok && s == "", "пустая строка — тоже string: ожидалось (\"\", true)")

	for _, v := range []any{1, nil, []byte("go")} {
		s, ok := asString(v)
		check(!ok && s == "", fmt.Sprintf("asString(%#v): ожидалось (\"\", false), получено (%q, %t)", v, s, ok))
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
