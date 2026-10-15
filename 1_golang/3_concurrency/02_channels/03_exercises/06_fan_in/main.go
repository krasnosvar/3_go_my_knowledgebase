// Упражнение 06: fan-in: слияние двух каналов.
//
// Задача: merge(a, b) возвращает канал, в который попадают все значения из
// a и b (fan-in). Выходной канал закрывается, когда закрыты ОБА входа.
//
// Подсказка: по горутине на каждый вход + WaitGroup; отдельная горутина
// ждёт wg.Wait() и закрывает выход. Порядок значений не гарантируется.
//
// Запуск: go run ./03_exercises/06_fan_in/main.go
package main

import (
	"fmt"
	"sort"
	"time"
)

// merge сливает a и b в один канал и закрывает его после обоих входов.
func merge(a, b <-chan int) <-chan int {
	// TODO: две горутины копируют a и b в out, третья закрывает out после wg.Wait().
	return nil
}

// --- вспомогательные функции проверки (не меняйте) ---

// collect читает канал до закрытия; ok=false, если канал не закрылся за d.
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
	fmt.Println("--- Проверка 06_fan_in ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	a := make(chan int, 3)
	b := make(chan int, 2)
	a <- 1
	a <- 3
	a <- 5
	b <- 2
	b <- 4
	close(a)
	close(b)

	got, closed := collect(merge(a, b), 300*time.Millisecond)
	sort.Ints(got)
	check(closed, "выходной канал не закрылся после закрытия обоих входов")
	check(fmt.Sprint(got) == "[1 2 3 4 5]", fmt.Sprintf("merge отдал %v, ожидалось [1 2 3 4 5] (в любом порядке)", got))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
