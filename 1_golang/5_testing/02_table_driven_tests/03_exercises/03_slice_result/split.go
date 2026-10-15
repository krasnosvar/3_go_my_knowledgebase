// Package split — код под тестом для упражнения 03.
package split

import "strings"

// SplitTrim делит s по sep, обрезает пробелы у частей и выбрасывает пустые.
func SplitTrim(s, sep string) []string {
	var out []string
	for _, p := range strings.Split(s, sep) {
		if p = strings.TrimSpace(p); p != "" {
			out = append(out, p)
		}
	}
	return out
}
