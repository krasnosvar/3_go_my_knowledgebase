// Упражнение 04: конвейер parse → square → sum.
//
// Задача: соберите конвейер из трёх стадий:
//   - parse(in <-chan string) <-chan int — strconv.Atoi, некорректные
//     строки пропускаются;
//   - square(in <-chan int) <-chan int — квадрат каждого числа;
//   - sum(in <-chan int) int — сумма всех значений (финальная стадия,
//     выполняется в вызывающей горутине).
//
// Промежуточные стадии — горутины, закрывающие свой выход.
//
// Запуск: go run ./03_exercises/04_pipeline/main.go
package main

import (
	"fmt"
	"time"
)

// parse превращает строки в числа, пропуская некорректные.
func parse(in <-chan string) <-chan int {
	// TODO
	return nil
}

// square возводит значения в квадрат.
func square(in <-chan int) <-chan int {
	// TODO
	return nil
}

// sum суммирует значения до закрытия канала.
func sum(in <-chan int) int {
	// TODO
	return 0
}

func main() {
	fmt.Println("--- Проверка 04_pipeline ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	src := make(chan string, 5)
	for _, s := range []string{"1", "x", "2", "", "3"} {
		src <- s
	}
	close(src)

	res := make(chan int, 1)
	go func() {
		sq := square(parse(src))
		if sq == nil {
			res <- -1
			return
		}
		res <- sum(sq)
	}()
	select {
	case got := <-res:
		check(got == 14, fmt.Sprintf("сумма квадратов = %d, ожидалось 14 (1+4+9)", got))
	case <-time.After(300 * time.Millisecond):
		check(false, "конвейер завис: стадии должны закрывать свои выходы")
	}

	ready := make(chan int, 2)
	ready <- 5
	ready <- 6
	close(ready)
	check(sum(ready) == 11, "sum([5 6]) должна вернуть 11")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
