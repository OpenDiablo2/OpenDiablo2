package d2ui

import (
	"image/color"
	"regexp"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
)

// Label represents a user interface label
type Label struct {
	*BaseWidget
	text            string
	Alignment       HorizontalAlign
	font            *d2asset.Font
	Color           map[int]color.Color
	backgroundColor color.Color

	*d2util.Logger
}

// NewLabel creates a new instance of a UI label
func (ui *UIManager) NewLabel(fontPath, palettePath string) *Label {
	font, err := ui.asset.LoadFont(fontPath+".tbl", fontPath+".dc6", palettePath)
	if err != nil {
		ui.Error(err.Error())
		return nil
	}

	base := NewBaseWidget(ui)

	result := &Label{
		BaseWidget: base,
		Alignment:  HorizontalAlignLeft,
		Color:      map[int]color.Color{0: color.White},
		font:       font,
		Logger:     ui.Logger,
	}

	result.bindManager(ui)

	return result
}

// Render draws the label on the screen, respliting the lines to allow for other alignments.
func (v *Label) Render(target d2interface.Surface) {
	target.PushTranslation(v.GetPosition())

	lines := strings.Split(v.text, "\n")
	yOffset := 0

	lastColor := v.Color[0]
	v.font.SetColor(lastColor)

	for _, line := range lines {
		lw, lh := v.GetTextMetrics(line)
		characters := []rune(line)

		target.PushTranslation(v.getAlignOffset(lw), yOffset)

		for idx := range characters {
			character := string(characters[idx])
			charWidth, charHeight := v.GetTextMetrics(character)

			if v.Color[idx] != nil {
				lastColor = v.Color[idx]
				v.font.SetColor(lastColor)
			}

			if v.backgroundColor != nil {
				target.DrawRect(charWidth, charHeight, v.backgroundColor)
			}

			err := v.font.RenderText(character, target)
			if err != nil {
				v.Error(err.Error())
			}

			target.PushTranslation(charWidth, 0)
		}

		target.PopN(len(characters))

		yOffset += lh

		target.Pop()
	}

	target.Pop()
}

// GetSize returns the size of the label
func (v *Label) GetSize() (width, height int) {
	return v.font.GetTextMetrics(v.text)
}

// GetTextMetrics returns the width and height of the enclosing rectangle in Pixels.
func (v *Label) GetTextMetrics(text string) (width, height int) {
	return v.font.GetTextMetrics(text)
}

// SetText sets the label's text
func (v *Label) SetText(newText string) {
	v.text = v.processColorTokens(newText)
}

// SetBackgroundColor sets the background highlight color
func (v *Label) SetBackgroundColor(c color.Color) {
	v.backgroundColor = c
}

func (v *Label) processColorTokens(str string) string {
	tokenMatch := regexp.MustCompile(ColorTokenMatch)
	tokenStrMatch := regexp.MustCompile(ColorStrMatch)
	empty := []byte("")

	tokenPosition := 0

	withoutTokens := string(tokenMatch.ReplaceAll([]byte(str), empty)) // remove tokens from string

	matches := tokenStrMatch.FindAll([]byte(str), -1)

	if len(matches) == 0 {
		v.Color[0] = getColor(ColorTokenWhite)
	}

	// we find the index of each token and update the color map.
	// the key in the map is the starting index of each color token, the value is the color
	for idx := range matches {
		match := matches[idx]
		matchToken := tokenMatch.Find(match)
		matchStr := string(tokenMatch.ReplaceAll(match, empty))
		token := ColorToken(matchToken)

		theColor := getColor(token)
		if theColor == nil {
			continue
		}

		if v.Color == nil {
			v.Color = make(map[int]color.Color)
		}

		v.Color[tokenPosition] = theColor

		tokenPosition += len(matchStr)
	}

	return withoutTokens
}

func (v *Label) getAlignOffset(textWidth int) int {
	switch v.Alignment {
	case HorizontalAlignLeft:
		return 0
	case HorizontalAlignCenter:
		return -textWidth / 2
	case HorizontalAlignRight:
		return -textWidth
	default:
		v.Fatal("Invalid Alignment")
		return 0
	}
}

// Advance is a no-op
func (v *Label) Advance(elapsed float64) error {
	return nil
}

func getColor(token ColorToken) color.Color {
	// https://github.com/OpenDiablo2/OpenDiablo2/issues/823
	colors := map[ColorToken]color.Color{
		ColorTokenGrey:   d2util.Color(ColorGrey100Alpha),
		ColorTokenWhite:  d2util.Color(ColorWhite100Alpha),
		ColorTokenBlue:   d2util.Color(ColorBlue100Alpha),
		ColorTokenYellow: d2util.Color(ColorYellow100Alpha),
		ColorTokenGreen:  d2util.Color(ColorGreen100Alpha),
		ColorTokenGold:   d2util.Color(ColorGold100Alpha),
		ColorTokenOrange: d2util.Color(ColorOrange100Alpha),
		ColorTokenRed:    d2util.Color(ColorRed100Alpha),
		ColorTokenBlack:  d2util.Color(ColorBlack100Alpha),
	}

	chosen := colors[token]

	if chosen == nil {
		return nil
	}

	return chosen
}
