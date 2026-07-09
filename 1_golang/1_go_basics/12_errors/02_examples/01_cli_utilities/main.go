package main

import (
	"errors"
	"fmt"
)

var (
	ErrConfigNotFound = errors.New("configuration file not found")
	ErrPermission     = errors.New("permission denied")
)

// readConfigFile симулирует чтение файла конфигурации
func readConfigFile(path string) error {
	if path == "" {
		return ErrConfigNotFound
	}
	return ErrPermission
}

// loadSettings симулирует загрузку настроек
func loadSettings(path string) error {
	err := readConfigFile(path)
	if err != nil {
		// Обертываем ошибку для сохранения контекста
		return fmt.Errorf("failed to read settings from %q: %w", path, err)
	}
	return nil
}

func main() {
	fmt.Println("--- Каскадная обработка и логирование ошибок ---")

	err := loadSettings("/etc/app/config.yaml")
	if err != nil {
		fmt.Printf("Полный лог ошибки:\n  -> %v\n\n", err)

		// Проверяем конкретные типы ошибок в цепочке через errors.Is
		if errors.Is(err, ErrConfigNotFound) {
			fmt.Println("Решение: Пожалуйста, проверьте наличие файла конфигурации.")
		} else if errors.Is(err, ErrPermission) {
			fmt.Println("Решение: Запустите программу с повышенными правами или проверьте права доступа.")
		}
	}
}
