// Упражнение 05: подстановка токенов в шаблон (Решение).
package main

import (
	"fmt"
	"strings"
)

// renderTemplate подставляет значения vars в шаблон tpl.
func renderTemplate(tpl string, vars map[string]string) string {
	pairs := make([]string, 0, len(vars)*2)
	for k, v := range vars {
		pairs = append(pairs, "{{"+k+"}}", v)
	}
	return strings.NewReplacer(pairs...).Replace(tpl)
}

func main() {
	fmt.Println("--- Проверка 05_replace_token (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	got := renderTemplate("Hello, {{name}}! You have {{n}} messages.", map[string]string{"name": "Ann", "n": "3"})
	check(got == "Hello, Ann! You have 3 messages.", fmt.Sprintf("получено %q", got))

	got = renderTemplate("{{a}}{{a}} {{b}}", map[string]string{"a": "x"})
	check(got == "xx {{b}}", fmt.Sprintf("повторы заменяются, неизвестные токены остаются: получено %q", got))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
