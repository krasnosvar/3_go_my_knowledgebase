// Упражнение 06: стоимость копирования большого receiver (Решение).
package main

import (
	"fmt"
)

// Big — «тяжёлая» структура.
type Big struct {
	data [1024]int
}

func (b *Big) Fill(v int) {
	for i := range b.data {
		b.data[i] = v
	}
}

func (b *Big) Sum() int {
	s := 0
	for _, x := range b.data {
		s += x
	}
	return s
}

func main() {
	fmt.Println("--- Проверка 06_copy_cost (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	b := &Big{}
	b.Fill(2)
	check(b.Sum() == 2048, fmt.Sprintf("после Fill(2): ожидалось Sum()=2048, получено %d", b.Sum()))

	var v any = Big{}
	_, hasFill := v.(interface{ Fill(int) })
	_, hasSum := v.(interface{ Sum() int })
	check(!hasFill, "Fill должен быть с pointer receiver (иначе копируется 8 КБ и изменения теряются)")
	check(!hasSum, "Sum должен быть с pointer receiver (иначе копируется 8 КБ на каждый вызов)")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
