// Упражнение 04: конструктор и методы (Решение).
package main

import (
	"fmt"
)

// User — пользователь.
type User struct {
	Name string
	Role string
}

func NewUser(name string) *User {
	return &User{Name: name, Role: "user"}
}

func (u *User) Greeting() string {
	return fmt.Sprintf("Привет, %s (%s)", u.Name, u.Role)
}

func main() {
	fmt.Println("--- Проверка 04_constructor_usage (Решение) ---")

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
