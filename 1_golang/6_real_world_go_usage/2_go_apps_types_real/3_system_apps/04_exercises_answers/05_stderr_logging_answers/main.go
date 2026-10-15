// Упражнение 05: логгер демона в stderr (Решение).
package main

import (
	"bytes"
	"fmt"
	"io"
	"log/slog"
	"strings"
)

func newLogger(w io.Writer, debug bool) *slog.Logger {
	level := slog.LevelInfo
	if debug {
		level = slog.LevelDebug
	}
	opts := &slog.HandlerOptions{
		Level: level,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if len(groups) == 0 && a.Key == slog.TimeKey {
				return slog.Attr{}
			}
			return a
		},
	}
	return slog.New(slog.NewTextHandler(w, opts))
}

func main() {
	fmt.Println("--- Проверка 05_stderr_logging (Решение) ---")

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
