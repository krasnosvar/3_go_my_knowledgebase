// Упражнение 02: методы с pointer receiver (Решение).
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

func (c *Counter) Inc() {
	c.n++
}

func (c *Counter) Add(d int) {
	c.n += d
}

func (c *Counter) Reset() {
	c.n = 0
}

func main() {
	fmt.Println("--- Проверка 02_pointer_receiver (Решение) ---")

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
