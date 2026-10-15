// Упражнение 03: предвыделение ёмкости слайса.
//
// Задача: перепишите Squares в squares.go так, чтобы был ровно один
// make с нужной ёмкостью (make([]int, 0, n) или make([]int, n)).
// Сейчас слайс растёт через append с nil — runtime несколько раз
// выделяет новый массив и копирует данные.
// Тест менять не нужно — правьте только squares.go.
//
// Запуск:
//
//	go test ./03_exercises/03_prealloc_slice/
//	go test -bench=. -benchmem -run='^$' ./03_exercises/03_prealloc_slice/
package prealloc

import "testing"

func TestSquaresResult(t *testing.T) {
	got := Squares(4)
	want := []int{0, 1, 4, 9}
	if len(got) != len(want) {
		t.Fatalf("len = %d, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("Squares(4)[%d] = %d, want %d", i, got[i], want[i])
		}
	}
}

func TestSquaresAllocs(t *testing.T) {
	allocs := testing.AllocsPerRun(100, func() { sink = Squares(1000) }) // sink: результат «утекает», как в реальном коде
	if allocs > 1 {
		t.Errorf("Squares(1000) делает %.0f аллокаций, нужно не больше 1", allocs)
	}
}

var sink []int

func BenchmarkSquares(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sink = Squares(1000)
	}
}
