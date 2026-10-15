// Упражнение 04: нормализация пользовательского ввода (Решение).
package main

import (
	"fmt"
	"strings"
)

// normalizeInputs возвращает очищенные значения в нижнем регистре.
func normalizeInputs(in []string) []string {
	out := make([]string, 0, len(in))
	for _, s := range in {
		s = strings.ToLower(strings.TrimSpace(s))
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func main() {
	fmt.Println("--- Проверка 04_trim_inputs (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	in := []string{"  Alice ", "BOB", "   ", "\tCarol\n"}
	got := normalizeInputs(in)
	want := []string{"alice", "bob", "carol"}
	check(fmt.Sprint(got) == fmt.Sprint(want), fmt.Sprintf("ожидалось %q, получено %q", want, got))
	check(in[0] == "  Alice ", "исходный срез не должен меняться")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
