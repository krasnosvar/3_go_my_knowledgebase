// Упражнение 07: явный Close вместо финализатора (Решение).
package main

import (
	"errors"
	"fmt"
)

// ErrAlreadyClosed возвращается при повторном Close.
var ErrAlreadyClosed = errors.New("resource already closed")

// openResources — счётчик неосвобождённых ресурсов (как открытые файлы/сокеты).
var openResources int

type Resource struct {
	closed bool
}

func Open() *Resource {
	openResources++
	return &Resource{}
}

// Close освобождает ресурс ровно один раз.
func (r *Resource) Close() error {
	if r.closed {
		return ErrAlreadyClosed
	}
	r.closed = true
	openResources--
	return nil
}

func main() {
	fmt.Println("--- Проверка 07_finalizer_risk (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	r1 := Open()
	r2 := Open()
	check(openResources == 2, fmt.Sprintf("после двух Open ожидалось 2 ресурса, получено %d", openResources))

	check(r1.Close() == nil, "первый Close должен вернуть nil")
	check(openResources == 1, fmt.Sprintf("после Close ожидался 1 ресурс, получено %d", openResources))
	check(errors.Is(r1.Close(), ErrAlreadyClosed), "повторный Close должен вернуть ErrAlreadyClosed")
	check(openResources == 1, "повторный Close не должен ничего освобождать")

	func() {
		defer r2.Close() // идиома: закрываем через defer сразу после Open
	}()
	check(openResources == 0, fmt.Sprintf("после defer r2.Close() ожидалось 0 ресурсов, получено %d", openResources))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
