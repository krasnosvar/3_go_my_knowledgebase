// Упражнение 06: строгий разбор и валидация JSON-запроса.
//
// Задача: реализуйте parseRequest(data) для тела запроса CreateUser:
//   - декодируйте через json.NewDecoder(bytes.NewReader(data)) с
//     DisallowUnknownFields() — лишние поля (опечатки клиента) → ошибка;
//   - после декодирования проверьте бизнес-правила:
//     Name не пустое (после TrimSpace), Age в диапазоне 1..150;
//   - при любой ошибке возвращайте пустой CreateUser и ошибку.
//
// Запуск: go run ./03_exercises/06_json_validation/main.go
package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

var _ = json.NewDecoder // держим импорты живыми в заготовке
var _ = bytes.NewReader
var _ = errors.New
var _ = strings.TrimSpace

type CreateUser struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// parseRequest строго разбирает и валидирует запрос.
func parseRequest(data []byte) (CreateUser, error) {
	// TODO: Decoder + DisallowUnknownFields + проверки Name/Age.
	return CreateUser{}, nil
}

func main() {
	fmt.Println("--- Проверка 06_json_validation ---")

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
