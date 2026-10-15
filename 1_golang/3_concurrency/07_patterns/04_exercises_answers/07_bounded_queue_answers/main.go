// Упражнение 07: ограниченная очередь с отказом при переполнении (Решение).
package main

import (
	"fmt"
	"time"
)

type Queue struct {
	ch chan int
}

func newQueue(capacity int) *Queue {
	return &Queue{ch: make(chan int, capacity)}
}

func (q *Queue) TryPush(v int) bool {
	select {
	case q.ch <- v:
		return true
	default:
		return false
	}
}

func (q *Queue) Pop(timeout time.Duration) (int, bool) {
	select {
	case v := <-q.ch:
		return v, true
	case <-time.After(timeout):
		return 0, false
	}
}

func main() {
	fmt.Println("--- Проверка 07_bounded_queue (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	q := newQueue(2)
	check(q.TryPush(1), "первый TryPush должен пройти")
	check(q.TryPush(2), "второй TryPush должен пройти")
	check(!q.TryPush(3), "третий TryPush в полную очередь должен вернуть false")

	v, ok := q.Pop(50 * time.Millisecond)
	check(ok && v == 1, fmt.Sprintf("Pop = (%d, %t), ожидалось (1, true)", v, ok))
	check(q.TryPush(3), "после Pop место освободилось")
	v, _ = q.Pop(50 * time.Millisecond)
	check(v == 2, fmt.Sprintf("FIFO: ожидалось 2, получено %d", v))
	v, _ = q.Pop(50 * time.Millisecond)
	check(v == 3, fmt.Sprintf("FIFO: ожидалось 3, получено %d", v))

	start := time.Now()
	_, ok = q.Pop(20 * time.Millisecond)
	check(!ok, "Pop из пустой очереди должен вернуть false по таймауту")
	check(time.Since(start) >= 15*time.Millisecond, "Pop вернулся раньше таймаута")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
