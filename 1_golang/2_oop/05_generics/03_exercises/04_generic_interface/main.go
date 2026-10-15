// Упражнение 04: обобщённый интерфейс.
//
// Задача: реализуйте in-memory хранилище memRepo[T], удовлетворяющее
// обобщённому интерфейсу Repository[T]:
//   - Save(id, v) — сохранить/перезаписать;
//   - Get(id) (T, bool) — достать;
//   - Count() int — количество записей.
//
// NewMemRepo[T]() должен вернуть готовый к работе репозиторий (с map).
//
// Запуск: go run ./03_exercises/04_generic_interface/main.go
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

// NewMemRepo создаёт пустое хранилище.
func NewMemRepo[T any]() *memRepo[T] {
	// TODO: инициализируйте map.
	return &memRepo[T]{}
}

func (r *memRepo[T]) Save(id string, v T) {
	// TODO: r.m[id] = v
}

func (r *memRepo[T]) Get(id string) (T, bool) {
	// TODO: v, ok := r.m[id]
	var zero T
	return zero, false
}

func (r *memRepo[T]) Count() int {
	// TODO: len(r.m)
	return 0
}

func main() {
	fmt.Println("--- Проверка 04_generic_interface ---")

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
