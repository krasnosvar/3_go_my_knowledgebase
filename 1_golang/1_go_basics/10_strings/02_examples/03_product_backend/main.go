package main

import (
	"fmt"
	"unicode"
)

// ValidatePassword проверяет пароль на требования безопасности:
// - Минимум 8 символов
// - Хотя бы одна заглавная буква
// - Хотя бы одна цифра
func ValidatePassword(password string) bool {
	// 1. Быстрая проверка длины в рунах (символах)
	runes := []rune(password)
	if len(runes) < 8 {
		return false
	}

	hasUpper := false
	hasDigit := false

	// 2. Проверяем типы рун
	for _, r := range runes {
		if unicode.IsUpper(r) {
			hasUpper = true
		}
		if unicode.IsDigit(r) {
			hasDigit = true
		}
	}

	return hasUpper && hasDigit
}

func main() {
	passwords := map[string]bool{
		"123":          false, // Слишком короткий
		"abcdefgh":     false, // Нет заглавных и цифр
		"Abcdefgh":     false, // Нет цифры
		"Abcdefgh1":    true,  // Валидный
		"парольПароль": false, // Нет цифры
		"Пароль123":    true,  // Валидный с кириллицей!
	}

	for pass, expected := range passwords {
		ok := ValidatePassword(pass)
		fmt.Printf("Пароль: %-15q | Валидный: %-5t | Корректно: %t\n", pass, ok, ok == expected)
	}
}
