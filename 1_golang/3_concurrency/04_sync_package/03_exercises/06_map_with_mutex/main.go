// Упражнение 06: конкурентный подсчёт слов в map.
//
// Задача: обычная map не безопасна для конкурентной записи (runtime
// падает с "concurrent map writes"). Реализуйте WordCount с мьютексом:
//   - Add(word) увеличивает счётчик слова;
//   - Get(word) возвращает текущее значение;
//   - Len() — число различных слов.
//
// Запуск: go run ./03_exercises/06_map_with_mutex/main.go
package main

import (
	"fmt"
	"sync"
)

// WordCount — потокобезопасный счётчик слов.
type WordCount struct {
	mu sync.Mutex
	m  map[string]int
}

func newWordCount() *WordCount { return &WordCount{m: map[string]int{}} }

// Add увеличивает счётчик word на 1.
func (w *WordCount) Add(word string) {
	// TODO: под мьютексом w.m[word]++
}

// Get возвращает счётчик word.
func (w *WordCount) Get(word string) int {
	// TODO
	return 0
}

// Len возвращает число различных слов.
func (w *WordCount) Len() int {
	// TODO
	return 0
}

func main() {
	fmt.Println("--- Проверка 06_map_with_mutex ---")

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
