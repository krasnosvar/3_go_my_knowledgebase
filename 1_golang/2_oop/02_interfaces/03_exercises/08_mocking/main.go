// Упражнение 08: подмена зависимости (mock) в тесте.
//
// Задача: greet зависит от интерфейса Clock, а не от time.Now() напрямую —
// поэтому в тесте время можно зафиксировать.
//   - реализуйте fixedClock.Now(): верните сохранённое поле t;
//   - реализуйте greet: час < 12 -> "Доброе утро", < 18 -> "Добрый день",
//     иначе -> "Добрый вечер".
//
// Запуск: go run ./03_exercises/08_mocking/main.go
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

// Now возвращает зафиксированное время.
func (f fixedClock) Now() time.Time {
	// TODO: верните f.t.
	return time.Time{}
}

// greet возвращает приветствие по времени суток.
func greet(c Clock) string {
	// TODO: h := c.Now().Hour(); ветвление по часу.
	return ""
}

func main() {
	fmt.Println("--- Проверка 08_mocking ---")

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
