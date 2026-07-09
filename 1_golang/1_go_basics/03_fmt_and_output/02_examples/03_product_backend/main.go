package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

var ErrRecordNotFound = errors.New("record not found")

// Имитация базы данных с форматированием ошибок
func getUserByID(id int) error {
	// Оборачиваем ошибку ErrRecordNotFound через %w, чтобы сохранить ее тип
	return fmt.Errorf("database query failed for user ID %d: %w", id, ErrRecordNotFound)
}

func main() {
	fmt.Println("--- Имитация ошибок бэкенда (fmt.Errorf) ---")
	err := getUserByID(42)
	fmt.Println("Ошибка:", err)

	// Проверяем, что исходная ошибка ErrRecordNotFound сохранилась внутри (unwrap)
	if errors.Is(err, ErrRecordNotFound) {
		fmt.Println("Успешно определено: Запись не найдена в базе данных!")
	}

	fmt.Println("\n--- Сравнение производительности конкатенации ---")

	// Сравним скорость склеивания 10,000 строк
	iterations := 10000

	// Вариант 1: Через fmt.Sprintf
	start := time.Now()
	var s1 string
	for i := 0; i < iterations; i++ {
		s1 = fmt.Sprintf("%sitem%d,", s1, i)
	}
	durationSprintf := time.Since(start)

	// Вариант 2: Через strings.Builder (минимизация аллокаций в куче)
	start = time.Now()
	var builder strings.Builder
	// Предаллокация буфера под строки (для максимальной скорости)
	builder.Grow(iterations * 10)
	for i := 0; i < iterations; i++ {
		builder.WriteString("item")
		// Использование встроенного WriteString или strconv.Itoa быстрее чем fmt
		builder.WriteString(fmt.Sprint(i))
		builder.WriteString(",")
	}
	s2 := builder.String()
	durationBuilder := time.Since(start)

	fmt.Printf("Время работы fmt.Sprintf:   %v (Длина строки: %d)\n", durationSprintf, len(s1))
	fmt.Printf("Время работы strings.Builder: %v (Длина строки: %d)\n", durationBuilder, len(s2))
	fmt.Printf("strings.Builder быстрее в %.1f раз!\n", float64(durationSprintf)/float64(durationBuilder))
}
