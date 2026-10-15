// Упражнение 02: ошибки в таблице.
//
// Задача: напишите табличный TestDivide, где у каждого случая есть поля
// a, b, want int и wantErr error. В цикле:
//   - если wantErr != nil — проверьте errors.Is(err, tt.wantErr) и не
//     сравнивайте результат;
//   - иначе — err == nil и got == want.
//
// Случаи: 10/2, 7/2 (целочисленное деление), -9/3, 1/0.
//
// Запуск: go test ./03_exercises/02_error_cases/
package divide

import "testing"

func TestDivide(t *testing.T) {
	t.Fatal("TODO: таблица с полем wantErr")
}
