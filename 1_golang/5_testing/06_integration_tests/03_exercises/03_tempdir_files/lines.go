// Package lines — код под тестом для упражнения 03.
package lines

import (
	"os"
	"strings"
)

// Save записывает строки в файл, по одной на строку.
func Save(path string, lines []string) error {
	return os.WriteFile(path, []byte(strings.Join(lines, "\n")+"\n"), 0o600)
}

// Load читает строки из файла, записанного Save.
func Load(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return strings.Split(strings.TrimSuffix(string(data), "\n"), "\n"), nil
}
