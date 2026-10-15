// Упражнение 04: правило уникальности и зависимость от часов.
//
// Задача: реализуйте (s *signupService) SignUp(email string) (user, error):
//   - email приводится к нижнему регистру (strings.ToLower) перед проверкой;
//   - если такой email уже есть в репозитории → errConflict;
//   - иначе создать user{Email, CreatedAt: s.now()} и сохранить в s.repo.
//
// Часы приходят снаружи (поле now) — так сервис детерминированно тестируется.
//
// Запуск: go run ./03_exercises/04_unique_email
package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

var errConflict = errors.New("email already registered")

var _ = strings.ToLower

type user struct {
	Email     string
	CreatedAt time.Time
}

type signupService struct {
	repo map[string]user
	now  func() time.Time
}

func (s *signupService) SignUp(email string) (user, error) {
	// TODO: нормализуйте email, проверьте конфликт, сохраните с CreatedAt = s.now().
	return user{}, nil
}

func main() {
	fmt.Println("--- Проверка 04_unique_email ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	fixed := time.Date(2024, 1, 2, 3, 4, 5, 0, time.UTC)
	s := &signupService{repo: map[string]user{}, now: func() time.Time { return fixed }}

	u, err := s.SignUp("Den@Example.com")
	check(err == nil && u.Email == "den@example.com", "email нормализован к нижнему регистру")
	check(u.CreatedAt.Equal(fixed), "CreatedAt берётся из инъектированных часов")
	_, err = s.SignUp("DEN@example.COM")
	check(errors.Is(err, errConflict), "повтор в другом регистре → errConflict")
	check(len(s.repo) == 1, "в репозитории один пользователь")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
