package main

import "fmt"

type UserRepo interface{ GetName(id int) (string, bool) }

type memRepo struct{ m map[int]string }

func (r memRepo) GetName(id int) (string, bool) { v, ok := r.m[id]; return v, ok }

func main() {
	repo := memRepo{m: map[int]string{1: "Den"}}
	if name, ok := repo.GetName(1); ok {
		fmt.Println(name)
	}
}
