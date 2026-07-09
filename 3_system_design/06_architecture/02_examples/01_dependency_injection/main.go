// Пример 01: Dependency Injection через интерфейс.
//
// Usecase зависит не от конкретной БД, а от интерфейса Repository.
// В main подставляется in-memory реализация. Показана подмена
// зависимости: тот же usecase работает поверх второй реализации
// (с префиксом) без единой правки самого usecase.
//
// Демонстрирует: constructor injection, "accept interfaces",
// composition root в main, заменяемость реализации.
package main

import "fmt"

// Repository — порт (интерфейс), объявленный на стороне потребителя.
// Usecase знает только про этот контракт, но не про конкретное хранилище.
type Repository interface {
	Save(key, value string) error
	Get(key string) (string, bool)
}

// Usecase — бизнес-сценарий, зависит от абстракции Repository.
type Usecase struct {
	repo Repository
}

// NewUsecase — конструктор: зависимость приходит извне (DI).
func NewUsecase(repo Repository) *Usecase {
	return &Usecase{repo: repo}
}

// Store сохраняет значение, Load читает — вся работа идёт через порт.
func (u *Usecase) Store(key, value string) error  { return u.repo.Save(key, value) }
func (u *Usecase) Load(key string) (string, bool) { return u.repo.Get(key) }

// --- Реализация 1: обычное in-memory хранилище. ---

type InMemoryRepo struct {
	data map[string]string
}

func NewInMemoryRepo() *InMemoryRepo {
	return &InMemoryRepo{data: make(map[string]string)}
}

func (r *InMemoryRepo) Save(key, value string) error {
	r.data[key] = value
	return nil
}

func (r *InMemoryRepo) Get(key string) (string, bool) {
	v, ok := r.data[key]
	return v, ok
}

// --- Реализация 2: другое in-memory хранилище, добавляющее префикс.
// Нужна лишь для демонстрации подмены зависимости. ---

type PrefixRepo struct {
	prefix string
	data   map[string]string
}

func NewPrefixRepo(prefix string) *PrefixRepo {
	return &PrefixRepo{prefix: prefix, data: make(map[string]string)}
}

func (r *PrefixRepo) Save(key, value string) error {
	r.data[key] = r.prefix + value
	return nil
}

func (r *PrefixRepo) Get(key string) (string, bool) {
	v, ok := r.data[key]
	return v, ok
}

func main() {
	fmt.Println("--- Dependency Injection через интерфейс ---")

	// Composition root: собираем граф зависимостей здесь.
	// Подставляем первую реализацию.
	uc := NewUsecase(NewInMemoryRepo())
	_ = uc.Store("user:1", "Alice")
	if v, ok := uc.Load("user:1"); ok {
		fmt.Printf("InMemoryRepo: user:1 = %q\n", v)
	}

	// Подмена зависимости: тот же usecase, другая реализация.
	// Код usecase не меняется — меняется только то, что передали в конструктор.
	uc2 := NewUsecase(NewPrefixRepo("v2:"))
	_ = uc2.Store("user:1", "Alice")
	if v, ok := uc2.Load("user:1"); ok {
		fmt.Printf("PrefixRepo:   user:1 = %q\n", v)
	}

	fmt.Println("Готово, программа завершается.")
}
