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
	OutputRaw(text string, category d2enum.TermCategory)
	Outputf(format string, params ...interface{})
	OutputInfof(format string, params ...interface{})
	OutputWarningf(format string, params ...interface{})
	OutputErrorf(format string, params ...interface{})
	OutputClear()
	IsVisible() bool
	Hide()
	Show()
	BindAction(name, description string, action interface{}) error
	UnbindAction(name string) error
}

// TerminalLogger is used tomake the Terminal write out
// (eg. to the system shell or to a file)
type TerminalLogger interface {
	Write(p []byte) (int, error)
}
