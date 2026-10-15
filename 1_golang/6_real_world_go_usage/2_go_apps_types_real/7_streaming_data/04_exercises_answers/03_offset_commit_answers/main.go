// Упражнение 03: коммит оффсетов при параллельной обработке (Решение).
package main

import "fmt"

type tracker struct {
	start int64
	done  map[int64]bool
}

func (t *tracker) Done(offset int64) {
	t.done[offset] = true
}

func (t *tracker) Committable() int64 {
	n := t.start - 1
	for t.done[n+1] {
		n++
	}
	return n
}

func main() {
	fmt.Println("--- Проверка 03_offset_commit (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	t := &tracker{start: 100, done: map[int64]bool{}}
	check(t.Committable() == 99, "ничего не готово → start-1")
	t.Done(101)
	t.Done(102)
	check(t.Committable() == 99, "100 не готов → коммитить нечего")
	t.Done(100)
	check(t.Committable() == 102, "100..102 готовы → 102")
	t.Done(104)
	check(t.Committable() == 102, "дыра на 103 → остаёмся на 102")
	t.Done(103)
	check(t.Committable() == 104, "дыра закрыта → 104")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
