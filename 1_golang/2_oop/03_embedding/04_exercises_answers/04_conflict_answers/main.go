// Упражнение 04: конфликт имён при встраивании (Решение).
package main

import (
	"fmt"
)

// A и B имеют метод с одинаковым именем.
type A struct{}
type B struct{}

func (A) Name() string { return "A" }
func (B) Name() string { return "B" }

// C встраивает оба.
type C struct {
	A
	B
}

func (c C) Names() string { return c.A.Name() + "," + c.B.Name() }

func (c C) Name() string { return "C" }

func main() {
	fmt.Println("--- Проверка 04_conflict (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	c := C{}
	check(c.Names() == "A,B", fmt.Sprintf("Names: ожидалось \"A,B\", получено %q", c.Names()))

	var v any = c
	n, ok := v.(interface{ Name() string })
	check(ok, "у C должен быть собственный Name() (иначе селектор неоднозначен)")
	if ok {
		check(n.Name() == "C", fmt.Sprintf("C.Name: ожидалось \"C\", получено %q", n.Name()))
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
