// Упражнение 05: сигнал о завершении через канал (Решение).
package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func startWorker(work func()) <-chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		work()
	}()
	return done
}

func main() {
	fmt.Println("--- Проверка 05_channel_signal (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	var ran atomic.Bool
	done := startWorker(func() {
		time.Sleep(10 * time.Millisecond)
		ran.Store(true)
	})

	select {
	case <-done:
		check(ran.Load(), "канал закрыт до того, как work завершилась")
	case <-time.After(300 * time.Millisecond):
		check(false, "канал done так и не закрылся")
	}

	// закрытый канал должен читаться повторно без блокировки
	if done != nil {
		select {
		case <-done:
		default:
			check(false, "done нужно закрывать, а не отправлять в него значение")
		}
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
