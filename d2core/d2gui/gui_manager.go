package d2gui

import (
	"image/color"
	"math"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
)

const (
	logPrefix = "GUI Manager"
)

// GuiManager is a GUI widget manager that handles dynamic layout/positioning of widgets
type GuiManager struct {
	asset         *d2asset.AssetManager
	layout        *Layout
	cursorAnim    d2interface.Animation
	cursorX       int
	cursorY       int
	loadingAnim   d2interface.Animation
	cursorVisible bool
	loading       bool

	*d2util.Logger
}

// CreateGuiManager creates an instance of the GuiManager
func CreateGuiManager(asset *d2asset.AssetManager, l d2util.LogLevel, inputManager d2interface.InputManager) (*GuiManager, error) {
	cursorAnim, err := asset.LoadAnimation(d2resource.CursorDefault, d2resource.PaletteUnits)
	if err != nil {
		return nil, err
	}

	loadingAnim, err := asset.LoadAnimation(d2resource.LoadingScreen, d2resource.PaletteLoading)
	if err != nil {
		return nil, err
	}

	manager := &GuiManager{
		asset:         asset,
		cursorAnim:    cursorAnim,
		loadingAnim:   loadingAnim,
		cursorVisible: true,
	}

	manager.Logger = d2util.NewLogger()
	manager.Logger.SetPrefix(logPrefix)
	manager.Logger.SetLevel(l)

	manager.clear()

	if err := inputManager.BindHandler(manager); err != nil {
		return nil, err
	}

	return manager, nil
}

// SetLayout sets the layout of the GuiManager
func (m *GuiManager) SetLayout(layout *Layout) {
	m.layout = layout
	if m.layout != nil {
		m.layout.AdjustEntryPlacement()
	}
}

// OnMouseButtonDown handles mouse button click events
func (m *GuiManager) OnMouseButtonDown(event d2interface.MouseEvent) bool {
	if m.layout == nil {
		return false
	}

	return m.layout.onMouseButtonDown(event)
}

// OnMouseButtonUp handles the mouse button release events
func (m *GuiManager) OnMouseButtonUp(event d2interface.MouseEvent) bool {
	if m.layout == nil {
		return false
	}

	return m.layout.onMouseButtonUp(event)
}

// OnMouseMove handles mouse movement events
func (m *GuiManager) OnMouseMove(event d2interface.MouseMoveEvent) bool {
	m.cursorX = event.X()
	m.cursorY = event.Y()

	if m.layout == nil {
		return false
	}

	return m.layout.onMouseMove(event)
}

// Render renders the GuiManager to the given surface
func (m *GuiManager) Render(target d2interface.Surface) error {
	if m.loading {
		m.renderLoadScreen(target)
	} else if m.layout != nil {
		m.layout.SetSize(target.GetSize())
		m.layout.render(target)
	}

	if m.cursorVisible {
		m.renderCursor(target)
	}

	return nil
}

func (m *GuiManager) renderLoadScreen(target d2interface.Surface) {
	target.Clear(color.Black)

	pushCount := 0

	screenWidth, screenHeight := target.GetSize()
	animWidth, animHeight := m.loadingAnim.GetCurrentFrameSize()

	target.PushTranslation(screenWidth/2-animWidth/2, screenHeight/2+animHeight/2)
	pushCount++

	target.PushTranslation(0, -animHeight)
	pushCount++

	defer target.PopN(pushCount)

	m.loadingAnim.Render(target)
}

func (m *GuiManager) renderCursor(target d2interface.Surface) {
	_, height := m.cursorAnim.GetCurrentFrameSize()
	pushCount := 0

	target.PushTranslation(m.cursorX, m.cursorY)
	pushCount++

	target.PushTranslation(0, -height)
	pushCount++

	defer target.PopN(pushCount)

	m.cursorAnim.Render(target)
}

// Advance advances the GuiManager state
func (m *GuiManager) Advance(elapsed float64) error {
	if !m.loading && m.layout != nil {
		if err := m.layout.advance(elapsed); err != nil {
			return err
		}
	}

	return nil
}

// ShowLoadScreen shows the loading screen with the given progress
func (m *GuiManager) ShowLoadScreen(progress float64) {
	progress = math.Min(progress, 1.0)
	progress = math.Max(progress, 0.0)

	animation := m.loadingAnim
	frameCount := animation.GetFrameCount()

	err := animation.SetCurrentFrame(int(float64(frameCount-1) * progress))
	if err != nil {
		m.Error(err.Error())
	}

	m.loading = true
}

// HideLoadScreen hides the load screen
func (m *GuiManager) HideLoadScreen() {
	m.loading = false
}

// ShowCursor makes the cursor visible
func (m *GuiManager) ShowCursor() {
	m.cursorVisible = true
}

// HideCursor hides the cursor
func (m *GuiManager) HideCursor() {
	m.cursorVisible = false
}

func (m *GuiManager) clear() {
	m.SetLayout(nil)
	m.HideLoadScreen()
}
