// Упражнение 04: конфликт имён при встраивании.
//
// Задача: C встраивает A и B, у обоих есть Name(). Селектор c.Name()
// неоднозначен — поэтому у C вообще нет метода Name в method set.
//   - реализуйте C.Names(): "A,B" — вызовы через c.A.Name() и c.B.Name();
//   - объявите у C собственный Name(), возвращающий "C", — он снимает
//     неоднозначность (метод внешнего типа побеждает).
//
// Запуск: go run ./03_exercises/04_conflict/main.go
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

// Names возвращает имена обоих встроенных типов через запятую.
func (c C) Names() string {
	// TODO: c.A.Name() + "," + c.B.Name()
	return ""
}

// TODO: объявите func (c C) Name() string, возвращающий "C".

func main() {
	fmt.Println("--- Проверка 04_conflict ---")

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
