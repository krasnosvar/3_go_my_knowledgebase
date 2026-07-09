package main

import "fmt"

type UserRepo interface {
	GetName(id int) (string, error)
}

type memoryRepo struct {
	data map[int]string
}

func (m memoryRepo) GetName(id int) (string, error) {
	v, ok := m.data[id]
	if !ok {
		return "", fmt.Errorf("not found")
	}
	return v, nil
}

func handle(repo UserRepo, id int) {
	name, err := repo.GetName(id)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("user:", name)
}

func main() {
	handle(memoryRepo{data: map[int]string{1: "Den"}}, 1)
}
