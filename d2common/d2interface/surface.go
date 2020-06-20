package d2interface

import (
	"image"
	"image/color"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

// Surface represents a renderable surface.
type Surface interface {
	Clear(color color.Color) error
	DrawRect(width, height int, color color.Color)
	DrawLine(x, y int, color color.Color)
	DrawText(format string, params ...interface{})
	GetSize() (width, height int)
	GetDepth() int
	Pop()
	PopN(n int)
	PushColor(color color.Color)
	PushCompositeMode(mode d2enum.CompositeMode)
	PushFilter(filter Filter)
	PushTranslation(x, y int)
	PushBrightness(brightness float64)
	Render(surface Surface) error
	// Renders a section of the surface enclosed by bounds
	RenderSection(surface Surface, bound image.Rectangle) error
	ReplacePixels(pixels []byte) error
	Screenshot() *image.RGBA
}
