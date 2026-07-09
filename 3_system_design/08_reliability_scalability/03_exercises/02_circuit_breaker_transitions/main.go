// Упражнение 02: переходы состояний circuit breaker.
//
// Задача: реализуйте метод OnResult(success bool) у предохранителя.
// Он вызывается после каждого вызова зависимости и обновляет состояние:
//
//	CLOSED:
//	  success -> сбросить счётчик ошибок (failures = 0);
//	  ошибка  -> failures++; если failures >= threshold -> OPEN.
//	HALF-OPEN:
//	  success -> перейти в CLOSED (и сбросить failures);
//	  ошибка  -> перейти в OPEN.
//	OPEN:
//	  OnResult в этом состоянии не вызывается (вызовы отклоняются),
//	  поэтому обрабатывать его тут не нужно.
//
// Переход CLOSED->OPEN и OPEN->HALF-OPEN по времени здесь не требуется —
// только реакция на результат вызова.
//
// Запуск: go run ./03_exercises/02_circuit_breaker_transitions/main.go
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

// OnResult обновляет состояние по результату вызова.
func (b *Breaker) OnResult(success bool) {
	// TODO: реализуйте переходы для CLOSED и HALF-OPEN (см. комментарий выше).
}

var _ = NewBreaker

func main() {
	fmt.Println("--- Проверка 02_circuit_breaker_transitions ---")

	type step struct {
		// setState: если >=0, принудительно ставим состояние перед вызовом
		// (эмулируем переход по таймауту в HALF-OPEN).
		setState int
		success  bool
		want     State
	}
	steps := []step{
		{-1, false, Closed},           // 1 ошибка, порог 3
		{-1, false, Closed},           // 2 ошибки
		{-1, false, Open},             // 3 ошибки -> OPEN
		{int(HalfOpen), false, Open},  // half-open, провал пробы -> OPEN
		{int(HalfOpen), true, Closed}, // half-open, успех пробы -> CLOSED
		{-1, false, Closed},           // снова копим в closed
		{-1, true, Closed},            // успех сбрасывает счётчик
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
