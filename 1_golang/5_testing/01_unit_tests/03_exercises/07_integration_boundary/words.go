// Package words — код под тестом для упражнения 07.
package words

import (
	"bufio"
	"io"
)

// CountWords считает слова в r. Функция принимает io.Reader, а не путь к
// файлу: благодаря этому её можно тестировать без диска и сети.
func CountWords(r io.Reader) (int, error) {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)
	n := 0
	for sc.Scan() {
		n++
	}
	return n, sc.Err()
}
