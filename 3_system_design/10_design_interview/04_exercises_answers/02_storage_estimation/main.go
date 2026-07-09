// Упражнение 02: оценка объёма хранилища (Решение).
package main

import "fmt"

func storageBytes(recordBytes, recordsPerDay, days int) int64 {
	return int64(recordBytes) * int64(recordsPerDay) * int64(days)
}

func main() {
	fmt.Println("--- Проверка 02_storage_estimation (Решение) ---")

	cases := []struct {
		recordBytes   int
		recordsPerDay int
		days          int
		want          int64
	}{
		{500, 100_000_000, 1, 50_000_000_000},
		{500, 100_000_000, 365 * 5, 91_250_000_000_000},
		{1000, 10_000_000, 365, 3_650_000_000_000},
		{200, 0, 100, 0},
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
