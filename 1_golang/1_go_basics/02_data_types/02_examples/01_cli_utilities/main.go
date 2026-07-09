package main

import (
	"fmt"
	"os"
	"strconv"
)

// Config представляет настройки CLI-приложения.
// Обратите внимание на выбор типов: порт ограничен диапазоном uint16,
// а режим verbose - это просто логический флаг bool.
type Config struct {
	Port    uint16
	Verbose bool
	Mode    string
}

func main() {
	// Имитация аргументов командной строки: --port 8080 --verbose
	args := []string{"8080", "true", "production"}

	if len(args) < 3 {
		fmt.Println("Ошибка: недостаточно аргументов.")
		os.Exit(1) // Код выхода - обычный int
	}

	// Парсим порт: порт должен быть uint16 (0-65535)
	portVal, err := strconv.ParseUint(args[0], 10, 16)
	if err != nil {
		fmt.Printf("Некорректный порт: %v\n", err)
		os.Exit(1)
	}

	// Парсим verbose флаг: конвертируем строку в bool
	verboseVal, err := strconv.ParseBool(args[1])
	if err != nil {
		fmt.Printf("Некорректный флаг verbose: %v\n", err)
		os.Exit(1)
	}

	// Создаем конфигурацию
	cfg := Config{
		Port:    uint16(portVal), // Явное приведение uint64 -> uint16
		Verbose: verboseVal,
		Mode:    args[2],
	}

	fmt.Println("--- Инициализация CLI-утилиты ---")
	fmt.Printf("Запуск на порту: %d\n", cfg.Port)
	fmt.Printf("Отладочный режим: %t\n", cfg.Verbose)
	fmt.Printf("Режим окружения: %s\n", cfg.Mode)
}
