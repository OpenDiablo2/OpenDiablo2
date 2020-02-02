package d2ui

import (
	"strings"
	"time"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"

	"github.com/hajimehoshi/ebiten/inpututil"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/hajimehoshi/ebiten"
)

// TextBox represents a text input box
type TextBox struct {
	text      string
	x         int
	y         int
	visible   bool
	enabled   bool
	bgSprite  *Sprite
	textLabel Label
	lineBar   Label
}

func CreateTextbox() TextBox {
	animation, _ := d2asset.LoadAnimation(d2resource.TextBox2, d2resource.PaletteUnits)
	bgSprite, _ := LoadSprite(animation)
	result := TextBox{
		bgSprite:  bgSprite,
		textLabel: CreateLabel(d2resource.FontFormal11, d2resource.PaletteUnits),
		lineBar:   CreateLabel(d2resource.FontFormal11, d2resource.PaletteUnits),
		enabled:   true,
		visible:   true,
	}
	result.lineBar.SetText("_")
	return result
}

func repeatingKeyPressed(key ebiten.Key) bool {
	const (
		delay    = 30
		interval = 3
	)
	d := inpututil.KeyPressDuration(key)
	if d == 1 {
		return true
	}
	if d >= delay && (d-delay)%interval == 0 {
		return true
	}
	return false
}

func (v *TextBox) Render(target d2render.Surface) {
	if !v.visible {
		return
	}
	v.bgSprite.Render(target)
	v.textLabel.Render(target)
	if (time.Now().UnixNano()/1e6)&(1<<8) > 0 {
		v.lineBar.Render(target)
	}
}

func (v *TextBox) Advance(elapsed float64) {

}

func (v *TextBox) Update() {
	if !v.visible || !v.enabled {
		return
	}
	newText := string(ebiten.InputChars())
	if len(newText) > 0 {
		v.text += newText
		v.SetText(v.text)
	}
	if repeatingKeyPressed(ebiten.KeyBackspace) {
		if len(v.text) >= 1 {
			v.text = v.text[:len(v.text)-1]
		}
		v.SetText(v.text)
	}
}

func (v *TextBox) GetText() string {
	return v.text
}

func (v *TextBox) SetText(newText string) {
	result := ""
	for _, c := range newText {
		if !strings.Contains("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", string(c)) {
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

func (v *TextBox) GetSize() (width, height int) {
	return v.bgSprite.GetCurrentFrameSize()
}

func (v *TextBox) SetPosition(x, y int) {
	v.x = x
	v.y = y
	v.textLabel.SetPosition(v.x+6, v.y+3)
	v.lineBar.SetPosition(v.x+6+v.textLabel.Width, v.y+3)
	v.bgSprite.SetPosition(v.x, v.y+26)
}

func (v *TextBox) GetPosition() (x, y int) {
	return v.x, v.y
}

func (v *TextBox) GetVisible() bool {
	return v.visible
}

func (v *TextBox) SetVisible(visible bool) {
	v.visible = visible
}

func (v *TextBox) GetEnabled() bool {
	return v.enabled
}

func (v *TextBox) SetEnabled(enabled bool) {
	v.enabled = enabled
}

func (v *TextBox) SetPressed(pressed bool) {
	// no op
}

func (v *TextBox) GetPressed() bool {
	return false
}

func (v *TextBox) OnActivated(callback func()) {
	// no op
}

func (v *TextBox) Activate() {
	//no op
}
