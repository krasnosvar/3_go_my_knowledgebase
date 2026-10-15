// Упражнение 01: базовый select (Решение).
package main

import (
	"fmt"
	"time"
)

func firstOf(a, b <-chan string) string {
	select {
	case v := <-a:
		return v
	case v := <-b:
		return v
	}
}

func main() {
	fmt.Println("--- Проверка 01_basic_select (Решение) ---")

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
