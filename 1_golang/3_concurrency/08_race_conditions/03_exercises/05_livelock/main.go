// Упражнение 05: livelock и случайная задержка.
//
// Задача: acquireBoth(first, second) должна захватить оба мьютекса и
// вернуть true (вызывающий потом сам их отпустит). Две горутины берут
// одни и те же мьютексы в РАЗНОМ порядке.
//
// «Вежливый» алгоритм: Lock(first); если second.TryLock() не удался —
// Unlock(first) и повторить. Если обе горутины делают это синхронно,
// они бесконечно уступают друг другу — livelock (все заняты, прогресса нет).
// Лекарство — случайная пауза перед повтором (jitter):
//
//	time.Sleep(time.Duration(rand.Intn(100)) * time.Microsecond)
//
// Запуск: go run ./03_exercises/05_livelock/main.go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// acquireBoth захватывает оба мьютекса, уступая при конфликте.
func acquireBoth(first, second *sync.Mutex) bool {
	// TODO: цикл Lock(first) / TryLock(second) / Unlock(first) + jitter.
	return false
}

func main() {
	fmt.Println("--- Проверка 05_livelock ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	var a, b sync.Mutex
	var success atomic.Int64
	done := make(chan struct{})
	go func() {
		var wg sync.WaitGroup
		worker := func(x, y *sync.Mutex) {
			defer wg.Done()
			for i := 0; i < 200; i++ {
				if acquireBoth(x, y) {
					success.Add(1)
					y.Unlock()
					x.Unlock()
				}
			}
		}
		wg.Add(2)
		go worker(&a, &b)
		go worker(&b, &a)
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(2 * time.Second):
		check(false, "горутины не завершились — livelock или deadlock")
	}
	check(success.Load() == 400, fmt.Sprintf("успешных захватов %d, ожидалось 400", success.Load()))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
