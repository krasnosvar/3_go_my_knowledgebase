// Упражнение 08: метод валидации.
//
// Задача: реализуйте метод Validate у User. Он возвращает ПЕРВУЮ найденную
// ошибку в таком порядке проверок:
//   - пустое Name           -> ошибка "name is required";
//   - Age < 0 или Age > 150 -> ошибка "invalid age";
//   - Email без "@"         -> ошибка "invalid email";
//
// Если всё корректно — nil.
//
// Запуск: go run ./03_exercises/08_validation_method/main.go
package main

import (
	"errors"
	"fmt"
	"strings"
)

// User — данные пользователя.
type User struct {
	Name  string
	Email string
	Age   int
}

// Держим импорты живыми в заготовке.
var (
	_ = errors.New
	_ = strings.Contains
)

// Validate проверяет поля пользователя.
func (u User) Validate() error {
	// TODO: проверки по порядку из условия; errors.New для ошибок.
	return nil
}

func main() {
	fmt.Println("--- Проверка 08_validation_method ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	cases := []struct {
		u    User
		want string // "" — ошибки быть не должно
	}{
		{User{Name: "Den", Email: "den@example.com", Age: 30}, ""},
		{User{Name: "", Email: "x@y", Age: 30}, "name is required"},
		{User{Name: "Den", Email: "x@y", Age: -1}, "invalid age"},
		{User{Name: "Den", Email: "x@y", Age: 200}, "invalid age"},
		{User{Name: "Den", Email: "no-at", Age: 20}, "invalid email"},
		{User{Name: "", Email: "no-at", Age: -5}, "name is required"},
	}
	for i, c := range cases {
		err := c.u.Validate()
		got := ""
		if err != nil {
			got = err.Error()
		}
		check(got == c.want, fmt.Sprintf("случай %d: ожидалось %q, получено %q", i, c.want, got))
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
