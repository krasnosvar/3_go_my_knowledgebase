// Упражнение 06: конкурентный подсчёт слов в map (Решение).
package main

import (
	"fmt"
	"sync"
)

type WordCount struct {
	mu sync.Mutex
	m  map[string]int
}

func newWordCount() *WordCount { return &WordCount{m: map[string]int{}} }

func (w *WordCount) Add(word string) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.m[word]++
}

func (w *WordCount) Get(word string) int {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.m[word]
}

func (w *WordCount) Len() int {
	w.mu.Lock()
	defer w.mu.Unlock()
	return len(w.m)
}

func main() {
	fmt.Println("--- Проверка 06_map_with_mutex (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	wc := newWordCount()
	words := []string{"go", "chan", "go", "mutex", "go", "chan"}
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for _, w := range words {
				wc.Add(w)
			}
		}()
	}
	wg.Wait()
	check(wc.Get("go") == 150, fmt.Sprintf("go = %d, ожидалось 150", wc.Get("go")))
	check(wc.Get("chan") == 100, fmt.Sprintf("chan = %d, ожидалось 100", wc.Get("chan")))
	check(wc.Get("mutex") == 50, fmt.Sprintf("mutex = %d, ожидалось 50", wc.Get("mutex")))
	check(wc.Len() == 3, fmt.Sprintf("Len = %d, ожидалось 3", wc.Len()))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
