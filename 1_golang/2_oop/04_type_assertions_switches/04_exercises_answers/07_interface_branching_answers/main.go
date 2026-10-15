// Упражнение 07: ветвление по интерфейсам (Решение).
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

func render(v any) string {
	switch x := v.(type) {
	case error:
		return "error: " + x.Error()
	case fmt.Stringer:
		return "stringer: " + x.String()
	default:
		return "plain: " + fmt.Sprint(v)
	}
}

func main() {
	fmt.Println("--- Проверка 07_interface_branching (Решение) ---")

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
