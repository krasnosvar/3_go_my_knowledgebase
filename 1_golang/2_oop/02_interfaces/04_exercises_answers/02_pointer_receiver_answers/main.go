// Упражнение 02: pointer receiver и реализация интерфейса (Решение).
package main

import (
	"fmt"
)

// Counter — интерфейс счётчика.
type Counter interface {
	Inc()
	Value() int
}

type counter struct {
	v int
}

func (c *counter) Inc()       { c.v++ }
func (c *counter) Value() int { return c.v }

func main() {
	fmt.Println("--- Проверка 02_pointer_receiver (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	var p any = &counter{}
	cnt, ok := p.(Counter)
	check(ok, "*counter должен реализовывать Counter")
	if ok {
		cnt.Inc()
		cnt.Inc()
		check(cnt.Value() == 2, fmt.Sprintf("после 2x Inc: ожидалось 2, получено %d", cnt.Value()))
	}

	var v any = counter{}
	_, valueOK := v.(Counter)
	check(!valueOK, "значение counter{} не должно реализовывать Counter (методы — с pointer receiver)")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
