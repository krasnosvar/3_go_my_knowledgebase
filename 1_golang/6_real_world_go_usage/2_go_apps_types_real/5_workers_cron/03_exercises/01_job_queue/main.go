// Упражнение 01: ограниченная очередь задач.
//
// Задача: реализуйте queue на буферизованном канале:
//   - newQueue(capacity) создаёт канал нужной ёмкости;
//   - TryEnqueue(j) — неблокирующая постановка: true, если место было,
//     false, если очередь полна (select с default) — так продюсер
//     узнаёт о перегрузке, а не зависает;
//   - TryDequeue() (job, bool) — неблокирующее извлечение в порядке FIFO.
//
// Запуск: go run ./03_exercises/01_job_queue
package main

import "fmt"

type job struct{ ID int }

type queue struct{ ch chan job }

func newQueue(capacity int) *queue { return &queue{ch: make(chan job, capacity)} }

func (q *queue) TryEnqueue(j job) bool {
	// TODO: select { case q.ch <- j: ...; default: ... }.
	return false
}

func (q *queue) TryDequeue() (job, bool) {
	// TODO: неблокирующее чтение из q.ch.
	return job{}, false
}

func main() {
	fmt.Println("--- Проверка 01_job_queue ---")

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
