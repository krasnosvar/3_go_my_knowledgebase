// Упражнение 05: очередь на двух стеках.
//
// Задача: реализуйте очередь (FIFO) через два стека (LIFO):
//   - Enqueue(v) кладёт в стек in;
//   - Dequeue() берёт из стека out; если out пуст — перекладывает ВСЕ
//     элементы из in в out (порядок разворачивается) и берёт вершину out.
//
// Амортизированно каждая операция O(1). Пустая очередь: (0, false).
//
// Запуск: go run ./03_exercises/05_queue_two_stacks/main.go
package main

import "fmt"

// TwoStackQueue — очередь на двух стеках.
type TwoStackQueue struct {
	in  []int
	out []int
}

func (q *TwoStackQueue) Enqueue(v int) {
	// TODO: положите v в q.in.
}

func (q *TwoStackQueue) Dequeue() (int, bool) {
	// TODO: при пустом out перелейте in в out, затем снимите вершину out.
	return 0, false
}

func main() {
	fmt.Println("--- Проверка 05_queue_two_stacks ---")

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
