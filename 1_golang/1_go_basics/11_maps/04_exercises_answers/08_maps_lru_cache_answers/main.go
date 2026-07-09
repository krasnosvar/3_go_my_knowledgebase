package main

import (
	"container/list"
	"fmt"
)

/*
	ОТВЕТЫ К ЗАДАНИЮ 8 (LRU Cache)
*/

// ================= РЕШЕНИЕ =================

type LRUCache struct {
	capacity int
	list     *list.List
	cache    map[int]*list.Element // Храним указатель на элемент списка для быстрого доступа
}

// Элемент, который мы будем хранить внутри Value списка (чтобы при удалении с конца знать, КАКОЙ ключ удалить из мапы)
type entry struct {
	key   int
	value int
}

func NewLRU(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		list:     list.New(),
		cache:    make(map[int]*list.Element),
	}
}

func (l *LRUCache) Get(key int) int {
	if element, ok := l.cache[key]; ok {
		// Ключ найден. Перемещаем элемент списка вперед (помечаем как "свежий")
		l.list.MoveToFront(element)
		// Возвращаем само значение из структуры entry
		return element.Value.(*entry).value
	}
	return -1
}

func (l *LRUCache) Put(key int, value int) {
	if element, ok := l.cache[key]; ok {
		// Ключ уже есть. Обновляем значение и делаем "свежим".
		l.list.MoveToFront(element)
		element.Value.(*entry).value = value
		return
	}

	// Ключа нет. Создаем новый элемент в начале списка (самый "свежий")
	newEntry := &entry{key, value}
	element := l.list.PushFront(newEntry)
	l.cache[key] = element

	// Проверяем capacity
	if l.list.Len() > l.capacity {
		// Если переполнен, удаляем самый старый элемент (с конца списка)
		oldest := l.list.Back()
		if oldest != nil {
			// Удаляем из списка
			l.list.Remove(oldest)
			// Извлекаем ключ этого старого элемента и удаляем его же из мапы
			deleteKey := oldest.Value.(*entry).key
			delete(l.cache, deleteKey)
		}
	}
}

// ===========================================

func main() {
	lru := NewLRU(2)
	lru.Put(1, 1)           // cache is {1=1}
	lru.Put(2, 2)           // cache is {1=1, 2=2}
	fmt.Println(lru.Get(1)) // return 1, cache is {2=2, 1=1}

	lru.Put(3, 3)           // evicts key 2, cache is {1=1, 3=3}
	fmt.Println(lru.Get(2)) // return -1 (not found)

	lru.Put(4, 4)           // evicts key 1, cache is {3=3, 4=4}
	fmt.Println(lru.Get(1)) // return -1 (not found)
	fmt.Println(lru.Get(3)) // return 3
	fmt.Println(lru.Get(4)) // return 4
}
