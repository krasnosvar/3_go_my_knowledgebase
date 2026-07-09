package main

import (
	"fmt"
	"runtime"
)

// printMemStats выводит текущее количество памяти в куче (heap allocations)
func printMemStats(stage string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("[%s] Heap Alloc: %7.2f MB\n", stage, float64(m.Alloc)/1024/1024)
}

func main() {
	fmt.Println("--- Симуляция утечки памяти в Map ---")
	printMemStats("Старт")

	// 1. Инициализируем мапу и наполняем 100 000 больших структур
	m := make(map[int][100]int)
	for i := 0; i < 100_000; i++ {
		m[i] = [100]int{}
	}
	printMemStats("Заполнение (100k элементов)")

	// 2. Удаляем 99 000 элементов
	for i := 0; i < 99_000; i++ {
		delete(m, i)
	}

	// Вызываем сборщик мусора
	runtime.GC()
	printMemStats("После delete (удалено 99k)")

	// ВАЖНО: Память не уменьшилась, так как мапа не сжимает свои бакеты в памяти!

	// 3. Решение: Копируем оставшиеся элементы в новую мапу
	newMap := make(map[int][100]int, len(m))
	for k, v := range m {
		newMap[k] = v
	}
	m = nil // Отпускаем старую мапу для GC
	_ = newMap

	// Вызываем GC повторно
	runtime.GC()
	printMemStats("После переноса в новую мапу")
}
