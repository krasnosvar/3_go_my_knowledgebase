package main

import (
	"fmt"
)

// TODO: Объявите функцию greeting, которая принимает имя (string) и возраст (int),
// и возвращает строку формата: "Привет, <имя>! Тебе <возраст> лет."
// ВНИМАНИЕ: ниже заглушка, чтобы программа компилировалась — замените тело функции.
// Подсказка: fmt.Sprintf из темы 03_fmt_and_output.
func greeting(name string, age int) string {
	return ""
}

func main() {
	fmt.Println("--- Проверка 01_basic_declaration ---")

	res := greeting("Алексей", 25)
	if res == "Привет, Алексей! Тебе 25 лет." {
		fmt.Println("Тест greeting: PASS")
	} else {
		fmt.Printf("Тест greeting: FAIL (получено %q)\n", res)
	}
}
