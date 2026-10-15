// Упражнение 02: тест веток ошибок.
//
// Задача: напишите тесты для ParseAge (age.go), покрывающие ВСЕ ветки:
//   - TestParseAgeOK: "42" -> 42, err == nil;
//   - TestParseAgeNotNumber: "abc" -> errors.Is(err, ErrNotNumber);
//   - TestParseAgeOutOfRange: "-1" и "151" -> errors.Is(err, ErrOutOfRange).
//
// Сравнивайте ошибки через errors.Is, а не по тексту: текст обёрнут fmt.Errorf.
// Проверить покрытие: go test -cover ./03_exercises/02_error_branch/
//
// Запуск: go test ./03_exercises/02_error_branch/
package age

import "testing"

func TestParseAge(t *testing.T) {
	t.Fatal("TODO: напишите тесты на успешный разбор и обе ветки ошибок")
}
