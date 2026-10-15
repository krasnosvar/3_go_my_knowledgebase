// Упражнение 05: очередь на двух стеках (Решение).
package main

import "fmt"

// TwoStackQueue — очередь на двух стеках.
type TwoStackQueue struct {
	in  []int
	out []int
}

func (q *TwoStackQueue) Enqueue(v int) {
	q.in = append(q.in, v)
}

func (q *TwoStackQueue) Dequeue() (int, bool) {
	if len(q.out) == 0 {
		for len(q.in) > 0 {
			last := q.in[len(q.in)-1]
			q.in = q.in[:len(q.in)-1]
			q.out = append(q.out, last)
		}
	}
	if len(q.out) == 0 {
		return 0, false
	}
	v := q.out[len(q.out)-1]
	q.out = q.out[:len(q.out)-1]
	return v, true
}

func main() {
	fmt.Println("--- Проверка 05_queue_two_stacks (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	var q TwoStackQueue
	q.Enqueue(1)
	q.Enqueue(2)
	a, _ := q.Dequeue()
	q.Enqueue(3)
	b, _ := q.Dequeue()
	c, _ := q.Dequeue()
	got := []int{a, b, c}
	check(fmt.Sprint(got) == "[1 2 3]", fmt.Sprintf("чередование Enqueue/Dequeue: ожидалось [1 2 3], получено %v", got))

	_, ok := q.Dequeue()
	check(!ok, "Dequeue на пустой очереди должен вернуть ok=false")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
