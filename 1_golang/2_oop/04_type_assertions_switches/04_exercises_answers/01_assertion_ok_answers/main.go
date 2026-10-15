// Упражнение 01: type assertion с проверкой ok (Решение).
package main

import (
	"fmt"
)

func asString(v any) (string, bool) {
	s, ok := v.(string)
	return s, ok
}

func main() {
	fmt.Println("--- Проверка 01_assertion_ok (Решение) ---")

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
