// Упражнение 03: setup и cleanup.
//
// Задача: напишите TestConfig с двумя подтестами:
//   - "roundtrip": Save во временный файл, затем Load -> та же map (maps.Equal);
//   - "missing file": Load несуществующего пути -> errors.Is(err, fs.ErrNotExist).
//
// Для setup используйте хелпер newConfigPath(t) с t.Helper(), который
// берёт каталог из t.TempDir() (удаляется автоматически после теста) и
// регистрирует t.Cleanup с t.Log("cleanup: ...") — увидите порядок в -v.
//
// Запуск: go test ./03_exercises/03_cleanup/
package config

import "testing"

func TestConfig(t *testing.T) {
	t.Fatal("TODO: подтесты с t.TempDir и t.Cleanup")
}
