// Упражнение 05: ожидание события через sync.Cond.
//
// Задача: Ready позволяет нескольким горутинам ждать публикации данных.
//   - Wait() блокируется, пока ready == false, затем возвращает value;
//   - Publish(v) сохраняет value, ставит ready = true и будит ВСЕХ ждущих
//     (Broadcast).
//
// Ожидание всегда в цикле: for !r.ready { r.cond.Wait() } — под r.mu.
//
// Запуск: go run ./03_exercises/05_cond/main.go
package main

import (
	"fmt"
	"sync"
	"time"
)

// Ready — «доска объявлений» с ожиданием публикации.
type Ready struct {
	mu    sync.Mutex
	cond  *sync.Cond
	ready bool
	value string
}

func newReady() *Ready {
	r := &Ready{}
	r.cond = sync.NewCond(&r.mu)
	return r
}

// Wait ждёт публикации и возвращает значение.
func (r *Ready) Wait() string {
	// TODO: Lock; for !r.ready { r.cond.Wait() }; вернуть r.value; Unlock.
	return ""
}

// Publish публикует значение и будит всех ждущих.
func (r *Ready) Publish(v string) {
	// TODO: Lock; записать; ready = true; Unlock; r.cond.Broadcast().
}

func main() {
	fmt.Println("--- Проверка 05_cond ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	r := newReady()
	results := make(chan string, 3)
	for i := 0; i < 3; i++ {
		go func() { results <- r.Wait() }()
	}

	time.Sleep(10 * time.Millisecond)
	r.Publish("data")

	for i := 0; i < 3; i++ {
		select {
		case v := <-results:
			check(v == "data", fmt.Sprintf("ждущий получил %q, ожидалось \"data\"", v))
		case <-time.After(300 * time.Millisecond):
			check(false, "ждущая горутина не проснулась после Publish")
		}
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
