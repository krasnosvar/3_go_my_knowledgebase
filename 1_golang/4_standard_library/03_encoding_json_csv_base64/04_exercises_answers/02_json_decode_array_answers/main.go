// Упражнение 02: декодирование JSON-массива (Решение).
package main

import (
	"encoding/json"
	"fmt"
)

type Item struct {
	SKU string `json:"sku"`
	Qty int    `json:"qty"`
}

// decodeItems разбирает JSON-массив товаров.
func decodeItems(data []byte) ([]Item, error) {
	var items []Item
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("decode items: %w", err)
	}
	return items, nil
}

func main() {
	fmt.Println("--- Проверка 02_json_decode_array (Решение) ---")

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
