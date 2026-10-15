// Упражнение 02: fan-out: N выходных каналов из одного входа.
//
// Задача: fanOut(in, n, f) запускает n worker'ов, которые конкурентно
// читают ОДИН входной канал in, применяют f и пишут каждый в СВОЙ выходной
// канал. Каждый worker закрывает свой выход, когда in закрыт.
// Вернуть срез из n выходных каналов.
//
// (Слияние этих каналов обратно в один — fan-in — следующая задача.)
//
// Запуск: go run ./03_exercises/02_fan_out/main.go
package main

import (
	"fmt"
	"time"
)

// fanOut раздаёт значения из in n worker'ам, у каждого свой выход.
func fanOut(in <-chan int, n int, f func(int) int) []<-chan int {
	// TODO: n каналов, по горутине на каждый, close(out) по окончании in.
	return nil
}

// --- вспомогательные функции проверки (не меняйте) ---

// collect читает канал до закрытия; ok=false, если не закрылся за d.
func collect(ch <-chan int, d time.Duration) ([]int, bool) {
	res := []int{}
	deadline := time.After(d)
	for {
		select {
		case v, open := <-ch:
			if !open {
				return res, true
			}
			res = append(res, v)
		case <-deadline:
			return res, false
		}
	}
}

func main() {
	fmt.Println("--- Проверка 02_fan_out ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	in := make(chan int, 10)
	for i := 1; i <= 10; i++ {
		in <- i
	}
	close(in)

	outs := fanOut(in, 3, func(x int) int { return x * 2 })
	check(len(outs) == 3, fmt.Sprintf("получено %d выходных каналов, ожидалось 3", len(outs)))

	sum, results := 0, make(chan []int, len(outs))
	for _, o := range outs {
		go func(o <-chan int) {
			vals, _ := collect(o, 300*time.Millisecond)
			results <- vals
		}(o)
	}
	for range outs {
		for _, v := range <-results {
			sum += v
		}
	}
	check(sum == 110, fmt.Sprintf("сумма по всем выходам = %d, ожидалось 110", sum))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
