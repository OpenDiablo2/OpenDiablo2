package d2ui

import (
	"strings"
	"time"

	"github.com/hajimehoshi/ebiten/inpututil"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2render"
	"github.com/hajimehoshi/ebiten"
)

// TextBox represents a text input box
type TextBox struct {
	text      string
	x         int
	y         int
	visible   bool
	enabled   bool
	bgSprite  d2render.Sprite
	textLabel Label
	lineBar   Label
}

func CreateTextbox(fileProvider d2interface.FileProvider) TextBox {
	result := TextBox{
		bgSprite:  d2render.CreateSprite(fileProvider.LoadFile(d2resource.TextBox2), d2datadict.Palettes[d2enum.Units]),
		textLabel: CreateLabel(fileProvider, d2resource.FontFormal11, d2enum.Units),
		lineBar:   CreateLabel(fileProvider, d2resource.FontFormal11, d2enum.Units),
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

func (v TextBox) Draw(target *ebiten.Image) {
	if !v.visible {
		return
	}
	v.bgSprite.Draw(target)
	v.textLabel.Draw(target)
	if (time.Now().UnixNano()/1e6)&(1<<8) > 0 {
		v.lineBar.Draw(target)
	}
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

func (v TextBox) GetText() string {
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
		v.lineBar.MoveTo(v.x+6+int(tw), v.y+3)
		v.textLabel.SetText(result)
		break
	}
}

func (v TextBox) GetSize() (width, height uint32) {
	return v.bgSprite.GetSize()
}

func (v *TextBox) MoveTo(x, y int) {
	v.x = x
	v.y = y
	v.textLabel.MoveTo(v.x+6, v.y+3)
	v.lineBar.MoveTo(v.x+6+int(v.textLabel.Width), v.y+3)
	v.bgSprite.MoveTo(v.x, v.y+26)
}

func (v TextBox) GetLocation() (x, y int) {
	return v.x, v.y
}

func (v TextBox) GetVisible() bool {
	return v.visible
}

func (v *TextBox) SetVisible(visible bool) {
	v.visible = visible
}

func (v TextBox) GetEnabled() bool {
	return v.enabled
}

func (v *TextBox) SetEnabled(enabled bool) {
	v.enabled = enabled
}

func (v *TextBox) SetPressed(pressed bool) {
	// no op
}

func (v TextBox) GetPressed() bool {
	return false
}

func (v *TextBox) OnActivated(callback func()) {
	// no op
}

func (v TextBox) Activate() {
	//no op
}
