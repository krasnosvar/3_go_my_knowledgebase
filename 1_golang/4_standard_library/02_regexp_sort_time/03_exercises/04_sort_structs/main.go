// Упражнение 04: сортировка структур по нескольким ключам.
//
// Задача: реализуйте sortUsers(users) — сортировка на месте:
//   - по Age по возрастанию;
//   - при равном Age — по Name по возрастанию (tie-break).
//
// Используйте sort.Slice с функцией less, учитывающей оба ключа.
//
// Запуск: go run ./03_exercises/04_sort_structs/main.go
package main

import (
	"fmt"
	"sort"
)

var _ = sort.Slice // держим импорт живым в заготовке

type User struct {
	Name string
	Age  int
}

// sortUsers сортирует пользователей по (Age, Name).
func sortUsers(users []User) {
	// TODO: sort.Slice(users, func(i, j int) bool { ... }) с tie-break по Name.
}

func main() {
	fmt.Println("--- Проверка 04_sort_structs ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	users := []User{{"Zoe", 30}, {"Bob", 25}, {"Ann", 30}, {"Eve", 25}, {"Max", 20}}
	sortUsers(users)
	got := fmt.Sprint(users)
	want := "[{Max 20} {Bob 25} {Eve 25} {Ann 30} {Zoe 30}]"
	check(got == want, fmt.Sprintf("ожидалось %s, получено %s", want, got))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
