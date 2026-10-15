// Упражнение 02: pointer receiver и реализация интерфейса.
//
// Задача: реализуйте методы Inc и Value у типа counter с pointer receiver,
// чтобы *counter реализовывал Counter. Обратите внимание: значение counter{}
// при этом интерфейс НЕ реализует — method set значения не включает методы
// с pointer receiver.
//
// Запуск: go run ./03_exercises/02_pointer_receiver/main.go
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

// TODO: реализуйте (c *counter) Inc() и (c *counter) Value() int.

func main() {
	fmt.Println("--- Проверка 02_pointer_receiver ---")

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
