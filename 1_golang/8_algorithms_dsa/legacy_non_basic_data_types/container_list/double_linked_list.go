package main

import (
	"container/list"
	"fmt"
)

func main() {
	// Создаем новый двусвязный список
	l := list.New()

	// PushBack добавляет элемент в конец списка
	// Мы можем добавлять любые типы данных (interface{}), например, строки
	l.PushBack("golang")
	l.PushBack("docker")

	// PushFront добавляет элемент в начало списка
	l.PushFront("kubernetes")

	// Front() возвращает указатель на первый элемент (*list.Element)
	// Элемент имеет методы Next() и Prev() для перемещения по списку
	fmt.Println("Проход по списку от начала к концу:")
	for e := l.Front(); e != nil; e = e.Next() {
		// e.Value содержит само значение элемента
		fmt.Println(e.Value)
	}

	// Мы также можем пройти список в обратном порядке
	fmt.Println("\nПроход по списку от конца к началу:")
	for e := l.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}

	// Пример удаления элемента
	// Сначала найдем элемент "docker"
	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value == "docker" {
			l.Remove(e) // Удаляем элемент из списка
			fmt.Println("\nЭлемент 'docker' был удален")
		}
	}

	fmt.Println("\nСписок после удаления:")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
