package main

import (
	"fmt"
)

// Ситуация из жизни: при старте сервиса подключаемся к базе данных.
// Внешние переменные dbStatus и attempts должны получить результат подключения,
// но из-за затенения (shadowing) внутри блока if они остаются нетронутыми.
// (Классический вариант этого бага — `db, err := connectDB()` — разобран в теории 05_shadowing.md.)

func main() {
	fmt.Println("--- Проверка 07_shadowing ---")

	dbStatus := "nil"
	attempts := 0

	hasConfig := true

	if hasConfig {
		// TODO: Исправьте эту строчку. Здесь допущена ошибка затенения (shadowing):
		// из-за := внутри блока объявляются НОВЫЕ локальные dbStatus и attempts,
		// а внешние переменные остаются со старыми значениями.
		dbStatus, attempts := "connected", 1
		_, _ = dbStatus, attempts
	}

	// Проверяем внешние переменные
	if dbStatus == "connected" && attempts == 1 {
		fmt.Println("Тест инициализации DB: PASS")
	} else {
		fmt.Printf("Тест инициализации DB: FAIL (база данных осталась %q, attempts=%d — ошибка shadowing!)\n", dbStatus, attempts)
	}
}
