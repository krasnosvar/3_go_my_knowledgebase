package main

import "fmt"

type User struct{ Name string }

func (u User) IsValid() bool { return u.Name != "" }
func main()                  { fmt.Println(User{Name: "Den"}.IsValid()) }
