package main

import (
	"fmt"
)

// Групповое объявление переменных
var (
	Host   = "localhost"
	Port   = 80
	Active = true
)

func main() {
	fmt.Println("--- Проверка 08_group_decl (Решение) ---")

	if Host == "localhost" && Port == 80 && Active == true {
		fmt.Println("Тест группового объявления: PASS")
	} else {
		fmt.Printf("Тест группового объявления: FAIL (Host=%q, Port=%d, Active=%t)\n", Host, Port, Active)
	}
}
