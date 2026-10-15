// Упражнение 01: стек на срезе (Push/Pop/Peek).
//
// Задача: реализуйте методы стека (LIFO) на срезе:
//   - Push(v) кладёт значение на вершину (в конец среза);
//   - Pop() снимает и возвращает вершину: (v, true); для пустого стека (0, false);
//   - Peek() возвращает вершину без снятия: (v, true) или (0, false);
//   - Len() — количество элементов.
//
// Запуск: go run ./03_exercises/01_push_pop/main.go
package main

import "fmt"

// Stack — стек целых чисел на срезе.
type Stack struct {
	items []int
}

func (s *Stack) Push(v int) {
	// TODO: добавьте v в конец s.items.
}

func (s *Stack) Pop() (int, bool) {
	// TODO: верните последний элемент и укоротите срез.
	return 0, false
}

func (s *Stack) Peek() (int, bool) {
	// TODO: верните последний элемент без удаления.
	return 0, false
}

func (s *Stack) Len() int {
	// TODO: верните количество элементов.
	return 0
}

func main() {
	fmt.Println("--- Проверка 01_push_pop ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	var s Stack
	_, ok := s.Pop()
	check(!ok, "Pop на пустом стеке должен вернуть ok=false")

	s.Push(1)
	s.Push(2)
	s.Push(3)
	check(s.Len() == 3, fmt.Sprintf("после трёх Push ожидался Len()=3, получено %d", s.Len()))

	v, ok := s.Peek()
	check(ok && v == 3, fmt.Sprintf("Peek: ожидалось (3, true), получено (%d, %t)", v, ok))

	var popped []int
	for s.Len() > 0 && len(popped) < 10 {
		v, _ := s.Pop()
		popped = append(popped, v)
	}
	check(fmt.Sprint(popped) == "[3 2 1]", fmt.Sprintf("порядок Pop: ожидалось [3 2 1], получено %v", popped))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
