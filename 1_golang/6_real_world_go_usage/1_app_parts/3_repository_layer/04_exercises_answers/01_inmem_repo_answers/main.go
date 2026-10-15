// Упражнение 01: in-memory репозиторий по контракту (Решение).
package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

type user struct {
	ID   int
	Name string
}

type UserRepo interface {
	Create(u user) int
	Get(id int) (user, error)
	Delete(id int) error
}

type memRepo struct {
	data   map[int]user
	nextID int
}

func newMemRepo() *memRepo { return &memRepo{data: map[int]user{}} }

func (r *memRepo) Create(u user) int {
	r.nextID++
	u.ID = r.nextID
	r.data[u.ID] = u
	return u.ID
}

func (r *memRepo) Get(id int) (user, error) {
	u, ok := r.data[id]
	if !ok {
		return user{}, ErrNotFound
	}
	return u, nil
}

func (r *memRepo) Delete(id int) error {
	if _, ok := r.data[id]; !ok {
		return ErrNotFound
	}
	delete(r.data, id)
	return nil
}

func main() {
	fmt.Println("--- Проверка 01_inmem_repo (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	var repo UserRepo = newMemRepo()
	id1 := repo.Create(user{Name: "a"})
	id2 := repo.Create(user{Name: "b"})
	check(id1 == 1 && id2 == 2, "ID выдаются по порядку 1, 2")
	u, err := repo.Get(2)
	check(err == nil && u.Name == "b" && u.ID == 2, "Get(2) возвращает b с ID=2")
	_, err = repo.Get(99)
	check(errors.Is(err, ErrNotFound), "Get(99) → ErrNotFound")
	check(repo.Delete(1) == nil, "Delete(1) успешен")
	check(errors.Is(repo.Delete(1), ErrNotFound), "повторный Delete(1) → ErrNotFound")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
