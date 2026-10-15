package main

import (
	"fmt"
)

// greeting принимает имя (string) и возраст (int), и возвращает форматированную строку
func greeting(name string, age int) string {
	return fmt.Sprintf("Привет, %s! Тебе %d лет.", name, age)
}

func main() {
	fmt.Println("--- Проверка 01_basic_declaration (Решение) ---")

	res := greeting("Алексей", 25)
	if res == "Привет, Алексей! Тебе 25 лет." {
		fmt.Println("Тест greeting: PASS")
	} else {
		fmt.Printf("Тест greeting: FAIL (получено %q)\n", res)
	}
}
