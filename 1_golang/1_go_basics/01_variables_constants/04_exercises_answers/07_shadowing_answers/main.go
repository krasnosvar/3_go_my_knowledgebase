package main

import (
	"fmt"
)

// Ситуация из жизни: при старте сервиса подключаемся к базе данных.
// Внешние переменные dbStatus и attempts должны получить результат подключения.

func main() {
	fmt.Println("--- Проверка 07_shadowing (Решение) ---")

	dbStatus := "nil"
	attempts := 0

	hasConfig := true

	if hasConfig {
		// Используем присвоение (=) вместо объявления (:=),
		// так как обе переменные уже объявлены во внешнем блоке.
		dbStatus, attempts = "connected", 1
	}

	// Проверяем внешние переменные
	if dbStatus == "connected" && attempts == 1 {
		fmt.Println("Тест инициализации DB: PASS")
	} else {
		fmt.Printf("Тест инициализации DB: FAIL (база данных осталась %q, attempts=%d — ошибка shadowing!)\n", dbStatus, attempts)
	}
}
