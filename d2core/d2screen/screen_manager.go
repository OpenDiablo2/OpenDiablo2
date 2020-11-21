package d2screen

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2gui"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
)

const (
	logPrefix = "Screen Manager"
)

// ScreenManager manages game screens (main menu, credits, character select, game, etc)
type ScreenManager struct {
	uiManager     *d2ui.UIManager
	nextScreen    Screen
	loadingScreen Screen
	loadingState  LoadingState
	currentScreen Screen
	guiManager    *d2gui.GuiManager

	*d2util.Logger
}

// NewScreenManager creates a screen manager
func NewScreenManager(ui *d2ui.UIManager, l d2util.LogLevel, guiManager *d2gui.GuiManager) *ScreenManager {
	sm := &ScreenManager{
		uiManager:  ui,
		guiManager: guiManager,
	}

	sm.Logger = d2util.NewLogger()
	sm.Logger.SetPrefix(logPrefix)
	sm.Logger.SetLevel(l)

	return sm
}

// SetNextScreen is about to set a given screen as next
func (sm *ScreenManager) SetNextScreen(screen Screen) {
	sm.nextScreen = screen
}

// Advance updates the UI on every frame
func (sm *ScreenManager) Advance(elapsed float64) error {
	switch {
	case sm.loadingScreen != nil:
		// this call blocks execution and could lead to deadlock if a screen implements OnLoad incorreclty
		load, ok := <-sm.loadingState.updates
		if !ok {
			sm.Warning("loadingState chan should not be closed while in a loading screen")
		}

		if load.err != nil {
			sm.Errorf("PROBLEM LOADING THE SCREEN: %v", load.err)
			return load.err
		}

		sm.guiManager.ShowLoadScreen(load.progress)

		if load.done {
			sm.currentScreen = sm.loadingScreen
			sm.loadingScreen = nil

			sm.guiManager.ShowCursor()
			sm.guiManager.HideLoadScreen()
		}
	case sm.nextScreen != nil:
		if handler, ok := sm.currentScreen.(ScreenUnloadHandler); ok {
			if err := handler.OnUnload(); err != nil {
				return err
			}
		}

		sm.uiManager.Reset()
		sm.guiManager.SetLayout(nil)

		if handler, ok := sm.nextScreen.(ScreenLoadHandler); ok {
			sm.guiManager.ShowLoadScreen(0)
			sm.guiManager.HideCursor()

			sm.loadingState = LoadingState{updates: make(chan loadingUpdate)}

			go func() {
				handler.OnLoad(sm.loadingState)
				sm.loadingState.Done()
			}()

			sm.currentScreen = nil
			sm.loadingScreen = sm.nextScreen
		} else {
			sm.currentScreen = sm.nextScreen
			sm.loadingScreen = nil
		}

		sm.nextScreen = nil
	case sm.currentScreen != nil:
		if handler, ok := sm.currentScreen.(ScreenAdvanceHandler); ok {
			if err := handler.Advance(elapsed); err != nil {
				return err
			}
		}
	}

	return nil
}

// Render renders the UI by a given surface
func (sm *ScreenManager) Render(surface d2interface.Surface) {
	if handler, ok := sm.currentScreen.(ScreenRenderHandler); ok {
		handler.Render(surface)
	}
}
