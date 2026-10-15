// Упражнение 01: ограниченная очередь задач (Решение).
package main

import "fmt"

type job struct{ ID int }

type queue struct{ ch chan job }

func newQueue(capacity int) *queue { return &queue{ch: make(chan job, capacity)} }

func (q *queue) TryEnqueue(j job) bool {
	select {
	case q.ch <- j:
		return true
	default:
		return false
	}
}

func (q *queue) TryDequeue() (job, bool) {
	select {
	case j := <-q.ch:
		return j, true
	default:
		return job{}, false
	}
}

func main() {
	fmt.Println("--- Проверка 01_job_queue (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	q := newQueue(2)
	check(q.TryEnqueue(job{1}) && q.TryEnqueue(job{2}), "две задачи помещаются")
	check(!q.TryEnqueue(job{3}), "третья не помещается (очередь полна)")
	j, ok := q.TryDequeue()
	check(ok && j.ID == 1, "FIFO: первой выходит задача 1")
	check(q.TryEnqueue(job{3}), "после извлечения место освободилось")
	j, _ = q.TryDequeue()
	j2, _ := q.TryDequeue()
	check(j.ID == 2 && j2.ID == 3, "дальше 2 и 3")
	_, ok = q.TryDequeue()
	check(!ok, "пустая очередь → false без блокировки")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
