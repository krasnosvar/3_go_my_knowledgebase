// Package greeting — код под тестом для упражнения 03.
package greeting

import "time"

// Clock — источник времени. В проде — realClock, в тестах — фиксированное время.
type Clock interface {
	Now() time.Time
}

type realClock struct{}

func (realClock) Now() time.Time { return time.Now() }

// RealClock возвращает системные часы.
func RealClock() Clock { return realClock{} }

// Greeting выбирает приветствие по часу: <12 — утро, <18 — день, иначе вечер.
func Greeting(c Clock) string {
	h := c.Now().Hour()
	switch {
	case h < 12:
		return "Доброе утро"
	case h < 18:
		return "Добрый день"
	}
	return "Добрый вечер"
}
