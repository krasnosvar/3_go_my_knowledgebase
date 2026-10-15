// Упражнение 08: текстовый отчёт через strings.Builder (Решение).
package main

import (
	"fmt"
	"strings"
)

type item struct {
	Name string
	Qty  int
}

// formatReport строит многострочный отчёт с итогом.
func formatReport(items []item) string {
	var b strings.Builder
	total := 0
	for _, it := range items {
		fmt.Fprintf(&b, "%s: %d\n", it.Name, it.Qty)
		total += it.Qty
	}
	fmt.Fprintf(&b, "total: %d\n", total)
	return b.String()
}

func main() {
	fmt.Println("--- Проверка 08_format_report (Решение) ---")

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
