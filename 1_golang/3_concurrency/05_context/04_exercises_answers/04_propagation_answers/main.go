// Упражнение 04: распространение отмены от родителя к потомку (Решение).
package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func derive(parent context.Context) (context.Context, context.CancelFunc) {
	return context.WithCancel(parent)
}

func main() {
	fmt.Println("--- Проверка 04_propagation (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	isDone := func(ctx context.Context) bool {
		select {
		case <-ctx.Done():
			return true
		case <-time.After(50 * time.Millisecond):
			return false
		}
	}

	parent, cancelParent := context.WithCancel(context.Background())
	child, cancelChild := derive(parent)
	if child == nil {
		check(false, "derive вернула nil")
	} else {
		cancelChild()
		check(isDone(child), "отмена потомка должна завершить потомка")
		check(!isDone(parent), "отмена потомка не должна затрагивать родителя")
	}
	cancelParent()

	parent2, cancelParent2 := context.WithCancel(context.Background())
	child2, cancelChild2 := derive(parent2)
	defer cancelChild2()
	cancelParent2()
	if child2 != nil {
		check(isDone(child2), "отмена родителя должна отменить потомка")
		check(errors.Is(child2.Err(), context.Canceled), fmt.Sprintf("Err потомка = %v, ожидалось context.Canceled", child2.Err()))
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
