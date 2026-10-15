// Упражнение 05: идемпотентность по ключу.
//
// Очереди доставляют at-least-once, значит одна задача может прийти дважды.
// Побочный эффект (списание, письмо) должен выполниться один раз.
//
// Задача: реализуйте (p *processor) Process(key string, effect func() error) error:
//   - ключ уже успешно обработан → вернуть nil, effect НЕ вызывать;
//   - иначе вызвать effect; при успехе запомнить ключ;
//     при ошибке ключ НЕ запоминать (повтор должен выполниться);
//   - Process вызывается конкурентно: защитите состояние мьютексом и не
//     допускайте параллельного выполнения effect для одного ключа
//     (простейший вариант — держать мьютекс на время effect).
//
// Запуск: go run ./03_exercises/05_idempotent_key
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
	// TODO: проверка ключа, effect, запоминание при успехе — под мьютексом.
	return effect()
}

func main() {
	fmt.Println("--- Проверка 05_idempotent_key ---")

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
