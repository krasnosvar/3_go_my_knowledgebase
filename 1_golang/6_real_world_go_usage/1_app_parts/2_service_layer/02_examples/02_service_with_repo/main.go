package main

import "fmt"

type Repo interface{ Save(string) error }
type memRepo struct{}

func (memRepo) Save(string) error { return nil }

type Service struct{ r Repo }

func (s Service) Register(name string) error { return s.r.Save(name) }

func main() { fmt.Println(Service{r: memRepo{}}.Register("x") == nil) }
