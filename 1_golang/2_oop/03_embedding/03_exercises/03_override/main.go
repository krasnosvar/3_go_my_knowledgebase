// Упражнение 03: «переопределение» метода встроенного типа.
//
// Задача: определите у Child собственный метод Print, который «перекрывает»
// Parent.Print и возвращает "child: " + результат Parent.Print().
// Это не наследование: Parent.Print по-прежнему доступен явно через c.Parent.
//
// Запуск: go run ./03_exercises/03_override/main.go
package main

import (
	"fmt"
)

// Parent — встраиваемый тип.
type Parent struct{}

// Print возвращает имя.
func (Parent) Print() string { return "parent" }

// Child встраивает Parent.
type Child struct {
	Parent
}

// TODO: объявите func (c Child) Print() string, вызывающий c.Parent.Print().

func main() {
	fmt.Println("--- Проверка 03_override ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	c := Child{}
	check(c.Print() == "child: parent", fmt.Sprintf("Child.Print: ожидалось \"child: parent\", получено %q", c.Print()))
	check(c.Parent.Print() == "parent", "Parent.Print должен оставаться доступен через c.Parent")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
