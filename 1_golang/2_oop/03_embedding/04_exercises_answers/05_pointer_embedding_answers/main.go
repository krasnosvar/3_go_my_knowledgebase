// Упражнение 05: встраивание указателя (Решение).
package main

import (
	"fmt"
)

// Counter — счётчик.
type Counter struct {
	n int
}

func (c *Counter) Inc()       { c.n++ }
func (c *Counter) Value() int { return c.n }

// Wrap встраивает указатель на Counter.
type Wrap struct {
	*Counter
}

func NewWrap() Wrap {
	return Wrap{Counter: &Counter{}}
}

func main() {
	fmt.Println("--- Проверка 05_pointer_embedding (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	w := NewWrap()
	check(w.Counter != nil, "встроенный *Counter не инициализирован (nil) — Inc упадёт с паникой")
	if w.Counter != nil {
		w.Inc()
		w2 := w // копия Wrap копирует указатель, а не сам Counter
		w2.Inc()
		check(w.Value() == 2, fmt.Sprintf("копии делят один Counter: ожидалось 2, получено %d", w.Value()))
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
