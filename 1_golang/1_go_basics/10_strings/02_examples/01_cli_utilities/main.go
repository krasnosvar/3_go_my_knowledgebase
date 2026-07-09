package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// Кириллические символы и эмодзи занимают более 1 байта в UTF-8
	filename := "отчёт_2026_📊.txt"

	// 1. Длина строки через len() возвращает байты
	byteLen := len(filename)

	// 2. Длина строки через RuneCountInString возвращает реальные символы (руны)
	runeLen := utf8.RuneCountInString(filename)

	fmt.Printf("Имя файла: %s\n", filename)
	fmt.Printf("Длина в байтах (len):       %d\n", byteLen)
	fmt.Printf("Количество символов (runes): %d\n", runeLen)

	// 3. Разница при обходе циклом
	fmt.Println("\n--- Побайтовый обход (обычный for): ---")
	for i := 0; i < 6; i++ { // Выведем только первые несколько байт
		fmt.Printf("Байт %d: 0x%x\n", i, filename[i])
	}

	fmt.Println("\n--- Порунно с декодированием (for range): ---")
	count := 0
	for idx, r := range filename {
		if count >= 4 {
			break
		}
		fmt.Printf("Индекс байта: %d, Символ: %c (код: %U)\n", idx, r, r)
		count++
	}
}
