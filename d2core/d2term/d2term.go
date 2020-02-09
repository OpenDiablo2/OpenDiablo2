package d2term

import (
	"errors"
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2input"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
)

var (
	ErrWasInit = errors.New("terminal system is already initialized")
	ErrNotInit = errors.New("terminal system is not initialized")
)

var singleton *terminal

func Initialize() error {
	verifyNotInit()

	terminal, err := createTerminal()
	if err != nil {
		return err
	}

	if err := d2input.BindHandlerWithPriority(terminal, d2input.PriorityHigh); err != nil {
		return err
	}

	singleton = terminal
	return nil
}

func Advance(elapsed float64) error {
	verifyWasInit()
	return singleton.advance(elapsed)
}

func Output(format string, params ...interface{}) {
	verifyWasInit()
	singleton.output(format, params...)
}

func OutputInfo(format string, params ...interface{}) {
	verifyWasInit()
	singleton.outputInfo(format, params...)
}

func OutputWarning(format string, params ...interface{}) {
	verifyWasInit()
	singleton.outputWarning(format, params...)
}

func OutputError(format string, params ...interface{}) {
	verifyWasInit()
	singleton.outputError(format, params...)
}

func BindAction(name, description string, action interface{}) error {
	verifyWasInit()
	return singleton.bindAction(name, description, action)
}

func UnbindAction(name string) error {
	verifyWasInit()
	return singleton.unbindAction(name)
}

func Render(surface d2render.Surface) error {
	verifyWasInit()
	return singleton.render(surface)
}

func BindLogger() {
	log.SetOutput(&terminalLogger{writer: log.Writer()})
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
