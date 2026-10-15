// Упражнение 03: setup и cleanup (Решение).
package config

import (
	"errors"
	"io/fs"
	"maps"
	"path/filepath"
	"testing"
)

// newConfigPath — setup: путь к файлу во временном каталоге теста.
func newConfigPath(t *testing.T) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "app.conf") // каталог удалит сам testing
	t.Cleanup(func() { t.Log("cleanup:", path) })
	return path
}

func TestConfig(t *testing.T) {
	t.Run("roundtrip", func(t *testing.T) {
		path := newConfigPath(t)
		want := map[string]string{"host": "localhost", "port": "8080"}
		if err := Save(path, want); err != nil {
			t.Fatalf("Save: %v", err)
		}
		got, err := Load(path)
		if err != nil {
			t.Fatalf("Load: %v", err)
		}
		if !maps.Equal(got, want) {
			t.Errorf("Load = %v, want %v", got, want)
		}
	})

	t.Run("missing file", func(t *testing.T) {
		_, err := Load(filepath.Join(t.TempDir(), "nope.conf"))
		if !errors.Is(err, fs.ErrNotExist) {
			t.Errorf("err = %v, want fs.ErrNotExist", err)
		}
	})
}
