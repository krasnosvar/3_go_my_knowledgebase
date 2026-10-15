// Упражнение 07: ожидание канала с таймаутом (Решение).
package main

import (
	"fmt"
	"time"
)

func recvWithTimeout(ch <-chan int, d time.Duration) (int, bool) {
	select {
	case v := <-ch:
		return v, true
	case <-time.After(d):
		return 0, false
	}
}

func main() {
	fmt.Println("--- Проверка 07_timeout (Решение) ---")

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
