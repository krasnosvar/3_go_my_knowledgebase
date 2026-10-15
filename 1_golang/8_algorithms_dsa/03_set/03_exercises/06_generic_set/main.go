// Упражнение 06: generic-множество Set[T comparable].
//
// Задача: реализуйте методы generic-множества Set[T comparable]:
//   - Add(v), Remove(v), Has(v) bool, Len() int.
//
// Ограничение comparable обязательно: ключ map должен поддерживать ==.
// Конструктор NewSet уже написан.
//
// Запуск: go run ./03_exercises/06_generic_set/main.go
package main

import "fmt"

// Set — generic-множество на map[T]struct{}.
type Set[T comparable] struct {
	m map[T]struct{}
}

// NewSet создаёт пустое множество.
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{m: make(map[T]struct{})}
}

func (s *Set[T]) Add(v T) {
	// TODO: добавьте v в s.m.
}

func (s *Set[T]) Remove(v T) {
	// TODO: удалите v из s.m.
}

func (s *Set[T]) Has(v T) bool {
	// TODO: comma-ok проверка.
	return false
}

func (s *Set[T]) Len() int {
	// TODO: размер множества.
	return 0
}

func main() {
	fmt.Println("--- Проверка 06_generic_set ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	ints := NewSet[int]()
	ints.Add(1)
	ints.Add(2)
	ints.Add(2)
	check(ints.Len() == 2, fmt.Sprintf("Set[int]: ожидался Len()=2, получено %d", ints.Len()))
	check(ints.Has(1) && !ints.Has(3), "Set[int]: Has(1)=true, Has(3)=false")

	ints.Remove(1)
	check(!ints.Has(1) && ints.Len() == 1, "после Remove(1) элемента 1 быть не должно")

	words := NewSet[string]()
	words.Add("go")
	check(words.Has("go") && words.Len() == 1, "Set[string]: ожидался один элемент \"go\"")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
