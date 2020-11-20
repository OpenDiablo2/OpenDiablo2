package d2ui

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
)

const (
	scrollbarOffsetY       = 30
	halfScrollbarOffsetY   = scrollbarOffsetY / 2
	scrollbarSpriteOffsetY = 10
	scrollbarFrameOffset   = 4
	scrollbarWidth         = 10
)

// static check that Scrollbar implements widget
var _ Widget = &Scrollbar{}

// Scrollbar is a vertical slider ui element
type Scrollbar struct {
	*BaseWidget
	enabled         bool
	currentOffset   int
	maxOffset       int
	lastDirChange   int
	onActivate      func()
	scrollbarSprite *Sprite
}

// NewScrollbar creates a scrollbar instance
func (ui *UIManager) NewScrollbar(x, y, height int) *Scrollbar {
	scrollbarSprite, err := ui.NewSprite(d2resource.Scrollbar, d2resource.PaletteSky)
	if err != nil {
		ui.Error(err.Error())
		return nil
	}

	base := NewBaseWidget(ui)
	base.SetPosition(x, y)
	base.height = height

	result := &Scrollbar{
		BaseWidget:      base,
		enabled:         true,
		scrollbarSprite: scrollbarSprite,
	}

	ui.addWidget(result)

	return result
}

// GetEnabled returns whether or not the scrollbar is enabled
func (v *Scrollbar) GetEnabled() bool {
	return v.enabled
}

// SetEnabled sets the enabled state
func (v *Scrollbar) SetEnabled(enabled bool) {
	v.enabled = enabled
}

// SetPressed is not used by the scrollbar, but is present to satisfy the ui widget interface
func (v *Scrollbar) SetPressed(_ bool) {}

// GetPressed is not used by the scrollbar, but is present to satisfy the ui widget interface
func (v *Scrollbar) GetPressed() bool { return false }

// OnActivated sets the onActivate callback function for the scrollbar
func (v *Scrollbar) OnActivated(callback func()) {
	v.onActivate = callback
}

func (v *Scrollbar) getBarPosition() int {
	maxOffset := float32(v.maxOffset) * float32(v.height-scrollbarOffsetY)
	return int(float32(v.currentOffset) / maxOffset)
}

// Activate will call the onActivate callback (if set)
func (v *Scrollbar) Activate() {
	_, my := v.manager.CursorPosition()
	barPosition := v.getBarPosition()

	if my <= v.y+barPosition+halfScrollbarOffsetY {
		if v.currentOffset > 0 {
			v.currentOffset--
			v.lastDirChange = -1
		}
	} else {
		if v.currentOffset < v.maxOffset {
			v.currentOffset++
			v.lastDirChange = 1
		}
	}

	if v.onActivate != nil {
		v.onActivate()
	}
}

// GetLastDirChange get the last direction change
func (v *Scrollbar) GetLastDirChange() int {
	return v.lastDirChange
}

// Render renders the scrollbar to the given surface
func (v *Scrollbar) Render(target d2interface.Surface) {
	if !v.visible || v.maxOffset == 0 {
		return
	}

	offset := 0

	if !v.enabled {
		offset = 2
	}

	v.scrollbarSprite.SetPosition(v.x, v.y)

	v.scrollbarSprite.RenderSegmented(target, 1, 1, 0+offset)

	v.scrollbarSprite.SetPosition(v.x, v.y+v.height-scrollbarSpriteOffsetY) // what is the magic?

	v.scrollbarSprite.RenderSegmented(target, 1, 1, 1+offset)

	if v.maxOffset == 0 || v.currentOffset < 0 || v.currentOffset > v.maxOffset {
		return
	}

	v.scrollbarSprite.SetPosition(v.x, v.y+10+v.getBarPosition())

	offset = scrollbarFrameOffset

	if !v.enabled {
		offset++
	}

	v.scrollbarSprite.RenderSegmented(target, 1, 1, offset)
}

// Advance advances the scrollbar sprite
func (v *Scrollbar) Advance(elapsed float64) error {
	return v.scrollbarSprite.Advance(elapsed)
}

// GetSize returns the scrollbar width and height
func (v *Scrollbar) GetSize() (width, height int) {
	return scrollbarWidth, v.height
}

// SetMaxOffset sets the maximum offset of the scrollbar
func (v *Scrollbar) SetMaxOffset(maxOffset int) {
	v.maxOffset = maxOffset
	if v.maxOffset < 0 {
		v.maxOffset = 0
	}

	if v.currentOffset > v.maxOffset {
		v.currentOffset = v.maxOffset
	}

	if v.maxOffset == 0 {
		v.currentOffset = 0
	}
}

// SetCurrentOffset sets the scrollbar's current offset
func (v *Scrollbar) SetCurrentOffset(currentOffset int) {
	v.currentOffset = currentOffset
}

// GetMaxOffset returns the max offset
func (v *Scrollbar) GetMaxOffset() int {
	return v.maxOffset
}

// GetCurrentOffset gets the current max offset of the scrollbar
func (v *Scrollbar) GetCurrentOffset() int {
	return v.currentOffset
}
