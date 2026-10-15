// Упражнение 06: очередь с приоритетами на container/heap.
//
// Задача: допишите очередь с приоритетами задач на container/heap:
//   - Less(i, j): выше приоритет (больше Priority) — раньше; при равенстве
//     приоритетов раньше идёт задача с меньшим Name (tie-breaking для
//     детерминизма);
//   - drain(tasks): положите все задачи в PQ (heap.Push) и извлекайте
//     (heap.Pop), возвращая имена в порядке извлечения.
//
// Len, Swap, Push, Pop уже реализованы.
//
// Запуск: go run ./03_exercises/06_priority_queue/main.go
package main

import (
	"container/heap"
	"fmt"
)

// Держим импорт container/heap живым в заготовке.
var _ = heap.Init

// Task — задача с приоритетом.
type Task struct {
	Name     string
	Priority int
}

// PQ — очередь с приоритетами (heap.Interface).
type PQ []*Task

func (pq PQ) Len() int      { return len(pq) }
func (pq PQ) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PQ) Push(x any)   { *pq = append(*pq, x.(*Task)) }
func (pq *PQ) Pop() any {
	old := *pq
	t := old[len(old)-1]
	*pq = old[:len(old)-1]
	return t
}

// Less задаёт порядок: больший Priority раньше, при равенстве — меньший Name.
func (pq PQ) Less(i, j int) bool {
	// TODO: сравните Priority, затем Name.
	return false
}

// drain возвращает имена задач в порядке извлечения из очереди с приоритетами.
func drain(tasks []Task) []string {
	// TODO: heap.Push всех задач, затем heap.Pop до опустошения.
	return nil
}

func main() {
	fmt.Println("--- Проверка 06_priority_queue ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	tasks := []Task{
		{"email", 1},
		{"deploy", 5},
		{"backup", 3},
		{"alert", 5},
		{"report", 1},
	}
	got := drain(tasks)
	check(fmt.Sprint(got) == "[alert deploy backup email report]",
		fmt.Sprintf("ожидалось [alert deploy backup email report], получено %v", got))

	check(len(drain(nil)) == 0, "пустой вход: ожидался пустой результат")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
