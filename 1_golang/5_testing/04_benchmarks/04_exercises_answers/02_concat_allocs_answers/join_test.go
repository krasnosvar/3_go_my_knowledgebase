// Упражнение 02: конкатенация без лишних аллокаций.
//
// Задача: перепишите Join в join.go так, чтобы он делал не больше одной
// аллокации: strings.Builder + b.Grow(итоговая длина) и запись частей.
// Сейчас строка собирается через += — каждая итерация создаёт новую строку.
// Тест менять не нужно — правьте только join.go.
//
// Запуск:
//
//	go test ./03_exercises/02_concat_allocs/
//	go test -bench=. -benchmem -run='^$' ./03_exercises/02_concat_allocs/
package concat

import (
	"strings"
	"testing"
)

var parts = strings.Fields("the quick brown fox jumps over the lazy dog again and again")

func TestJoinResult(t *testing.T) {
	if got, want := Join(parts, " "), strings.Join(parts, " "); got != want {
		t.Errorf("Join = %q, want %q", got, want)
	}
	if got := Join(nil, ","); got != "" {
		t.Errorf("Join(nil) = %q, want empty", got)
	}
}

func TestJoinAllocs(t *testing.T) {
	allocs := testing.AllocsPerRun(100, func() { sink = Join(parts, " ") })
	if allocs > 1 {
		t.Errorf("Join делает %.0f аллокаций, нужно не больше 1", allocs)
	}
}

var sink string

func BenchmarkJoin(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sink = Join(parts, " ")
	}
}
