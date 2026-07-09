package main

import (
	"fmt"
	"runtime"
)

// generateHugeLogData возвращает слайс большого объема (например, 10 млн байт)
func generateHugeLogData() []byte {
	return make([]byte, 10_000_000)
}

func main() {
	// 1. Опасный путь: получаем срез огромного слайса
	hugeSlice := generateHugeLogData()

	// Оставляем только последние 2 байта
	badClip := hugeSlice[len(hugeSlice)-2:]

	// Удаляем ссылку на исходный слайс
	hugeSlice = nil

	// Вызываем сборщик мусора
	runtime.GC()

	// Измеряем память
	var m1 runtime.MemStats
	runtime.ReadMemStats(&m1)
	fmt.Printf("Память удерживается (плохой путь): %.2f MB\n", float64(m1.Alloc)/1024/1024)
	_ = badClip // держим ссылку, чтобы GC не собрал badClip

	// 2. Безопасный путь: копируем нужные байты в новый маленький слайс
	hugeSlice2 := generateHugeLogData()

	goodClip := make([]byte, 2)
	copy(goodClip, hugeSlice2[len(hugeSlice2)-2:])

	hugeSlice2 = nil
	runtime.GC()

	var m2 runtime.MemStats
	runtime.ReadMemStats(&m2)
	fmt.Printf("Память удерживается (хороший путь): %.2f MB\n", float64(m2.Alloc)/1024/1024)
	_ = goodClip
}
