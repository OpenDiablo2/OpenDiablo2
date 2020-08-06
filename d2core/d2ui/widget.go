package d2ui

// Widget defines an object that is a UI widget
type Widget interface {
	Drawable
	bindManager(ui *UIManager)
	GetEnabled() bool
	SetEnabled(enabled bool)
	SetPressed(pressed bool)
	GetPressed() bool
	OnActivated(callback func())
	Activate()
}
