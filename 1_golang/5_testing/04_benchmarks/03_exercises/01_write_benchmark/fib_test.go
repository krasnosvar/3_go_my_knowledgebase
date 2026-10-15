// Упражнение 01: первый бенчмарк.
//
// Задача: напишите BenchmarkFib с под-бенчмарками для n = 10, 20, 40:
//
//	for _, n := range []int{10, 20, 40} {
//	    b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) { ... })
//	}
//
// Внутри — b.ReportAllocs() и цикл `for i := 0; i < b.N; i++`. Результат
// сохраняйте в переменную пакета (sink), чтобы компилятор не выкинул вызов.
// Затем удалите TestTODO.
//
// Запуск:
//
//	go test ./03_exercises/01_write_benchmark/                  — тесты
//	go test -bench=. -run='^$' ./03_exercises/01_write_benchmark/ — бенчмарки
package fib

import "testing"

func TestFib(t *testing.T) {
	if got := Fib(10); got != 55 {
		t.Errorf("Fib(10) = %d, want 55", got)
	}
}

func TestTODO(t *testing.T) {
	t.Fatal("TODO: напишите BenchmarkFib с под-бенчмарками и удалите TestTODO")
}
