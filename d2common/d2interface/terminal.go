package d2interface

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2input"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
)

type TermCategory int

const (
	TermCategoryNone TermCategory = iota
	TermCategoryInfo
	TermCategoryWarning
	TermCategoryError
)

const (
	termCharWidth   = 6
	termCharHeight  = 16
	termRowCount    = 24
	termRowCountMax = 32
	termColCountMax = 128
	termAnimLength  = 0.5
)

type termVis int

const (
	termVisHidden termVis = iota
	termVisShowing
	termVisShown
	termVisHiding
)

type Terminal interface {
	BindLogger()

	Advance(elapsed float64) error
	OnKeyDown(event d2input.KeyEvent) bool
	OnKeyChars(event d2input.KeyCharsEvent) bool
	Render(surface d2render.Surface) error
	Execute(command string) error
	OutputRaw(text string, category TermCategory)
	Output(format string, params ...interface{})
	OutputInfo(format string, params ...interface{})
	OutputWarning(format string, params ...interface{})
	OutputError(format string, params ...interface{})
	OutputClear()
	IsVisible() bool
	Hide()
	Show()
	BindAction(name, description string, action interface{}) error
	UnbindAction(name string) error
}

type TerminalLogger interface {
	Write(p []byte) (int, error)
}
