// Упражнение 08: множество флагов на битовой маске.
//
// Задача: маленькое множество (до 64 элементов) удобно хранить битовой маской.
// Реализуйте операции над правами Perm:
//   - grant(p, f) — добавить флаг(и) f (p | f);
//   - revoke(p, f) — убрать флаг(и) f (p &^ f);
//   - has(p, f) — true, если ВСЕ биты f установлены в p.
//
// Запуск: go run ./03_exercises/08_permission_flags/main.go
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
	// TODO: побитовое ИЛИ.
	return p
}

// revoke убирает флаги f.
func revoke(p, f Perm) Perm {
	// TODO: оператор AND NOT (&^).
	return p
}

// has сообщает, установлены ли все флаги f.
func has(p, f Perm) bool {
	// TODO: проверьте p & f == f.
	return false
}

func main() {
	fmt.Println("--- Проверка 08_permission_flags ---")

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
