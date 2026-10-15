// Упражнение 01: валидация входа в service layer (Решение).
package main

import (
	"errors"
	"fmt"
	"strings"
)

var errInvalid = errors.New("invalid input")

type userService struct {
	saved []string
}

func (s *userService) Register(email string, age int) error {
	if strings.Count(email, "@") != 1 {
		return fmt.Errorf("email %q: %w", email, errInvalid)
	}
	if age < 18 || age > 120 {
		return fmt.Errorf("age %d: %w", age, errInvalid)
	}
	s.saved = append(s.saved, email)
	return nil
}

func main() {
	fmt.Println("--- Проверка 01_validate_input (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	s := &userService{}
	check(s.Register("a@b.c", 30) == nil, "валидный пользователь регистрируется")
	check(errors.Is(s.Register("no-at", 30), errInvalid), "email без @ → errInvalid")
	check(errors.Is(s.Register("a@@b", 30), errInvalid), "email с двумя @ → errInvalid")
	check(errors.Is(s.Register("x@y.z", 10), errInvalid), "age=10 → errInvalid")
	check(len(s.saved) == 1, "сохранён только валидный пользователь")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
