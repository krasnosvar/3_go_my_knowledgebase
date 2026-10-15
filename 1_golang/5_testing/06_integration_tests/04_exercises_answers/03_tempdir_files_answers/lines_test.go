// Упражнение 03: реальные файлы во временном каталоге (Решение).
package lines

import (
	"errors"
	"io/fs"
	"path/filepath"
	"slices"
	"testing"
)

func TestRoundTrip(t *testing.T) {
	path := filepath.Join(t.TempDir(), "data.txt")
	want := []string{"first", "second", "третья"}
	if err := Save(path, want); err != nil {
		t.Fatalf("Save: %v", err)
	}
	got, err := Load(path)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if !slices.Equal(got, want) {
		t.Errorf("Load = %q, want %q", got, want)
	}
}

func TestLoadMissing(t *testing.T) {
	_, err := Load(filepath.Join(t.TempDir(), "missing.txt"))
	if !errors.Is(err, fs.ErrNotExist) {
		t.Errorf("err = %v, want fs.ErrNotExist", err)
	}
}
