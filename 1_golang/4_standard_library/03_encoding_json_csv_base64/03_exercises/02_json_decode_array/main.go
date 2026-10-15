// Упражнение 02: декодирование JSON-массива.
//
// Задача: реализуйте decodeItems(data), которая разбирает JSON-массив
// объектов {"sku": "...", "qty": N} в []Item через json.Unmarshal.
//   - невалидный JSON → (nil, ошибка);
//   - "[]" → пустой срез без ошибки.
//
// Запуск: go run ./03_exercises/02_json_decode_array/main.go
package main

import (
	"encoding/json"
	"fmt"
)

var _ = json.Unmarshal // держим импорт живым в заготовке

type Item struct {
	SKU string `json:"sku"`
	Qty int    `json:"qty"`
}

// decodeItems разбирает JSON-массив товаров.
func decodeItems(data []byte) ([]Item, error) {
	// TODO: var items []Item; json.Unmarshal(data, &items).
	return nil, nil
}

func main() {
	fmt.Println("--- Проверка 02_json_decode_array ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	items, err := decodeItems([]byte(`[{"sku":"A-1","qty":2},{"sku":"B-7","qty":10}]`))
	check(err == nil && len(items) == 2 && items[0] == Item{"A-1", 2} && items[1] == Item{"B-7", 10},
		fmt.Sprintf("получено %+v, %v", items, err))

	items, err = decodeItems([]byte(`[]`))
	check(err == nil && len(items) == 0, fmt.Sprintf("\"[]\": получено %+v, %v", items, err))

	_, err = decodeItems([]byte(`[{"sku":"A-1",`))
	check(err != nil, "битый JSON должен давать ошибку")

	_, err = decodeItems([]byte(`{"sku":"A-1"}`))
	check(err != nil, "объект вместо массива должен давать ошибку")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
