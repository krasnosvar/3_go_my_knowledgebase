package main

import (
	"fmt"
)

// TODO: Сгруппируйте объявление переменных ниже с помощью var ( ... ).
// Объявите:
// - Host со значением "localhost"
// - Port со значением 80
// - Active со значением true
//
// ВНИМАНИЕ: ниже стоят заглушки, чтобы программа компилировалась.
// Замените их одним групповым объявлением var ( ... ) с нужными значениями.
var Host string
var Port int
var Active bool

func main() {
	fmt.Println("--- Проверка 08_group_decl ---")

	if Host == "localhost" && Port == 80 && Active == true {
		fmt.Println("Тест группового объявления: PASS")
	} else {
		fmt.Printf("Тест группового объявления: FAIL (Host=%q, Port=%d, Active=%t)\n", Host, Port, Active)
	}
}
