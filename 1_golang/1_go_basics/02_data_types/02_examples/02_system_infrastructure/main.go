package main

import (
	"fmt"
	"math"
)

func main() {
	// Объем памяти ОС в байтах (uint64 отлично подходит для больших размеров памяти)
	var totalMemory uint64 = 16 * 1024 * 1024 * 1024 // 16 GB
	var usedMemory uint64 = 6 * 1024 * 1024 * 1024   // 6 GB

	// Загрузка CPU: float64 для точных вычислений
	var cpuLoad float64 = 0.7385

	// Ловушка: деление целых чисел (totalMemory / usedMemory) в Go отбросит дробную часть.
	// Чтобы получить корректный процент, нужно сначала привести значения к float64.
	memoryUsagePercent := (float64(usedMemory) / float64(totalMemory)) * 100

	fmt.Println("--- Системные метрики ---")
	fmt.Printf("Общая память: %d ГБ\n", totalMemory/(1024*1024*1024))
	fmt.Printf("Используемая память: %.2f%%\n", memoryUsagePercent)

	// Округляем нагрузку процессора до процентов
	cpuPercent := math.Round(cpuLoad * 100)
	fmt.Printf("Загрузка CPU: %.0f%%\n", cpuPercent)

	// Пример сравнения float: имитация порога предупреждения
	const alertThreshold float64 = 0.70
	if cpuLoad > alertThreshold {
		fmt.Printf("ВНИМАНИЕ: Нагрузка CPU (%.2f) превышает лимит (%.2f)!\n", cpuLoad, alertThreshold)
	}
}
