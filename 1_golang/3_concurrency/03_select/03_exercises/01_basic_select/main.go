// Упражнение 01: базовый select.
//
// Задача: firstOf(a, b) ждёт первое значение из любого из двух каналов
// и возвращает его. Используйте select с двумя ветками case.
//
// Запуск: go run ./03_exercises/01_basic_select/main.go
package main

import (
	"fmt"
	"time"
)

// firstOf возвращает первое пришедшее значение из a или b.
func firstOf(a, b <-chan string) string {
	// TODO: select { case v := <-a: ...; case v := <-b: ... }
	return ""
}

func main() {
	fmt.Println("--- Проверка 01_basic_select ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	a := make(chan string, 1)
	b := make(chan string, 1)
	a <- "from-a"
	check(firstOf(a, b) == "from-a", "значение готово только в a — ожидалось \"from-a\"")

	b <- "from-b"
	check(firstOf(a, b) == "from-b", "значение готово только в b — ожидалось \"from-b\"")

	slow := make(chan string)
	go func() {
		time.Sleep(10 * time.Millisecond)
		slow <- "late"
	}()
	check(firstOf(a, slow) == "late", "select должен дождаться значения из медленного канала")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
