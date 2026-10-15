// Package store — пример: интеграционный тест с файловой системой.
package store

import (
	"encoding/json"
	"os"
)

// Settings — сохраняемые настройки.
type Settings struct {
	Theme string `json:"theme"`
	Font  int    `json:"font"`
}

// Save пишет настройки в JSON-файл.
func Save(path string, s Settings) error {
	data, err := json.Marshal(s)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o600)
}

// Load читает настройки из JSON-файла.
func Load(path string) (Settings, error) {
	var s Settings
	data, err := os.ReadFile(path)
	if err != nil {
		return s, err
	}
	return s, json.Unmarshal(data, &s)
}
