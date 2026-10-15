// Упражнение 05: идемпотентность по ключу (Решение).
package main

import (
	"errors"
	"fmt"
	"sync"
)

type processor struct {
	mu   sync.Mutex
	done map[string]bool
}

func (p *processor) Process(key string, effect func() error) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.done[key] {
		return nil
	}
	if err := effect(); err != nil {
		return err
	}
	p.done[key] = true
	return nil
}

func main() {
	fmt.Println("--- Проверка 05_idempotent_key (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	p := &processor{done: map[string]bool{}}
	var mu sync.Mutex
	charges := 0
	charge := func() error { mu.Lock(); charges++; mu.Unlock(); return nil }

	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() { defer wg.Done(); _ = p.Process("order-1", charge) }()
	}
	wg.Wait()
	check(charges == 1, fmt.Sprintf("20 доставок одного ключа → 1 списание (было %d)", charges))

	fails := 0
	err := p.Process("order-2", func() error { fails++; return errors.New("gateway down") })
	check(err != nil, "ошибка effect возвращается")
	err = p.Process("order-2", charge)
	check(err == nil && charges == 2, "после ошибки повтор выполняется")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
