// Упражнение 03: склейка слов через strings.Join.
//
// Задача: реализуйте joinWords(words, sep):
//   - пустые строки и строки из одних пробелов пропускаются;
//   - у остальных слов обрезаются пробелы по краям;
//   - результат склеивается через sep (strings.Join);
//   - если слов не осталось — вернуть "".
//
// Запуск: go run ./03_exercises/03_join_words/main.go
package main

import (
	"fmt"
	"strings"
)

var _ = strings.Join // держим импорт живым в заготовке

// joinWords склеивает непустые слова через sep.
func joinWords(words []string, sep string) string {
	// TODO: отфильтруйте пустые слова, обрежьте пробелы, склейте strings.Join.
	return "x"
}

func main() {
	fmt.Println("--- Проверка 03_join_words ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	got := joinWords([]string{"go", "is", "fun"}, " ")
	check(got == "go is fun", fmt.Sprintf("ожидалось \"go is fun\", получено %q", got))

	got = joinWords([]string{" a ", "", "  ", "b"}, ",")
	check(got == "a,b", fmt.Sprintf("пустые слова пропускаются: ожидалось \"a,b\", получено %q", got))

	got = joinWords(nil, "-")
	check(got == "", fmt.Sprintf("для nil ожидалась пустая строка, получено %q", got))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
