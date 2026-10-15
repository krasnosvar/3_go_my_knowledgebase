// Упражнение 08: разбор JSON-подобной структуры (Решение).
package main

import (
	"fmt"
	"strings"
)

func getPath(m map[string]any, path string) (any, bool) {
	var cur any = m
	for _, key := range strings.Split(path, ".") {
		obj, ok := cur.(map[string]any)
		if !ok {
			return nil, false
		}
		cur, ok = obj[key]
		if !ok {
			return nil, false
		}
	}
	return cur, true
}

func getInt(m map[string]any, path string) (int, bool) {
	v, ok := getPath(m, path)
	if !ok {
		return 0, false
	}
	f, ok := v.(float64)
	if !ok || f != float64(int(f)) {
		return 0, false
	}
	return int(f), true
}

func main() {
	fmt.Println("--- Проверка 08_json_like (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	doc := map[string]any{
		"user": map[string]any{
			"name": "Den",
			"age":  30.0,
			"address": map[string]any{
				"city": "Belgrade",
			},
		},
		"score": 9.5,
	}

	v, ok := getPath(doc, "user.address.city")
	check(ok && v == "Belgrade", fmt.Sprintf("user.address.city: получено (%v, %t)", v, ok))

	_, ok = getPath(doc, "user.phone")
	check(!ok, "user.phone: ключа нет, ожидалось ok=false")

	_, ok = getPath(doc, "user.name.first")
	check(!ok, "user.name.first: name — строка, а не объект, ожидалось ok=false")

	n, ok := getInt(doc, "user.age")
	check(ok && n == 30, fmt.Sprintf("getInt user.age: ожидалось (30, true), получено (%d, %t)", n, ok))

	_, ok = getInt(doc, "score")
	check(!ok, "getInt score=9.5: дробное, ожидалось ok=false")

	_, ok = getInt(doc, "user.name")
	check(!ok, "getInt user.name: строка, ожидалось ok=false")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
