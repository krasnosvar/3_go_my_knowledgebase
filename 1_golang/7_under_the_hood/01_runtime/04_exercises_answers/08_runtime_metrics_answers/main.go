// Упражнение 08: runtime/metrics (Решение).
package main

import (
	"fmt"
	"runtime"
	"runtime/metrics"
)

// readUint64Metric читает одну Uint64-метрику runtime по её имени.
func readUint64Metric(name string) (uint64, bool) {
	s := []metrics.Sample{{Name: name}}
	metrics.Read(s) // в отличие от ReadMemStats, не останавливает мир
	if s[0].Value.Kind() != metrics.KindUint64 {
		return 0, false
	}
	return s[0].Value.Uint64(), true
}

func main() {
	fmt.Println("--- Проверка 08_runtime_metrics (Решение) ---")

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
