// Упражнение 05: запись событий с дедупликацией.
//
// Kubernetes склеивает одинаковые события: вместо тысячи записей
// "BackOff" одна запись с Count=1000.
//
// Задача: реализуйте (r *recorder) Event(object, typ, reason, message string):
//   - ключ события — (object, typ, reason, message);
//   - новое событие → добавить в r.events с Count = 1;
//   - повтор → увеличить Count существующего (порядок r.events не меняется).
//
// Запуск: go run ./03_exercises/05_event_recording
package main

import "fmt"

type event struct {
	Object, Type, Reason, Message string
	Count                         int
}

type recorder struct {
	events []*event
	index  map[string]*event
}

func newRecorder() *recorder { return &recorder{index: map[string]*event{}} }

func (r *recorder) Event(object, typ, reason, message string) {
	// TODO: найдите событие по ключу в r.index, иначе создайте.
}

func main() {
	fmt.Println("--- Проверка 05_event_recording ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	r := newRecorder()
	for i := 0; i < 5; i++ {
		r.Event("pod/web-0", "Warning", "BackOff", "restarting failed container")
	}
	r.Event("pod/web-0", "Normal", "Pulled", "image pulled")
	r.Event("pod/web-1", "Warning", "BackOff", "restarting failed container")

	check(len(r.events) == 3, fmt.Sprintf("три разных события (получено %d)", len(r.events)))
	if len(r.events) == 3 {
		check(r.events[0].Count == 5 && r.events[0].Reason == "BackOff", "BackOff web-0 склеен: Count=5")
		check(r.events[1].Count == 1 && r.events[2].Object == "pod/web-1", "остальные события отдельно, порядок сохранён")
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
