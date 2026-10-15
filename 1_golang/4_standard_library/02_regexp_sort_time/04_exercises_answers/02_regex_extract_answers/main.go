// Упражнение 02: извлечение данных группами регулярного выражения (Решение).
package main

import (
	"fmt"
	"regexp"
	"strconv"
)

var issueRe = regexp.MustCompile(`#(\d+)`)

// extractIDs возвращает номера задач из текста.
func extractIDs(text string) []int {
	matches := issueRe.FindAllStringSubmatch(text, -1)
	ids := make([]int, 0, len(matches))
	for _, m := range matches {
		n, err := strconv.Atoi(m[1])
		if err != nil {
			continue // \d+ может не влезть в int — пропускаем
		}
		ids = append(ids, n)
	}
	return ids
}

func main() {
	fmt.Println("--- Проверка 02_regex_extract (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	got := extractIDs("fix #12 and #7, see also #300")
	check(fmt.Sprint(got) == "[12 7 300]", fmt.Sprintf("ожидалось [12 7 300], получено %v", got))

	got = extractIDs("no issues here, # alone")
	check(got != nil && len(got) == 0, fmt.Sprintf("без совпадений ожидался пустой (не nil) срез, получено %v", got))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
