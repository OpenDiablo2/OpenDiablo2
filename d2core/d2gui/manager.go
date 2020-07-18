package d2gui

import (
	"image/color"
	"math"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
)

type manager struct {
	layout        *Layout
	cursorAnim    d2interface.Animation
	cursorX       int
	cursorY       int
	loadingAnim   d2interface.Animation
	cursorVisible bool
	loading       bool
}

func createGuiManager(inputManager d2interface.InputManager) (*manager, error) {
	cursorAnim, err := d2asset.LoadAnimation(d2resource.CursorDefault, d2resource.PaletteUnits)
	if err != nil {
		return nil, err
	}

	loadingAnim, err := d2asset.LoadAnimation(d2resource.LoadingScreen, d2resource.PaletteLoading)
	if err != nil {
		return nil, err
	}

	manager := &manager{
		cursorAnim:    cursorAnim,
		loadingAnim:   loadingAnim,
		cursorVisible: true,
	}

	manager.clear()

	if err := inputManager.BindHandler(manager); err != nil {
		return nil, err
	}

	return manager, nil
}

func (m *manager) SetLayout(layout *Layout) {
	m.layout = layout
	if m.layout != nil {
		m.layout.AdjustEntryPlacement()
	}
}

func (m *manager) OnMouseButtonDown(event d2interface.MouseEvent) bool {
	if m.layout == nil {
		return false
	}

	return m.layout.onMouseButtonDown(event)
}

func (m *manager) OnMouseButtonUp(event d2interface.MouseEvent) bool {
	if m.layout == nil {
		return false
	}

	return m.layout.onMouseButtonUp(event)
}

func (m *manager) OnMouseMove(event d2interface.MouseMoveEvent) bool {
	m.cursorX = event.X()
	m.cursorY = event.Y()

	if m.layout == nil {
		return false
	}

	return m.layout.onMouseMove(event)
}

func (m *manager) render(target d2interface.Surface) error {
	if m.loading {
		if err := m.renderLoadScreen(target); err != nil {
			return err
		}
	} else if m.layout != nil {
		m.layout.SetSize(target.GetSize())
		if err := m.layout.render(target); err != nil {
			return err
		}
	}

	if m.cursorVisible {
		if err := m.renderCursor(target); err != nil {
			return err
		}
	}

	return nil
}

func (m *manager) renderLoadScreen(target d2interface.Surface) error {
	if clearErr := target.Clear(color.Black); clearErr != nil {
		return clearErr
	}

	pushCount := 0

	screenWidth, screenHeight := target.GetSize()
	animWidth, animHeight := m.loadingAnim.GetCurrentFrameSize()

	target.PushTranslation(screenWidth/2-animWidth/2, screenHeight/2+animHeight/2)
	pushCount++

	target.PushTranslation(0, -animHeight)
	pushCount++

	defer target.PopN(pushCount)

	return m.loadingAnim.Render(target)
}

func (m *manager) renderCursor(target d2interface.Surface) error {
	_, height := m.cursorAnim.GetCurrentFrameSize()
	pushCount := 0

	target.PushTranslation(m.cursorX, m.cursorY)
	pushCount++

	target.PushTranslation(0, -height)
	pushCount++

	defer target.PopN(pushCount)

	return m.cursorAnim.Render(target)
}

func (m *manager) advance(elapsed float64) error {
	if !m.loading && m.layout != nil {
		if err := m.layout.advance(elapsed); err != nil {
			return err
		}
	}

	return nil
}

func (m *manager) showLoadScreen(progress float64) {
	progress = math.Min(progress, 1.0)
	progress = math.Max(progress, 0.0)

	animation := m.loadingAnim
	frameCount := animation.GetFrameCount()

	_ = animation.SetCurrentFrame(int(float64(frameCount-1) * progress))

	m.loading = true
}

func (m *manager) hideLoadScreen() {
	m.loading = false
}

func (m *manager) showCursor() {
	m.cursorVisible = true
}

func (m *manager) hideCursor() {
	m.cursorVisible = false
}

func (m *manager) clear() {
	m.SetLayout(nil)
	m.hideLoadScreen()
}
