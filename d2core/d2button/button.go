package d2button

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

type buttonCallback = func() (preventPropagation bool)

// Button defines a standard wide UI button
type Button struct {
	Layout   ButtonLayout
	Sprite   d2interface.Sprite
	Surfaces struct {
		Normal         d2interface.Surface
		Pressed        d2interface.Surface
		Toggled        d2interface.Surface
		PressedToggled d2interface.Surface
		Disabled       d2interface.Surface
	}
	callback      buttonCallback
	width, height int
	enabled       bool
	pressed       bool
	toggled       bool
}

// New creates an instance of Button
func New() *Button {
	btn := &Button{
		enabled: true,
	}

	//buttonLayout := GetLayout(t)
	//btn.Layout = buttonLayout
	//
	//btn.normalSurface = ui.renderer.NewSurface(btn.width, btn.height)
	//
	//buttonSprite.SetPosition(0, 0)
	//buttonSprite.SetEffect(d2enum.DrawEffectModulate)
	//
	//btn.createTooltip()
	//
	//ui.addWidget(btn) // important that this comes before prerenderStates!
	//
	//btn.prerenderStates(buttonSprite, &buttonLayout, lbl)

	return btn
}

type buttonStateDescriptor struct {
	baseFrame            int
	offsetX, offsetY     int
	prerenderdestination *d2interface.Surface
	fmtErr               string
}

// this is some jank shit, and if things go wrong you should suspect this func first
func (v *Button) GetButtonSize() (w, h int) {
	if v.Sprite == nil {
		return 0, 0
	}

	base := v.Layout.BaseFrame
	sx, sy := v.Layout.XSegments, v.Layout.YSegments

	if sx < 1 {
		sx = 1
	}

	if sy < 1 {
		sy = 1
	}

	for idx := base; idx < (base + sx*sy); idx++ {
		row, column := idx/sx, idx%sx

		err := v.Sprite.SetCurrentFrame(idx)
		if err != nil {
			continue
		}

		fw, fh := v.Sprite.GetCurrentFrameSize()

		if row == 0 {
			w += fw
		}

		if column == 0 {
			h += fh
		}
	}

	return w, h
}

// OnActivated defines the callback handler for the activate event
func (v *Button) OnActivated(callback buttonCallback) {
	v.callback = callback
}

// Activate calls the on activated callback handler, if any
func (v *Button) Activate() {
	if v.callback == nil {
		return
	}

	v.callback()
}

// Toggle negates the toggled state of the button
func (v *Button) Toggle() {
	v.toggled = !v.toggled
}

// GetToggled returns the toggled state of the button
func (v *Button) GetToggled() bool {
	return v.toggled
}

// Advance advances the button state
func (v *Button) GetCurrentTexture() d2interface.Surface {
	if !v.enabled {
		return v.Surfaces.Disabled
	}

	if v.pressed && v.toggled {
		return v.Surfaces.PressedToggled
	}

	if v.pressed {
		return v.Surfaces.Pressed
	}

	return v.Surfaces.Normal
}

// GetEnabled returns the enabled state
func (v *Button) GetEnabled() bool {
	return v.enabled
}

// SetEnabled sets the enabled state
func (v *Button) SetEnabled(enabled bool) {
	v.enabled = enabled
}

// GetPressed returns the pressed state of the button
func (v *Button) GetPressed() bool {
	return v.pressed
}

// SetPressed sets the pressed state of the button
func (v *Button) SetPressed(pressed bool) {
	v.pressed = pressed
}
