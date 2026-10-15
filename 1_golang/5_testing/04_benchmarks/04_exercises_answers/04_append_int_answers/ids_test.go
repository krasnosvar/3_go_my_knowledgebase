// Упражнение 04: strconv.Append* вместо fmt.
//
// Задача: перепишите FormatIDs в ids.go так, чтобы аллокаций было не
// больше двух: один []byte с запасом ёмкости + итоговое string(buf).
// Числа дописывайте через strconv.AppendInt. Сейчас на каждый id
// создаётся строка через fmt.Sprint, а потом ещё strings.Join.
// Тест менять не нужно — правьте только ids.go.
//
// Запуск:
//
//	go test ./03_exercises/04_append_int/
//	go test -bench=. -benchmem -run='^$' ./03_exercises/04_append_int/
package ids

import "testing"

func TestFormatIDsResult(t *testing.T) {
	if got := FormatIDs([]int{1, 22, 333}); got != "1,22,333" {
		t.Errorf("FormatIDs = %q, want %q", got, "1,22,333")
	}
	if got := FormatIDs(nil); got != "" {
		t.Errorf("FormatIDs(nil) = %q, want empty", got)
	}
}

func TestFormatIDsAllocs(t *testing.T) {
	ids := make([]int, 100)
	for i := range ids {
		ids[i] = 1000 + i
	}
	allocs := testing.AllocsPerRun(100, func() { _ = FormatIDs(ids) })
	if allocs > 2 {
		t.Errorf("FormatIDs делает %.0f аллокаций, нужно не больше 2", allocs)
	}
}
