package d2ui

import (
	"fmt"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2gui"
	"image/color"
	"log"
	"regexp"
	"strings"
)

// ColorToken is a string which is used inside of label strings to set font color.
type ColorToken string

const (
	colorTokenFmt   = `%s%s`
	colorTokenMatch = `\[[^\]]+\]`
	colorStrMatch   = colorTokenMatch + `[^\[]+`
)

const (
	ColorTokenGrey   ColorToken = "[grey]"
	ColorTokenRed    ColorToken = "[red]"
	ColorTokenWhite  ColorToken = "[white]"
	ColorTokenBlue   ColorToken = "[blue]"
	ColorTokenYellow ColorToken = "[yellow]"
	ColorTokenGreen  ColorToken = "[green]"
	ColorTokenGold   ColorToken = "[gold]"
	ColorTokenOrange ColorToken = "[orange]"
	ColorTokenBlack  ColorToken = "[black]"
)

// Color tokens for item labels
const (
	ColorTokenSocketedItem = ColorTokenGrey
	ColorTokenNormalItem   = ColorTokenWhite
	ColorTokenMagicItem    = ColorTokenBlue
	ColorTokenRareItem     = ColorTokenYellow
	ColorTokenSetItem      = ColorTokenGreen
	ColorTokenUniqueItem   = ColorTokenGold
	ColorTokenCraftedItem  = ColorTokenOrange
)

const (
	ColorTokenServer = ColorTokenRed
	ColorTokenButton = ColorTokenBlack
)

const (
	ColorTokenCharacterName = ColorTokenGold
	ColorTokenCharacterDesc = ColorTokenWhite
	ColorTokenCharacterType = ColorTokenGreen
)

// ColorTokenize formats the string with the given color token
func ColorTokenize(s string, t ColorToken) string {
	return fmt.Sprintf(colorTokenFmt, t, s)
}

// Label represents a user interface label
type Label struct {
	text            string
	X               int
	Y               int
	Alignment       d2gui.HorizontalAlign
	font            d2interface.Font
	Color           map[int]color.Color
	backgroundColor color.Color
}

// CreateLabel creates a new instance of a UI label
func CreateLabel(fontPath, palettePath string) Label {
	font, _ := d2asset.LoadFont(fontPath+".tbl", fontPath+".dc6", palettePath)
	result := Label{
		Alignment: d2gui.HorizontalAlignLeft,
		Color:     map[int]color.Color{0: color.White},
		font:      font,
	}

	return result
}

// Render draws the label on the screen, respliting the lines to allow for other alignments.
func (v *Label) Render(target d2interface.Surface) {
	target.PushTranslation(v.X, v.Y)

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

			_ = v.font.RenderText(character, target)

			target.PushTranslation(charWidth, 0)
		}

		target.PopN(len(characters))

		yOffset += lh

		target.Pop()
	}

	target.Pop()
}

// SetPosition moves the label to the specified location
func (v *Label) SetPosition(x, y int) {
	v.X = x
	v.Y = y
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
func (v *Label) SetBackgroundColor(c color.RGBA) {
	v.backgroundColor = c
}

func (v *Label) processColorTokens(str string) string {
	tokenMatch := regexp.MustCompile(colorTokenMatch)
	tokenStrMatch := regexp.MustCompile(colorStrMatch)
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
	case d2gui.HorizontalAlignLeft:
		return 0
	case d2gui.HorizontalAlignCenter:
		return -textWidth / 2
	case d2gui.HorizontalAlignRight:
		return -textWidth
	default:
		log.Fatal("Invalid Alignment")
		return 0
	}
}

func getColor(token ColorToken) color.Color {
	alpha := uint8(255)

	// todo this should really come from the PL2 files
	colors := map[ColorToken]color.Color{
		ColorTokenGrey:   color.RGBA{105, 105, 105, alpha},
		ColorTokenWhite:  color.RGBA{255, 255, 255, alpha},
		ColorTokenBlue:   color.RGBA{105, 105, 255, alpha},
		ColorTokenYellow: color.RGBA{255, 255, 100, alpha},
		ColorTokenGreen:  color.RGBA{0, 255, 0, alpha},
		ColorTokenGold:   color.RGBA{199, 179, 119, alpha},
		ColorTokenOrange: color.RGBA{255, 168, 0, alpha},
		ColorTokenRed:    color.RGBA{255, 77, 77, alpha},
		ColorTokenBlack:    color.RGBA{0, 0, 0, alpha},
	}

	chosen := colors[token]

	if chosen == nil {
		return colors[ColorTokenWhite]
	}

	return chosen
}
