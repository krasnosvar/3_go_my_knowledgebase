package main

import (
	"fmt"
	"sort"
)

/*
	ОТВЕТЫ К ЗАДАНИЮ 4
*/

func main() {
	fmt.Println("--- Task 1: Вывод ключей по алфавиту ---")
	task1_sorted_keys()

	fmt.Println("\n--- Task 2: Подсчет частоты слов ---")
	task2_word_freq()
}

func task1_sorted_keys() {
	// ================= РЕШЕНИЕ: Task 1 =================
	scores := map[string]int{
		"Charlie": 80,
		"Alice":   95,
		"Bob":     85,
	}

	// 1. Создаем слайс нужной длины
	keys := make([]string, 0, len(scores))

	// 2. Заполняем его
	for k := range scores {
		keys = append(keys, k)
	}

	// 3. Сортируем
	sort.Strings(keys)

	// 4. Выводим
	for _, k := range keys {
		fmt.Printf("%s: %d\n", k, scores[k])
	}
	// ====================================================
}

func task2_word_freq() {
	// ================= РЕШЕНИЕ: Task 2 =================
	words := []string{"go", "rust", "go", "java", "go", "rust"}

	freq := make(map[string]int)
	for _, word := range words {
		freq[word]++
	}

	fmt.Println("Частоты:", freq)

	// Продвинутый уровень: Сортировка по значению
	type kv struct {
		Key   string
		Value int
	}
	var ss []kv
	for k, v := range freq {
		ss = append(ss, kv{k, v})
	}

	// Сортируем по убыванию значений
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	fmt.Println("Топ слова (сортировка по частоте):")
	for _, kv := range ss {
		fmt.Printf("%s: %d\n", kv.Key, kv.Value)
	}
	// ====================================================
}
