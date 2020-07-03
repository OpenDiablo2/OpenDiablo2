package d2interface

// Priority of the event handler
type Priority int

//noinspection GoUnusedConst // nothing is low priority yet
const (
	// PriorityLow is a low priority handler
	PriorityLow Priority = iota
	// PriorityDefault is a default priority handler
	PriorityDefault
	// PriorityHigh is a high priority handler
	PriorityHigh
)
