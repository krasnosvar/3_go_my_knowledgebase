// Упражнение 01: реализовать порт Repository и usecase (Решение).
package main

import "fmt"

type Repository interface {
	Save(key, value string) error
	Get(key string) (string, bool)
}

type Usecase struct {
	repo Repository
}

func NewUsecase(repo Repository) *Usecase {
	return &Usecase{repo: repo}
}

func (u *Usecase) StoreAndLoad(key, value string) (string, bool) {
	_ = u.repo.Save(key, value)
	return u.repo.Get(key)
}

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

func main() {
	fmt.Println("--- Проверка 01_implement_repository (Решение) ---")

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
