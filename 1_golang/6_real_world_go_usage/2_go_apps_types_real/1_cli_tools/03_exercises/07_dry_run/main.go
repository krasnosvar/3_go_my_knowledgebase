// Упражнение 07: режим --dry-run.
//
// Опасные CLI (удаление, миграции) должны уметь показать план без изменений.
//
// Задача: реализуйте cleanup(files []string, dryRun bool, remove func(string) error, out io.Writer) error:
//   - для каждого файла с суффиксом ".tmp":
//     dryRun → печатать "would remove <f>\n" и НЕ вызывать remove;
//     иначе → вызвать remove(f), печатать "removed <f>\n";
//   - остальные файлы пропускать молча;
//   - первая ошибка remove прерывает работу и возвращается обёрнутой:
//     fmt.Errorf("remove %s: %w", f, err).
//
// Запуск: go run ./03_exercises/07_dry_run
package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"
)

var _ = strings.HasSuffix

func cleanup(files []string, dryRun bool, remove func(string) error, out io.Writer) error {
	// TODO: отфильтруйте *.tmp, в dryRun только печатайте план.
	return nil
}

func main() {
	fmt.Println("--- Проверка 07_dry_run ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	files := []string{"a.tmp", "keep.go", "b.tmp"}
	var removed []string
	rm := func(f string) error { removed = append(removed, f); return nil }

	var out bytes.Buffer
	err := cleanup(files, true, rm, &out)
	check(err == nil && len(removed) == 0, "dry-run ничего не удаляет")
	check(out.String() == "would remove a.tmp\nwould remove b.tmp\n", "dry-run печатает план")

	out.Reset()
	err = cleanup(files, false, rm, &out)
	check(err == nil && fmt.Sprint(removed) == "[a.tmp b.tmp]", "реальный режим удаляет *.tmp")
	check(out.String() == "removed a.tmp\nremoved b.tmp\n", "реальный режим печатает отчёт")

	denied := errors.New("permission denied")
	err = cleanup(files, false, func(string) error { return denied }, io.Discard)
	check(errors.Is(err, denied) && err != nil && strings.Contains(err.Error(), "a.tmp"), "ошибка обёрнута с именем файла")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
