// Упражнение 02: методы с pointer receiver.
//
// Задача: реализуйте методы Counter, меняющие состояние, с pointer receiver:
//   - Inc() — увеличивает n на 1;
//   - Add(d) — увеличивает n на d;
//   - Reset() — обнуляет n.
//
// С value receiver изменения применились бы к копии и потерялись бы.
//
// Запуск: go run ./03_exercises/02_pointer_receiver/main.go
package main

import (
	"fmt"
)

// Counter — счётчик.
type Counter struct {
	n int
}

// Value возвращает текущее значение (только чтение — value receiver).
func (c Counter) Value() int { return c.n }

// Inc увеличивает счётчик на 1.
func (c *Counter) Inc() {
	// TODO: увеличьте c.n на 1.
}

// Add увеличивает счётчик на d.
func (c *Counter) Add(d int) {
	// TODO: увеличьте c.n на d.
}

// Reset обнуляет счётчик.
func (c *Counter) Reset() {
	// TODO: обнулите c.n.
}

func main() {
	fmt.Println("--- Проверка 02_pointer_receiver ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	var c Counter // адресуемая переменная: c.Inc() == (&c).Inc()
	c.Inc()
	c.Inc()
	c.Inc()
	check(c.Value() == 3, fmt.Sprintf("после 3x Inc: ожидалось 3, получено %d", c.Value()))

	c.Add(10)
	check(c.Value() == 13, fmt.Sprintf("после Add(10): ожидалось 13, получено %d", c.Value()))

	p := &c
	p.Reset()
	check(c.Value() == 0, fmt.Sprintf("после Reset через указатель: ожидалось 0, получено %d", c.Value()))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
