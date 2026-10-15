// ⏩ Забегание вперёд: горутины, sync.RWMutex и sync.WaitGroup разбираются в 3_concurrency.

package main

import (
	"fmt"
	"sync"
	"time"
)

// Session содержит данные сессии пользователя
type Session struct {
	UserID    string
	CreatedAt time.Time
}

// SessionStore — потокобезопасное хранилище сессий на базе RWMutex
type SessionStore struct {
	mu       sync.RWMutex
	sessions map[string]Session
}

func NewSessionStore() *SessionStore {
	return &SessionStore{
		sessions: make(map[string]Session),
	}
}

// Set сохраняет сессию (пишущая блокировка Lock)
func (store *SessionStore) Set(token string, session Session) {
	store.mu.Lock()
	defer store.mu.Unlock()
	store.sessions[token] = session
}

// Get возвращает сессию (читающая блокировка RLock)
func (store *SessionStore) Get(token string) (Session, bool) {
	store.mu.RLock()
	defer store.mu.RUnlock()
	session, exists := store.sessions[token]
	return session, exists
}

// Delete удаляет сессию (пишущая блокировка Lock)
func (store *SessionStore) Delete(token string) {
	store.mu.Lock()
	defer store.mu.Unlock()
	delete(store.sessions, token)
}

func main() {
	store := NewSessionStore()

	// Демонстрируем конкурентную запись и чтение
	var wg sync.WaitGroup
	wg.Add(3)

	// Воркер 1: запись сессии
	go func() {
		defer wg.Done()
		store.Set("token_alice", Session{UserID: "alice", CreatedAt: time.Now()})
	}()

	// Воркер 2: запись другой сессии
	go func() {
		defer wg.Done()
		store.Set("token_bob", Session{UserID: "bob", CreatedAt: time.Now()})
	}()

	// Воркер 3: чтение сессии в цикле
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			if sess, ok := store.Get("token_alice"); ok {
				fmt.Printf("Найдена сессия пользователя: %s\n", sess.UserID)
				return
			}
			time.Sleep(1 * time.Millisecond)
		}
	}()

	wg.Wait()
	fmt.Println("Сессии успешно проверены.")
}
