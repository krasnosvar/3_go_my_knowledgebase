// Упражнение 04: spy: проверка вызовов.
//
// Задача: напишите spyNotifier, который записывает получателей в слайс
// calls []string и может «сломаться» на пользователе failOn (вернуть ошибку).
// Тесты:
//   - TestBroadcastAll: users a, b, c -> calls == [a b c], err == nil;
//   - TestBroadcastStopsOnError: failOn = "b" -> err != nil и calls == [a b]
//     (c не вызывался — проверка поведения, а не только результата).
//
// Не переусердствуйте: проверяйте то, что важно для контракта.
//
// Запуск: go test ./03_exercises/04_spy_calls/
package broadcast

import "testing"

func TestBroadcast(t *testing.T) {
	t.Fatal("TODO: spy с записью вызовов")
}
