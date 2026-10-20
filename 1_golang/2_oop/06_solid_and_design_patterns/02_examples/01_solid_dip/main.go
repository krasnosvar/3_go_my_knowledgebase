// Пример 01: Dependency Inversion Principle (инверсия зависимостей).
//
// Показываем «до» и «после».
//
// ДО: высокоуровневый модуль Report жёстко создаёт конкретную зависимость
// внутри себя (ConsoleNotifier) — заменить её (на email, slack, заглушку)
// нельзя без правки Report. Это прямая зависимость от детали.
//
// ПОСЛЕ: Report зависит от абстракции (интерфейс Notifier), а конкретную
// реализацию передают снаружи через конструктор. Теперь любую реализацию
// подставляем без изменения Report — и код тестируем (можно подставить фейк).
//
// Программа детерминированно печатает работу обеих реализаций через
// один и тот же интерфейс и завершается.
//
// Запуск: go run ./02_examples/01_solid_dip/main.go
package main

import (
	"fmt"
	"strings"
)

// ---------------------------------------------------------------------------
// ДО (нарушение DIP): Report сам создаёт конкретную зависимость.
// ---------------------------------------------------------------------------

type ConsoleNotifierBad struct{}

func (ConsoleNotifierBad) Send(msg string) { fmt.Println("[console] " + msg) }

type ReportBad struct {
	notifier ConsoleNotifierBad // привязка к КОНКРЕТНОМУ типу
}

func (r ReportBad) Publish() {
	// Report знает про конкретный ConsoleNotifierBad и создаёт его сам.
	r.notifier.Send("отчёт готов (жёсткая зависимость)")
}

// ---------------------------------------------------------------------------
// ПОСЛЕ (соблюдение DIP): зависимость — это абстракция (интерфейс).
// ---------------------------------------------------------------------------

// Notifier — абстракция, от которой зависит высокоуровневый модуль.
type Notifier interface {
	Send(msg string) error
}

// Report — высокоуровневый модуль. Зависит только от интерфейса Notifier.
type Report struct {
	notifier Notifier
}

// NewReport внедряет конкретную реализацию извне (dependency injection).
func NewReport(n Notifier) *Report { return &Report{notifier: n} }

func (r *Report) Publish() error {
	return r.notifier.Send("отчёт готов")
}

// --- Две разные реализации одной абстракции ---

// ConsoleNotifier печатает сообщение в консоль.
type ConsoleNotifier struct{}

func (ConsoleNotifier) Send(msg string) error {
	fmt.Println("[console] " + msg)
	return nil
}

// UppercaseNotifier печатает сообщение заглавными буквами (имитируем
// другой «канал» доставки с иным поведением).
type UppercaseNotifier struct{}

func (UppercaseNotifier) Send(msg string) error {
	fmt.Println("[UPPER] " + strings.ToUpper(msg))
	return nil
}

func main() {
	fmt.Println("=== ДО: жёсткая зависимость от конкретного типа ===")
	bad := ReportBad{}
	bad.Publish()

	fmt.Println()
	fmt.Println("=== ПОСЛЕ: Report зависит от абстракции Notifier ===")

	// Один и тот же Report работает с разными реализациями,
	// подставленными снаружи через один интерфейс.
	notifiers := []Notifier{ConsoleNotifier{}, UppercaseNotifier{}}
	for _, n := range notifiers {
		r := NewReport(n)
		if err := r.Publish(); err != nil {
			fmt.Println("ошибка:", err)
		}
	}

	fmt.Println()
	fmt.Println("Готово, программа завершается.")
}
