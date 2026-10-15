// Package counter — пример: setup/teardown через хелпер и t.Cleanup.
package counter

import "os"

// AppendLine дописывает строку в файл, создавая его при необходимости.
func AppendLine(path, line string) error {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o600)
	if err != nil {
		return err
	}
	if _, err := f.WriteString(line + "\n"); err != nil {
		f.Close()
		return err
	}
	return f.Close()
}
