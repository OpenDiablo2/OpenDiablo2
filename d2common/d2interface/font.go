package d2interface

import (
	"image/color"
)

// Font is a graphical representation associated with a set of glyphs.
type Font interface {
	SetColor(c color.Color)
	GetTextMetrics(text string) (width, height int)
	RenderText(text string, target Surface) error
}
