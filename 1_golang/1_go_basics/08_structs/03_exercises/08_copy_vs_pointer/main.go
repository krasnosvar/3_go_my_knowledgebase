package main

import (
	"fmt"
)

type Config struct {
	MaxConn int
}

// ModifyConfigByCopy принимает структуру по значению (копию) и меняет MaxConn на 10.
func ModifyConfigByCopy(cfg Config) {
	// TODO: Измените MaxConn.
}

// ModifyConfigByPointer принимает указатель на структуру и меняет MaxConn на 20.
func ModifyConfigByPointer(cfg *Config) {
	// TODO: Измените MaxConn.
}

func main() {
	fmt.Println("--- Проверка 08_copy_vs_pointer ---")

	c := Config{MaxConn: 5}

	// 1. Попытка изменить копию
	ModifyConfigByCopy(c)
	t1 := c.MaxConn == 5 // Значение в оригинале не должно измениться!

	// 2. Изменение через указатель
	ModifyConfigByPointer(&c)
	t2 := c.MaxConn == 20 // Значение должно обновиться!

	if t1 && t2 {
		fmt.Println("Тест копирования vs указателя: PASS")
	} else {
		fmt.Printf("Тест копирования vs указателя: FAIL (c.MaxConn=%d)\n", c.MaxConn)
	}
}
