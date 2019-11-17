package d2ui

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2corecommon/d2coreinterface"
)

// Widget defines an object that is a UI widget
type Widget interface {
	d2coreinterface.Drawable
	GetEnabled() bool
	SetEnabled(enabled bool)
	SetPressed(pressed bool)
	GetPressed() bool
	OnActivated(callback func())
	Activate()
}
