// Упражнение 05: логгер демона в stderr.
//
// Под systemd/Docker всё, что процесс пишет в stderr, попадает в journald/
// docker logs — с собственной меткой времени. Свою метку добавлять не нужно.
//
// Задача: реализуйте newLogger(w io.Writer, debug bool) *slog.Logger:
//   - slog.NewTextHandler(w, opts);
//   - уровень: Debug, если debug == true, иначе Info;
//   - через ReplaceAttr убрать атрибут времени (a.Key == slog.TimeKey
//     на верхнем уровне, т.е. len(groups) == 0 → вернуть пустой slog.Attr{}).
//
// Запуск: go run ./03_exercises/05_stderr_logging
package main

import (
	"bytes"
	"fmt"
	"io"
	"log/slog"
	"strings"
)

func newLogger(w io.Writer, debug bool) *slog.Logger {
	// TODO: HandlerOptions с Level и ReplaceAttr.
	return slog.New(slog.NewTextHandler(w, nil))
}

func main() {
	fmt.Println("--- Проверка 05_stderr_logging ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	var buf bytes.Buffer
	l := newLogger(&buf, false)
	l.Debug("hidden")
	l.Info("started", "pid", 42)
	out := buf.String()
	check(!strings.Contains(out, "hidden"), "debug скрыт при debug=false")
	check(strings.Contains(out, "msg=started pid=42"), "info пишется с полями")
	check(!strings.Contains(out, "time="), "метки времени нет")

	buf.Reset()
	newLogger(&buf, true).Debug("visible")
	check(strings.Contains(buf.String(), "level=DEBUG msg=visible"), "debug виден при debug=true")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
