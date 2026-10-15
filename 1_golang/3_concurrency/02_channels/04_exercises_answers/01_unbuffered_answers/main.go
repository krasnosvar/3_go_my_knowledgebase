// Упражнение 01: передача значения через небуферизованный канал (Решение).
package main

import "fmt"

func transfer(v int) int {
	ch := make(chan int)
	go func() {
		ch <- v
	}()
	return <-ch
}

func main() {
	fmt.Println("--- Проверка 01_unbuffered (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	check(transfer(42) == 42, fmt.Sprintf("transfer(42) = %d, ожидалось 42", transfer(42)))
	check(transfer(-7) == -7, fmt.Sprintf("transfer(-7) = %d, ожидалось -7", transfer(-7)))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
