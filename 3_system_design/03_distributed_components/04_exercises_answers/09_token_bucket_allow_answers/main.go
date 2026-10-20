// Упражнение 09: Allow() для token bucket (Решение).
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

func (b *TokenBucket) Allow() bool {
	if b.tokens >= 1 {
		b.tokens--
		return true
	}
	return false
}

var _ = NewTokenBucket

func main() {
	fmt.Println("--- Проверка 09_token_bucket_allow (Решение) ---")

	type step struct {
		op   string
		want bool
	}
	steps := []step{
		{"r", true}, {"r", true}, {"r", true},
		{"r", false},
		{"t", false},
		{"r", true}, {"r", false},
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
