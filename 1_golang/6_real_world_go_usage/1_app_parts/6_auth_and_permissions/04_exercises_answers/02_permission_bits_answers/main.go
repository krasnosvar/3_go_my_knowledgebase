// Упражнение 02: права как битовые флаги и роли (Решение).
package main

import "fmt"

type perm uint8

const (
	Read perm = 1 << iota
	Write
	Delete
)

var roles = map[string]perm{
	"viewer": Read,
	"editor": Read | Write,
	"admin":  Read | Write | Delete,
}

func can(role string, need perm) bool {
	have, ok := roles[role]
	return ok && have&need == need
}

func main() {
	fmt.Println("--- Проверка 02_permission_bits (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	check(can("viewer", Read), "viewer может Read")
	check(!can("viewer", Write), "viewer не может Write")
	check(can("editor", Read|Write), "editor может Read|Write")
	check(!can("editor", Write|Delete), "editor не может Write|Delete (нет Delete)")
	check(can("admin", Delete), "admin может Delete")
	check(!can("guest", Read), "неизвестная роль → false")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
