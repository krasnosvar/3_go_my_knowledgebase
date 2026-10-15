package counter

import (
	"os"
	"path/filepath"
	"testing"
)

// setupFile — хелпер: создаёт путь во временном каталоге.
func setupFile(t *testing.T) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "log.txt")
	t.Cleanup(func() { t.Logf("teardown %s", path) }) // выполнится после подтеста
	return path
}

func TestAppendLine(t *testing.T) {
	t.Run("creates file", func(t *testing.T) {
		path := setupFile(t)
		if err := AppendLine(path, "a"); err != nil {
			t.Fatal(err)
		}
		data, _ := os.ReadFile(path)
		if string(data) != "a\n" {
			t.Errorf("content = %q", data)
		}
	})
	t.Run("appends", func(t *testing.T) {
		path := setupFile(t)
		_ = AppendLine(path, "a")
		_ = AppendLine(path, "b")
		data, _ := os.ReadFile(path)
		if string(data) != "a\nb\n" {
			t.Errorf("content = %q", data)
		}
	})
}
