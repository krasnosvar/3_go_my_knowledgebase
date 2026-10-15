// Упражнение 06: мультиплексирование N источников (Решение).
package main

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

func multiplex(srcs ...<-chan string) <-chan string {
	out := make(chan string)
	var wg sync.WaitGroup
	for _, src := range srcs {
		wg.Add(1)
		go func(src <-chan string) {
			defer wg.Done()
			for v := range src {
				out <- v
			}
		}(src)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// --- вспомогательные функции проверки (не меняйте) ---

// collectStrings читает канал до закрытия; ok=false, если не закрылся за d.
func collectStrings(ch <-chan string, d time.Duration) ([]string, bool) {
	res := []string{}
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
	fmt.Println("--- Проверка 06_multiplex (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	mk := func(vals ...string) <-chan string {
		ch := make(chan string, len(vals))
		for _, v := range vals {
			ch <- v
		}
		close(ch)
		return ch
	}

	got, closed := collectStrings(multiplex(mk("a1", "a2"), mk("b1"), mk("c1", "c2", "c3")), 300*time.Millisecond)
	sort.Strings(got)
	check(closed, "выходной канал не закрылся")
	check(fmt.Sprint(got) == "[a1 a2 b1 c1 c2 c3]", fmt.Sprintf("получено %v, ожидалось все 6 значений", got))

	got, closed = collectStrings(multiplex(), 300*time.Millisecond)
	check(closed && len(got) == 0, "без источников выход должен сразу закрыться")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
