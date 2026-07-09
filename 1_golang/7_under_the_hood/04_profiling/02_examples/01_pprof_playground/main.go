package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof" // Подключаем пакет для раздачи профилей по HTTP
	"runtime"
	"sync"
	"time"
)

func main() {
	// Включаем встроенное отслеживание блокировок и конкуренции за мьютексы
	// В продакшене их стоит включать с осторожностью, но для "полигона" - самое то
	runtime.SetBlockProfileRate(1)
	runtime.SetMutexProfileFraction(1)

	fmt.Println("🚀 Запуск сервера. Статистика доступна здесь:")
	fmt.Println("   http://localhost:8080/debug/pprof/")
	fmt.Println("\nДля снятия профилей используйте консоль, например:")
	fmt.Println("   go tool pprof http://localhost:8080/debug/pprof/profile?seconds=5")

	// 1. Запускаем симуляцию нагрузки на CPU
	go cpuHeavy()

	// 2. Запускаем симуляцию неочищаемой памяти
	go memoryLeaky()

	// 3. Запускаем "зомби" горутины, которые висят навсегда
	go goroutineLeak()

	// 4. Запускаем войну горутин за один замок (Mutex)
	go mutexContention()

	// Запускаем HTTP сервер на 8080 (он заблокирует поток main и не даст программе завершиться)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// ----------------------------------------------------
// 1. Нагрузка на CPU: бесконечные тяжелые вычисления
// ----------------------------------------------------
func cpuHeavy() {
	for {
		_ = fibonacci(30)
	}
}

// Рекурсивный Фибоначчи сильно грузит процессор
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// ----------------------------------------------------
// 2. Утечка памяти (Memory Leak)
// ----------------------------------------------------
// Глобальная переменная, которая будет держать объекты в памяти,
// не давая Сборщику Мусора (GC) их очистить
var leakyData [][]byte

func memoryLeaky() {
	for {
		// Выделяем по 1 Мегабайту каждые 100 миллисекунд
		data := make([]byte, 1024*1024)
		leakyData = append(leakyData, data)
		time.Sleep(100 * time.Millisecond)
	}
}

// ----------------------------------------------------
// 3. Утечка Горутин (Goroutine Leak)
// ----------------------------------------------------
func goroutineLeak() {
	for i := 0; i < 100; i++ {
		go func() {
			ch := make(chan int)
			<-ch // Горутина блокируется раз и навсегда (ожидает из пустого канала без писателей)
		}()
	}
}

// ----------------------------------------------------
// 4. Очередь за Мьютексом (Mutex Contention)
// ----------------------------------------------------
func mutexContention() {
	var mu sync.Mutex

	for i := 0; i < 10; i++ {
		go func() {
			for {
				mu.Lock()
				// Горутина нагло держит замок 10мс, пока остальные 9 ждут в очереди
				time.Sleep(10 * time.Millisecond)
				mu.Unlock()
			}
		}()
	}
}
