package d2ui

var _ Widget = &Button{}

type SwitchableButton struct {
	*BaseWidget
	active        *Button
	inactive      *Button
	onActivate    func()
	onDezactivate func()
	state         bool
}

func (ui *UIManager) NewSwitchableButton(active, inactive *Button, state bool) *SwitchableButton {
	base := NewBaseWidget(ui)
	base.SetVisible(true)

	sbtn := &SwitchableButton{
		BaseWidget: base,
		active:     active,
		inactive:   inactive,
		state:      state,
	}

	sbtn.OnActivated(func() {})
	sbtn.OnDezactivated(func() {})

	return sbtn
}

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

func (sbtn *SwitchableButton) OnActivated(cb func()) {
	sbtn.active.OnActivated(func() {
		cb()
		sbtn.state = false
		sbtn.SetVisible(sbtn.GetVisible())
	})
}

func (sbtn *SwitchableButton) Activate() {
	sbtn.onActivate()
}

func (sbtn *SwitchableButton) OnDezactivated(cb func()) {
	sbtn.inactive.OnActivated(func() {
		cb()
		sbtn.state = true
		sbtn.SetVisible(sbtn.GetVisible())
	})
}

func (sbtn *SwitchableButton) Dezactivate() {
	sbtn.onDezactivate()
}

func (sbtn *SwitchableButton) SetPosition(x, y int) {
	sbtn.BaseWidget.SetPosition(x, y)
	sbtn.active.SetPosition(x, y)
	sbtn.inactive.SetPosition(x, y)
}
