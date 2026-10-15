// Упражнение 03: коммит оффсетов при параллельной обработке.
//
// Сообщения партиции обрабатываются параллельно и завершаются не по порядку.
// Коммитить можно только «непрерывный префикс»: если 5 готов, а 4 нет,
// закоммитить 5 нельзя — при падении 4 потеряется.
//
// Задача: реализуйте tracker:
//   - Done(offset int64) — отметить оффсет обработанным;
//   - Committable() int64 — наибольший оффсет N такой, что все оффсеты
//     от start до N обработаны; если start не обработан — start-1.
//
// (В Kafka коммитят N+1 — «следующий для чтения», здесь для простоты N.)
//
// Запуск: go run ./03_exercises/03_offset_commit
package main

import "fmt"

type tracker struct {
	start int64
	done  map[int64]bool
}

func (t *tracker) Done(offset int64) {
	// TODO: отметьте оффсет.
}

func (t *tracker) Committable() int64 {
	// TODO: идите от start, пока оффсеты подряд обработаны.
	return t.start - 1
}

func main() {
	fmt.Println("--- Проверка 03_offset_commit ---")

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
