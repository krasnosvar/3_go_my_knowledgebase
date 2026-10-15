// Упражнение 01: передача значения через небуферизованный канал.
//
// Задача: transfer(v) создаёт небуферизованный канал (make(chan int)),
// запускает горутину, которая отправляет в него v, и читает значение
// в вызывающей горутине. Возвращает прочитанное.
//
// Небуферизованный канал — точка встречи: отправка завершается только
// когда кто-то читает. Поэтому отправлять и читать в одной горутине
// нельзя (deadlock) — нужна вторая горутина.
//
// Запуск: go run ./03_exercises/01_unbuffered/main.go
package main

import "fmt"

// transfer передаёт v через небуферизованный канал и возвращает полученное.
func transfer(v int) int {
	// TODO: ch := make(chan int); go func() { ch <- v }(); return <-ch
	return 0
}

func main() {
	fmt.Println("--- Проверка 01_unbuffered ---")

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
