// Упражнение 01: struct-теги encoding/json.
//
// Задача: расставьте теги полям структуры User так, чтобы json.Marshal давал
//
//	{"id":7,"name":"Ann","email":"ann@example.com"}
//
// и при пустом Email поле email вообще не выводилось. Правила:
//   - ID → "id", Name → "name";
//   - Email → "email" с опцией omitempty;
//   - Password никогда не попадает в JSON (тег "-").
//
// Менять нужно только теги — код main трогать не требуется.
//
// Запуск: go run ./03_exercises/01_json_tags/main.go
package main

import (
	"encoding/json"
	"fmt"
)

// TODO: добавьте json-теги полям User.
type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

func main() {
	fmt.Println("--- Проверка 01_json_tags ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	b, err := json.Marshal(User{ID: 7, Name: "Ann", Email: "ann@example.com", Password: "secret"})
	want := `{"id":7,"name":"Ann","email":"ann@example.com"}`
	check(err == nil && string(b) == want, fmt.Sprintf("получено %s, ожидалось %s", b, want))

	b, _ = json.Marshal(User{ID: 1, Name: "Bob"})
	want = `{"id":1,"name":"Bob"}`
	check(string(b) == want, fmt.Sprintf("пустой Email не должен выводиться: получено %s", b))

	var u User
	err = json.Unmarshal([]byte(`{"id":3,"name":"Eve","password":"x"}`), &u)
	check(err == nil && u.ID == 3 && u.Name == "Eve" && u.Password == "",
		fmt.Sprintf("Unmarshal: получено %+v, %v (Password должен игнорироваться)", u, err))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
