// Упражнение 07: потоковое чтение JSON через json.Decoder.
//
// Задача: реализуйте sumStream(r), которая читает из r последовательность
// JSON-объектов {"value": N} (по одному на строку, как в NDJSON-логах)
// и возвращает сумму value.
//   - читайте json.NewDecoder(r).Decode(&ev) в цикле до io.EOF;
//   - io.EOF — нормальный конец потока, а не ошибка;
//   - любая другая ошибка → (0, ошибка).
//
// Decoder не загружает весь поток в память — это важно для больших файлов
// и тел HTTP-запросов.
//
// Запуск: go run ./03_exercises/07_stream_decoder/main.go
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
)

var _ = json.NewDecoder // держим импорты живыми в заготовке
var _ = errors.Is

type event struct {
	Value int `json:"value"`
}

// sumStream суммирует value из потока JSON-объектов.
func sumStream(r io.Reader) (int, error) {
	// TODO: цикл dec.Decode(&ev) до errors.Is(err, io.EOF).
	return -1, nil
}

func main() {
	fmt.Println("--- Проверка 07_stream_decoder ---")

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
