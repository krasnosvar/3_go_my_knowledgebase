// Упражнение 04: чтение двух каналов до их закрытия.
//
// Задача: drainBoth(nums, words) читает оба канала через select, пока
// ОБА не закрыты, и возвращает сумму чисел и список слов (в порядке
// поступления из words).
//
// Приём: когда канал закрыт (v, ok := <-ch; !ok), присвойте переменной
// канала nil — чтение из nil-канала блокируется навсегда, и select
// перестанет выбирать эту ветку. Цикл идёт, пока хотя бы один канал не nil.
//
// Запуск: go run ./03_exercises/04_two_channels/main.go
package main

import (
	"fmt"
	"time"
)

// drainBoth вычитывает оба канала до закрытия.
func drainBoth(nums <-chan int, words <-chan string) (int, []string) {
	// TODO: for nums != nil || words != nil { select { ... } }
	return 0, nil
}

func main() {
	fmt.Println("--- Проверка 04_two_channels ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	nums := make(chan int)
	words := make(chan string)
	go func() {
		for i := 1; i <= 4; i++ {
			nums <- i
		}
		close(nums)
	}()
	go func() {
		for _, w := range []string{"go", "select", "nil"} {
			words <- w
		}
		close(words)
	}()

	type result struct {
		sum   int
		words []string
	}
	res := make(chan result, 1)
	go func() {
		s, w := drainBoth(nums, words)
		res <- result{s, w}
	}()

	select {
	case r := <-res:
		check(r.sum == 10, fmt.Sprintf("сумма = %d, ожидалось 10", r.sum))
		check(fmt.Sprint(r.words) == "[go select nil]", fmt.Sprintf("слова = %v, ожидалось [go select nil]", r.words))
	case <-time.After(300 * time.Millisecond):
		check(false, "drainBoth зависла: закрытый канал нужно заменять на nil")
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
