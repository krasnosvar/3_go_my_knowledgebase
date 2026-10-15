// Упражнение 01: копирование потока с ограничением размера (Решение).
package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

// copyLimited копирует не больше max байт из src в dst.
func copyLimited(dst io.Writer, src io.Reader, max int64) (int64, error) {
	return io.Copy(dst, io.LimitReader(src, max))
}

func main() {
	fmt.Println("--- Проверка 01_reader_copy (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	var buf bytes.Buffer
	n, err := copyLimited(&buf, strings.NewReader("hello, world"), 5)
	check(err == nil && n == 5 && buf.String() == "hello", fmt.Sprintf("max=5: n=%d, buf=%q, err=%v", n, buf.String(), err))

	buf.Reset()
	n, err = copyLimited(&buf, strings.NewReader("short"), 100)
	check(err == nil && n == 5 && buf.String() == "short", fmt.Sprintf("max больше данных: n=%d, buf=%q, err=%v", n, buf.String(), err))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
