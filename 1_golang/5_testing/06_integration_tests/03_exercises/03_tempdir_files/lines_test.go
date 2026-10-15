// Упражнение 03: реальные файлы во временном каталоге.
//
// Задача: напишите тесты с настоящей файловой системой:
//   - TestRoundTrip: Save в filepath.Join(t.TempDir(), "data.txt"), затем
//     Load -> тот же слайс (slices.Equal);
//   - TestLoadMissing: Load несуществующего файла -> errors.Is(err, fs.ErrNotExist).
//
// t.TempDir() уникален для теста и удаляется автоматически — тесты не
// мешают друг другу и не оставляют мусора.
//
// Запуск: go test ./03_exercises/03_tempdir_files/
package lines

import "testing"

func TestLines(t *testing.T) {
	t.Fatal("TODO: round-trip через t.TempDir")
}
