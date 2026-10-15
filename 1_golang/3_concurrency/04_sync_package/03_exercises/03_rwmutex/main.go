// Упражнение 03: конфиг под sync.RWMutex.
//
// Задача: Config хранит map[string]string и читается гораздо чаще, чем
// пишется. Реализуйте:
//   - Get(key) — под RLock/RUnlock (читатели не блокируют друг друга);
//   - Set(key, val) — под Lock/Unlock (писатель эксклюзивен).
//
// Запуск: go run ./03_exercises/03_rwmutex/main.go
package main

import (
	"fmt"
	"sync"
)

// Config — потокобезопасное хранилище настроек.
type Config struct {
	mu sync.RWMutex
	m  map[string]string
}

func newConfig() *Config { return &Config{m: map[string]string{}} }

// Get читает значение под RLock.
func (c *Config) Get(key string) (string, bool) {
	// TODO: c.mu.RLock(); defer c.mu.RUnlock(); ...
	return "", false
}

// Set записывает значение под Lock.
func (c *Config) Set(key, val string) {
	// TODO: c.mu.Lock(); defer c.mu.Unlock(); ...
}

func main() {
	fmt.Println("--- Проверка 03_rwmutex ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	cfg := newConfig()
	cfg.Set("mode", "dev")

	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 50; j++ {
				cfg.Get("mode")
			}
		}()
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for j := 0; j < 50; j++ {
			cfg.Set(fmt.Sprintf("k%d", j), "v")
		}
		cfg.Set("mode", "prod")
	}()
	wg.Wait()

	v, ok := cfg.Get("mode")
	check(ok && v == "prod", fmt.Sprintf("Get(mode) = (%q, %t), ожидалось (\"prod\", true)", v, ok))
	_, ok = cfg.Get("k49")
	check(ok, "ключ k49 должен быть записан")
	_, ok = cfg.Get("missing")
	check(!ok, "отсутствующий ключ не должен находиться")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
