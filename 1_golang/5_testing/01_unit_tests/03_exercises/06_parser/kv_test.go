// Упражнение 06: тест парсера.
//
// Задача: напишите тесты для Parse (kv.go):
//   - "a=1; b = 2" -> map[a:1 b:2] (сравнение через maps.Equal);
//   - "" -> пустая map без ошибки;
//   - "a=1;;b=2" -> пустой сегмент пропускается;
//   - "novalue" и "=1" -> errors.Is(err, ErrBadPair).
//
// Запуск: go test ./03_exercises/06_parser/
package kv

import "testing"

func TestParse(t *testing.T) {
	t.Fatal("TODO: проверьте успешный разбор и ошибочные входы")
}
