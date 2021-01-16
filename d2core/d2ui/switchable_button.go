package d2ui

import "github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"

// static check if SwitchableButton implemented widget
var _ Widget = &SwitchableButton{}

// SwitchableButton represents switchable button widget
type SwitchableButton struct {
	*BaseWidget
	active        *Button
	inactive      *Button
	onActivate    func()
	onDezactivate func()
	state         bool
}

// NewSwitchableButton creates new switchable button
func (ui *UIManager) NewSwitchableButton(active, inactive *Button, state bool) *SwitchableButton {
	base := NewBaseWidget(ui)

	sbtn := &SwitchableButton{
		BaseWidget: base,
		active:     active,
		inactive:   inactive,
		state:      state,
	}
	sbtn.bindManager(ui)
	sbtn.SetVisible(false)

	sbtn.OnActivated(func() {})
	sbtn.OnDezactivated(func() {})

	ui.addWidget(sbtn)

	return sbtn
}

// SetVisible sets widget's visibility
func (sbtn *SwitchableButton) SetVisible(visible bool) {
	if !visible {
		sbtn.active.SetVisible(false)
		sbtn.inactive.SetVisible(false)

		return
	}

	if sbtn.state {
		sbtn.active.SetVisible(true)
		sbtn.inactive.SetVisible(false)
	} else {
		sbtn.active.SetVisible(false)
		sbtn.inactive.SetVisible(true)
	}
}

// OnActivated sets onActivate callback
func (sbtn *SwitchableButton) OnActivated(cb func()) {
	sbtn.active.OnActivated(func() {
		cb()
		sbtn.state = false
		sbtn.SetVisible(sbtn.GetVisible())
	})
}

// Activate switches widget into active state
func (sbtn *SwitchableButton) Activate() {
	sbtn.onActivate()
}

// OnDezactivated sets onDezactivate callback
func (sbtn *SwitchableButton) OnDezactivated(cb func()) {
	sbtn.inactive.OnActivated(func() {
		cb()
		sbtn.state = true
		sbtn.SetVisible(sbtn.GetVisible())
	})
}

// Dezactivate switch widget to inactive state
func (sbtn *SwitchableButton) Dezactivate() {
	sbtn.onDezactivate()
}

// SetPosition sets widget's position
func (sbtn *SwitchableButton) SetPosition(x, y int) {
	sbtn.BaseWidget.SetPosition(x, y)
	sbtn.active.SetPosition(x, y)
	sbtn.inactive.SetPosition(x, y)
}

// Advance advances widget
func (sbtn *SwitchableButton) Advance(_ float64) error {
	// noop
	return nil
}

// Render renders widget
func (sbtn *SwitchableButton) Render(target d2interface.Surface) {
	if sbtn.active.GetVisible() {
		sbtn.active.Render(target)
	}

	if sbtn.inactive.GetVisible() {
		sbtn.inactive.Render(target)
	}
}
