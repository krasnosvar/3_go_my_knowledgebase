// Упражнение 01: горутины в состоянии ожидания (waiting) (Решение).
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// spawnBlocked запускает n горутин, ждущих закрытия release.
// Пока горутина ждёт канал, планировщик держит её в состоянии waiting:
// она не занимает ни M, ни P и стоит лишь её стек (от 2 KB).
func spawnBlocked(n int, release <-chan struct{}, wg *sync.WaitGroup) {
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			<-release
		}()
	}
}

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
	fmt.Println("--- Проверка 01_goroutine_states (Решение) ---")

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
