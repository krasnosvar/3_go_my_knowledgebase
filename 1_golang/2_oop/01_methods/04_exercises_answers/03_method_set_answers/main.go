// Упражнение 03: method set значения и указателя (Решение).
package main

import (
	"fmt"
)

// Namer умеет отдавать имя.
type Namer interface {
	Name() string
}

// Setter умеет менять имя.
type Setter interface {
	SetName(n string)
}

// T — тип с именем.
type T struct {
	name string
}

func (t T) Name() string { return t.name }

func (t *T) SetName(n string) {
	t.name = n
}

func main() {
	fmt.Println("--- Проверка 03_method_set (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	var v any = T{name: "go"}
	n, ok := v.(Namer)
	check(ok, "значение T{} должно реализовывать Namer (нужен value receiver у Name)")
	if ok {
		check(n.Name() == "go", fmt.Sprintf("Name(): ожидалось \"go\", получено %q", n.Name()))
	}

	_, valueIsSetter := v.(Setter)
	check(!valueIsSetter, "значение T{} НЕ должно реализовывать Setter (SetName — pointer receiver)")

	var p any = &T{}
	s, ok := p.(Setter)
	check(ok, "*T должен реализовывать Setter")
	if ok {
		s.SetName("gopher")
		check(p.(*T).name == "gopher", fmt.Sprintf("после SetName: ожидалось \"gopher\", получено %q", p.(*T).name))
	}
	_, ptrIsNamer := p.(Namer)
	check(ptrIsNamer, "*T тоже должен реализовывать Namer (method set *T включает value-методы)")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
