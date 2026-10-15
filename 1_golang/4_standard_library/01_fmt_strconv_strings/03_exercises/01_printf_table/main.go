// Упражнение 01: строка таблицы через fmt.Sprintf.
//
// Задача: реализуйте formatRow(name, qty, price), которая возвращает строку
// фиксированной ширины для табличного вывода:
//   - name выровнено влево в поле шириной 10 символов (%-10s);
//   - qty выровнено вправо в поле шириной 5 (%5d);
//   - price выровнено вправо в поле шириной 8 с 2 знаками после точки (%8.2f);
//   - колонки разделены символом '|'.
//
// Пример: formatRow("apple", 3, 1.5) == "apple     |    3|    1.50"
//
// Запуск: go run ./03_exercises/01_printf_table/main.go
package main

import "fmt"

// formatRow форматирует одну строку таблицы.
func formatRow(name string, qty int, price float64) string {
	// TODO: верните fmt.Sprintf с глаголами %-10s, %5d и %8.2f через '|'.
	return ""
}

func main() {
	fmt.Println("--- Проверка 01_printf_table ---")

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
