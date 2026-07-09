// Упражнение 02: внедрить зависимость через конструктор.
//
// Задача: реализуйте конструктор NewService так, чтобы зависимость Clock
// (источник времени) приходила извне (constructor injection), а метод
// Timestamp возвращал строку, полученную от этого Clock.
//
// Смысл: заменяемость зависимости. В тесте подставляется детерминированный
// FixedClock, поэтому результат предсказуем и не зависит от реального времени.
//
// Запуск: go run ./03_exercises/02_constructor_injection/main.go
package main

import "fmt"

// Clock — порт: источник времени (не менять).
type Clock interface {
	Now() string
}

// FixedClock — детерминированная реализация для проверки (не менять).
type FixedClock struct {
	value string
}

func (c FixedClock) Now() string { return c.value }

// Service зависит от Clock.
type Service struct {
	clock Clock
}

// NewService — конструктор с внедрением зависимости.
func NewService(clock Clock) *Service {
	// TODO: сохраните переданный clock в поле структуры.
	return &Service{}
}

// Timestamp возвращает "ts=" + текущее время от clock.
func (s *Service) Timestamp() string {
	// TODO: верните "ts=" + s.clock.Now().
	return ""
}

func main() {
	fmt.Println("--- Проверка 02_constructor_injection ---")

	svc := NewService(FixedClock{value: "2026-01-01T00:00:00Z"})

	allOK := true
	got := svc.Timestamp()
	want := "ts=2026-01-01T00:00:00Z"
	if got != want {
		allOK = false
		fmt.Printf("FAIL: Timestamp() = %q, ожидалось %q\n", got, want)
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
