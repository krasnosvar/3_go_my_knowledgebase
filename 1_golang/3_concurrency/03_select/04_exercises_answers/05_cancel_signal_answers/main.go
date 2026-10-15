// Упражнение 05: выход из цикла по сигналу отмены (Решение).
package main

import (
	"fmt"
	"time"
)

func loop(done <-chan struct{}, tick <-chan int) int {
	n := 0
	for {
		select {
		case <-tick:
			n++
		case <-done:
			return n
		}
	}
}

func main() {
	fmt.Println("--- Проверка 05_cancel_signal (Решение) ---")

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
