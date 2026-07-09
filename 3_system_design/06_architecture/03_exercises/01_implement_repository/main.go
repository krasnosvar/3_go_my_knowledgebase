// Упражнение 01: реализовать порт Repository и usecase поверх него.
//
// Задача:
//  1. Реализуйте in-memory тип InMemoryRepo, удовлетворяющий интерфейсу
//     Repository (методы Save и Get).
//  2. Реализуйте метод (*Usecase).StoreAndLoad: сохранить (key, value)
//     через repo, затем прочитать по key и вернуть значение.
//
// Usecase уже зависит от интерфейса Repository — ваша задача подставить
// рабочую реализацию и связать их.
//
// Запуск: go run ./03_exercises/01_implement_repository/main.go
package main

import "fmt"

// Repository — порт (не менять).
type Repository interface {
	Save(key, value string) error
	Get(key string) (string, bool)
}

// Usecase зависит от порта (не менять поле).
type Usecase struct {
	repo Repository
}

func NewUsecase(repo Repository) *Usecase {
	return &Usecase{repo: repo}
}

// StoreAndLoad сохраняет пару и читает её обратно.
func (u *Usecase) StoreAndLoad(key, value string) (string, bool) {
	// TODO: сохраните value по key через u.repo, затем верните u.repo.Get(key).
	return "", false
}

// InMemoryRepo — ваша реализация Repository.
type InMemoryRepo struct {
	// TODO: добавьте поле-хранилище (например, map[string]string).
}

func NewInMemoryRepo() *InMemoryRepo {
	// TODO: инициализируйте хранилище.
	return &InMemoryRepo{}
}

// Save и Get — заготовки: сейчас ничего не делают (тест печатает FAIL).
// TODO: реализуйте их поверх вашего хранилища.
func (r *InMemoryRepo) Save(key, value string) error {
	// TODO: сохраните value по key.
	return nil
}

func (r *InMemoryRepo) Get(key string) (string, bool) {
	// TODO: верните сохранённое значение и признак наличия.
	return "", false
}

func main() {
	fmt.Println("--- Проверка 01_implement_repository ---")

	uc := NewUsecase(NewInMemoryRepo())

	allOK := true
	if v, ok := uc.StoreAndLoad("k1", "v1"); !ok || v != "v1" {
		allOK = false
		fmt.Printf("FAIL: StoreAndLoad(k1,v1) = (%q,%v), ожидалось (\"v1\",true)\n", v, ok)
	}
	if v, ok := uc.StoreAndLoad("k2", "v2"); !ok || v != "v2" {
		allOK = false
		fmt.Printf("FAIL: StoreAndLoad(k2,v2) = (%q,%v), ожидалось (\"v2\",true)\n", v, ok)
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
