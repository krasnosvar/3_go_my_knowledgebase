// Упражнение 08: map под мьютексом + метрики на атомиках.
//
// Задача: Cache защищает map мьютексом (map нельзя обновлять атомарно),
// а счётчики hits/misses — атомики (чтобы читать статистику, не трогая
// мьютекс). Реализуйте:
//   - Set(k, v) — запись под mu;
//   - Get(k) — чтение под mu; при попадании hits.Add(1), иначе misses.Add(1);
//   - Stats() — (hits, misses) через Load.
//
// Запуск: go run ./03_exercises/08_hybrid/main.go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Cache — map под мьютексом со счётчиками на атомиках.
type Cache struct {
	mu     sync.Mutex
	m      map[string]string
	hits   atomic.Int64
	misses atomic.Int64
}

func newCache() *Cache { return &Cache{m: map[string]string{}} }

// Set записывает значение.
func (c *Cache) Set(k, v string) {
	// TODO
}

// Get читает значение и обновляет статистику.
func (c *Cache) Get(k string) (string, bool) {
	// TODO
	return "", false
}

// Stats возвращает (hits, misses).
func (c *Cache) Stats() (int64, int64) {
	// TODO
	return 0, 0
}

func main() {
	fmt.Println("--- Проверка 08_hybrid ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	c := newCache()
	c.Set("a", "A")
	c.Set("b", "B")

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Get("a")
		}()
	}
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Get("zzz")
		}()
	}
	wg.Wait()

	hits, misses := c.Stats()
	check(hits == 100, fmt.Sprintf("hits = %d, ожидалось 100", hits))
	check(misses == 50, fmt.Sprintf("misses = %d, ожидалось 50", misses))
	v, ok := c.Get("b")
	check(ok && v == "B", "Get(b) должен вернуть \"B\"")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
