// Упражнение 01: валидация входа в service layer.
//
// Задача: реализуйте (s *userService) Register(email string, age int) error:
//   - email должен содержать "@" ровно один раз и не быть пустым;
//   - age — в диапазоне 18..120;
//   - при нарушении вернуть ошибку, обёрнутую вокруг errInvalid
//     (fmt.Errorf("...: %w", errInvalid)), и НЕ сохранять пользователя;
//   - при успехе добавить email в s.saved.
//
// Запуск: go run ./03_exercises/01_validate_input
package main

import (
	"errors"
	"fmt"
	"strings"
)

var errInvalid = errors.New("invalid input")

var _ = strings.Count

type userService struct {
	saved []string
}

func (s *userService) Register(email string, age int) error {
	// TODO: провалидируйте email и age, оберните ошибку вокруг errInvalid.
	return nil
}

func main() {
	fmt.Println("--- Проверка 01_validate_input ---")

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
