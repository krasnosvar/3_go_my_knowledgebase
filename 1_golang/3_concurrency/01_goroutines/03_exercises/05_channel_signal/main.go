// Упражнение 05: сигнал о завершении через канал.
//
// Задача: startWorker(work) запускает work в отдельной горутине и сразу
// возвращает канал done, который ЗАКРЫВАЕТСЯ, когда work завершилась.
// Вызывающий код ждёт завершения через <-done.
//
// Почему close, а не отправка значения: закрытие «видят» все читатели
// канала, а отправку — только один.
//
// Запуск: go run ./03_exercises/05_channel_signal/main.go
package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

// startWorker запускает work в горутине и возвращает канал завершения.
func startWorker(work func()) <-chan struct{} {
	// TODO: создайте канал, запустите горутину, которая выполнит work
	// и закроет канал (defer close(done)), верните канал.
	return nil
}

func main() {
	fmt.Println("--- Проверка 05_channel_signal ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	var ran atomic.Bool
	done := startWorker(func() {
		time.Sleep(10 * time.Millisecond)
		ran.Store(true)
	})

	select {
	case <-done:
		check(ran.Load(), "канал закрыт до того, как work завершилась")
	case <-time.After(300 * time.Millisecond):
		check(false, "канал done так и не закрылся")
	}

	// закрытый канал должен читаться повторно без блокировки
	if done != nil {
		select {
		case <-done:
		default:
			check(false, "done нужно закрывать, а не отправлять в него значение")
		}
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
