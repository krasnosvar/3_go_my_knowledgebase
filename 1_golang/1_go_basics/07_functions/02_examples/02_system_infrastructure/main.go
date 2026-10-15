// ⏩ Забегание вперёд: горутины (go func) разбираются в 3_concurrency/01_goroutines.

package main

import (
	"fmt"
	"time"
)

// SafeGo запускает переданную функцию в горутине и предотвращает краш приложения,
// если функция вызовет панику.
func SafeGo(task func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("[RECOVERY] Перехвачена паника в фоновой задаче: %v\n", r)
			}
		}()

		task()
	}()
}

func main() {
	fmt.Println("--- Старт главного приложения ---")

	// Запускаем фоновую задачу, которая завершится паникой
	SafeGo(func() {
		fmt.Println("  [Worker] Начало вычислений...")
		time.Sleep(10 * time.Millisecond)
		panic("деление на ноль во встроенной библиотеке")
	})

	// Даем фоновой задаче время выполниться
	time.Sleep(50 * time.Millisecond)

	fmt.Println("--- Приложение продолжает жить и работать! ---")
}
