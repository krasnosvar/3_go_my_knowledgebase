package main

import (
	"encoding/json"
	"fmt"
)

// TODO: Добавьте к структуре Product json-тэги:
// ID -> "product_id"
// Name -> "product_name"
type Product struct {
	ID   int
	Name string
}

func main() {
	fmt.Println("--- Проверка 06_struct_tags ---")

	p := Product{ID: 101, Name: "Laptop"}
	jsonData, _ := json.Marshal(p)

	// Ожидается JSON формата: {"product_id":101,"product_name":"Laptop"}
	expected := `{"product_id":101,"product_name":"Laptop"}`

	// Унмаршалинг для проверки: если тэги настроены верно, мы сможем распарсить JSON обратно.
	var parsed Product
	_ = json.Unmarshal([]byte(expected), &parsed)

	if parsed.ID == 101 && parsed.Name == "Laptop" {
		fmt.Println("Тест json-тэгов: PASS")
	} else {
		fmt.Printf("Тест json-тэгов: FAIL (получен JSON: %s)\n", string(jsonData))
	}
}
