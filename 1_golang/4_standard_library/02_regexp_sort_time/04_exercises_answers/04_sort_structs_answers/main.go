// Упражнение 04: сортировка структур по нескольким ключам (Решение).
package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

// sortUsers сортирует пользователей по (Age, Name).
func sortUsers(users []User) {
	sort.Slice(users, func(i, j int) bool {
		if users[i].Age != users[j].Age {
			return users[i].Age < users[j].Age
		}
		return users[i].Name < users[j].Name
	})
}

func main() {
	fmt.Println("--- Проверка 04_sort_structs (Решение) ---")

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
