// Упражнение 05: разбор и форматирование дат (Решение).
package main

import (
	"fmt"
	"time"
)

// parseDate разбирает дату ISO-формата "2006-01-02".
func parseDate(s string) (time.Time, error) {
	return time.Parse("2006-01-02", s)
}

// formatRU форматирует дату как "02.01.2006".
func formatRU(t time.Time) string {
	return t.Format("02.01.2006")
}

func main() {
	fmt.Println("--- Проверка 05_parse_date (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	t, err := parseDate("2024-03-09")
	check(err == nil && t.Year() == 2024 && t.Month() == time.March && t.Day() == 9,
		fmt.Sprintf("parseDate(\"2024-03-09\") = %v, %v", t, err))
	check(formatRU(t) == "09.03.2024", fmt.Sprintf("formatRU = %q, ожидалось \"09.03.2024\"", formatRU(t)))

	for _, bad := range []string{"2024-13-01", "09.03.2024", "2024-02-30", ""} {
		if _, err := parseDate(bad); err == nil {
			check(false, fmt.Sprintf("parseDate(%q) должен вернуть ошибку", bad))
		}
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
