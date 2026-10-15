// Упражнение 02: очередь на срезе (Enqueue/Dequeue) (Решение).
package main

import "fmt"

// Queue — очередь целых чисел на срезе.
type Queue struct {
	items []int
}

func (q *Queue) Enqueue(v int) {
	q.items = append(q.items, v)
}

func (q *Queue) Dequeue() (int, bool) {
	if len(q.items) == 0 {
		return 0, false
	}
	v := q.items[0]
	q.items = q.items[1:]
	return v, true
}

func (q *Queue) Len() int {
	return len(q.items)
}

func main() {
	fmt.Println("--- Проверка 02_enqueue_dequeue (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	var q Queue
	_, ok := q.Dequeue()
	check(!ok, "Dequeue на пустой очереди должен вернуть ok=false")

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	check(q.Len() == 3, fmt.Sprintf("ожидался Len()=3, получено %d", q.Len()))

	var got []int
	for q.Len() > 0 && len(got) < 10 {
		v, _ := q.Dequeue()
		got = append(got, v)
	}
	check(fmt.Sprint(got) == "[1 2 3]", fmt.Sprintf("порядок Dequeue: ожидалось [1 2 3], получено %v", got))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
