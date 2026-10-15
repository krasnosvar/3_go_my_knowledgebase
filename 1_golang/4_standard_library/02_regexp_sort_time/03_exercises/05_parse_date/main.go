// Упражнение 05: разбор и форматирование дат.
//
// Задача:
//   - parseDate(s) разбирает дату формата "YYYY-MM-DD" (layout "2006-01-02")
//     через time.Parse; неверная строка → ошибка;
//   - formatRU(t) возвращает дату в виде "DD.MM.YYYY" (layout "02.01.2006").
//
// Помните: в Go layout задаётся эталонной датой Mon Jan 2 15:04:05 MST 2006,
// а не буквами вроде YYYY.
//
// Запуск: go run ./03_exercises/05_parse_date/main.go
package main

import (
	"fmt"
	"time"
)

// parseDate разбирает дату ISO-формата "2006-01-02".
func parseDate(s string) (time.Time, error) {
	// TODO: time.Parse с layout "2006-01-02".
	return time.Time{}, nil
}

// formatRU форматирует дату как "02.01.2006".
func formatRU(t time.Time) string {
	// TODO: t.Format с layout "02.01.2006".
	return ""
}

func main() {
	fmt.Println("--- Проверка 05_parse_date ---")

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
