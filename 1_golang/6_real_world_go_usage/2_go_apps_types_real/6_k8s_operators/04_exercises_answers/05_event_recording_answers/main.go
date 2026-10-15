// Упражнение 05: запись событий с дедупликацией (Решение).
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
	key := object + "\x00" + typ + "\x00" + reason + "\x00" + message
	if e, ok := r.index[key]; ok {
		e.Count++
		return
	}
	e := &event{Object: object, Type: typ, Reason: reason, Message: message, Count: 1}
	r.index[key] = e
	r.events = append(r.events, e)
}

func main() {
	fmt.Println("--- Проверка 05_event_recording (Решение) ---")

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
