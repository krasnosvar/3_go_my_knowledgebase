package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("--- Task 1: Mutex ---")
	task1_mutex()

	fmt.Println("\n--- Task 2: sync.Map ---")
	task2_sync_map()
}

type SafeMap struct {
	mu sync.RWMutex
	m  map[string]int
}

func (sm *SafeMap) Set(k string, v int) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.m[k] = v
}

func (sm *SafeMap) Get(k string) (int, bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	val, ok := sm.m[k]
	return val, ok
}

func task1_mutex() {
	// ================= РЕШЕНИЕ: Task 1 =================
	sm := SafeMap{m: make(map[string]int)}

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			sm.Set("test", val) // Без мьютекса тут была бы гонка и паника
		}(i)
	}
	wg.Wait()

	val, _ := sm.Get("test")
	fmt.Println("Значение 'test' (одно из последних):", val)
	// ====================================================
}

func task2_sync_map() {
	// ================= РЕШЕНИЕ: Task 2 =================
	var m sync.Map

	// 1. Store
	m.Store("apple", 10)

	// 2. Load
	if val, ok := m.Load("apple"); ok {
		fmt.Println("Loaded apple:", val)
	}

	// 3. LoadOrStore
	actual, loaded := m.LoadOrStore("apple", 20)
	fmt.Printf("LoadOrStore apple: value=%v, was_loaded=%v\n", actual, loaded)

	// 4. Delete
	m.Delete("apple")
	_, ok := m.Load("apple")
	fmt.Println("Apple exists after delete?", ok)
	// ====================================================
}
