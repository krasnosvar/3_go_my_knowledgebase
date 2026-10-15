// Упражнение 05: подстановка токенов в шаблон.
//
// Задача: реализуйте renderTemplate(tpl, vars), которая заменяет каждое
// вхождение "{{key}}" на vars[key]. Используйте strings.NewReplacer
// (одна проходка по строке, все замены сразу).
// Токены, которых нет в vars, остаются как есть.
//
// Запуск: go run ./03_exercises/05_replace_token/main.go
package main

import (
	"fmt"
	"strings"
)

var _ = strings.NewReplacer // держим импорт живым в заготовке

// renderTemplate подставляет значения vars в шаблон tpl.
func renderTemplate(tpl string, vars map[string]string) string {
	// TODO: соберите пары "{{key}}", value и примените strings.NewReplacer.
	return tpl
}

func main() {
	fmt.Println("--- Проверка 05_replace_token ---")

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
