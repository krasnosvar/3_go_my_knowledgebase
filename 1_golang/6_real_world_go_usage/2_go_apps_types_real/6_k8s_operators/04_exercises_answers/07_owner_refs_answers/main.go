// Упражнение 07: ownerReferences и каскадное удаление (Решение).
package main

import (
	"fmt"
	"sort"
)

type obj struct {
	UID    string
	Owners []string
}

func orphans(objs map[string]obj, deleted string) []string {
	gone := map[string]bool{deleted: true}
	for changed := true; changed; {
		changed = false
		for uid, o := range objs {
			if gone[uid] || len(o.Owners) == 0 {
				continue
			}
			all := true
			for _, ow := range o.Owners {
				if !gone[ow] {
					all = false
					break
				}
			}
			if all {
				gone[uid] = true
				changed = true
			}
		}
	}
	var out []string
	for uid := range gone {
		if uid != deleted {
			out = append(out, uid)
		}
	}
	sort.Strings(out)
	return out
}

func main() {
	fmt.Println("--- Проверка 07_owner_refs (Решение) ---")

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
