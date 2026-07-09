package main

import "fmt"

// Node представляет отдельный узел списка
type Node struct {
	Value int   // Значение узла
	Next  *Node // Указатель на следующий узел
}

// SinglyLinkedList представляет сам список
type SinglyLinkedList struct {
	Head *Node // Указатель на начало списка
}

// Add добавляет новый элемент в конец списка
func (l *SinglyLinkedList) Add(value int) {
	newNode := &Node{Value: value}

	// Если список пуст, новый узел становится головой
	if l.Head == nil {
		l.Head = newNode
		return
	}

	// Ищем последний узел
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}

	// Присоединяем новый узел к последнему
	current.Next = newNode
}

// Print выводит все элементы списка
func (l *SinglyLinkedList) Print() {
	if l.Head == nil {
		fmt.Println("Список пуст")
		return
	}

	current := l.Head
	for current != nil {
		fmt.Printf("%d -> ", current.Value)
		current = current.Next
	}
	fmt.Println("nil")
}

func main() {
	// Инициализируем пустой список
	myList := &SinglyLinkedList{}

	fmt.Println("Добавляем элементы 10, 20, 30:")
	myList.Add(10)
	myList.Add(20)
	myList.Add(30)

	// Выводим результат
	myList.Print()

	// Пример ручного обхода
	fmt.Println("Ручной обход:")
	for node := myList.Head; node != nil; node = node.Next {
		fmt.Printf("Узел с данными: %d\n", node.Value)
	}
}
