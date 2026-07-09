package main

import (
	"fmt"
	"time"
)

type CommandFunc func(args []string) error

// WithLogging — middleware для логирования выполнения и замера времени работы команды.
// Демонстрирует функцию первого класса, принимающую другую функцию и возвращающую новую.
func WithLogging(name string, cmd CommandFunc) CommandFunc {
	return func(args []string) error {
		fmt.Printf("[CLI] Запуск команды %q с аргументами %v...\n", name, args)
		start := time.Now()

		err := cmd(args)

		duration := time.Since(start)
		if err != nil {
			fmt.Printf("[CLI] Ошибка выполнения %q: %v (время: %v)\n", name, err, duration)
		} else {
			fmt.Printf("[CLI] Команда %q успешно выполнена за %v\n", name, duration)
		}
		return err
	}
}

func main() {
	// Обычная функция команды
	buildCmd := func(args []string) error {
		fmt.Println("  -> Собираем бинарный файл...")
		time.Sleep(50 * time.Millisecond) // имитация работы
		return nil
	}

	// Оборачиваем ее логгером
	loggedBuild := WithLogging("build", buildCmd)

	// Запуск
	_ = loggedBuild([]string{"-o", "bin/app", "./cmd/app"})
}
