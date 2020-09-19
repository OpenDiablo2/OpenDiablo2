package d2records

// Events holds all of the event records from events.txt
type Events []*EventRecord

// EventRecord is a representation of a single row from events.txt
type EventRecord struct {
	Event string
}
