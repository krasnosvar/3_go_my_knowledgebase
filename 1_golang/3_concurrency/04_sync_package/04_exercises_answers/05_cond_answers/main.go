// Упражнение 05: ожидание события через sync.Cond (Решение).
package main

import (
	"fmt"
	"sync"
	"time"
)

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

func (r *Ready) Wait() string {
	r.mu.Lock()
	defer r.mu.Unlock()
	for !r.ready {
		r.cond.Wait()
	}
	return r.value
}

func (r *Ready) Publish(v string) {
	r.mu.Lock()
	r.value = v
	r.ready = true
	r.mu.Unlock()
	r.cond.Broadcast()
}

func main() {
	fmt.Println("--- Проверка 05_cond (Решение) ---")

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
