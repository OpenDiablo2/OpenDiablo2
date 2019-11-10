package ui

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

// Widget defines an object that is a UI widget
type Widget interface {
	d2interface.Drawable
	GetEnabled() bool
	SetEnabled(enabled bool)
	SetPressed(pressed bool)
	GetPressed() bool
	OnActivated(callback func())
	Activate()
}
