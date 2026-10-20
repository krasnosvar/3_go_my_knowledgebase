// Упражнение 10: семантика HTTP-методов (safe / idempotent).
//
// Задача: реализовать classify — по имени HTTP-метода вернуть два флага:
//
//	safe       — метод только читает, не меняет состояние сервера;
//	idempotent — повторный одинаковый вызов даёт то же состояние.
//
// Ориентир (стандартная семантика REST):
//
//	GET, HEAD, OPTIONS -> safe=true,  idempotent=true
//	PUT, DELETE        -> safe=false, idempotent=true
//	POST, PATCH        -> safe=false, idempotent=false
//	неизвестный метод  -> safe=false, idempotent=false
//
// Запуск: go run ./03_exercises/10_method_semantics/main.go
package main

import (
	"fmt"
	"net/http"
)

// classify — реализуйте согласно условию.
func classify(method string) (safe, idempotent bool) {
	// TODO:
	//   switch method {
	//   case http.MethodGet, http.MethodHead, http.MethodOptions: return
	// true, true
	//   case http.MethodPut, http.MethodDelete: return false, true
	//   default (POST, PATCH, неизвестное): return false, false
	//   }
	return false, false
}

// var _ удерживает импорт net/http используемым, пока TODO не сделан.
var _ = http.MethodGet

func main() {
	fmt.Println("--- Проверка 10_method_semantics ---")

	type want struct{ safe, idem bool }
	cases := map[string]want{
		http.MethodGet:     {true, true},
		http.MethodHead:    {true, true},
		http.MethodOptions: {true, true},
		http.MethodPut:     {false, true},
		http.MethodDelete:  {false, true},
		http.MethodPost:    {false, false},
		http.MethodPatch:   {false, false},
		"WAT":              {false, false},
	}

	allOK := true
	for m, w := range cases {
		safe, idem := classify(m)
		if safe != w.safe || idem != w.idem {
			allOK = false
			fmt.Printf("FAIL: %-7s -> safe=%v idem=%v, ожидалось safe=%v idem=%v\n",
				m, safe, idem, w.safe, w.idem)
		}
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
