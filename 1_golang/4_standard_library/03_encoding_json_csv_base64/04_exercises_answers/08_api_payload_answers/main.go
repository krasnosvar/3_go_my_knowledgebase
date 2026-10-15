// Упражнение 08: конверт ответа API (Решение).
package main

import (
	"encoding/json"
	"fmt"
)

type Item struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type meta struct {
	Total int `json:"total"`
	Count int `json:"count"`
}

type envelope struct {
	Data []Item `json:"data"`
	Meta meta   `json:"meta"`
}

// buildResponse собирает JSON-конверт ответа.
func buildResponse(items []Item, total int) ([]byte, error) {
	if items == nil {
		items = []Item{}
	}
	return json.Marshal(envelope{Data: items, Meta: meta{Total: total, Count: len(items)}})
}

func main() {
	fmt.Println("--- Проверка 08_api_payload (Решение) ---")

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
