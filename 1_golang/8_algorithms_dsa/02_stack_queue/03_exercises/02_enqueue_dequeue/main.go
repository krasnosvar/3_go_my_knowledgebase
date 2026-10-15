// Упражнение 02: очередь на срезе (Enqueue/Dequeue).
//
// Задача: реализуйте очередь (FIFO) на срезе:
//   - Enqueue(v) добавляет значение в хвост;
//   - Dequeue() забирает элемент из головы: (v, true) или (0, false), если пусто;
//   - Len() — количество элементов.
//
// Подсказка: q.items = q.items[1:] — простой вариант; помните, что такой
// сдвиг не освобождает память под массивом (см. теорию про ring buffer).
//
// Запуск: go run ./03_exercises/02_enqueue_dequeue/main.go
package main

import "fmt"

// Queue — очередь целых чисел на срезе.
type Queue struct {
	items []int
}

func (q *Queue) Enqueue(v int) {
	// TODO: добавьте v в хвост.
}

func (q *Queue) Dequeue() (int, bool) {
	// TODO: заберите первый элемент.
	return 0, false
}

func (q *Queue) Len() int {
	// TODO: верните количество элементов.
	return 0
}

func main() {
	fmt.Println("--- Проверка 02_enqueue_dequeue ---")

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
