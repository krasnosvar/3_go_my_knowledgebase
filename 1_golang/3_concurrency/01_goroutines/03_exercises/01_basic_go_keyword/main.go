// Упражнение 01: запуск горутин ключевым словом go.
//
// Задача: реализуйте launch(n, work) так, чтобы она запускала n горутин,
// каждая из которых вызывает work(id) со своим id от 0 до n-1.
//
// Важно: launch НЕ должна ждать завершения work — она только запускает
// горутины и сразу возвращается (именно это делает ключевое слово go).
// Проверка блокирует все work до тех пор, пока launch не вернётся.
//
// Запуск: go run ./03_exercises/01_basic_go_keyword/main.go
package main

import (
	"fmt"
	"time"
)

// launch запускает n горутин, каждая вызывает work(id).
func launch(n int, work func(id int)) {
	// TODO: в цикле запустите work(i) через go.
}

func main() {
	fmt.Println("--- Проверка 01_basic_go_keyword ---")

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
