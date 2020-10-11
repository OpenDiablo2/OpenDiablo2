package d2events

// NewEventEmitter initializes and returns an EventEmitter instance
func NewEventEmitter() *EventEmitter {
	ee := &EventEmitter{
		listeners: make(map[string][]*EventListener),
		count:     0,
	}

	return ee
}
