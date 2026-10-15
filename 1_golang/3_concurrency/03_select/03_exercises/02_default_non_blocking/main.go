// Упражнение 02: неблокирующие операции через default.
//
// Задача:
//   - trySend(ch, v) пытается отправить без блокировки: true, если удалось,
//     false, если буфер полон;
//   - tryRecv(ch) пытается прочитать без блокировки: (v, true) или (0, false).
//
// Используйте select с веткой default — она срабатывает, когда ни один
// case не готов немедленно.
//
// Запуск: go run ./03_exercises/02_default_non_blocking/main.go
package main

import "fmt"

// trySend отправляет v, если это возможно без ожидания.
func trySend(ch chan<- int, v int) bool {
	// TODO: select { case ch <- v: ...; default: ... }
	return false
}

// tryRecv читает значение, если оно уже есть.
func tryRecv(ch <-chan int) (int, bool) {
	// TODO: select { case v := <-ch: ...; default: ... }
	return 0, false
}

func main() {
	fmt.Println("--- Проверка 02_default_non_blocking ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	ch := make(chan int, 1)
	check(trySend(ch, 1), "первая отправка в пустой буфер должна удаться")
	check(!trySend(ch, 2), "вторая отправка в полный буфер должна вернуть false")

	v, ok := tryRecv(ch)
	check(ok && v == 1, fmt.Sprintf("tryRecv = (%d, %t), ожидалось (1, true)", v, ok))

	v, ok = tryRecv(ch)
	check(!ok && v == 0, "чтение из пустого канала должно вернуть (0, false)")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
