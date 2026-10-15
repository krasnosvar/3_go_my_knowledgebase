// Упражнение 01: строка таблицы через fmt.Sprintf (Решение).
package main

import "fmt"

// formatRow форматирует одну строку таблицы.
func formatRow(name string, qty int, price float64) string {
	return fmt.Sprintf("%-10s|%5d|%8.2f", name, qty, price)
}

func main() {
	fmt.Println("--- Проверка 01_printf_table (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	got := formatRow("apple", 3, 1.5)
	check(got == "apple     |    3|    1.50", fmt.Sprintf("formatRow(apple, 3, 1.5) = %q", got))

	got = formatRow("watermelon", 120, 99.999)
	check(got == "watermelon|  120|  100.00", fmt.Sprintf("formatRow(watermelon, 120, 99.999) = %q (округление до 2 знаков)", got))

	got = formatRow("kiwi", 0, 0)
	check(len(got) == 10+1+5+1+8, fmt.Sprintf("длина строки должна быть 25, получено %d", len(got)))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
