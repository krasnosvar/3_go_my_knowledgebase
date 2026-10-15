// Упражнение 01: неявная реализация интерфейса (Решение).
package main

import (
	"fmt"
	"strings"
)

// Speaker — всё, что умеет говорить.
type Speaker interface {
	Speak() string
}

// Cat — кошка.
type Cat struct{}

// Dog — собака.
type Dog struct{}

func (Cat) Speak() string { return "meow" }

func (Dog) Speak() string { return "woof" }

func chorus(ss []Speaker) string {
	parts := make([]string, 0, len(ss))
	for _, s := range ss {
		parts = append(parts, s.Speak())
	}
	return strings.Join(parts, " ")
}

func main() {
	fmt.Println("--- Проверка 01_basic_interface (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	// Проверяем через any: заготовка компилируется, даже если методов ещё нет.
	var animals []Speaker
	for _, v := range []any{Cat{}, Dog{}} {
		s, ok := v.(Speaker)
		check(ok, fmt.Sprintf("%T должен реализовывать Speaker", v))
		if ok {
			animals = append(animals, s)
		}
	}
	if len(animals) == 2 {
		check(animals[0].Speak() == "meow", fmt.Sprintf("Cat.Speak: получено %q", animals[0].Speak()))
		check(animals[1].Speak() == "woof", fmt.Sprintf("Dog.Speak: получено %q", animals[1].Speak()))
		got := chorus(animals)
		check(got == "meow woof", fmt.Sprintf("chorus: ожидалось \"meow woof\", получено %q", got))
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
