package main

import (
	"fmt"
	"strings"
	"time"
)

const iterations = 20_000

func concatWithPlus() string {
	var s string
	for i := 0; i < iterations; i++ {
		s += "log_line "
	}
	return s
}

func concatWithBuilder() string {
	var b strings.Builder
	// Оптимизация: резервируем память под итоговую строку (9 байт на слово)
	b.Grow(iterations * 9)
	for i := 0; i < iterations; i++ {
		b.WriteString("log_line ")
	}
	return b.String()
}

func main() {
	// Сравнение производительности
	start := time.Now()
	_ = concatWithPlus()
	durationPlus := time.Since(start)

	start = time.Now()
	_ = concatWithBuilder()
	durationBuilder := time.Since(start)

	fmt.Printf("Конкатенация через '+' (20k итераций): %v\n", durationPlus)
	fmt.Printf("Конкатенация через Builder (с Grow):   %v\n", durationBuilder)
	if durationBuilder > 0 {
		fmt.Printf("Ускорение в %.1f раз(а)!\n", float64(durationPlus)/float64(durationBuilder))
	}
}
