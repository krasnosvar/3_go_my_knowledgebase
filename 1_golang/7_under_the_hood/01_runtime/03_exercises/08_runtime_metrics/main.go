// Упражнение 08: runtime/metrics.
//
// Задача: реализуйте readUint64Metric(name) через пакет runtime/metrics:
// создайте []metrics.Sample с одним элементом Name: name, вызовите
// metrics.Read и верните (Value.Uint64(), true). Если метрика не
// поддерживается (Value.Kind() == metrics.KindBad) или имеет не Uint64-тип,
// верните (0, false).
//
// Проверка: "/sched/goroutines:goroutines" не меньше runtime.NumGoroutine()
// (метрика учитывает и системные горутины runtime: GC-воркеры и т.п.); "/gc/cycles/total:gc-cycles" растёт после runtime.GC();
// несуществующая метрика даёт ok == false.
//
// Запуск: go run ./03_exercises/08_runtime_metrics/main.go
package main

import (
	"fmt"
	"runtime"
	"runtime/metrics"
)

// readUint64Metric читает одну Uint64-метрику runtime по её имени.
func readUint64Metric(name string) (uint64, bool) {
	// TODO: metrics.Read со слайсом из одного metrics.Sample; проверьте Kind.
	_ = metrics.KindUint64
	return 0, false
}

func main() {
	fmt.Println("--- Проверка 08_runtime_metrics ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	g, ok := readUint64Metric("/sched/goroutines:goroutines")
	ng := runtime.NumGoroutine()
	check(ok && g >= uint64(ng), fmt.Sprintf("горутин по metrics=%d (ok=%t) должно быть >= NumGoroutine=%d", g, ok, ng))

	before, ok1 := readUint64Metric("/gc/cycles/total:gc-cycles")
	runtime.GC()
	after, ok2 := readUint64Metric("/gc/cycles/total:gc-cycles")
	check(ok1 && ok2 && after > before, fmt.Sprintf("циклы GC должны вырасти после runtime.GC(): было %d, стало %d", before, after))

	_, okBad := readUint64Metric("/no/such/metric:units")
	check(!okBad, "несуществующая метрика должна давать ok == false")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
