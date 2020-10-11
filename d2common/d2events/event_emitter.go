package d2events

type EventEmitter struct {
	listeners map[string][]*EventListener
	count     int
}

func (ee *EventEmitter) Emit(event string, args ...interface{}) {
	listeners := ee.listeners[event]

	if listeners == nil {
		return
	}

	for idx := range listeners {
		if listeners[idx].fn != nil {
			listeners[idx].fn(args...)
		}

		if listeners[idx].once {
			listeners = append(listeners[:idx], listeners[idx+1:]...)
		}
	}
}

func (ee *EventEmitter) On(event string, fn func(...interface{})) {
	ee.addListener(event, fn, false)
}

func (ee *EventEmitter) Off(event string, fn func(...interface{})) {
	ee.removeListener(event, fn)
}

func (ee *EventEmitter) Once(event string, fn func(...interface{})) {
	ee.addListener(event, fn, true)
}

func (ee *EventEmitter) addListener(event string, fn func(...interface{}), once bool) *EventEmitter {
	if fn == nil {
		return ee
	}

	listener := &EventListener{fn, once}

	if ee.listeners[event] == nil {
		ee.listeners[event] = []*EventListener{listener}
	} else {
		ee.listeners[event] = append(ee.listeners[event], listener)
	}

	return ee
}

func (ee *EventEmitter) removeListener(event string, fn func(...interface{})) {
	listeners := ee.listeners[event]

	if listeners == nil {
		return
	}

	for idx := range listeners {
		listenerFn := &listeners[idx].fn
		removeFn := &fn
		if listenerFn == removeFn {
			ee.listeners[event] = append(listeners[:idx], listeners[idx+1:]...)
		}
	}
}

func (ee *EventEmitter) eventNames() []string {
	names := make([]string, len(ee.listeners))

	idx := 0

	for event := range ee.listeners {
		names[idx] = event
		idx++
	}

	return names
}

func (ee *EventEmitter) clearEvent(event string) {
	ee.count--
	if ee.count <= 0 {
		ee.count = 0
		ee.listeners = make(map[string][]*EventListener)

		return
	}

	delete(ee.listeners, event)
}

func (ee *EventEmitter) getHandlers(event string) []func(...interface{}) {
	handlers := make([]func(...interface{}), 0)
	listeners := ee.listeners[event]

	if listeners != nil {
		for idx := range listeners {
			handlers = append(handlers, listeners[idx].fn)
		}
	}

	return handlers
}

func (ee *EventEmitter) getHandlerCount(event string) int {
	if ee.listeners == nil {
		return 0
	}

	listeners := ee.listeners[event]
	if listeners == nil {
		return 0
	}

	return len(listeners)
}

func (ee *EventEmitter) removeAllListeners(events ...string) {
	if events != nil {
		if len(events) > 0 {
			for idx := range events {
				ee.clearEvent(events[idx])
			}
		}

		return
	}

	ee.listeners = make(map[string][]*EventListener, 0)
}
