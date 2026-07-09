package main

import (
	"fmt"
	"reflect"
)

// TODO: Объявите нетипизированную константу `Pi` со значением 3.14
// TODO: Объявите типизированную константу `MaxCount` типа int32 со значением 100

var _ = reflect.TypeOf

func main() {
	fmt.Println("--- Проверка 03_const ---")

	// ВНИМАНИЕ: Для прохождения проверки объявите константы выше функции main, а здесь мы проверим их значения
	// (Если компилятор ругается на undefined, значит константы не объявлены)

	/*
		fmt.Printf("Pi: %v (тип: %s)\n", Pi, reflect.TypeOf(Pi))
		fmt.Printf("MaxCount: %v (тип: %s)\n", MaxCount, reflect.TypeOf(MaxCount))

		if Pi == 3.14 {
			fmt.Println("Тест Pi: PASS")
		} else {
			fmt.Println("Тест Pi: FAIL")
		}

		if MaxCount == 100 && reflect.TypeOf(MaxCount).String() == "int32" {
			fmt.Println("Тест MaxCount: PASS")
		} else {
			fmt.Println("Тест MaxCount: FAIL")
		}
	*/

	fmt.Println("Раскомментируйте блок кода в main для запуска тестов")
}
