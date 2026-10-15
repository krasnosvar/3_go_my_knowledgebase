// Упражнение 03: валидация запроса с перечнем нарушений.
//
// В gRPC принято отвечать InvalidArgument и перечислять нарушения по полям
// (google.rpc.BadRequest.FieldViolations). Смоделируем это.
//
// Задача: реализуйте validate(req createOrderRequest) []violation:
//   - customer_id пустой → {"customer_id", "required"};
//   - items пуст → {"items", "at least one item"};
//   - у каждой позиции i quantity <= 0 → {"items[i].quantity", "must be > 0"};
//
// Порядок: customer_id, items, затем позиции по порядку. Нет нарушений → nil.
//
// Запуск: go run ./03_exercises/03_request_validation
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
	// TODO: соберите все нарушения в срез.
	return nil
}

func main() {
	fmt.Println("--- Проверка 03_request_validation ---")

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
