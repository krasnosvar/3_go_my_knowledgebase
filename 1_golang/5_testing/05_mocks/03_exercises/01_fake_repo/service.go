// Package users — код под тестом для упражнения 01.
package users

import "errors"

// ErrEmptyName — пустое имя пользователя.
var ErrEmptyName = errors.New("empty name")

// User — сохраняемая сущность.
type User struct{ Name string }

// Repo — зависимость сервиса (в проде — БД).
type Repo interface {
	Save(u User) error
}

// Service — бизнес-логика регистрации.
type Service struct{ repo Repo }

// NewService создаёт сервис с переданным репозиторием.
func NewService(repo Repo) *Service { return &Service{repo: repo} }

// Register валидирует имя и сохраняет пользователя.
func (s *Service) Register(name string) error {
	if name == "" {
		return ErrEmptyName
	}
	return s.repo.Save(User{Name: name})
}
