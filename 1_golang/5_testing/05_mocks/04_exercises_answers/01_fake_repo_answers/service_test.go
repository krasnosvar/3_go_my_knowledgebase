// Упражнение 01: fake-репозиторий (Решение).
package users

import (
	"errors"
	"testing"
)

// fakeRepo — in-memory реализация Repo для тестов.
type fakeRepo struct {
	saved []User
}

func (f *fakeRepo) Save(u User) error {
	f.saved = append(f.saved, u)
	return nil
}

func TestRegisterSaves(t *testing.T) {
	repo := &fakeRepo{}
	if err := NewService(repo).Register("alice"); err != nil {
		t.Fatalf("Register: %v", err)
	}
	if len(repo.saved) != 1 || repo.saved[0].Name != "alice" {
		t.Errorf("saved = %v, want [{alice}]", repo.saved)
	}
}

func TestRegisterEmptyName(t *testing.T) {
	repo := &fakeRepo{}
	err := NewService(repo).Register("")
	if !errors.Is(err, ErrEmptyName) {
		t.Errorf("err = %v, want ErrEmptyName", err)
	}
	if len(repo.saved) != 0 {
		t.Errorf("Save не должен вызываться, saved = %v", repo.saved)
	}
}
