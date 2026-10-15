// Package grade — пример: таблица в виде map и подтесты.
package grade

// Grade переводит баллы 0..100 в оценку.
func Grade(score int) string {
	switch {
	case score >= 90:
		return "A"
	case score >= 75:
		return "B"
	case score >= 60:
		return "C"
	}
	return "F"
}
