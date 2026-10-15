// Упражнение 06: строгий разбор и валидация JSON-запроса (Решение).
package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type CreateUser struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// parseRequest строго разбирает и валидирует запрос.
func parseRequest(data []byte) (CreateUser, error) {
	var req CreateUser
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&req); err != nil {
		return CreateUser{}, fmt.Errorf("decode: %w", err)
	}
	if strings.TrimSpace(req.Name) == "" {
		return CreateUser{}, errors.New("name is required")
	}
	if req.Age < 1 || req.Age > 150 {
		return CreateUser{}, fmt.Errorf("age %d out of range 1..150", req.Age)
	}
	return req, nil
}

func main() {
	fmt.Println("--- Проверка 06_json_validation (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	req, err := parseRequest([]byte(`{"name":"Ann","age":30}`))
	check(err == nil && req == CreateUser{"Ann", 30}, fmt.Sprintf("валидный запрос: получено %+v, %v", req, err))

	bad := map[string]string{
		"лишнее поле":  `{"name":"Ann","age":30,"admin":true}`,
		"пустое имя":   `{"name":"  ","age":30}`,
		"возраст 0":    `{"name":"Ann"}`,
		"возраст 200":  `{"name":"Ann","age":200}`,
		"неверный тип": `{"name":"Ann","age":"30"}`,
		"битый JSON":   `{"name":`,
	}
	for name, body := range bad {
		r, err := parseRequest([]byte(body))
		check(err != nil && r == CreateUser{}, fmt.Sprintf("%s: ожидалась ошибка и пустой результат, получено %+v, %v", name, r, err))
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
