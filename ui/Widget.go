package ui

import (
	"github.com/OpenDiablo2/OpenDiablo2/common"
)

// Widget defines an object that is a UI widget
type Widget interface {
	common.Drawable
	GetEnabled() bool
	SetEnabled(enabled bool)
	SetPressed(pressed bool)
	GetPressed() bool
	OnActivated(callback func())
	Activate()
}
