// Упражнение 03: подмена часов.
//
// Задача: напишите fixedClock{t time.Time} с методом Now() и табличный
// TestGreeting для часов 0, 11, 12, 17, 18, 23 (обе стороны каждой границы).
// Время создавайте через time.Date(2024, 1, 1, hour, 0, 0, 0, time.UTC).
// Тест с time.Now() внутри был бы недетерминированным — зависел бы от
// времени запуска.
//
// Запуск: go test ./03_exercises/03_fake_clock/
package greeting

import "testing"

func TestGreeting(t *testing.T) {
	t.Fatal("TODO: fixedClock и таблица по часам")
}
