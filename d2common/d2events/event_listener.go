package d2events

type EventListener struct {
	fn   func(...interface{})
	once bool
}
