package main

import (
	"fmt"
)

type Config struct {
	Host string
}

// updateHost безопасно обновляет Host
func updateHost(cfg *Config, newHost string) {
	if cfg != nil {
		cfg.Host = newHost
	}
}

func main() {
	fmt.Println("--- Проверка 04_nil_pointer (Решение) ---")

	// Тест 1: Передаем валидный указатель
	c := &Config{Host: "localhost"}
	updateHost(c, "127.0.0.1")
	t1 := c.Host == "127.0.0.1"

	// Тест 2: Передаем nil-указатель
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Тест nil-safety: FAIL (произошла паника)")
		}
	}()
	updateHost(nil, "127.0.0.1")
	t2 := true

	if t1 && t2 {
		fmt.Println("Тест nil-safety: PASS")
	} else {
		fmt.Println("Тест nil-safety: FAIL")
	}
}
