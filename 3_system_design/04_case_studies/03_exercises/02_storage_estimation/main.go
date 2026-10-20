// Упражнение 02: оценка объёма хранилища.
//
// Задача: посчитайте требуемый объём хранилища в байтах по формуле
//
//	объём = размер_записи × записей_в_сутки × дней_хранения
//
// (см. 01_theory/02_estimations.md).
//
// Запуск: go run ./03_exercises/02_storage_estimation/main.go
package main

import "fmt"

// storageBytes возвращает объём хранилища в байтах.
// recordBytes — размер одной записи, recordsPerDay — записей в сутки,
// days — срок хранения в днях.
func storageBytes(recordBytes, recordsPerDay, days int) int64 {
	// TODO: перемножьте параметры (используйте int64, чтобы не переполнить).
	return 0
}

func main() {
	fmt.Println("--- Проверка 02_storage_estimation ---")

	cases := []struct {
		recordBytes   int
		recordsPerDay int
		days          int
		want          int64
	}{
		{500, 100_000_000, 1, 50_000_000_000},           // 50 GB за сутки
		{500, 100_000_000, 365 * 5, 91_250_000_000_000}, // ~91 TB за 5 лет
		{1000, 10_000_000, 365, 3_650_000_000_000},      // ~3.65 TB метаданных за год
		{200, 0, 100, 0}, // нет записей — ноль
	}

	allOK := true
	for _, c := range cases {
		got := storageBytes(c.recordBytes, c.recordsPerDay, c.days)
		if got != c.want {
			allOK = false
			fmt.Printf("FAIL: storageBytes(%d, %d, %d) = %d, ожидалось %d\n",
				c.recordBytes, c.recordsPerDay, c.days, got, c.want)
		}
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
