// Упражнение 06: маленькие интерфейсы.
//
// Задача: реализуйте copyAll, которая читает из Reader строки, пока Read не
// вернёт ok=false, и пишет каждую во Writer. Вернуть количество
// скопированных строк. Функция зависит от двух маленьких интерфейсов, а не
// от конкретных типов — поэтому работает с любыми источниками/приёмниками.
//
// Запуск: go run ./03_exercises/06_small_interfaces/main.go
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

// copyAll копирует все строки из r в w.
func copyAll(r Reader, w Writer) int {
	// TODO: цикл: s, ok := r.Read(); если !ok — выход; иначе w.Write(s).
	return 0
}

func main() {
	fmt.Println("--- Проверка 06_small_interfaces ---")

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
