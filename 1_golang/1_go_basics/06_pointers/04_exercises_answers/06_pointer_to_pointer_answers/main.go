package main

import (
	"fmt"
)

// getValueFromDoublePointer разыменовывает двойной указатель
func getValueFromDoublePointer(ptr **int) int {
	if ptr == nil || *ptr == nil {
		return -1
	}
	return **ptr
}

func main() {
	fmt.Println("--- Проверка 06_pointer_to_pointer (Решение) ---")

	x := 99
	px := &x
	ppx := &px

	v1 := getValueFromDoublePointer(ppx)
	v2 := getValueFromDoublePointer(nil)

	var pnil *int = nil
	v3 := getValueFromDoublePointer(&pnil)

	if v1 == 99 && v2 == -1 && v3 == -1 {
		fmt.Println("Тест double pointer: PASS")
	} else {
		fmt.Printf("Тест double pointer: FAIL (v1=%d, v2=%d, v3=%d)\n", v1, v2, v3)
	}
}
