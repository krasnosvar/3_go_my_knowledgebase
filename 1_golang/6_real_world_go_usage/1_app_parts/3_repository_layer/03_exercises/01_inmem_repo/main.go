// Упражнение 01: in-memory репозиторий по контракту.
//
// Задача: реализуйте методы memRepo, удовлетворяющие интерфейсу UserRepo:
//   - Create(u) присваивает u.ID = следующий номер (1, 2, ...), сохраняет и возвращает ID;
//   - Get(id) возвращает (user, nil) или (user{}, ErrNotFound);
//   - Delete(id) удаляет; удаление отсутствующего → ErrNotFound.
//
// Запуск: go run ./03_exercises/01_inmem_repo
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
	// TODO: выдайте следующий ID и сохраните пользователя.
	return 0
}

func (r *memRepo) Get(id int) (user, error) {
	// TODO: верните ErrNotFound для отсутствующего id.
	return user{}, nil
}

func (r *memRepo) Delete(id int) error {
	// TODO: удалите запись; отсутствующая → ErrNotFound.
	return nil
}

func main() {
	fmt.Println("--- Проверка 01_inmem_repo ---")

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
