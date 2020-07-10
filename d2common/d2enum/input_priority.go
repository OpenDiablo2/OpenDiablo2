package d2enum

// Priority of the event handler
type Priority int

// Priorities
const (
	PriorityLow Priority = iota
	PriorityDefault
	PriorityHigh
)
