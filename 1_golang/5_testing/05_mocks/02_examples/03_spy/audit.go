// Package audit — пример: spy фиксирует, как зависимость была вызвана.
package audit

// Logger — зависимость для аудита.
type Logger interface {
	Log(event string)
}

// Transfer выполняет перевод и пишет два события аудита.
func Transfer(l Logger, from, to string, amount int) {
	l.Log("debit:" + from)
	l.Log("credit:" + to)
}
