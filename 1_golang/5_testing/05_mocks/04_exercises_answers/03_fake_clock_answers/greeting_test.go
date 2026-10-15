// Упражнение 03: подмена часов (Решение).
package greeting

import (
	"testing"
	"time"
)

type fixedClock struct{ t time.Time }

func (f fixedClock) Now() time.Time { return f.t }

func at(hour int) Clock {
	return fixedClock{time.Date(2024, 1, 1, hour, 0, 0, 0, time.UTC)}
}

func TestGreeting(t *testing.T) {
	tests := []struct {
		hour int
		want string
	}{
		{0, "Доброе утро"},
		{11, "Доброе утро"},
		{12, "Добрый день"},
		{17, "Добрый день"},
		{18, "Добрый вечер"},
		{23, "Добрый вечер"},
	}
	for _, tt := range tests {
		if got := Greeting(at(tt.hour)); got != tt.want {
			t.Errorf("hour %d: Greeting = %q, want %q", tt.hour, got, tt.want)
		}
	}
}
