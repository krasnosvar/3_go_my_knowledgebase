// Упражнение 06: подсчёт вхождений подстроки (Решение).
package main

import (
	"fmt"
	"strings"
)

// countWord считает непересекающиеся вхождения word без учёта регистра.
func countWord(text, word string) int {
	if word == "" {
		return 0
	}
	return strings.Count(strings.ToLower(text), strings.ToLower(word))
}

func main() {
	fmt.Println("--- Проверка 06_count_substring (Решение) ---")

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
