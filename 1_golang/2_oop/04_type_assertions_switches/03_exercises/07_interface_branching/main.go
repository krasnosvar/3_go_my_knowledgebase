// Упражнение 07: ветвление по интерфейсам.
//
// Задача: реализуйте render(v any) string, проверяя, какие ИНТЕРФЕЙСЫ
// реализует значение (assertion к интерфейсу, а не к конкретному типу):
//   - error         -> "error: " + Error()
//   - fmt.Stringer  -> "stringer: " + String()
//   - иначе         -> "plain: " + fmt.Sprint(v)
//
// error проверяется первым (порядок case важен, если тип реализует оба).
//
// Запуск: go run ./03_exercises/07_interface_branching/main.go
package main

import (
	"errors"
	"fmt"
)

// Point реализует fmt.Stringer.
type Point struct{ X, Y int }

func (p Point) String() string { return fmt.Sprintf("(%d,%d)", p.X, p.Y) }

// BadPoint реализует и error, и fmt.Stringer.
type BadPoint struct{}

func (BadPoint) Error() string  { return "bad point" }
func (BadPoint) String() string { return "bp" }

// render описывает значение по реализуемым интерфейсам.
func render(v any) string {
	// TODO: switch x := v.(type) { case error: ... case fmt.Stringer: ... default: ... }
	return ""
}

func main() {
	fmt.Println("--- Проверка 07_interface_branching ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	cases := []struct {
		in   any
		want string
	}{
		{Point{1, 2}, "stringer: (1,2)"},
		{errors.New("boom"), "error: boom"},
		{BadPoint{}, "error: bad point"},
		{42, "plain: 42"},
	}
	for _, c := range cases {
		got := render(c.in)
		check(got == c.want, fmt.Sprintf("render(%T): ожидалось %q, получено %q", c.in, c.want, got))
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
