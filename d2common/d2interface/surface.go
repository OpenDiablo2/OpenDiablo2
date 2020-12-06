package d2interface

import (
	"image"
	"image/color"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

// Surface represents a renderable surface.
type Surface interface {
	Renderer() Renderer
	Clear(color color.Color)
	DrawRect(width, height int, color color.Color)
	DrawLine(x, y int, color color.Color)
	DrawTextf(format string, params ...interface{})
	GetSize() (width, height int)
	GetDepth() int
	Pop()
	PopN(n int)
	PushColor(color color.Color)
	PushEffect(effect d2enum.DrawEffect)
	PushFilter(filter d2enum.Filter)
	PushTranslation(x, y int)
	PushSkew(x, y float64)
	PushScale(x, y float64)
	PushRotate(theta float64)
	PushBrightness(brightness float64)
	PushSaturation(saturation float64)
	Render(surface Surface)
	// Renders a section of the surface enclosed by bounds
	RenderSection(surface Surface, bound image.Rectangle)
	ReplacePixels(pixels []byte)
	Screenshot() *image.RGBA
}
