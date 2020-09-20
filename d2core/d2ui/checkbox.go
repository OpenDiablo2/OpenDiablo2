package d2ui

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
)

// Checkbox represents a checkbox UI element
type Checkbox struct {
	manager      *UIManager
	Image        d2interface.Surface
	checkedImage d2interface.Surface
	x            int
	y            int
	width        int
	height       int
	onClick      func()
	checkState   bool
	visible      bool
	enabled      bool
}

// NewCheckbox creates a new instance of a checkbox
func (ui *UIManager) NewCheckbox(checkState bool) *Checkbox {
	var err error
	result := &Checkbox{
		checkState: checkState,
		visible:    true,
		width:      0,
		height:     0,
		enabled:    true,
	}

	checkboxSprite, err := ui.NewSprite(d2resource.Checkbox, d2resource.PaletteFechar)
	if err != nil {
		log.Print(err)
		return nil
	}

	result.width, result.height, err = checkboxSprite.GetFrameSize(0)
	if err != nil {
		log.Print(err)
		return nil
	}
	checkboxSprite.SetPosition(0, 0)

	result.Image, err = ui.renderer.NewSurface(result.width, result.height, d2enum.FilterNearest)
	if err != nil {
		log.Print(err)
		return nil
	}

	err = checkboxSprite.RenderSegmented(result.Image, 1, 1, 0)
	if err != nil {
		log.Print(err)
		return nil
	}

	result.checkedImage, err = ui.renderer.NewSurface(result.width, result.height, d2enum.FilterNearest)
	if err != nil {
		log.Print(err)
		return nil
	}

	err = checkboxSprite.RenderSegmented(result.checkedImage, 1, 1, 1)
	if err != nil {
		log.Print(err)
		return nil
	}

	ui.addWidget(result)

	return result
}

// bindManager binds the checkbox to the UI manager
func (v *Checkbox) bindManager(manager *UIManager) {
	v.manager = manager
}

// Render renders the checkbox
func (v *Checkbox) Render(target d2interface.Surface) error {
	var err error
	target.PushTranslation(v.x, v.y)
	defer target.Pop()

	target.PushFilter(d2enum.FilterNearest)
	defer target.Pop()

	if v.checkState {
		err = target.Render(v.checkedImage)
		if err != nil {
			return err
		}
	} else {
		err = target.Render(v.Image)
		if err != nil {
			return err
		}
	}

	return nil
}

// Advance does nothing for checkboxes
func (v *Checkbox) Advance(_ float64) error {
	return nil
}

// GetEnabled returns the enabled state of the checkbox
func (v *Checkbox) GetEnabled() bool {
	return v.enabled
}

// SetEnabled sets the enabled state of the checkbox
func (v *Checkbox) SetEnabled(enabled bool) {
	v.enabled = enabled
}

// SetPressed does nothing for checkboxes
func (v *Checkbox) SetPressed(_ bool) {
}

// SetCheckState sets the check state of the checkbox
func (v *Checkbox) SetCheckState(checkState bool) {
	v.checkState = checkState
}

// GetCheckState returns the check state of the checkbox
func (v *Checkbox) GetCheckState() bool {
	return v.checkState
}

// GetPressed returns the pressed state of the checkbox
func (v *Checkbox) GetPressed() bool {
	return v.checkState
}

// OnActivated sets the callback function of the click event for the checkbox
func (v *Checkbox) OnActivated(callback func()) {
	v.onClick = callback
}

// Activate activates the checkbox
func (v *Checkbox) Activate() {
	v.checkState = !v.checkState
	if v.onClick == nil {
		return
	}

	v.onClick()
}

// GetPosition returns the position of the checkbox
func (v *Checkbox) GetPosition() (x, y int) {
	return v.x, v.y
}

// GetSize returns the size of the checkbox
func (v *Checkbox) GetSize() (width, height int) {
	return v.width, v.height
}

// GetVisible returns the visibility state of the checkbox
func (v *Checkbox) GetVisible() bool {
	return v.visible
}

// SetPosition sets the position of the checkbox
func (v *Checkbox) SetPosition(x, y int) {
	v.x = x
	v.y = y
}

// SetVisible sets the visibility of the checkbox
func (v *Checkbox) SetVisible(visible bool) {
	v.visible = visible
}
