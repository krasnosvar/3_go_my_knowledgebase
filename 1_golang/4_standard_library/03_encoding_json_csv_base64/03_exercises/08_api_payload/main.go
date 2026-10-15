// Упражнение 08: конверт ответа API.
//
// Задача: реализуйте buildResponse(items, total), которая сериализует
// ответ вида
//
//	{"data":[{"id":1,"title":"a"}],"meta":{"total":10,"count":1}}
//
// Требования:
//   - опишите типы envelope/meta с json-тегами;
//   - count = len(items);
//   - nil-срез items должен давать "data":[] , а не "data":null
//     (клиенты на JS/TS часто падают на null вместо массива).
//
// Запуск: go run ./03_exercises/08_api_payload/main.go
package main

import (
	"encoding/json"
	"fmt"
)

var _ = json.Marshal // держим импорт живым в заготовке

type Item struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// buildResponse собирает JSON-конверт ответа.
func buildResponse(items []Item, total int) ([]byte, error) {
	// TODO: опишите типы envelope/meta (см. условие), заполните (nil → []Item{}) и вызовите json.Marshal.
	return nil, nil
}

func main() {
	fmt.Println("--- Проверка 08_api_payload ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	b, err := buildResponse([]Item{{1, "a"}, {2, "b"}}, 10)
	want := `{"data":[{"id":1,"title":"a"},{"id":2,"title":"b"}],"meta":{"total":10,"count":2}}`
	check(err == nil && string(b) == want, fmt.Sprintf("получено %s, ожидалось %s", b, want))

	b, err = buildResponse(nil, 0)
	want = `{"data":[],"meta":{"total":0,"count":0}}`
	check(err == nil && string(b) == want, fmt.Sprintf("nil items: получено %s, ожидалось %s", b, want))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
