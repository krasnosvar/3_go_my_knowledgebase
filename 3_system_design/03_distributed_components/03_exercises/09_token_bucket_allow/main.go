// Упражнение 09: Allow() для token bucket.
//
// Задача: реализуйте метод Allow() у token bucket. Корзина хранит
// текущее число токенов (tokens) и ёмкость (capacity). Метод Tick()
// уже реализован — он доливает refill токенов, не превышая capacity.
//
// Allow должен:
//   - если токенов >= 1: списать один токен и вернуть true (запрос пропущен);
//   - иначе: вернуть false (отказ), число токенов не менять.
//
// Запуск: go run ./03_exercises/09_token_bucket_allow/main.go
package main

import "fmt"

type TokenBucket struct {
	capacity float64
	refill   float64
	tokens   float64
}

func NewTokenBucket(capacity, refill float64) *TokenBucket {
	return &TokenBucket{capacity: capacity, refill: refill, tokens: capacity}
}

func (b *TokenBucket) Tick() {
	b.tokens += b.refill
	if b.tokens > b.capacity {
		b.tokens = b.capacity
	}
}

// Allow пытается забрать один токен.
func (b *TokenBucket) Allow() bool {
	// TODO: если токенов >= 1 — списать один и вернуть true; иначе false.
	return false
}

// Держим тип живым на случай пустой заготовки.
var _ = NewTokenBucket

func main() {
	fmt.Println("--- Проверка 09_token_bucket_allow ---")

	// Сценарий: "r" — запрос (ждём want=true/false), "t" — tick.
	type step struct {
		op   string
		want bool // для "r"
	}
	steps := []step{
		{"r", true}, {"r", true}, {"r", true}, // burst из 3 (capacity=3)
		{"r", false},              // токенов нет
		{"t", false},              // +1
		{"r", true}, {"r", false}, // 1 пройдёт, 2-й нет
	}

	b := NewTokenBucket(3, 1)
	allOK := true
	for i, s := range steps {
		switch s.op {
		case "t":
			b.Tick()
		case "r":
			got := b.Allow()
			if got != s.want {
				allOK = false
				fmt.Printf("FAIL: шаг %d Allow() = %v, ожидалось %v\n", i, got, s.want)
			}
		}
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
