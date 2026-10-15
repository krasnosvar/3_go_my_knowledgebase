// Упражнение 07: ограниченная очередь с отказом при переполнении.
//
// Задача: Queue — очередь фиксированной ёмкости на буферизованном канале:
//   - newQueue(capacity) создаёт канал make(chan int, capacity);
//   - TryPush(v) кладёт без ожидания; если очередь полна — false
//     (backpressure: вызывающий сам решает, отбросить или повторить);
//   - Pop(timeout) ждёт элемент не дольше timeout: (v, true) или (0, false).
//
// Запуск: go run ./03_exercises/07_bounded_queue/main.go
package main

import (
	"fmt"
	"time"
)

// Queue — очередь ограниченной ёмкости.
type Queue struct {
	ch chan int
}

func newQueue(capacity int) *Queue {
	// TODO: return &Queue{ch: make(chan int, capacity)}
	return &Queue{}
}

// TryPush добавляет v без блокировки.
func (q *Queue) TryPush(v int) bool {
	// TODO: select с default.
	return false
}

// Pop извлекает элемент, ожидая не дольше timeout.
func (q *Queue) Pop(timeout time.Duration) (int, bool) {
	// TODO: select с time.After.
	return 0, false
}

func main() {
	fmt.Println("--- Проверка 07_bounded_queue ---")

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
