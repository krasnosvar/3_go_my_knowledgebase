package main

import (
	"fmt"
)

// getAge возвращает возраст человека по его имени.
// Если человека нет в базе, верните -1 и false.
func getAge(db map[string]int, name string) (int, bool) {
	// TODO: Реализуйте поиск по мапе с помощью comma-ok.
	return 0, false
}

func main() {
	fmt.Println("--- Проверка 01_maps_basics ---")

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
