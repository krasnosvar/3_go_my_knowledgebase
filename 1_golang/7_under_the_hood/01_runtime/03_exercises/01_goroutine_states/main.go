// Упражнение 01: горутины в состоянии ожидания (waiting).
//
// Задача: реализуйте spawnBlocked(n, release, wg) так, чтобы:
//   - запускалось ровно n горутин;
//   - каждая горутина блокировалась на чтении из канала release
//     (состояние waiting: горутина жива, но не занимает P);
//   - после закрытия release каждая горутина завершалась и вызывала wg.Done().
//
// Проверка: пока release открыт, runtime.NumGoroutine() должен вырасти ровно
// на n; после close(release) и wg.Wait() — вернуться к исходному значению.
//
// Запуск: go run ./03_exercises/01_goroutine_states/main.go
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// spawnBlocked запускает n горутин, ждущих закрытия release.
func spawnBlocked(n int, release <-chan struct{}, wg *sync.WaitGroup) {
	// TODO: wg.Add(n); запустите n горутин, каждая ждёт <-release и вызывает wg.Done().
}

// waitGoroutines ждёт (до 1 с), пока число горутин не станет want.
func waitGoroutines(want int) int {
	deadline := time.Now().Add(time.Second)
	for time.Now().Before(deadline) {
		if runtime.NumGoroutine() == want {
			return want
		}
		time.Sleep(time.Millisecond)
	}
	return runtime.NumGoroutine()
}

func main() {
	fmt.Println("--- Проверка 01_goroutine_states ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	base := runtime.NumGoroutine()
	release := make(chan struct{})
	var wg sync.WaitGroup

	const n = 50
	spawnBlocked(n, release, &wg)

	got := waitGoroutines(base + n)
	check(got == base+n, fmt.Sprintf("пока release открыт, горутин должно быть %d, получено %d", base+n, got))

	close(release)
	wg.Wait()

	got = waitGoroutines(base)
	check(got == base, fmt.Sprintf("после close(release) горутин должно быть %d, получено %d", base, got))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
