// Упражнение 05: встраивание указателя.
//
// Задача: Wrap встраивает *Counter. Нулевое значение Wrap{} содержит nil
// указатель, и вызов Inc упадёт с паникой. Реализуйте NewWrap, который
// инициализирует встроенный указатель. Заметьте: копии Wrap разделяют один
// и тот же Counter.
//
// Запуск: go run ./03_exercises/05_pointer_embedding/main.go
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

// NewWrap создаёт Wrap с инициализированным счётчиком.
func NewWrap() Wrap {
	// TODO: верните Wrap{Counter: &Counter{}}.
	return Wrap{}
}

func main() {
	fmt.Println("--- Проверка 05_pointer_embedding ---")

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
