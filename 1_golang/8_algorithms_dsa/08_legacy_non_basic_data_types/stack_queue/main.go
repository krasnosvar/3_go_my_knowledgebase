package main

import "fmt"

// Stack реализует структуру данных LIFO (Last In, First Out) на базе слайса
type Stack []int

// Push добавляет элемент на вершину стека
func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

// Pop удаляет и возвращает верхний элемент стека
func (s *Stack) Pop() (int, bool) {
	if len(*s) == 0 {
		return 0, false
	}
	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element, true
}

// Queue реализует структуру данных FIFO (First In, First Out)
type Queue []int

// Enqueue добавляет элемент в конец очереди
func (q *Queue) Enqueue(v int) {
	*q = append(*q, v)
}

// Dequeue удаляет и возвращает первый элемент очереди
func (q *Queue) Dequeue() (int, bool) {
	if len(*q) == 0 {
		return 0, false
	}
	element := (*q)[0]
	*q = (*q)[1:]
	return element, true
}

func main() {
	// Пример работы со стеком
	fmt.Println("--- STACK ---")
	stack := Stack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	for {
		v, ok := stack.Pop()
		if !ok {
			break
		}
		fmt.Printf("Popped: %d\n", v)
	}

	// Пример работы с очередью
	fmt.Println("\n--- QUEUE ---")
	queue := Queue{}
	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)

	for {
		v, ok := queue.Dequeue()
		if !ok {
			break
		}
		fmt.Printf("Dequeued: %d\n", v)
	}
}
