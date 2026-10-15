// Упражнение 07: ожидание канала с таймаутом.
//
// Задача: recvWithTimeout(ch, d) ждёт значение из ch не дольше d.
// Если значение пришло — вернуть (v, true), если истёк таймаут — (0, false).
//
// Используйте select с веткой case <-time.After(d).
//
// Запуск: go run ./03_exercises/07_timeout/main.go
package main

import (
	"fmt"
	"time"
)

// recvWithTimeout читает из ch или сдаётся через d.
func recvWithTimeout(ch <-chan int, d time.Duration) (int, bool) {
	// TODO: select { case v := <-ch: ...; case <-time.After(d): ... }
	return 0, false
}

func main() {
	fmt.Println("--- Проверка 07_timeout ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	ready := make(chan int, 1)
	ready <- 99
	v, ok := recvWithTimeout(ready, 100*time.Millisecond)
	check(ok && v == 99, fmt.Sprintf("готовое значение: получено (%d, %t), ожидалось (99, true)", v, ok))

	silent := make(chan int)
	start := time.Now()
	v, ok = recvWithTimeout(silent, 20*time.Millisecond)
	elapsed := time.Since(start)
	check(!ok && v == 0, "для молчащего канала ожидалось (0, false)")
	check(elapsed >= 15*time.Millisecond, fmt.Sprintf("вернулись через %v — таймаут не выдержан", elapsed))
	check(elapsed < 300*time.Millisecond, fmt.Sprintf("ждали %v — таймаут не сработал", elapsed))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
