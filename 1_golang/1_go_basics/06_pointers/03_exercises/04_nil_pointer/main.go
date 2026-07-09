package main

import (
	"fmt"
)

type Config struct {
	Host string
}

// updateHost обновляет Host в конфиге. Функция должна быть безопасной к nil-указателям:
// если cfg равен nil, функция не должна падать с паникой, а должна просто проигнорировать вызов.
func updateHost(cfg *Config, newHost string) {
	// TODO: Реализуйте безопасную запись нового хоста.
}

func main() {
	fmt.Println("--- Проверка 04_nil_pointer ---")

	// Тест 1: Передаем валидный указатель
	c := &Config{Host: "localhost"}
	updateHost(c, "127.0.0.1")
	t1 := c.Host == "127.0.0.1"

	// Тест 2: Передаем nil-указатель (не должно падать с panic!)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Тест nil-safety: FAIL (произошла паника)")
		}
	}()
	updateHost(nil, "127.0.0.1")
	t2 := true // Дошли сюда без паники

	if t1 && t2 {
		fmt.Println("Тест nil-safety: PASS")
	} else {
		fmt.Println("Тест nil-safety: FAIL")
	}
}
