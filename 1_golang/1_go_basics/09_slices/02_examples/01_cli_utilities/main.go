package main

import (
	"fmt"
)

func main() {
	// Демонстрируем побочные эффекты при нарезке без копирования.
	args := []string{"run", "--verbose", "--output=json", "main.go"}

	// 1. Нарезаем аргументы
	opts := args[1:3] // ["--verbose", "--output=json"]
	fmt.Println("Исходные опции (opts):   ", opts)

	// Меняем первый элемент опций
	opts[0] = "--silent"

	// Изменение отразилось на оригинальных аргументах!
	fmt.Println("После изменения opts:")
	fmt.Println("  opts:                 ", opts)
	fmt.Println("  args (оригинал):      ", args)

	// 2. Безопасное копирование с помощью copy()
	args2 := []string{"run", "--verbose", "--output=json", "main.go"}

	// Создаем независимый слайс нужного размера
	opts2 := make([]string, 2)
	copy(opts2, args2[1:3])

	// Меняем элемент в opts2
	opts2[0] = "--silent"

	fmt.Println("\nПосле безопасного копирования (copy):")
	fmt.Println("  opts2:                ", opts2)
	fmt.Println("  args2 (оригинал):     ", args2) // оригинал остался неизменным!
}
