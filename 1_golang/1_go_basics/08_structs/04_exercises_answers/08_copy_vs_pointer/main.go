package main

import (
	"fmt"
)

type Config struct {
	MaxConn int
}

func ModifyConfigByCopy(cfg Config) {
	cfg.MaxConn = 10
}

func ModifyConfigByPointer(cfg *Config) {
	cfg.MaxConn = 20
}

func main() {
	fmt.Println("--- Проверка 08_copy_vs_pointer (Решение) ---")

	c := Config{MaxConn: 5}

	ModifyConfigByCopy(c)
	t1 := c.MaxConn == 5

	ModifyConfigByPointer(&c)
	t2 := c.MaxConn == 20

	if t1 && t2 {
		fmt.Println("Тест копирования vs указателя: PASS")
	} else {
		fmt.Printf("Тест копирования vs указателя: FAIL (c.MaxConn=%d)\n", c.MaxConn)
	}
}
