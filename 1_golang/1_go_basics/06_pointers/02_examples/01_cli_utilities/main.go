package main

import (
	"fmt"
)

type Config struct {
	Verbose bool
	Threads int
}

// Попытка настроить конфигурацию, передавая структуру ПО ЗНАЧЕНИЮ (копирование).
// Оригинальная структура не изменится!
func configureByValue(cfg Config) {
	cfg.Verbose = true
	cfg.Threads = 8
}

// Правильная настройка конфигурации, передавая структуру ПО УКАЗАТЕЛЮ.
// Все изменения сохранятся в оригинальном объекте.
func configureByPointer(cfg *Config) {
	cfg.Verbose = true
	cfg.Threads = 8
}

func main() {
	fmt.Println("--- CLI: Передача структуры по значению vs по указателю ---")

	// Создаем дефолтный конфиг
	appConfig := Config{Verbose: false, Threads: 1}

	// 1. Пытаемся настроить по значению
	configureByValue(appConfig)
	fmt.Printf("После configureByValue:   %+v (изменения потерялись!)\n", appConfig)

	// 2. Настраиваем по указателю (передаем адрес с помощью &)
	configureByPointer(&appConfig)
	fmt.Printf("После configureByPointer: %+v (изменения применились!)\n", appConfig)
}
