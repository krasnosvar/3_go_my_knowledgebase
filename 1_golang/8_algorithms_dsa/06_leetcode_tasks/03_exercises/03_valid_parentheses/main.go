// Упражнение 03: валидные скобки (LeetCode 20) и минимум удалений (LeetCode 1249).
//
// Задача:
//  1. isValid(s) — строка из ()[]{} валидна, если каждая открывающая скобка
//     закрыта скобкой того же типа в правильном порядке (стек).
//  2. minRemoveToMakeValid(s) — удалите минимальное число круглых скобок,
//     чтобы строка стала валидной (буквы сохраняются). Идея: стек индексов
//     незакрытых '(' + пометка лишних ')'; помеченные символы выбрасываем.
//
// Запуск: go run ./03_exercises/03_valid_parentheses/main.go
package main

import (
	"fmt"
	"strings"
)

// Держим импорт strings живым в заготовке.
var _ = strings.Builder{}

// isValid проверяет строку из скобок ()[]{}.
func isValid(s string) bool {
	// TODO: стек открывающих скобок.
	return false
}

// minRemoveToMakeValid удаляет минимум круглых скобок для валидности.
func minRemoveToMakeValid(s string) string {
	// TODO: стек индексов '(' + множество индексов на удаление.
	return s
}

func main() {
	fmt.Println("--- Проверка 03_valid_parentheses ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	for _, c := range []struct {
		in   string
		want bool
	}{{"()", true}, {"()[]{}", true}, {"{[]}", true}, {"(]", false}, {"([)]", false}, {"]", false}} {
		got := isValid(c.in)
		check(got == c.want, fmt.Sprintf("isValid(%q): ожидалось %t, получено %t", c.in, c.want, got))
	}

	for _, c := range []struct{ in, want string }{
		{"lee(t(c)o)de)", "lee(t(c)o)de"},
		{"a)b(c)d", "ab(c)d"},
		{"))((", ""},
	} {
		got := minRemoveToMakeValid(c.in)
		check(got == c.want, fmt.Sprintf("minRemoveToMakeValid(%q): ожидалось %q, получено %q", c.in, c.want, got))
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
