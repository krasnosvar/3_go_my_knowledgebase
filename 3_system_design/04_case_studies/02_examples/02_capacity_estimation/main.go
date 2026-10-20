// Пример 02: калькулятор оценок «на салфетке» (back-of-the-envelope).
//
// Считает по методике из 01_theory/02_estimations.md:
//   - средний и пиковый QPS из DAU и числа запросов на пользователя;
//   - объём хранилища из размера записи, числа записей в сутки и срока
//   хранения;
//   - пропускную способность (bandwidth) из QPS и размера ответа.
//
// Все числа зашиты как пример (детерминированно), программа печатает
// расчёт и завершается.
//
// Запуск: go run ./02_examples/02_capacity_estimation/main.go
package main

import "fmt"

const secondsPerDay = 86_400 // ≈ 10^5

// qps возвращает средний и пиковый QPS.
// dau — дневная аудитория, reqPerUser — запросов на пользователя в сутки,
// peakFactor — во сколько раз пик выше среднего (обычно 2..3).
func qps(dau, reqPerUser int, peakFactor float64) (avg, peak float64) {
	perDay := float64(dau) * float64(reqPerUser)
	avg = perDay / secondsPerDay
	peak = avg * peakFactor
	return avg, peak
}

// storageBytes возвращает объём хранилища в байтах:
// размер записи × записей в сутки × дней хранения.
func storageBytes(recordBytes, recordsPerDay, days int) int64 {
	return int64(recordBytes) * int64(recordsPerDay) * int64(days)
}

// bandwidthBytesPerSec возвращает пропускную способность: QPS × размер ответа.
func bandwidthBytesPerSec(reqPerSec float64, respBytes int) float64 {
	return reqPerSec * float64(respBytes)
}

// humanBytes форматирует байты в удобные единицы (KB/MB/GB/TB).
func humanBytes(b float64) string {
	const unit = 1000.0 // округляем 1024≈1000, как в прикидках
	units := []string{"B", "KB", "MB", "GB", "TB", "PB"}
	i := 0
	for b >= unit && i < len(units)-1 {
		b /= unit
		i++
	}
	return fmt.Sprintf("%.2f %s", b, units[i])
}

func main() {
	fmt.Println("--- Калькулятор оценок нагрузки ---")

	// Вводные (пример: сокращатель ссылок, read-heavy).
	const (
		dau           = 10_000_000  // 10 млн активных в сутки
		reqPerUser    = 20          // запросов на пользователя в сутки
		peakFactor    = 2.0         // пик = 2× среднего
		recordsPerDay = 100_000_000 // новых записей в сутки (100 млн)
		recordBytes   = 500         // размер одной записи, байт
		retentionDays = 5 * 365     // храним 5 лет
		respBytes     = 1000        // размер ответа на чтение, байт
	)

	avgQPS, peakQPS := qps(dau, reqPerUser, peakFactor)
	fmt.Printf("\n[QPS] DAU=%d, %d запр./польз./сутки, пик ×%.0f\n", dau, reqPerUser, peakFactor)
	fmt.Printf("  запросов в сутки: %d\n", dau*reqPerUser)
	fmt.Printf("  средний QPS: %.0f\n", avgQPS)
	fmt.Printf("  пиковый QPS: %.0f\n", peakQPS)

	total := storageBytes(recordBytes, recordsPerDay, retentionDays)
	perDay := storageBytes(recordBytes, recordsPerDay, 1)
	fmt.Printf("\n[Хранилище] запись=%d Б, %d записей/сутки, хранение=%d дней\n", recordBytes, recordsPerDay, retentionDays)
	fmt.Printf("  в сутки: %s\n", humanBytes(float64(perDay)))
	fmt.Printf("  всего:   %s\n", humanBytes(float64(total)))

	bw := bandwidthBytesPerSec(peakQPS, respBytes)
	fmt.Printf("\n[Bandwidth] пиковый QPS=%.0f, ответ=%d Б\n", peakQPS, respBytes)
	fmt.Printf("  исходящий трафик: %s/с (%.0f Mbps)\n", humanBytes(bw), bw*8/1e6)

	fmt.Println("\nГотово, программа завершается.")
}
