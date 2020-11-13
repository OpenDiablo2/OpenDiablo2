package d2player

import (
	"image/color"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2gui"
)

type bindingLayout struct {
	wrapperLayout   *d2gui.Layout
	descLayout      *d2gui.Layout
	descLabel       *d2gui.Label
	primaryLayout   *d2gui.Layout
	primaryLabel    *d2gui.Label
	secondaryLayout *d2gui.Layout
	secondaryLabel  *d2gui.Label

	binding   *KeyBinding
	gameEvent d2enum.GameEvent
}

func (l *bindingLayout) setTextAndColor(layout *d2gui.Label, text string, col color.RGBA) error {
	if err := layout.SetText(text); err != nil {
		return err
	}

	if err := layout.SetColor(col); err != nil {
		return err
	}

	return nil
}

func (l *bindingLayout) SetPrimaryBindingTextAndColor(text string, col color.RGBA) error {
	return l.setTextAndColor(l.primaryLabel, text, col)
}

func (l *bindingLayout) SetSecondaryBindingTextAndColor(text string, col color.RGBA) error {
	return l.setTextAndColor(l.secondaryLabel, text, col)
}

func (l *bindingLayout) Reset() error {
	if err := l.descLabel.SetIsHovered(false); err != nil {
		return err
	}

	if err := l.primaryLabel.SetIsHovered(false); err != nil {
		return err
	}

	if err := l.secondaryLabel.SetIsHovered(false); err != nil {
		return err
	}

	l.primaryLabel.SetIsBlinking(false)
	l.secondaryLabel.SetIsBlinking(false)

	return nil
}

func (l *bindingLayout) isInLayoutRect(x, y int, targetLayout *d2gui.Layout) bool {
	targetW, targetH := targetLayout.GetSize()
	targetX, targetY := targetLayout.Sx, targetLayout.Sy

	if x >= targetX && x <= targetX+targetW && y >= targetY && y <= targetY+targetH {
		return true
	}

	return false
}

func (l *bindingLayout) GetPointedLayoutAndLabel(x, y int) (d2enum.GameEvent, KeyBindingType) {
	if l.isInLayoutRect(x, y, l.descLayout) {
		return l.gameEvent, KeyBindingTypePrimary
	}

	if l.primaryLayout != nil {
		if l.isInLayoutRect(x, y, l.primaryLayout) {
			return l.gameEvent, KeyBindingTypePrimary
		}
	}

	if l.secondaryLayout != nil {
		if l.isInLayoutRect(x, y, l.secondaryLayout) {
			return l.gameEvent, KeyBindingTypeSecondary
		}
	}

	return defaultGameEvent, KeyBindingTypeNone
}
