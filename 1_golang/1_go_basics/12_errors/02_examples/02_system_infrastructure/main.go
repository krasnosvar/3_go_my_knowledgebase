package main

import (
	"errors"
	"fmt"
	"time"
)

var (
	ErrConnectionTimeout = errors.New("connection timeout")
	ErrDatabaseDown      = errors.New("database is down")
)

// makeDatabaseCall симулирует нестабильный вызов базы данных
func makeDatabaseCall(attempt int) error {
	if attempt < 3 {
		// Симулируем сетевую ошибку на первых попытках
		return fmt.Errorf("network layer failed: %w", ErrConnectionTimeout)
	}
	return nil
}

func main() {
	fmt.Println("--- Симуляция клиента с Retry на основе errors.Is ---")

	maxRetries := 5
	for attempt := 1; attempt <= maxRetries; attempt++ {
		fmt.Printf("Попытка %d...\n", attempt)
		err := makeDatabaseCall(attempt)

		if err == nil {
			fmt.Println("Успешно: Запрос к БД выполнен!")
			break
		}

		// Проверяем, является ли ошибка временной (timeout)
		if errors.Is(err, ErrConnectionTimeout) {
			fmt.Printf("Временная ошибка (%v). Повтор через 50мс...\n", err)
			time.Sleep(50 * time.Millisecond)
			continue
		}

		// Неизвестная критическая ошибка — прекращаем попытки
		fmt.Printf("Критический сбой: %v. Прекращаем операции.\n", err)
		break
	}
}
