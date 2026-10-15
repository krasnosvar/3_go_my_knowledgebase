// Упражнение 04: authz на уровне ресурса (ownership).
//
// Задача: реализуйте authorize(u user, d doc, action string) error:
//   - admin может всё;
//   - "read": владелец или документ публичный (d.Public);
//   - "edit", "delete": только владелец (d.OwnerID == u.ID);
//   - неизвестное действие → errUnknownAction;
//   - отказ → errForbidden (handler превратит его в 403, а не 404/500).
//
// Запуск: go run ./03_exercises/04_ownership
package main

import (
	"errors"
	"fmt"
)

var (
	errForbidden     = errors.New("forbidden")
	errUnknownAction = errors.New("unknown action")
)

type user struct {
	ID    int
	Admin bool
}

type doc struct {
	OwnerID int
	Public  bool
}

func authorize(u user, d doc, action string) error {
	// TODO: проверьте admin, затем правила для read/edit/delete.
	return nil
}

func main() {
	fmt.Println("--- Проверка 04_ownership ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	owner, other, admin := user{ID: 1}, user{ID: 2}, user{ID: 3, Admin: true}
	private, public := doc{OwnerID: 1}, doc{OwnerID: 1, Public: true}

	check(authorize(owner, private, "edit") == nil, "владелец редактирует")
	check(errors.Is(authorize(other, private, "read"), errForbidden), "чужой не читает приватный")
	check(authorize(other, public, "read") == nil, "чужой читает публичный")
	check(errors.Is(authorize(other, public, "delete"), errForbidden), "чужой не удаляет публичный")
	check(authorize(admin, private, "delete") == nil, "admin удаляет всё")
	check(errors.Is(authorize(owner, private, "share"), errUnknownAction), "неизвестное действие → errUnknownAction")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
