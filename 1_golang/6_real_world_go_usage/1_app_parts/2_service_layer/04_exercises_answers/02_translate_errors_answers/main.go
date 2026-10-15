// Упражнение 02: сервис зависит от интерфейса репозитория (Решение).
package main

import (
	"errors"
	"fmt"
)

var (
	errRepoNotFound = errors.New("repo: no rows")
	ErrUserNotFound = errors.New("user not found")
	errDBDown       = errors.New("db down")
)

// userRepo — узкий интерфейс, объявленный на стороне потребителя (сервиса).
type userRepo interface {
	Get(id int) (string, error)
}

type service struct{ repo userRepo }

func (s *service) UserName(id int) (string, error) {
	name, err := s.repo.Get(id)
	if errors.Is(err, errRepoNotFound) {
		return "", ErrUserNotFound
	}
	if err != nil {
		return "", fmt.Errorf("get user: %w", err)
	}
	return name, nil
}

// fakeRepo — тестовый двойник: данные в map, ошибка для id=500.
type fakeRepo map[int]string

func (f fakeRepo) Get(id int) (string, error) {
	if id == 500 {
		return "", errDBDown
	}
	n, ok := f[id]
	if !ok {
		return "", errRepoNotFound
	}
	return n, nil
}

func main() {
	fmt.Println("--- Проверка 02_translate_errors (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	s := &service{repo: fakeRepo{1: "den"}}
	n, err := s.UserName(1)
	check(err == nil && n == "den", "id=1 → \"den\"")
	_, err = s.UserName(2)
	check(errors.Is(err, ErrUserNotFound), "id=2 → ErrUserNotFound")
	check(!errors.Is(err, errRepoNotFound), "ошибка репозитория не протекает наружу")
	_, err = s.UserName(500)
	check(errors.Is(err, errDBDown) && err != errDBDown, "прочая ошибка обёрнута с %w")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
