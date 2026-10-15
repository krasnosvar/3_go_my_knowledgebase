// Упражнение 06: подсчёт вхождений подстроки.
//
// Задача: реализуйте countWord(text, word) — число НЕпересекающихся
// вхождений word в text без учёта регистра.
//   - используйте strings.ToLower и strings.Count;
//   - пустой word → 0 (обратите внимание: strings.Count(s, "") возвращает
//     len(s)+1 — это ловушка).
//
// Запуск: go run ./03_exercises/06_count_substring/main.go
package main

import (
	"fmt"
	"strings"
)

var _ = strings.Count // держим импорт живым в заготовке

// countWord считает непересекающиеся вхождения word без учёта регистра.
func countWord(text, word string) int {
	// TODO: обработайте пустой word, затем strings.Count по lower-версиям.
	return -1
}

func main() {
	fmt.Println("--- Проверка 06_count_substring ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	check(countWord("Go go GO gopher", "go") == 4, fmt.Sprintf("\"Go go GO gopher\"/\"go\": ожидалось 4, получено %d", countWord("Go go GO gopher", "go")))
	check(countWord("abababa", "aba") == 2, fmt.Sprintf("вхождения не пересекаются: ожидалось 2, получено %d", countWord("abababa", "aba")))
	check(countWord("hello", "") == 0, fmt.Sprintf("пустой word: ожидалось 0, получено %d", countWord("hello", "")))
	check(countWord("hello", "xyz") == 0, "отсутствующая подстрока → 0")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
