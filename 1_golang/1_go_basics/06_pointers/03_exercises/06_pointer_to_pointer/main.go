package main

import (
	"fmt"
)

// getValueFromDoublePointer разыменовывает двойной указатель и возвращает значение.
// Если ptr равен nil, или *ptr равен nil, верните -1.
func getValueFromDoublePointer(ptr **int) int {
	// TODO: Реализуйте безопасное разыменование двойного указателя.
	return 0
}

func main() {
	fmt.Println("--- Проверка 06_pointer_to_pointer ---")

	x := 99
	px := &x
	ppx := &px

	v1 := getValueFromDoublePointer(ppx) // 99
	v2 := getValueFromDoublePointer(nil) // -1

	var pnil *int = nil
	v3 := getValueFromDoublePointer(&pnil) // -1

	if v1 == 99 && v2 == -1 && v3 == -1 {
		fmt.Println("Тест double pointer: PASS")
	} else {
		fmt.Printf("Тест double pointer: FAIL (v1=%d, v2=%d, v3=%d)\n", v1, v2, v3)
	}
}
