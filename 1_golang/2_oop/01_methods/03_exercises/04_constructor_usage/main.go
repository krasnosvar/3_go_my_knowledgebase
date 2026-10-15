// Упражнение 04: конструктор и методы.
//
// Задача: реализуйте конструктор NewUser и метод Greeting.
//   - NewUser(name) возвращает *User с Name=name и ролью по умолчанию "user";
//   - Greeting() возвращает строку вида "Привет, Den (user)".
//
// В Go нет конструкторов в языке — это обычная функция NewXxx.
//
// Запуск: go run ./03_exercises/04_constructor_usage/main.go
package main

import (
	"fmt"
)

// User — пользователь.
type User struct {
	Name string
	Role string
}

// NewUser создаёт пользователя с ролью по умолчанию "user".
func NewUser(name string) *User {
	// TODO: верните &User{...} с Name=name и Role="user".
	return &User{}
}

// Greeting возвращает приветствие.
func (u *User) Greeting() string {
	// TODO: верните "Привет, <Name> (<Role>)" (подсказка: fmt.Sprintf).
	return ""
}

func main() {
	fmt.Println("--- Проверка 04_constructor_usage ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	u := NewUser("Den")
	check(u != nil, "NewUser не должен возвращать nil")
	if u == nil {
		fmt.Println("FAIL")
		return
	}
	check(u.Name == "Den", fmt.Sprintf("Name: ожидалось \"Den\", получено %q", u.Name))
	check(u.Role == "user", fmt.Sprintf("Role по умолчанию: ожидалось \"user\", получено %q", u.Role))
	check(u.Greeting() == "Привет, Den (user)", fmt.Sprintf("Greeting: получено %q", u.Greeting()))

	u.Role = "admin"
	check(u.Greeting() == "Привет, Den (admin)", fmt.Sprintf("Greeting после смены роли: получено %q", u.Greeting()))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
