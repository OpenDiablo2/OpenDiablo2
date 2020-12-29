package d2interface

const numColors = 256

// Color represents a color
type Color interface {
	R() uint8
	G() uint8
	B() uint8
	A() uint8
	RGBA() uint32
	SetRGBA(uint32)
	BGRA() uint32
	SetBGRA(uint32)
}

// Palette is a color palette
type Palette interface {
	NumColors() int
	GetColors() [numColors]Color
	GetColor(idx int) (Color, error)
}
