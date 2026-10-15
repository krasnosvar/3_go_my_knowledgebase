// Упражнение 01: запуск горутин ключевым словом go (Решение).
package main

import (
	"fmt"
	"time"
)

func launch(n int, work func(id int)) {
	for i := 0; i < n; i++ {
		go work(i)
	}
}

func main() {
	fmt.Println("--- Проверка 01_basic_go_keyword (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	gate := make(chan struct{})
	started := make(chan int, 10)
	work := func(id int) {
		<-gate // все горутины ждут, пока launch не вернётся
		started <- id
	}

	returned := make(chan struct{})
	go func() {
		launch(5, work)
		close(returned)
	}()

	select {
	case <-returned:
	case <-time.After(300 * time.Millisecond):
		check(false, "launch не вернулась: work вызывается синхронно, а не через go")
	}
	close(gate)

	seen := map[int]bool{}
	timeout := time.After(300 * time.Millisecond)
	for waiting := true; waiting && len(seen) < 5; {
		select {
		case id := <-started:
			seen[id] = true
		case <-timeout:
			waiting = false
		}
	}
	for i := 0; i < 5; i++ {
		check(seen[i], fmt.Sprintf("work(%d) не был вызван", i))
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
