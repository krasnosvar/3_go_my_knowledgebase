// Упражнение 03: обобщённая структура (Решение).
package main

import (
	"fmt"
)

// Optional хранит значение, которое может отсутствовать.
type Optional[T any] struct {
	val T
	set bool
}

func Some[T any](v T) Optional[T] {
	return Optional[T]{val: v, set: true}
}

func (o Optional[T]) Get() (T, bool) {
	return o.val, o.set
}

func (o Optional[T]) OrElse(def T) T {
	if o.set {
		return o.val
	}
	return def
}

func main() {
	fmt.Println("--- Проверка 03_generic_struct (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	o := Some(42)
	v, ok := o.Get()
	check(ok && v == 42, fmt.Sprintf("Some(42).Get: ожидалось (42, true), получено (%d, %t)", v, ok))
	check(o.OrElse(7) == 42, "Some(42).OrElse(7): ожидалось 42")

	var empty Optional[string]
	_, ok = empty.Get()
	check(!ok, "пустой Optional: ожидалось ok=false")
	check(empty.OrElse("default") == "default", "пустой Optional.OrElse: ожидалось \"default\"")

	zero := Some(0) // ноль — тоже значение
	_, ok = zero.Get()
	check(ok && zero.OrElse(5) == 0, "Some(0): значение есть, OrElse должен вернуть 0")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
