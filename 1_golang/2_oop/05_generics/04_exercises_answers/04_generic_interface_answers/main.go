// Упражнение 04: обобщённый интерфейс (Решение).
package main

import (
	"fmt"
)

// Repository — обобщённый интерфейс хранилища.
type Repository[T any] interface {
	Save(id string, v T)
	Get(id string) (T, bool)
	Count() int
}

// User — пример сущности.
type User struct {
	Name string
}

// memRepo — хранилище в памяти.
type memRepo[T any] struct {
	m map[string]T
}

func NewMemRepo[T any]() *memRepo[T] {
	return &memRepo[T]{m: make(map[string]T)}
}

func (r *memRepo[T]) Save(id string, v T) {
	r.m[id] = v
}

func (r *memRepo[T]) Get(id string) (T, bool) {
	v, ok := r.m[id]
	return v, ok
}

func (r *memRepo[T]) Count() int {
	return len(r.m)
}

func main() {
	fmt.Println("--- Проверка 04_generic_interface (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	var users Repository[User] = NewMemRepo[User]()
	users.Save("1", User{Name: "Den"})
	users.Save("2", User{Name: "Ann"})
	users.Save("1", User{Name: "Denis"}) // перезапись

	u, ok := users.Get("1")
	check(ok && u.Name == "Denis", fmt.Sprintf("Get(\"1\"): ожидалось Denis, получено (%+v, %t)", u, ok))
	_, ok = users.Get("404")
	check(!ok, "Get несуществующего: ожидалось ok=false")
	check(users.Count() == 2, fmt.Sprintf("Count: ожидалось 2, получено %d", users.Count()))

	var nums Repository[int] = NewMemRepo[int]()
	nums.Save("x", 10)
	n, ok := nums.Get("x")
	check(ok && n == 10, "Repository[int] тоже должен работать")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
