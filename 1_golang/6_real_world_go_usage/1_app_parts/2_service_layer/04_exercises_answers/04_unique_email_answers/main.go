// Упражнение 04: правило уникальности и зависимость от часов (Решение).
package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

var errConflict = errors.New("email already registered")

type user struct {
	Email     string
	CreatedAt time.Time
}

type signupService struct {
	repo map[string]user
	now  func() time.Time
}

func (s *signupService) SignUp(email string) (user, error) {
	email = strings.ToLower(email)
	if _, exists := s.repo[email]; exists {
		return user{}, errConflict
	}
	u := user{Email: email, CreatedAt: s.now()}
	s.repo[email] = u
	return u, nil
}

func main() {
	fmt.Println("--- Проверка 04_unique_email (Решение) ---")

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
