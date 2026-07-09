// Упражнение 03: исправить направление зависимостей.
//
// Проблема: usecase должен уметь логировать, но НЕ должен зависеть от
// конкретного логгера (детали инфраструктуры). Сейчас интерфейса нет,
// и вам нужно ввести порт в нужном слое, чтобы зависимость смотрела внутрь.
//
// Задача:
//  1. Объявите интерфейс-порт Logger с методом Log(msg string) — это порт
//     на стороне потребителя (usecase).
//  2. Сделайте так, чтобы ReportUsecase зависел от Logger (а не от
//     конкретного типа), и метод Do вызывал logger.Log(...).
//  3. StdoutLogger (адаптер) должен удовлетворять Logger.
//
// Правильное направление: адаптер (StdoutLogger) зависит от порта (Logger),
// а не usecase от адаптера.
//
// Запуск: go run ./03_exercises/03_fix_dependency_direction/main.go
package main

import (
	"fmt"
	"strings"
)

// TODO: объявите здесь интерфейс-порт Logger с методом Log(msg string).

// ReportUsecase — потребитель. Должен зависеть от порта Logger.
type ReportUsecase struct {
	// TODO: замените зависимость на интерфейс Logger.
	logger interface{}
}

// NewReportUsecase принимает Logger через конструктор.
// TODO: измените тип параметра на Logger и сохраните его в поле.
func NewReportUsecase(logger interface{}) *ReportUsecase {
	return &ReportUsecase{logger: logger}
}

// Do выполняет сценарий и логирует сообщение через порт.
func (u *ReportUsecase) Do(name string) {
	// TODO: вызовите u.logger.Log("report: " + name).
}

// StdoutLogger — адаптер. Пишет в переданный буфер (для детерминизма проверки).
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
	fmt.Println("--- Проверка 03_fix_dependency_direction ---")

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
