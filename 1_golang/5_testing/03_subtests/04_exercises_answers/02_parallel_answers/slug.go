// Package slug — код под тестом для упражнения 02.
package slug

import (
	"strings"
	"unicode"
)

// Slugify делает из заголовка slug: нижний регистр, пробелы -> '-',
// всё, кроме букв и цифр, выбрасывается.
func Slugify(s string) string {
	var b strings.Builder
	for _, w := range strings.Fields(strings.ToLower(s)) {
		var part strings.Builder
		for _, r := range w {
			if unicode.IsLetter(r) || unicode.IsDigit(r) {
				part.WriteRune(r)
			}
		}
		if part.Len() == 0 {
			continue
		}
		if b.Len() > 0 {
			b.WriteByte('-')
		}
		b.WriteString(part.String())
	}
	return b.String()
}
