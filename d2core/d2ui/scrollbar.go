package d2ui

import (
	"log"

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

// Scrollbar is a vertical slider ui element
type Scrollbar struct {
	manager         *UIManager
	x, y, height    int
	visible         bool
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
		log.Print(err)
		return nil
	}

	result := &Scrollbar{
		visible:         true,
		enabled:         true,
		x:               x,
		y:               y,
		height:          height,
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
func (v *Scrollbar) Render(target d2interface.Surface) error {
	if !v.visible || v.maxOffset == 0 {
		return nil
	}

	offset := 0

	if !v.enabled {
		offset = 2
	}

	v.scrollbarSprite.SetPosition(v.x, v.y)

	if err := v.scrollbarSprite.RenderSegmented(target, 1, 1, 0+offset); err != nil {
		return err
	}

	v.scrollbarSprite.SetPosition(v.x, v.y+v.height-scrollbarSpriteOffsetY) // what is the magic?

	if err := v.scrollbarSprite.RenderSegmented(target, 1, 1, 1+offset); err != nil {
		return err
	}

	if v.maxOffset == 0 || v.currentOffset < 0 || v.currentOffset > v.maxOffset {
		return nil
	}

	v.scrollbarSprite.SetPosition(v.x, v.y+10+v.getBarPosition())

	offset = scrollbarFrameOffset

	if !v.enabled {
		offset++
	}

	if err := v.scrollbarSprite.RenderSegmented(target, 1, 1, offset); err != nil {
		return err
	}

	return nil
}

// bindManager binds the scrollbar to the UI manager
func (v *Scrollbar) bindManager(manager *UIManager) {
	v.manager = manager
}

// Advance advances the scrollbar sprite
func (v *Scrollbar) Advance(elapsed float64) error {
	return v.scrollbarSprite.Advance(elapsed)
}

// GetSize returns the scrollbar width and height
func (v *Scrollbar) GetSize() (width, height int) {
	return scrollbarWidth, v.height
}

// SetPosition sets the scrollbar x,y position
func (v *Scrollbar) SetPosition(x, y int) {
	v.x = x
	v.y = y
}

// GetPosition returns the scrollbar x,y position
func (v *Scrollbar) GetPosition() (x, y int) {
	return v.x, v.y
}

// GetVisible returns whether or not the scrollbar is visible
func (v *Scrollbar) GetVisible() bool {
	return v.visible
}

// SetVisible sets the scrollbar visibility state
func (v *Scrollbar) SetVisible(visible bool) {
	v.visible = visible
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
