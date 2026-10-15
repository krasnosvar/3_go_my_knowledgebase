// Упражнение 02: извлечение данных группами регулярного выражения.
//
// Задача: реализуйте extractIDs(text), которая находит все упоминания
// задач вида "#123" и возвращает номера как []int в порядке появления.
//   - используйте группу захвата: `#(\d+)` и FindAllStringSubmatch;
//   - номер берётся из m[1] и переводится в int через strconv.Atoi;
//   - если совпадений нет — пустой срез (len 0).
//
// Запуск: go run ./03_exercises/02_regex_extract/main.go
package main

import (
	"fmt"
	"regexp"
	"strconv"
)

var _ = regexp.MustCompile // держим импорты живыми в заготовке
var _ = strconv.Atoi

// extractIDs возвращает номера задач из текста.
func extractIDs(text string) []int {
	// TODO: FindAllStringSubmatch(text, -1), из каждого совпадения m[1] → int.
	return nil
}

func main() {
	fmt.Println("--- Проверка 02_regex_extract ---")

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
