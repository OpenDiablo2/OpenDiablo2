package d2ui

import (
	"strings"
	"time"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
)

// TextBox represents a text input box
type TextBox struct {
	textLabel Label
	lineBar   Label
	text      string
	filter    string
	x         int
	y         int
	bgSprite  *Sprite
	visible   bool
	enabled   bool
	isFocused bool
}

// CreateTextbox creates a new instance of a text box
func CreateTextbox() TextBox {
	animation, _ := d2asset.LoadAnimation(d2resource.TextBox2, d2resource.PaletteUnits)
	bgSprite, _ := LoadSprite(animation)
	tb := TextBox{
		filter:    "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
		bgSprite:  bgSprite,
		textLabel: CreateLabel(d2resource.FontFormal11, d2resource.PaletteUnits),
		lineBar:   CreateLabel(d2resource.FontFormal11, d2resource.PaletteUnits),
		enabled:   true,
		visible:   true,
	}
	tb.lineBar.SetText("_")

	return tb
}

// SetFilter sets the text box filter
func (v *TextBox) SetFilter(filter string) {
	v.filter = filter
}

// Render renders the text box
func (v *TextBox) Render(target d2interface.Surface) error {
	if !v.visible {
		return nil
	}

	if err := v.bgSprite.Render(target); err != nil {
		return err
	}

	v.textLabel.Render(target)

	if (time.Now().UnixNano()/1e6)&(1<<8) > 0 {
		v.lineBar.Render(target)
	}

	return nil
}

// OnKeyChars handles key character events
func (v *TextBox) OnKeyChars(event d2interface.KeyCharsEvent) bool {
	if !v.isFocused || !v.visible || !v.enabled {
		return false
	}

	newText := string(event.Chars())

	if len(newText) > 0 {
		v.text += newText

		v.SetText(v.text)

		return true
	}

	return false
}

// OnKeyRepeat handles key repeat events
func (v *TextBox) OnKeyRepeat(event d2interface.KeyEvent) bool {
	if event.Key() == d2enum.KeyBackspace && debounceEvents(event.Duration()) {
		if len(v.text) >= 1 {
			v.text = v.text[:len(v.text)-1]
		}

		v.SetText(v.text)
	}

	return false
}

func debounceEvents(numFrames int) bool {
	const (
		delay    = 30
		interval = 3
	)

	if numFrames == 1 {
		return true
	}

	if numFrames >= delay && (numFrames-delay)%interval == 0 {
		return true
	}

	return false
}

// Advance updates the text box
func (v *TextBox) Advance(_ float64) {
	if !v.visible || !v.enabled {
		return
	}
}

// Update updates the textbox
func (v *TextBox) Update() {
}

// GetText returns the text box's text
func (v *TextBox) GetText() string {
	return v.text
}

// SetText sets the text box's text
//nolint:gomnd // Built-in values
func (v *TextBox) SetText(newText string) {
	result := ""

	for _, c := range newText {
		if !strings.Contains(v.filter, string(c)) {
			continue
		}

		result += string(c)
	}

	if len(result) > 15 {
		result = result[0:15]
	}

	v.text = result

	for {
		tw, _ := v.textLabel.GetTextMetrics(result)

		if tw > 150 {
			result = result[1:]
			continue
		}

		v.lineBar.SetPosition(v.x+6+tw, v.y+3)
		v.textLabel.SetText(result)

		break
	}
}

// GetSize returns the size of the text box
func (v *TextBox) GetSize() (width, height int) {
	return v.bgSprite.GetCurrentFrameSize()
}

// SetPosition sets the position of the text box
//nolint:gomnd // Built-in values
func (v *TextBox) SetPosition(x, y int) {
	lw, _ := v.textLabel.GetSize()

	v.x = x
	v.y = y

	v.textLabel.SetPosition(v.x+6, v.y+3)
	v.lineBar.SetPosition(v.x+6+lw, v.y+3)
	v.bgSprite.SetPosition(v.x, v.y+26)
}

// GetPosition returns the position of the text box
func (v *TextBox) GetPosition() (x, y int) {
	return v.x, v.y
}

// GetVisible returns the visibility of the text box
func (v *TextBox) GetVisible() bool {
	return v.visible
}

// SetVisible sets the visibility of the text box
func (v *TextBox) SetVisible(visible bool) {
	v.visible = visible
}

// GetEnabled returns the enabled state of the text box
func (v *TextBox) GetEnabled() bool {
	return v.enabled
}

// SetEnabled sets the enabled state of the text box
func (v *TextBox) SetEnabled(enabled bool) {
	v.enabled = enabled
}

// SetPressed does nothing for text boxes
func (v *TextBox) SetPressed(_ bool) {
	// no op
}

// GetPressed does nothing for text boxes
func (v *TextBox) GetPressed() bool {
	return false
}

// OnActivated handles activation events for the text box
func (v *TextBox) OnActivated(_ func()) {
	// no op
}

// Activate activates the text box
func (v *TextBox) Activate() {
	v.isFocused = true
}
