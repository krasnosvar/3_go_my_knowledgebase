// Package token — пример: подмена времени через функцию now.
package token

import "time"

// Token живёт до ExpiresAt.
type Token struct{ ExpiresAt time.Time }

// Validator проверяет токены. now — инъекция времени (в проде time.Now).
type Validator struct{ now func() time.Time }

// NewValidator создаёт валидатор с системными часами.
func NewValidator() *Validator { return &Validator{now: time.Now} }

// Valid сообщает, не истёк ли токен.
func (v *Validator) Valid(t Token) bool { return v.now().Before(t.ExpiresAt) }
