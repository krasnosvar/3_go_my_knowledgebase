// Package stack — код под тестом для упражнения 03.
package stack

// Stack — стек целых чисел. Нулевое значение готово к использованию.
type Stack struct {
	items []int
}

// Push кладёт v на вершину стека.
func (s *Stack) Push(v int) { s.items = append(s.items, v) }

// Pop снимает элемент с вершины. Для пустого стека возвращает (0, false).
func (s *Stack) Pop() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	v := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return v, true
}

// Len возвращает число элементов.
func (s *Stack) Len() int { return len(s.items) }
