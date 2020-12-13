package d2label

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"image/color"
	"regexp"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2bitmapfont"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
)

// New creates a new label, initializing the unexported fields
func New() *Label {
	return &Label{
		colors: map[int]color.Color{0: color.White},
		backgroundColor: color.Transparent,
	}
}

// Label represents a user interface label
type Label struct {
	dirty           bool // used to flag when to re-render the label
	text            string // has color tokens
	rawText         string // unmodified text
	Alignment       d2ui.HorizontalAlign
	Font            *d2bitmapfont.BitmapFont
	colors          map[int]color.Color
	backgroundColor color.Color
}

func (v *Label) Render(target d2interface.Surface) {
	lines := strings.Split(v.text, "\n")
	yOffset := 0

	lastColor := v.colors[0]
	v.Font.SetColor(lastColor)

	for _, line := range lines {
		lw, lh := v.GetTextMetrics(line)
		characters := []rune(line)

		if v.backgroundColor != nil {
			target.Clear(v.backgroundColor)
		}

		target.PushTranslation(v.GetAlignOffset(lw), yOffset)

		for idx := range characters {
			character := string(characters[idx])
			charWidth, _ := v.GetTextMetrics(character)

			if v.colors[idx] != nil {
				lastColor = v.colors[idx]
				v.Font.SetColor(lastColor)
			}

			_ = v.Font.RenderText(character, target)

			target.PushTranslation(charWidth, 0)
		}

		target.PopN(len(characters))

		yOffset += lh

		target.Pop()
	}

	v.dirty = false
}

// IsDirty returns if the label needs to be re-rendered
func (v *Label) IsDirty() bool {
	return v.dirty
}

// GetSize returns the size of the label
func (v *Label) GetSize() (width, height int) {
	return v.Font.GetTextMetrics(v.text)
}

// GetTextMetrics returns the width and height of the enclosing rectangle in Pixels.
func (v *Label) GetTextMetrics(text string) (width, height int) {
	return v.Font.GetTextMetrics(text)
}

// SetText sets the label's text
func (v *Label) SetText(newText string) {
	if v.rawText == newText {
		return
	}

	v.rawText = newText
	v.dirty = true
	v.text = v.processColorTokens(newText)
}

// GetText returns the label's rawText
func (v *Label) GetText() string {
	return v.rawText
}

// SetBackgroundColor sets the background highlight color
func (v *Label) SetBackgroundColor(c color.Color) {
	r1, g1, b1, a1 := c.RGBA()
	r2, g2, b2, a2 := v.backgroundColor.RGBA()

	if (r1==r2) && (g1==g2) && (b1==b2) && (a1==a2) {
		return
	}

	v.dirty = true
	v.backgroundColor = c
}

func (v *Label) processColorTokens(str string) string {
	tokenMatch := regexp.MustCompile(d2ui.ColorTokenMatch)
	tokenStrMatch := regexp.MustCompile(d2ui.ColorStrMatch)
	empty := []byte("")

	tokenPosition := 0

	withoutTokens := string(tokenMatch.ReplaceAll([]byte(str), empty)) // remove tokens from string

	matches := tokenStrMatch.FindAll([]byte(str), -1)

	if len(matches) == 0 {
		v.colors[0] = getColor(d2ui.ColorTokenWhite)
	}

	// we find the index of each token and update the color map.
	// the key in the map is the starting index of each color token, the value is the color
	for idx := range matches {
		match := matches[idx]
		matchToken := tokenMatch.Find(match)
		matchStr := string(tokenMatch.ReplaceAll(match, empty))
		token := d2ui.ColorToken(matchToken)

		theColor := getColor(token)
		if theColor == nil {
			continue
		}

		if v.colors == nil {
			v.colors = make(map[int]color.Color)
		}

		v.colors[tokenPosition] = theColor

		tokenPosition += len(matchStr)
	}

	return withoutTokens
}

func (v *Label) GetAlignOffset(textWidth int) int {
	switch v.Alignment {
	case d2ui.HorizontalAlignLeft:
		return 0
	case d2ui.HorizontalAlignCenter:
		return -textWidth / 2
	case d2ui.HorizontalAlignRight:
		return -textWidth
	default:
		return 0
	}
}

func getColor(token d2ui.ColorToken) color.Color {
	// https://github.com/OpenDiablo2/OpenDiablo2/issues/823
	colors := map[d2ui.ColorToken]color.Color{
		d2ui.ColorTokenGrey:   d2util.Color(d2ui.ColorGrey100Alpha),
		d2ui.ColorTokenWhite:  d2util.Color(d2ui.ColorWhite100Alpha),
		d2ui.ColorTokenBlue:   d2util.Color(d2ui.ColorBlue100Alpha),
		d2ui.ColorTokenYellow: d2util.Color(d2ui.ColorYellow100Alpha),
		d2ui.ColorTokenGreen:  d2util.Color(d2ui.ColorGreen100Alpha),
		d2ui.ColorTokenGold:   d2util.Color(d2ui.ColorGold100Alpha),
		d2ui.ColorTokenOrange: d2util.Color(d2ui.ColorOrange100Alpha),
		d2ui.ColorTokenRed:    d2util.Color(d2ui.ColorRed100Alpha),
		d2ui.ColorTokenBlack:  d2util.Color(d2ui.ColorBlack100Alpha),
	}

	chosen := colors[token]

	if chosen == nil {
		return nil
	}

	return chosen
}
