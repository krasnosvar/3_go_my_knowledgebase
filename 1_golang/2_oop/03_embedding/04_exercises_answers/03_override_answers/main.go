// Упражнение 03: «переопределение» метода встроенного типа (Решение).
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

func (c Child) Print() string { return "child: " + c.Parent.Print() }

func main() {
	fmt.Println("--- Проверка 03_override (Решение) ---")

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
