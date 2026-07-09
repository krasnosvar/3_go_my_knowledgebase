package main

import (
	"fmt"
)

func getAge(db map[string]int, name string) (int, bool) {
	age, ok := db[name]
	if !ok {
		return -1, false
	}
	return age, true
}

func main() {
	fmt.Println("--- Проверка 01_maps_basics (Решение) ---")

	db := map[string]int{
		"Alice": 25,
		"Bob":   30,
	}

	age1, ok1 := getAge(db, "Alice")
	age2, ok2 := getAge(db, "Charlie")

	t1 := ok1 && age1 == 25
	t2 := !ok2 && age2 == -1

	if t1 && t2 {
		fmt.Println("Тест getAge: PASS")
	} else {
		fmt.Printf("Тест getAge: FAIL (age1=%d, ok1=%t, age2=%d, ok2=%t)\n", age1, ok1, age2, ok2)
	}
}
