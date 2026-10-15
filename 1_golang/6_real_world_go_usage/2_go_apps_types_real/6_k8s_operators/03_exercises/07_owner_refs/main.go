// Упражнение 07: ownerReferences и каскадное удаление.
//
// Garbage collector Kubernetes удаляет объекты, все владельцы которых
// удалены (Deployment → ReplicaSet → Pod).
//
// Задача: реализуйте orphans(objs map[string]obj, deleted string) []string —
// UID всех объектов, которые нужно удалить вслед за deleted:
//   - объект удаляется, если ВСЕ его владельцы (Owners) удалены
//     (объект без владельцев не трогаем);
//   - каскад транзитивный: удалённый ReplicaSet тянет за собой поды;
//   - результат отсортирован (sort.Strings), сам deleted не включать.
//
// Запуск: go run ./03_exercises/07_owner_refs
package main

import (
	"fmt"
	"sort"
)

type obj struct {
	UID    string
	Owners []string
}

var _ = sort.Strings

func orphans(objs map[string]obj, deleted string) []string {
	// TODO: повторяйте проход, пока находятся новые объекты, все владельцы которых удалены.
	return nil
}

func main() {
	fmt.Println("--- Проверка 07_owner_refs ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	objs := map[string]obj{
		"deploy": {UID: "deploy"},
		"rs-1":   {UID: "rs-1", Owners: []string{"deploy"}},
		"pod-a":  {UID: "pod-a", Owners: []string{"rs-1"}},
		"pod-b":  {UID: "pod-b", Owners: []string{"rs-1"}},
		"shared": {UID: "shared", Owners: []string{"rs-1", "other"}},
		"other":  {UID: "other"},
		"lonely": {UID: "lonely"},
	}
	got := fmt.Sprint(orphans(objs, "deploy"))
	check(got == "[pod-a pod-b rs-1]", "каскад deploy → rs-1 → pod-a,pod-b: "+got)
	check(fmt.Sprint(orphans(objs, "lonely")) == "[]", "у lonely нет зависимых")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
