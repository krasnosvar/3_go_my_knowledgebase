package main

import "fmt"

func main() {
	// В Go нет встроенного типа Set, его реализуют через map
	// Использование struct{} экономит память, так как пустая структура занимает 0 байт
	set := make(map[string]struct{})

	// Добавление элементов
	set["golang"] = struct{}{}
	set["docker"] = struct{}{}
	set["kubernetes"] = struct{}{}
	set["golang"] = struct{}{} // Дубликат не добавится

	// Проверка наличия
	val := "docker"
	if _, exists := set[val]; exists {
		fmt.Printf("Элемент '%s' есть в множестве\n", val)
	}

	// Удаление
	delete(set, "docker")

	// Итерация
	fmt.Println("\nТекущие элементы множества:")
	for item := range set {
		fmt.Println("-", item)
	}

	fmt.Printf("\nРазмер множества: %d\n", len(set))
}
