// Упражнение 06: очередь с приоритетами на container/heap (Решение).
package main

import (
	"container/heap"
	"fmt"
)

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
	if pq[i].Priority != pq[j].Priority {
		return pq[i].Priority > pq[j].Priority
	}
	return pq[i].Name < pq[j].Name
}

// drain возвращает имена задач в порядке извлечения из очереди с приоритетами.
func drain(tasks []Task) []string {
	pq := &PQ{}
	for i := range tasks {
		heap.Push(pq, &tasks[i])
	}
	var out []string
	for pq.Len() > 0 {
		out = append(out, heap.Pop(pq).(*Task).Name)
	}
	return out
}

func main() {
	fmt.Println("--- Проверка 06_priority_queue (Решение) ---")

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
