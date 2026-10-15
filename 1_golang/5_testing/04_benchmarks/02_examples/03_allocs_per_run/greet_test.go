package greet

import "testing"

// AllocsPerRun превращает «бенчмарк» в обычный тест: регрессия по
// аллокациям ломает go test, а не только цифры в -bench.
func TestGreetAllocs(t *testing.T) {
	if got := Greet("Go"); got != "Hello, Go!" {
		t.Fatalf("Greet = %q", got)
	}
	allocs := testing.AllocsPerRun(100, func() { _ = Greet("Go") })
	if allocs > 1 {
		t.Errorf("Greet: %.0f аллокаций, ожидалась 1", allocs)
	}
}
