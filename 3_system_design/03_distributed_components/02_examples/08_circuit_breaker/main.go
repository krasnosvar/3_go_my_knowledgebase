// Пример 08: circuit breaker с состояниями closed/open/half-open.
//
// Логика:
//   - closed: пропускаем вызовы; failThreshold ошибок ПОДРЯД -> open.
//   - open: вызовы не идут (fail fast); после cooldown логических тиков
//     -> half-open.
//   - half-open: пропускаем пробные вызовы; успех -> closed (счётчики
//     сброшены), провал -> снова open.
//
// Время детерминированное: cooldown отсчитывается методом Tick(), а не
// реальными часами. Печатаем все переходы состояний.
//
// Запуск: go run ./02_examples/08_circuit_breaker/main.go
package main

import (
	"errors"
	"fmt"
)

type State int

const (
	Closed State = iota
	Open
	HalfOpen
)

func (s State) String() string {
	switch s {
	case Closed:
		return "CLOSED"
	case Open:
		return "OPEN"
	case HalfOpen:
		return "HALF-OPEN"
	default:
		return "?"
	}
}

var errOpen = errors.New("circuit breaker открыт: вызов отклонён (fail fast)")

// CircuitBreaker — предохранитель на логических тиках.
type CircuitBreaker struct {
	failThreshold int // ошибок подряд для перехода closed -> open
	cooldown      int // тиков в open перед переходом в half-open

	state     State
	failures  int // ошибок подряд в closed
	sinceOpen int // тиков с момента открытия
}

func NewCircuitBreaker(failThreshold, cooldown int) *CircuitBreaker {
	return &CircuitBreaker{
		failThreshold: failThreshold,
		cooldown:      cooldown,
		state:         Closed,
	}
}

// Tick продвигает время: если мы в Open и cooldown истёк — переходим в
// HalfOpen (пробуем пропустить разведочный вызов).
func (cb *CircuitBreaker) Tick() {
	if cb.state == Open {
		cb.sinceOpen++
		if cb.sinceOpen >= cb.cooldown {
			cb.transition(HalfOpen)
		}
	}
}

func (cb *CircuitBreaker) transition(to State) {
	fmt.Printf("        [переход] %s -> %s\n", cb.state, to)
	cb.state = to
	switch to {
	case Open:
		cb.sinceOpen = 0
	case Closed:
		cb.failures = 0
	}
}

// Call выполняет fn через предохранитель.
func (cb *CircuitBreaker) Call(fn func() error) error {
	if cb.state == Open {
		return errOpen // fail fast, зависимость не трогаем
	}

	err := fn()
	if err != nil {
		// В closed копим ошибки; в half-open один провал сразу открывает.
		if cb.state == HalfOpen {
			cb.transition(Open)
			return err
		}
		cb.failures++
		if cb.failures >= cb.failThreshold {
			cb.transition(Open)
		}
		return err
	}

	// Успех.
	if cb.state == HalfOpen {
		cb.transition(Closed) // зависимость ожила
	} else {
		cb.failures = 0
	}
	return nil
}

func main() {
	fmt.Println("--- circuit breaker (порог=3 ошибки, cooldown=2 тика) ---")
	cb := NewCircuitBreaker(3, 2)

	// Детерминированная «зависимость»: сначала падает, с 6-го вызова жива.
	callNo := 0
	dep := func() error {
		callNo++
		if callNo < 6 {
			return errors.New("зависимость недоступна")
		}
		return nil
	}

	// Управляемый сценарий действий.
	steps := []string{
		"call", "call", "call", // 3 ошибки подряд -> OPEN
		"call",         // OPEN: fail fast, зависимость не вызывается
		"tick", "tick", // cooldown истёк -> HALF-OPEN
		"call",         // пробный вызов: ещё падает (callNo=4<6) -> снова OPEN
		"tick", "tick", // -> HALF-OPEN
		"call",         // пробный: callNo=5<6, всё ещё падает -> OPEN
		"tick", "tick", // -> HALF-OPEN
		"call", // пробный: callNo=6, успех -> CLOSED
		"call", // норм работа в CLOSED
	}

	for i, s := range steps {
		switch s {
		case "tick":
			cb.Tick()
			fmt.Printf("шаг %2d: TICK   [%s]\n", i, cb.state)
		case "call":
			err := cb.Call(dep)
			if errors.Is(err, errOpen) {
				fmt.Printf("шаг %2d: CALL   -> отклонён предохранителем [%s]\n", i, cb.state)
			} else if err != nil {
				fmt.Printf("шаг %2d: CALL   -> ошибка зависимости [%s]\n", i, cb.state)
			} else {
				fmt.Printf("шаг %2d: CALL   -> успех [%s]\n", i, cb.state)
			}
		}
	}

	fmt.Printf("\nФинальное состояние: %s\n", cb.state)
	fmt.Println("Готово, программа завершается.")
}
