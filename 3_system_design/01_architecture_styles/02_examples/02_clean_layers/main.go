// Пример 02: маленький срез Clean Architecture.
//
// Слои в одном файле (для наглядности), зависимости направлены внутрь:
//
//	Entity (домен) <- Usecase (приложение) <- Adapter (in-memory) <- main
// (composition root)
//
// Repository — порт (интерфейс) в слое usecase; InMemoryUserRepo — адаптер,
// реализующий порт. Сценарий: создать пользователя и получить его обратно.
//
// Демонстрирует: правило зависимостей (внутрь), порт/адаптер,
// изоляцию бизнес-логики от хранилища.
package main

import (
	"errors"
	"fmt"
)

// ===== Слой 1: Entity (домен). Ничего не знает о БД и о usecase. =====

type User struct {
	ID   string
	Name string
}

// Validate — бизнес-правило домена.
func (u User) Validate() error {
	if u.ID == "" {
		return errors.New("user: пустой ID")
	}
	if u.Name == "" {
		return errors.New("user: пустое имя")
	}
	return nil
}

// ===== Слой 2: Usecase (приложение) + порт. =====
// Порт объявлен здесь, на стороне потребителя (инверсия зависимостей).

var ErrNotFound = errors.New("user: не найден")

// UserRepository — порт. Usecase зависит от этого интерфейса, а не от адаптера.
type UserRepository interface {
	Save(u User) error
	FindByID(id string) (User, error)
}

// UserUsecase оркеструет домен и порт; про конкретное хранилище не знает.
type UserUsecase struct {
	repo UserRepository
}

func NewUserUsecase(repo UserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}

// Create — сценарий создания: валидирует домен и сохраняет через порт.
func (uc *UserUsecase) Create(id, name string) (User, error) {
	u := User{ID: id, Name: name}
	if err := u.Validate(); err != nil {
		return User{}, err
	}
	if err := uc.repo.Save(u); err != nil {
		return User{}, err
	}
	return u, nil
}

// GetByID — сценарий чтения через порт.
func (uc *UserUsecase) GetByID(id string) (User, error) {
	return uc.repo.FindByID(id)
}

// ===== Слой 3: Adapter (in-memory). Реализует порт из слоя usecase. =====

type InMemoryUserRepo struct {
	data map[string]User
}

func NewInMemoryUserRepo() *InMemoryUserRepo {
	return &InMemoryUserRepo{data: make(map[string]User)}
}

func (r *InMemoryUserRepo) Save(u User) error {
	r.data[u.ID] = u
	return nil
}

func (r *InMemoryUserRepo) FindByID(id string) (User, error) {
	u, ok := r.data[id]
	if !ok {
		return User{}, ErrNotFound
	}
	return u, nil
}

// ===== main: composition root. Связывает слои снаружи внутрь. =====

func main() {
	fmt.Println("--- Clean Architecture: срез слоёв ---")

	// Собираем граф: адаптер -> usecase.
	uc := NewUserUsecase(NewInMemoryUserRepo())

	// Сценарий 1: создать пользователя.
	created, err := uc.Create("u1", "Alice")
	if err != nil {
		fmt.Printf("создание не удалось: %v\n", err)
		return
	}
	fmt.Printf("создан: %+v\n", created)

	// Сценарий 2: получить пользователя обратно.
	got, err := uc.GetByID("u1")
	if err != nil {
		fmt.Printf("чтение не удалось: %v\n", err)
		return
	}
	fmt.Printf("получен: %+v\n", got)

	// Сценарий 3: домен отвергает невалидные данные (бизнес-правило).
	if _, err := uc.Create("", "Bob"); err != nil {
		fmt.Printf("валидация сработала: %v\n", err)
	}

	fmt.Println("Готово, программа завершается.")
}
