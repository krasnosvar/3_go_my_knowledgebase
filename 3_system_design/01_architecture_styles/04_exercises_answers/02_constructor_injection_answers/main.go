// Упражнение 02: внедрить зависимость через конструктор (Решение).
package main

import "fmt"

type Clock interface {
	Now() string
}

type FixedClock struct {
	value string
}

func (c FixedClock) Now() string { return c.value }

type Service struct {
	clock Clock
}

func NewService(clock Clock) *Service {
	return &Service{clock: clock}
}

func (s *Service) Timestamp() string {
	return "ts=" + s.clock.Now()
}

func main() {
	fmt.Println("--- Проверка 02_constructor_injection (Решение) ---")

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
