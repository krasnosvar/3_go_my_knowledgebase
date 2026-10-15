// Упражнение 07: потоковое чтение JSON через json.Decoder (Решение).
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
)

type event struct {
	Value int `json:"value"`
}

// sumStream суммирует value из потока JSON-объектов.
func sumStream(r io.Reader) (int, error) {
	dec := json.NewDecoder(r)
	sum := 0
	for {
		var ev event
		err := dec.Decode(&ev)
		if errors.Is(err, io.EOF) {
			return sum, nil
		}
		if err != nil {
			return 0, fmt.Errorf("decode event: %w", err)
		}
		sum += ev.Value
	}
}

func main() {
	fmt.Println("--- Проверка 07_stream_decoder (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	in := "{\"value\": 1}\n{\"value\": 2}\n{\"value\": 39}\n"
	sum, err := sumStream(strings.NewReader(in))
	check(err == nil && sum == 42, fmt.Sprintf("ожидалось 42, nil; получено %d, %v", sum, err))

	sum, err = sumStream(strings.NewReader(""))
	check(err == nil && sum == 0, fmt.Sprintf("пустой поток: ожидалось 0, nil; получено %d, %v", sum, err))

	_, err = sumStream(strings.NewReader("{\"value\": 1}\n{\"value\": oops}\n"))
	check(err != nil, "битый объект в потоке должен давать ошибку")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
