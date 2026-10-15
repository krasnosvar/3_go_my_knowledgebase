// Упражнение 04: однонаправленные каналы (Решение).
package main

import (
	"fmt"
	"time"
)

func sender(out chan<- int, vals []int) {
	defer close(out)
	for _, v := range vals {
		out <- v
	}
}

func receiver(in <-chan int) []int {
	res := []int{}
	for v := range in {
		res = append(res, v)
	}
	return res
}

func main() {
	fmt.Println("--- Проверка 04_directional (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	ch := make(chan int)
	go sender(ch, []int{1, 2, 3})

	result := make(chan []int, 1)
	go func() { result <- receiver(ch) }()

	select {
	case got := <-result:
		check(fmt.Sprint(got) == "[1 2 3]", fmt.Sprintf("receiver вернул %v, ожидалось [1 2 3]", got))
	case <-time.After(300 * time.Millisecond):
		check(false, "receiver не завершился: sender должен закрыть канал")
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
