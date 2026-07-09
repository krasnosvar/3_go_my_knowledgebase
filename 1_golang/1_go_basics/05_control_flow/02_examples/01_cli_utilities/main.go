package main

import (
	"fmt"
)

func main() {
	fmt.Println("--- CLI: Обработка команд ---")

	// Симулируем ввод аргументов команды
	args := []string{"deploy", "run"}

	for _, cmd := range args {
		fmt.Printf("\nОбрабатываем команду: %q\n", cmd)

		switch cmd {
		case "deploy":
			fmt.Println("  [Step 1] Сборка проекта...")
			fmt.Println("  [Step 2] Проверка конфигурации...")
			// Используем fallthrough, чтобы после deploy автоматически выполнился шаг run
			fallthrough

		case "run":
			fmt.Println("  [Step 3] Запуск приложения...")

		case "clean":
			fmt.Println("  Очистка временных файлов...")

		default:
			fmt.Println("  Неизвестная команда. Введите help для списка команд.")
		}
	}
}
