// Упражнение 01: идемпотентный reconcile (Решение).
package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type cluster struct{ pods map[string]bool }

// podIndex возвращает индекс пода приложения app или -1, если под чужой.
func podIndex(app, pod string) int {
	s, ok := strings.CutPrefix(pod, app+"-")
	if !ok {
		return -1
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		return -1
	}
	return n
}

func reconcile(c *cluster, app string, replicas int) (created, deleted int) {
	for i := 0; i < replicas; i++ {
		name := app + "-" + strconv.Itoa(i)
		if !c.pods[name] {
			c.pods[name] = true
			created++
		}
	}
	for name := range c.pods {
		if idx := podIndex(app, name); idx >= replicas {
			delete(c.pods, name)
			deleted++
		}
	}
	return created, deleted
}

func names(c *cluster) string {
	var out []string
	for n := range c.pods {
		out = append(out, n)
	}
	sort.Strings(out)
	return strings.Join(out, ",")
}

func main() {
	fmt.Println("--- Проверка 01_reconcile_idempotent (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	c := &cluster{pods: map[string]bool{"web-0": true, "db-0": true}}
	cr, del := reconcile(c, "web", 3)
	check(cr == 2 && del == 0 && names(c) == "db-0,web-0,web-1,web-2", "scale up до 3: "+names(c))
	cr, del = reconcile(c, "web", 3)
	check(cr == 0 && del == 0, "повторный reconcile ничего не меняет")
	cr, del = reconcile(c, "web", 1)
	check(cr == 0 && del == 2 && names(c) == "db-0,web-0", "scale down до 1: "+names(c))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
