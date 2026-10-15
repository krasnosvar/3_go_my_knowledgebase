// Упражнение 08: конструктор со встроенной структурой.
//
// Задача: реализуйте NewChild(id, name). Встроенное поле инициализируется
// по имени типа: Child{Base: Base{ID: id}, Name: name}. После этого поле ID
// и метод Describe() у Child доступны напрямую.
//
// Запуск: go run ./03_exercises/08_constructor/main.go
package main

import (
	"fmt"
)

// Base — общие поля сущностей.
type Base struct {
	ID int
}

// Describe возвращает "#<ID>".
func (b Base) Describe() string { return fmt.Sprintf("#%d", b.ID) }

// Child — сущность со встроенным Base.
type Child struct {
	Base
	Name string
}

// NewChild создаёт Child с заполненным Base.
func NewChild(id int, name string) Child {
	// TODO: заполните и Base, и Name.
	return Child{}
}

func main() {
	fmt.Println("--- Проверка 08_constructor ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	c := NewChild(7, "Den")
	check(c.ID == 7, fmt.Sprintf("ID (продвинутое поле): ожидалось 7, получено %d", c.ID))
	check(c.Name == "Den", fmt.Sprintf("Name: ожидалось \"Den\", получено %q", c.Name))
	check(c.Describe() == "#7", fmt.Sprintf("Describe: ожидалось \"#7\", получено %q", c.Describe()))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
