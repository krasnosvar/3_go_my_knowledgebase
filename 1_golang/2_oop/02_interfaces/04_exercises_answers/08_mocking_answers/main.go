// Упражнение 08: подмена зависимости (mock) в тесте (Решение).
package main

import (
	"fmt"
	"time"
)

// Clock — источник текущего времени.
type Clock interface {
	Now() time.Time
}

// fixedClock всегда возвращает одно и то же время.
type fixedClock struct {
	t time.Time
}

func (f fixedClock) Now() time.Time { return f.t }

func greet(c Clock) string {
	h := c.Now().Hour()
	switch {
	case h < 12:
		return "Доброе утро"
	case h < 18:
		return "Добрый день"
	default:
		return "Добрый вечер"
	}
}

func main() {
	fmt.Println("--- Проверка 08_mocking (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	at := func(hour int) Clock {
		return fixedClock{t: time.Date(2024, 1, 1, hour, 0, 0, 0, time.UTC)}
	}
	cases := []struct {
		hour int
		want string
	}{
		{9, "Доброе утро"},
		{15, "Добрый день"},
		{20, "Добрый вечер"},
	}
	for _, c := range cases {
		got := greet(at(c.hour))
		check(got == c.want, fmt.Sprintf("час %d: ожидалось %q, получено %q", c.hour, c.want, got))
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
