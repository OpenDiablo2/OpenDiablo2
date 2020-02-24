package d2event

import (
	"errors"
	"log"
	"github.com/kataras/go-events"
)

var (
	ErrWasInit         = errors.New("Event bus is already initialized")
	ErrNotInit         = errors.New("Event bus has not been initialized")
)

var singleton events.EventEmmiter

func Initialize() error {
	verifyNotInit()

	singleton = events.New()
	log.Printf("Initialized the Event Bus...")
	return nil
}

func Emit(key string, data ...interface{}) {
	singleton.Emit(events.EventName(key), data...)
}

func On(key string, fn func(payload ...interface{})) {
	singleton.On(events.EventName(key), fn)
}

func verifyWasInit() {
	if singleton == nil {
		panic(ErrNotInit)
	}
}

func verifyNotInit() {
	if singleton != nil {
		panic(ErrWasInit)
	}
}
