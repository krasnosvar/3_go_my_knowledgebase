package main

import (
	"container/list"
	"fmt"
)

func main() {
	// Инициализация
	l := list.New()

	// Наполняем список
	e1 := l.PushBack("Первый")
	e2 := l.PushBack("Второй")
	e3 := l.PushBack("Третий")

	fmt.Println("Исходный список:")
	printList(l)

	// InsertAfter: вставляем новый элемент после e2 ("Второй")
	l.InsertAfter("2.5 (вставка после Второго)", e2)
	fmt.Println("\nПосле InsertAfter:")
	printList(l)

	// MoveToFront: перемещаем e3 ("Третий") в самое начало
	l.MoveToFront(e3)
	fmt.Println("\nПосле MoveToFront (Третий в начало):")
	printList(l)

	// MoveAfter: перемещаем e1 ("Первый") после e2 ("Второй")
	l.MoveAfter(e1, e2)
	fmt.Println("\nПосле MoveAfter (Первый после Второго):")
	printList(l)

	// Очистка списка
	fmt.Printf("\nДлина списка до Init(): %d\n", l.Len())
	l.Init() // Init() очищает список или инициализирует новый
	fmt.Printf("Длина списка после Init(): %d\n", l.Len())
}

// Вспомогательная функция для печати списка
func printList(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("[%v] ", e.Value)
	}
	fmt.Println()
}
