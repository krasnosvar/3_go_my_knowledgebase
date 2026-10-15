// Упражнение 02: переходы состояний circuit breaker (Решение).
package main

import "fmt"

type State int

const (
	Closed State = iota
	Open
	HalfOpen
)

func (s State) String() string {
	return [...]string{"CLOSED", "OPEN", "HALF-OPEN"}[s]
}

type Breaker struct {
	threshold int
	state     State
	failures  int
}

func NewBreaker(threshold int) *Breaker {
	return &Breaker{threshold: threshold, state: Closed}
}

func (b *Breaker) OnResult(success bool) {
	switch b.state {
	case Closed:
		if success {
			b.failures = 0
			return
		}
		b.failures++
		if b.failures >= b.threshold {
			b.state = Open
		}
	case HalfOpen:
		if success {
			b.state = Closed
			b.failures = 0
		} else {
			b.state = Open
		}
	case Open:
		// В Open вызовы отклоняются, OnResult не вызывается.
	}
}

var _ = NewBreaker

func main() {
	fmt.Println("--- Проверка 02_circuit_breaker_transitions (Решение) ---")

	type step struct {
		setState int
		success  bool
		want     State
	}
	steps := []step{
		{-1, false, Closed},
		{-1, false, Closed},
		{-1, false, Open},
		{int(HalfOpen), false, Open},
		{int(HalfOpen), true, Closed},
		{-1, false, Closed},
		{-1, true, Closed},
	}

	b := NewBreaker(3)
	allOK := true
	for i, s := range steps {
		if s.setState >= 0 {
			b.state = State(s.setState)
			if State(s.setState) == HalfOpen {
				b.failures = 0
			}
		}
		b.OnResult(s.success)
		if b.state != s.want {
			allOK = false
			fmt.Printf("FAIL: шаг %d состояние = %s, ожидалось %s\n", i, b.state, s.want)
		}
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
