// Упражнение 02: права как битовые флаги и роли.
//
// Задача:
//   - заполните roles: viewer = Read; editor = Read|Write; admin = Read|Write|Delete;
//   - реализуйте can(role string, need perm) bool: true, если у роли есть
//     ВСЕ биты из need (need может содержать несколько прав сразу);
//     неизвестная роль → false.
//
// Запуск: go run ./03_exercises/02_permission_bits
package main

import "fmt"

type perm uint8

const (
	Read perm = 1 << iota
	Write
	Delete
)

var roles = map[string]perm{
	// TODO: viewer, editor, admin.
}

func can(role string, need perm) bool {
	// TODO: проверьте, что все биты need установлены у роли.
	return false
}

func main() {
	fmt.Println("--- Проверка 02_permission_bits ---")

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
