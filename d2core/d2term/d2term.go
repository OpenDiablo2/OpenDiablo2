package d2term

import (
	"errors"
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2input"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
)

var (
	ErrHasInit = errors.New("terminal system is already initialized")
	ErrNotInit = errors.New("terminal system is not initialized")
)

var singleton *terminal

func Initialize() error {
	if singleton != nil {
		return ErrHasInit
	}

	terminal, err := createTerminal()
	if err != nil {
		return err
	}

	if err := d2input.BindHandlerWithPriority(terminal, d2input.PriorityHigh); err != nil {
		log.Println(err)
		return err
	}

	singleton = terminal
	return nil
}

func Shutdown() {
	if singleton != nil {
		d2input.UnbindHandler(singleton)
		singleton = nil
	}
}

func Advance(elapsed float64) error {
	if singleton == nil {
		return ErrNotInit
	}

	if singleton != nil {
		return singleton.advance(elapsed)
	}

	return ErrNotInit
}

func Output(format string, params ...interface{}) error {
	if singleton == nil {
		return ErrNotInit
	}

	return singleton.output(format, params...)
}

func OutputInfo(format string, params ...interface{}) error {
	if singleton == nil {
		return ErrNotInit
	}

	return singleton.outputInfo(format, params...)
}

func OutputWarning(format string, params ...interface{}) error {
	if singleton == nil {
		return ErrNotInit
	}

	return singleton.outputWarning(format, params...)
}

func OutputError(format string, params ...interface{}) error {
	if singleton == nil {
		return ErrNotInit
	}

	return singleton.outputError(format, params...)
}

func BindAction(name, description string, action interface{}) error {
	if singleton == nil {
		return ErrNotInit
	}

	return singleton.bindAction(name, description, action)
}

func UnbindAction(name string) error {
	if singleton == nil {
		return ErrNotInit
	}

	return singleton.unbindAction(name)
}

func Render(surface d2render.Surface) error {
	if singleton == nil {
		return ErrNotInit
	}

	return singleton.render(surface)
}

func BindLogger() {
	log.SetOutput(&terminalLogger{writer: log.Writer()})
}
