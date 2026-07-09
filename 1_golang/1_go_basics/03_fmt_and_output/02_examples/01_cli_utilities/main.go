package main

import (
	"fmt"
)

type Config struct {
	Env   string
	Port  int
	Debug bool
}

func main() {
	cfg := Config{
		Env:   "production",
		Port:  9000,
		Debug: false,
	}

	fmt.Println("--- Использование флагов форматирования fmt ---")

	// 1. Обычный флаг %v (только значения полей)
	fmt.Printf("Флаг %%v:   %v\n", cfg)

	// 2. Подробный флаг %+v (поля с именами)
	fmt.Printf("Флаг %%+v:  %+v\n", cfg)

	// 3. Go-синтаксис %#v (структура с типом и именами)
	fmt.Printf("Флаг %%#v:  %#v\n", cfg)

	// 4. Тип переменной %T
	fmt.Printf("Флаг %%T:   %T\n", cfg)
}
