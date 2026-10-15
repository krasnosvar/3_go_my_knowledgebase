// Упражнение 08: метод валидации (Решение).
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

func (u User) Validate() error {
	if u.Name == "" {
		return errors.New("name is required")
	}
	if u.Age < 0 || u.Age > 150 {
		return errors.New("invalid age")
	}
	if !strings.Contains(u.Email, "@") {
		return errors.New("invalid email")
	}
	return nil
}

func main() {
	fmt.Println("--- Проверка 08_validation_method (Решение) ---")

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
