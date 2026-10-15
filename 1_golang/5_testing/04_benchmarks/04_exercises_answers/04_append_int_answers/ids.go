// Package ids — код для упражнения 04 (Решение).
package ids

import "strconv"

// FormatIDs возвращает id через запятую: "1,22,333".
func FormatIDs(ids []int) string {
	if len(ids) == 0 {
		return ""
	}
	buf := make([]byte, 0, len(ids)*8) // запас: ~7 цифр + запятая на id
	for i, id := range ids {
		if i > 0 {
			buf = append(buf, ',')
		}
		buf = strconv.AppendInt(buf, int64(id), 10)
	}
	return string(buf) // вторая аллокация: копия в неизменяемую строку
}
