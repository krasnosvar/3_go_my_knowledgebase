// Упражнение 03: разбор CSV через encoding/csv.
//
// Задача: реализуйте parseCSV(r), которая читает CSV вида
//
//	name,age
//	Ann,30
//	"Smith, John",41
//
// и возвращает []Person. Правила:
//   - первая строка — заголовок, её пропускаем;
//   - age переводится через strconv.Atoi; ошибка → (nil, ошибка с номером строки);
//   - кавычки и запятые внутри полей разбирает csv.Reader — не делите строку вручную.
//
// Запуск: go run ./03_exercises/03_csv_parse/main.go
package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"strings"
)

var _ = csv.NewReader // держим импорты живыми в заготовке
var _ = strconv.Atoi

type Person struct {
	Name string
	Age  int
}

// parseCSV читает людей из CSV с заголовком.
func parseCSV(r io.Reader) ([]Person, error) {
	// TODO: csv.NewReader(r).ReadAll(), пропустить заголовок, Atoi для age.
	return nil, nil
}

func main() {
	fmt.Println("--- Проверка 03_csv_parse ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	in := "name,age\nAnn,30\n\"Smith, John\",41\n"
	people, err := parseCSV(strings.NewReader(in))
	check(err == nil && len(people) == 2 && people[0] == Person{"Ann", 30} && people[1] == Person{"Smith, John", 41},
		fmt.Sprintf("получено %+v, %v", people, err))

	_, err = parseCSV(strings.NewReader("name,age\nBob,abc\n"))
	check(err != nil && strings.Contains(err.Error(), "line 2"), fmt.Sprintf("ожидалась ошибка с \"line 2\", получено %v", err))

	_, err = parseCSV(strings.NewReader("name,age\nBob\n"))
	check(err != nil, "строка с неверным числом полей должна давать ошибку (csv.ErrFieldCount)")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
