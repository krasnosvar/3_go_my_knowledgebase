// Упражнение 02: валидация числа через strconv.Atoi.
//
// Задача: реализуйте parsePort(s), которая разбирает номер порта из строки:
//   - пробелы по краям допускаются (strings.TrimSpace);
//   - нечисловая строка → ошибка (ошибку strconv оберните через %w);
//   - число вне диапазона 1..65535 → ошибка;
//   - при ошибке возвращайте 0.
//
// Запуск: go run ./03_exercises/02_atoi_validation/main.go
package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Держим импорты живыми в заготовке (в решении они нужны).
var _ = strconv.Atoi
var _ = strings.TrimSpace

// parsePort разбирает порт из строки и проверяет диапазон.
func parsePort(s string) (int, error) {
	// TODO: TrimSpace → strconv.Atoi → проверка диапазона 1..65535.
	return 0, nil
}

func main() {
	fmt.Println("--- Проверка 02_atoi_validation ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	if p, err := parsePort("8080"); err != nil || p != 8080 {
		check(false, fmt.Sprintf("parsePort(\"8080\") = %d, %v; ожидалось 8080, nil", p, err))
	}
	if p, err := parsePort("  443 "); err != nil || p != 443 {
		check(false, fmt.Sprintf("parsePort(\"  443 \") = %d, %v; пробелы по краям допустимы", p, err))
	}
	for _, bad := range []string{"abc", "", "0", "65536", "-1"} {
		if p, err := parsePort(bad); err == nil || p != 0 {
			check(false, fmt.Sprintf("parsePort(%q) = %d, %v; ожидалась ошибка и 0", bad, p, err))
		}
	}
	_, err := parsePort("12x")
	var numErr *strconv.NumError
	check(errors.As(err, &numErr), "ошибка разбора должна оборачивать *strconv.NumError (используйте %w)")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
