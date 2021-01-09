package d2interface

import "github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"

// Terminal is a drop-down terminal and shell
// It is used throughout the codebase, most parts of the engine will
// `bind` commands, which are available for use in the shell
type Terminal interface {
	BindLogger()

	Advance(elapsed float64) error
	OnKeyDown(event KeyEvent) bool
	OnKeyChars(event KeyCharsEvent) bool
	Render(surface Surface) error
	Execute(command string) error
	Rawf(category d2enum.TermCategory, format string, params ...interface{})
	Printf(format string, params ...interface{})
	Infof(format string, params ...interface{})
	Warningf(format string, params ...interface{})
	Errorf(format string, params ...interface{})
	Clear()
	Visible() bool
	Hide()
	Show()
	Bind(name, description string, arguments []string, fn func(args []string) error) error
	Unbind(name ...string) error
}

// TerminalLogger is used tomake the Terminal write out
// (eg. to the system shell or to a file)
type TerminalLogger interface {
	Write(p []byte) (int, error)
}
