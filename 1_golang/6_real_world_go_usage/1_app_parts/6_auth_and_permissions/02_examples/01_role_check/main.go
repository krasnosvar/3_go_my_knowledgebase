package main

import "fmt"

func can(role, action string) bool { return role == "admin" || action == "read" }

func main() { fmt.Println(can("user", "read"), can("user", "delete")) }
