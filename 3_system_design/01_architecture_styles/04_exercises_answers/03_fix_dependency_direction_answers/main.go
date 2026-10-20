// Упражнение 03: исправить направление зависимостей (Решение).
package main

import (
	"fmt"
	"strings"
)

// Logger — порт на стороне потребителя (usecase зависит от него).
type Logger interface {
	Log(msg string)
}

// ReportUsecase зависит от порта Logger, а не от конкретного адаптера.
type ReportUsecase struct {
	logger Logger
}

func NewReportUsecase(logger Logger) *ReportUsecase {
	return &ReportUsecase{logger: logger}
}

func (u *ReportUsecase) Do(name string) {
	u.logger.Log("report: " + name)
}

// StdoutLogger — адаптер, зависит от порта (реализует Logger).
type StdoutLogger struct {
	sink *strings.Builder
}

func NewStdoutLogger(sink *strings.Builder) *StdoutLogger {
	return &StdoutLogger{sink: sink}
}

func (l *StdoutLogger) Log(msg string) {
	l.sink.WriteString(msg + "\n")
}

func main() {
	fmt.Println("--- Проверка 03_fix_dependency_direction (Решение) ---")

	var sink strings.Builder
	uc := NewReportUsecase(NewStdoutLogger(&sink))
	uc.Do("daily")

	allOK := true
	got := sink.String()
	want := "report: daily\n"
	if got != want {
		allOK = false
		fmt.Printf("FAIL: лог = %q, ожидалось %q\n", got, want)
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
