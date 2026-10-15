// Упражнение 02: неблокирующие операции через default (Решение).
package main

import "fmt"

func trySend(ch chan<- int, v int) bool {
	select {
	case ch <- v:
		return true
	default:
		return false
	}
}

func tryRecv(ch <-chan int) (int, bool) {
	select {
	case v := <-ch:
		return v, true
	default:
		return 0, false
	}
}

func main() {
	fmt.Println("--- Проверка 02_default_non_blocking (Решение) ---")

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
