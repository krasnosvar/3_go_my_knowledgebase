// Упражнение 01: подтесты через t.Run.
//
// Задача: напишите TestClass: таблица случаев + t.Run(tt.name, ...) на
// каждый. Внутри подтеста используйте t.Errorf. Случаи: 100, 200, 301,
// 404, 503, 99, 600.
// Проверьте, что можно запустить один подтест:
//
//	go test -run 'TestClass/not_found' -v ./03_exercises/01_t_run/
//
// (пробелы в имени подтеста Go заменяет на '_').
//
// Запуск: go test ./03_exercises/01_t_run/
package status

import "testing"

func TestClass(t *testing.T) {
	t.Fatal("TODO: подтесты через t.Run")
}
