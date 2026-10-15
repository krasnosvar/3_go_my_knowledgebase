package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	ID   int    `json:"product_id"`
	Name string `json:"product_name"`
}

func main() {
	fmt.Println("--- Проверка 06_struct_tags (Решение) ---")

	p := Product{ID: 101, Name: "Laptop"}
	jsonData, _ := json.Marshal(p)

	expected := `{"product_id":101,"product_name":"Laptop"}`

	var parsed Product
	_ = json.Unmarshal([]byte(expected), &parsed)

	if parsed.ID == 101 && parsed.Name == "Laptop" {
		fmt.Println("Тест json-тэгов: PASS")
	} else {
		fmt.Printf("Тест json-тэгов: FAIL (получен JSON: %s)\n", string(jsonData))
	}
}
