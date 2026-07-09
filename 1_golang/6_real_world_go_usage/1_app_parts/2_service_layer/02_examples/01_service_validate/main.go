package main

import "fmt"

type Service struct{}

func (Service) CreateUser(name string) error {
	if name == "" {
		return fmt.Errorf("name required")
	}
	return nil
}

func main() {
	fmt.Println(Service{}.CreateUser("den") == nil)
}
