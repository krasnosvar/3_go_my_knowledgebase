package main

import "fmt"

type UserRepo interface{ Get(int) (string, bool) }
type mem struct{ m map[int]string }

func (r mem) Get(id int) (string, bool) { v, ok := r.m[id]; return v, ok }

func main() { name, ok := mem{m: map[int]string{1: "den"}}.Get(1); fmt.Println(name, ok) }
