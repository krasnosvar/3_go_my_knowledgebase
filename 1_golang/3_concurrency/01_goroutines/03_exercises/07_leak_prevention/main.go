// Упражнение 07: горутина без утечки.
//
// Задача: generator(stop) возвращает канал, в который фоновая горутина
// пишет 0, 1, 2, ... бесконечно. Без пути остановки такая горутина
// навсегда зависнет на отправке, когда читатель уйдёт, — это утечка.
//
// Сделайте так, чтобы горутина завершалась, когда закрыт stop, и при
// выходе закрывала выходной канал. Отправку оформите через select:
//
//	select {
//	case out <- i:
//	case <-stop:
//	    return
//	}
//
// Запуск: go run ./03_exercises/07_leak_prevention/main.go
package main

import (
	"fmt"
	"time"
)

// generator пишет 0,1,2,... в канал, пока не закрыт stop.
func generator(stop <-chan struct{}) <-chan int {
	// TODO: создайте канал, запустите горутину с select по out<-i и <-stop,
	// закройте out при выходе (defer close(out)).
	return nil
}

func main() {
	fmt.Println("--- Проверка 07_leak_prevention ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	stop := make(chan struct{})
	ch := generator(stop)

	first := time.After(300 * time.Millisecond)
	for want := 0; want < 3; want++ {
		select {
		case v := <-ch:
			check(v == want, fmt.Sprintf("получено %d, ожидалось %d", v, want))
		case <-first:
			check(false, fmt.Sprintf("значение %d не пришло", want))
			want = 3
		}
	}

	close(stop)
	closed := false
	deadline := time.After(300 * time.Millisecond)
	for !closed {
		select {
		case _, ok := <-ch:
			if !ok {
				closed = true
			}
		case <-deadline:
			check(false, "после закрытия stop выходной канал не закрылся — горутина утекла")
			closed = true
		}
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
