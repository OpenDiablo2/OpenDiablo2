package mouse

type MouseButton int

const (
	ButtonLeft = MouseButton(iota)
	ButtonMiddle
	ButtonRight
)

var MouseButtons = [...]MouseButton{
	ButtonLeft,
	ButtonMiddle,
	ButtonRight,
}
