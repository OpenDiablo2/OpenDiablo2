package d2interface

import (
	"image/color"
)

// Font is a font
type Font interface {
	SetColor(c color.Color)
	GetTextMetrics(text string) (width, height int)
	Clone() Font
	RenderText(text string, target Surface) error
}
