// Пример 07: рейт-лимитер token bucket (детерминированный).
//
// Корзина ёмкостью capacity наполняется refillRate токенами за один
// логический тик. Каждый Allow() тратит токен: есть токен -> запрос
// пропущен, нет -> отказ (429).
//
// Чтобы пример был детерминированным и не флакал, время НЕ реальное:
// вместо time.Now()/time.Tick мы двигаем «часы» вручную методом Tick().
// Прогоняем фиксированный сценарий и печатаем пропуск/отказ.
//
// Запуск: go run ./02_examples/07_token_bucket/main.go
package main

import "fmt"

// TokenBucket — token bucket на логических тиках (без реального времени).
type TokenBucket struct {
	capacity   float64 // максимум токенов в корзине
	refillRate float64 // сколько токенов добавляется за один Tick()
	tokens     float64 // текущее число токенов
}

// NewTokenBucket создаёт полную корзину.
func NewTokenBucket(capacity, refillRate float64) *TokenBucket {
	return &TokenBucket{
		capacity:   capacity,
		refillRate: refillRate,
		tokens:     capacity,
	}
}

// Tick продвигает время на один логический шаг: доливаем токены,
// но не выше ёмкости.
func (b *TokenBucket) Tick() {
	b.tokens += b.refillRate
	if b.tokens > b.capacity {
		b.tokens = b.capacity
	}
}

// Allow пытается забрать один токен. true -> запрос пропущен.
func (b *TokenBucket) Allow() bool {
	if b.tokens >= 1 {
		b.tokens--
		return true
	}
	return false
}

func main() {
	fmt.Println("--- token bucket (capacity=3, refill=1/тик) ---")

	// Сценарий: список шагов. "req" — попытка запроса, "tick" — долив.
	steps := []string{
		"req", "req", "req", "req", // 4 подряд: первые 3 пройдут (burst), 4-й — отказ
		"tick",       // +1 токен
		"req", "req", // 1-й пройдёт, 2-й — отказ
		"tick", "tick", "tick", "tick", // доливаем сверх ёмкости (упрёмся в 3)
		"req", "req", "req", "req", // снова burst: 3 пройдут, 4-й — отказ
	}

	bucket := NewTokenBucket(3, 1)

	allowed, denied := 0, 0
	for i, s := range steps {
		switch s {
		case "tick":
			bucket.Tick()
			fmt.Printf("шаг %2d: TICK   -> токенов=%.0f\n", i, bucket.tokens)
		case "req":
			if bucket.Allow() {
				allowed++
				fmt.Printf("шаг %2d: ЗАПРОС -> ALLOW (осталось %.0f)\n", i, bucket.tokens)
			} else {
				denied++
				fmt.Printf("шаг %2d: ЗАПРОС -> DENY  (токенов %.0f) [429]\n", i, bucket.tokens)
			}
		}
	}

	fmt.Printf("\nИтого: пропущено=%d, отказано=%d\n", allowed, denied)
	fmt.Println("Готово, программа завершается.")
}
