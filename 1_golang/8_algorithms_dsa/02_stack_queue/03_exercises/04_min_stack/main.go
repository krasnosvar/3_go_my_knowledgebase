// Упражнение 04: стек с минимумом за O(1).
//
// Задача: реализуйте MinStack — стек, который умеет возвращать минимум за O(1):
//   - Push(v), Pop() (int, bool), Min() (int, bool);
//   - для пустого стека Pop/Min возвращают (0, false).
//
// Подсказка: храните второй стек mins, где mins[i] — минимум среди items[0..i].
//
// Запуск: go run ./03_exercises/04_min_stack/main.go
package main

import "fmt"

// MinStack — стек с поддержкой минимума за O(1).
type MinStack struct {
	items []int
	mins  []int
}

func (s *MinStack) Push(v int) {
	// TODO: положите v в items и актуальный минимум в mins.
}

func (s *MinStack) Pop() (int, bool) {
	// TODO: снимите вершину из обоих стеков.
	return 0, false
}

func (s *MinStack) Min() (int, bool) {
	// TODO: верните вершину mins.
	return 0, false
}

func main() {
	fmt.Println("--- Проверка 04_min_stack ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	var s MinStack
	_, ok := s.Min()
	check(!ok, "Min на пустом стеке должен вернуть ok=false")

	s.Push(5)
	s.Push(3)
	s.Push(7)
	s.Push(2)
	m, ok := s.Min()
	check(ok && m == 2, fmt.Sprintf("после 5,3,7,2: ожидался Min=2, получено (%d, %t)", m, ok))

	s.Pop() // убрали 2
	m, ok = s.Min()
	check(ok && m == 3, fmt.Sprintf("после Pop: ожидался Min=3, получено (%d, %t)", m, ok))

	s.Pop() // убрали 7
	s.Pop() // убрали 3
	m, ok = s.Min()
	check(ok && m == 5, fmt.Sprintf("остался 5: ожидался Min=5, получено (%d, %t)", m, ok))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
