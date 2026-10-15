// Упражнение 01: встраивание структуры.
//
// Задача: встройте A в B (поле без имени: просто `A`), чтобы метод Hello
// стал доступен у B напрямую: B{}.Hello().
//
// Запуск: go run ./03_exercises/01_basic_embedding/main.go
package main

import (
	"fmt"
)

// A — базовый тип с методом.
type A struct{}

// Hello возвращает приветствие.
func (A) Hello() string { return "hello" }

// B должен получить метод Hello через встраивание.
type B struct {
	// TODO: встройте A.
}

func main() {
	fmt.Println("--- Проверка 01_basic_embedding ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	var v any = B{}
	h, ok := v.(interface{ Hello() string })
	check(ok, "B должен иметь метод Hello (встройте A)")
	if ok {
		check(h.Hello() == "hello", fmt.Sprintf("Hello: получено %q", h.Hello()))
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
