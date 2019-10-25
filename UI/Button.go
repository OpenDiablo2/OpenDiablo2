package UI

// Button defines an object that acts like a button
type Button interface {
	Widget
	isPressed() bool
	setPressed(bool)
}
