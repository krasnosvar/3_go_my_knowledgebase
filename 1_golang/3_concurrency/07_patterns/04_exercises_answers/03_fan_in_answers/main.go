// Упражнение 03: fan-in с остановкой по done (Решение).
package main

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

func fanIn(done <-chan struct{}, chans ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	for _, in := range chans {
		wg.Add(1)
		go func(in <-chan int) {
			defer wg.Done()
			for {
				select {
				case v, ok := <-in:
					if !ok {
						return
					}
					select {
					case out <- v:
					case <-done:
						return
					}
				case <-done:
					return
				}
			}
		}(in)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
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
	fmt.Println("--- Проверка 03_fan_in (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	mk := func(vals ...int) <-chan int {
		ch := make(chan int, len(vals))
		for _, v := range vals {
			ch <- v
		}
		close(ch)
		return ch
	}

	got, closed := collect(fanIn(make(chan struct{}), mk(1, 2), mk(3), mk(4, 5)), 300*time.Millisecond)
	sort.Ints(got)
	check(closed, "выход не закрылся после закрытия всех входов")
	check(fmt.Sprint(got) == "[1 2 3 4 5]", fmt.Sprintf("получено %v, ожидалось [1 2 3 4 5]", got))

	done := make(chan struct{})
	neverClosed := make(chan int)
	out := fanIn(done, neverClosed)
	close(done)
	_, closed = collect(out, 300*time.Millisecond)
	check(closed, "после close(done) выход должен закрыться, даже если входы открыты")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
