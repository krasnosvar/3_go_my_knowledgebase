// Упражнение 08: множество флагов на битовой маске (Решение).
package main

import "fmt"

// Perm — набор прав в виде битовой маски.
type Perm uint8

const (
	Read Perm = 1 << iota
	Write
	Exec
)

// grant добавляет флаги f.
func grant(p, f Perm) Perm {
	return p | f
}

// revoke убирает флаги f.
func revoke(p, f Perm) Perm {
	return p &^ f
}

// has сообщает, установлены ли все флаги f.
func has(p, f Perm) bool {
	return p&f == f
}

func main() {
	fmt.Println("--- Проверка 08_permission_flags (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	var p Perm
	p = grant(p, Read|Write)
	check(has(p, Read) && has(p, Write), "после grant(Read|Write) должны быть Read и Write")
	check(!has(p, Exec), "Exec не выдавали")
	check(has(p, Read|Write), "has(Read|Write) должен вернуть true")
	check(!has(p, Read|Exec), "has(Read|Exec) должен вернуть false: Exec не выдан")

	p = revoke(p, Write)
	check(has(p, Read) && !has(p, Write), "после revoke(Write) должен остаться только Read")
	check(p == Read, fmt.Sprintf("ожидалось значение %d (Read), получено %d", Read, p))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
