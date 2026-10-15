// Упражнение 08: текстовый отчёт через strings.Builder.
//
// Задача: реализуйте formatReport(items), которая строит отчёт:
//
//	apple: 3
//	pear: 10
//	total: 13
//
// Требования:
//   - строки идут в порядке items, каждая заканчивается '\n';
//   - последняя строка — "total: <сумма Qty>\n";
//   - собирайте результат в strings.Builder (fmt.Fprintf(&b, ...)),
//     а не конкатенацией строк в цикле.
//
// Запуск: go run ./03_exercises/08_format_report/main.go
package main

import (
	"fmt"
	"strings"
)

var _ strings.Builder // держим импорт живым в заготовке

type item struct {
	Name string
	Qty  int
}

// formatReport строит многострочный отчёт с итогом.
func formatReport(items []item) string {
	// TODO: strings.Builder + fmt.Fprintf для каждой строки и строки total.
	return ""
}

func main() {
	fmt.Println("--- Проверка 08_format_report ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	got := formatReport([]item{{"apple", 3}, {"pear", 10}})
	want := "apple: 3\npear: 10\ntotal: 13\n"
	check(got == want, fmt.Sprintf("ожидалось %q, получено %q", want, got))

	got = formatReport(nil)
	check(got == "total: 0\n", fmt.Sprintf("пустой список: ожидалось \"total: 0\\n\", получено %q", got))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
