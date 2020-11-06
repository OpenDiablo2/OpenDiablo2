package d2ui

// Widget defines an object that is a UI widget
type Widget interface {
	Drawable
	bindManager(ui *UIManager)
}

// ClickableWidget defines an object that can be clicked
type ClickableWidget interface {
	Widget
	SetEnabled(enabled bool)
	SetPressed(pressed bool)
	GetEnabled() bool
	GetPressed() bool
	OnActivated(callback func())
	Activate()
}
