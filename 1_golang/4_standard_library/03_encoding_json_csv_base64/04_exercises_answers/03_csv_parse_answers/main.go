// Упражнение 03: разбор CSV через encoding/csv (Решение).
package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Person struct {
	Name string
	Age  int
}

// parseCSV читает людей из CSV с заголовком.
func parseCSV(r io.Reader) ([]Person, error) {
	rows, err := csv.NewReader(r).ReadAll()
	if err != nil {
		return nil, fmt.Errorf("read csv: %w", err)
	}
	if len(rows) == 0 {
		return []Person{}, nil
	}
	people := make([]Person, 0, len(rows)-1)
	for i, row := range rows[1:] {
		age, err := strconv.Atoi(row[1])
		if err != nil {
			return nil, fmt.Errorf("line %d: bad age %q: %w", i+2, row[1], err)
		}
		people = append(people, Person{Name: row[0], Age: age})
	}
	return people, nil
}

func main() {
	fmt.Println("--- Проверка 03_csv_parse (Решение) ---")

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
