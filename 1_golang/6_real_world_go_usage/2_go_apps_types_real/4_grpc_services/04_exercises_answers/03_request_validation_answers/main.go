// Упражнение 03: валидация запроса с перечнем нарушений (Решение).
package main

import "fmt"

type item struct {
	SKU      string
	Quantity int
}

type createOrderRequest struct {
	CustomerID string
	Items      []item
}

type violation struct{ Field, Desc string }

func validate(req createOrderRequest) []violation {
	var v []violation
	if req.CustomerID == "" {
		v = append(v, violation{"customer_id", "required"})
	}
	if len(req.Items) == 0 {
		v = append(v, violation{"items", "at least one item"})
	}
	for i, it := range req.Items {
		if it.Quantity <= 0 {
			v = append(v, violation{fmt.Sprintf("items[%d].quantity", i), "must be > 0"})
		}
	}
	return v
}

func main() {
	fmt.Println("--- Проверка 03_request_validation (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	check(validate(createOrderRequest{CustomerID: "c1", Items: []item{{"a", 1}}}) == nil, "валидный запрос → nil")
	v := validate(createOrderRequest{})
	check(fmt.Sprint(v) == "[{customer_id required} {items at least one item}]", "пустой запрос: "+fmt.Sprint(v))
	v = validate(createOrderRequest{CustomerID: "c1", Items: []item{{"a", 1}, {"b", 0}, {"c", -2}}})
	check(fmt.Sprint(v) == "[{items[1].quantity must be > 0} {items[2].quantity must be > 0}]", "нарушения по позициям: "+fmt.Sprint(v))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
