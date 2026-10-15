// Упражнение 06: маленькие интерфейсы (Решение).
package main

import (
	"fmt"
)

// Reader отдаёт строки по одной.
type Reader interface {
	Read() (string, bool)
}

// Writer принимает строки.
type Writer interface {
	Write(s string)
}

// sliceReader читает из слайса.
type sliceReader struct {
	items []string
	pos   int
}

func (r *sliceReader) Read() (string, bool) {
	if r.pos >= len(r.items) {
		return "", false
	}
	s := r.items[r.pos]
	r.pos++
	return s, true
}

// collectWriter запоминает строки с префиксом "> ".
type collectWriter struct {
	got []string
}

func (w *collectWriter) Write(s string) { w.got = append(w.got, "> "+s) }

func copyAll(r Reader, w Writer) int {
	n := 0
	for {
		s, ok := r.Read()
		if !ok {
			return n
		}
		w.Write(s)
		n++
	}
}

func main() {
	fmt.Println("--- Проверка 06_small_interfaces (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	r := &sliceReader{items: []string{"a", "b", "c"}}
	w := &collectWriter{}
	n := copyAll(r, w)
	check(n == 3, fmt.Sprintf("количество: ожидалось 3, получено %d", n))
	check(fmt.Sprint(w.got) == "[> a > b > c]", fmt.Sprintf("записано: получено %v", w.got))

	n = copyAll(&sliceReader{}, &collectWriter{})
	check(n == 0, fmt.Sprintf("пустой источник: ожидалось 0, получено %d", n))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
