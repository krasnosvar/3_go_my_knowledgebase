// Package bytesize — код под тестом для упражнения 05.
package bytesize

import "fmt"

// FormatBytes форматирует размер в байтах: "512 B", "1.5 KB", "2.0 MB" (основание 1024).
func FormatBytes(n int64) string {
	const unit = 1024
	if n < unit {
		return fmt.Sprintf("%d B", n)
	}
	div, exp := int64(unit), 0
	for m := n / unit; m >= unit && exp < 3; m /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(n)/float64(div), "KMGT"[exp])
}
