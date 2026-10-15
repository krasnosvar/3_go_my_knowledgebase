// Упражнение 01: неявная реализация интерфейса.
//
// Задача: сделайте так, чтобы Cat и Dog реализовывали интерфейс Speaker.
// В Go нет слова implements: тип реализует интерфейс, если у него есть все
// его методы.
//   - Cat.Speak() возвращает "meow";
//   - Dog.Speak() возвращает "woof";
//   - chorus(ss) склеивает ответы всех Speaker через пробел.
//
// Запуск: go run ./03_exercises/01_basic_interface/main.go
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

// TODO: добавьте метод Speak() string для Cat и для Dog.

// chorus склеивает голоса через пробел.
func chorus(ss []Speaker) string {
	// TODO: соберите s.Speak() всех элементов и склейте через strings.Join.
	_ = strings.Join
	return ""
}

func main() {
	fmt.Println("--- Проверка 01_basic_interface ---")

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
