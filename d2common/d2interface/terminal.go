package d2interface

// TermCategory applies styles to the lines in the  Terminal
type TermCategory int

// Terminal Category types
const (
	TermCategoryNone TermCategory = iota
	TermCategoryInfo
	TermCategoryWarning
	TermCategoryError
)

// Terminal is a drop-down terminal and shell
// It is used throughout the codebase, most parts of the engine will
// `bind` commands, which are available for use in the shell
type Terminal interface {
	AppComponent
	BindLogger()
	OnKeyDown(event KeyEvent) bool
	OnKeyChars(event KeyCharsEvent) bool
	Execute(command string) error
	OutputRaw(text string, category TermCategory)
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
