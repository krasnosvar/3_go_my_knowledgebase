// Упражнение 05: выход из цикла по сигналу отмены.
//
// Задача: loop(done, tick) в цикле читает значения из tick и считает их.
// Когда закрывается done — функция возвращает количество обработанных.
//
// Используйте for { select { case <-tick: ...; case <-done: return n } }.
//
// Запуск: go run ./03_exercises/05_cancel_signal/main.go
package main

import (
	"fmt"
	"time"
)

// loop считает значения из tick до закрытия done.
func loop(done <-chan struct{}, tick <-chan int) int {
	// TODO: бесконечный цикл с select по tick и done.
	return 0
}

func main() {
	fmt.Println("--- Проверка 05_cancel_signal ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	done := make(chan struct{})
	tick := make(chan int)
	go func() {
		for i := 0; i < 3; i++ {
			select {
			case tick <- i: // небуферизованный: отправка завершится только после чтения
			case <-time.After(200 * time.Millisecond):
			}
		}
		close(done)
	}()

	res := make(chan int, 1)
	go func() { res <- loop(done, tick) }()

	select {
	case n := <-res:
		check(n == 3, fmt.Sprintf("обработано %d значений, ожидалось 3", n))
	case <-time.After(time.Second):
		check(false, "loop не завершилась после закрытия done")
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
