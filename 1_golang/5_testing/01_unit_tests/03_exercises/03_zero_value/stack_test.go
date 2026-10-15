// Упражнение 03: нулевое значение типа.
//
// Задача: напишите TestStackZeroValue, который проверяет, что
// `var s Stack` (без конструктора) пригоден к работе:
//   - s.Len() == 0;
//   - s.Pop() на пустом стеке возвращает ok == false и не паникует;
//   - после Push(1), Push(2) порядок Pop — 2, затем 1 (LIFO), Len уменьшается.
//
// Запуск: go test ./03_exercises/03_zero_value/
package stack

import "testing"

func TestStackZeroValue(t *testing.T) {
	t.Fatal("TODO: проверьте работу нулевого значения Stack")
}
