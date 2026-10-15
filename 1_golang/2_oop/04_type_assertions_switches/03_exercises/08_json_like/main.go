// Упражнение 08: разбор JSON-подобной структуры.
//
// Задача: данные как после json.Unmarshal в any: объекты — map[string]any,
// числа — float64. Реализуйте getPath(m, "user.address.city"):
//   - путь делится по "."; на каждом шаге текущее значение должно быть
//     map[string]any и содержать ключ, иначе (nil, false);
//
// и getInt(m, path) — getPath + проверка, что значение float64 без дробной
// части; возвращает (int, true) или (0, false).
//
// Запуск: go run ./03_exercises/08_json_like/main.go
package main

import (
	"fmt"
	"strings"
)

// getPath достаёт значение по пути через точку.
func getPath(m map[string]any, path string) (any, bool) {
	// TODO: strings.Split(path, "."); идите по ключам, на каждом шаге
	// проверяя cur.(map[string]any).
	_ = strings.Split
	return nil, false
}

// getInt достаёт целое число по пути.
func getInt(m map[string]any, path string) (int, bool) {
	// TODO: getPath + v.(float64) + проверка целого.
	return 0, false
}

func main() {
	fmt.Println("--- Проверка 08_json_like ---")

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
