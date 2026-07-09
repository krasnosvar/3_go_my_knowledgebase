// Пример 03: retry с экспоненциальным backoff + jitter.
//
// Оборачиваем «нестабильную» функцию, которая падает первые K раз, и
// повторяем её с растущей задержкой. Backoff: base * 2^attempt (с
// потолком maxDelay). Jitter (случайный разброс) детерминирован —
// используется локальный math/rand с ФИКСИРОВАННЫМ seed, переданным
// параметром, а НЕ глобальный источник. Поэтому запуск воспроизводим.
//
// Мы не спим по-настоящему (чтобы пример завершался мгновенно): вместо
// time.Sleep задержка вычисляется и печатается.
//
// Запуск: go run ./02_examples/03_retry_backoff/main.go
package main

import (
	"errors"
	"fmt"
	"math/rand"
)

// backoffDelay возвращает базовую экспоненциальную задержку в мс
// (без jitter): base * 2^attempt, но не больше maxDelay.
func backoffDelay(attempt, base, maxDelay int) int {
	d := base
	for i := 0; i < attempt; i++ {
		d *= 2
		if d >= maxDelay {
			return maxDelay
		}
	}
	if d > maxDelay {
		return maxDelay
	}
	return d
}

// withJitter применяет "full jitter": случайное значение в [0, delay].
// rng — детерминированный источник (фиксированный seed).
func withJitter(rng *rand.Rand, delay int) int {
	if delay <= 0 {
		return 0
	}
	return rng.Intn(delay + 1)
}

// retry повторяет fn до maxAttempts раз с backoff+jitter.
// Возвращает номер успешной попытки (1-based) или ошибку.
func retry(fn func() error, maxAttempts, base, maxDelay int, rng *rand.Rand) (int, error) {
	var lastErr error
	for attempt := 0; attempt < maxAttempts; attempt++ {
		lastErr = fn()
		if lastErr == nil {
			return attempt + 1, nil
		}
		if attempt == maxAttempts-1 {
			break // повторов больше не будет
		}
		base := backoffDelay(attempt, base, maxDelay)
		sleep := withJitter(rng, base)
		fmt.Printf("попытка %d: ошибка (%v); backoff=%dмс, sleep+jitter=%dмс\n",
			attempt+1, lastErr, base, sleep)
	}
	return maxAttempts, lastErr
}

func main() {
	fmt.Println("--- retry + exponential backoff + jitter ---")

	// Детерминированный источник случайности: seed фиксирован.
	rng := rand.New(rand.NewSource(42))

	// «Нестабильная» функция: падает первые 3 раза, потом успех.
	failFirst := 3
	calls := 0
	unstable := func() error {
		calls++
		if calls <= failFirst {
			return errors.New("временный сбой")
		}
		return nil
	}

	// base=100мс, потолок 1000мс, до 5 попыток.
	attempt, err := retry(unstable, 5, 100, 1000, rng)
	if err != nil {
		fmt.Printf("итог: провал после %d попыток: %v\n", attempt, err)
	} else {
		fmt.Printf("итог: успех с попытки %d (всего вызовов: %d)\n", attempt, calls)
	}

	fmt.Println("Готово, программа завершается.")
}
