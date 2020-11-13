package d2ui

import (
	"log"
	"sort"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
)

// UIManager manages a collection of UI elements (buttons, textboxes, labels)
type UIManager struct {
	asset            *d2asset.AssetManager
	renderer         d2interface.Renderer
	inputManager     d2interface.InputManager
	audio            d2interface.AudioProvider
	widgets          []Widget
	widgetsGroups    []*WidgetGroup
	clickableWidgets []ClickableWidget
	cursorButtons    CursorButton
	CursorX          int
	CursorY          int
	pressedWidget    ClickableWidget
	clickSfx         d2interface.SoundEffect
}

// Note: methods for creating buttons and stuff are in their respective files

// Initialize is meant to be called after the game loads all of the necessary files
// for sprites and audio
func (ui *UIManager) Initialize() {
	sfx, err := ui.audio.LoadSound(d2resource.SFXButtonClick, false, false)
	if err != nil {
		log.Fatalf("failed to initialize ui: %v", err)
	}

	ui.clickSfx = sfx

	if err := ui.inputManager.BindHandler(ui); err != nil {
		log.Fatalf("failed to initialize ui: %v", err)
	}
}

// Reset resets the state of the UI manager. Typically called for new screens
func (ui *UIManager) Reset() {
	ui.widgets = nil
	ui.clickableWidgets = nil
	ui.pressedWidget = nil
}

// addWidgetGroup adds a widgetGroup to the UI manager and sorts by priority
func (ui *UIManager) addWidgetGroup(group *WidgetGroup) {
	ui.widgetsGroups = append(ui.widgetsGroups, group)

	sort.SliceStable(ui.widgetsGroups, func(i, j int) bool {
		return ui.widgetsGroups[i].renderPriority < ui.widgetsGroups[j].renderPriority
	})
}

// addWidget adds a widget to the UI manager
func (ui *UIManager) addWidget(widget Widget) {
	err := ui.inputManager.BindHandler(widget)
	if err != nil {
		log.Print(err)
	}

	clickable, ok := widget.(ClickableWidget)
	if ok {
		ui.clickableWidgets = append(ui.clickableWidgets, clickable)
	}

	ui.widgets = append(ui.widgets, widget)
	widget.bindManager(ui)
}

// OnMouseButtonUp is an event handler for input
func (ui *UIManager) OnMouseButtonUp(event d2interface.MouseEvent) bool {
	ui.CursorX, ui.CursorY = event.X(), event.Y()
	if event.Button() == d2enum.MouseButtonLeft {
		ui.cursorButtons |= CursorButtonLeft
		// activate previously pressed widget if cursor is still hovering
		w := ui.pressedWidget

		if w != nil && w.Contains(ui.CursorX, ui.CursorY) && w.GetVisible() && w.GetEnabled() {
			w.Activate()
		}

		// unpress all widgets that are pressed
		for _, w := range ui.clickableWidgets {
			w.SetPressed(false)
		}
	}

	return false
}

// OnMouseMove is the mouse move event handler
func (ui *UIManager) OnMouseMove(event d2interface.MouseMoveEvent) bool {
	for _, w := range ui.widgetsGroups {
		if w.GetVisible() {
			w.OnMouseMove(event.X(), event.Y())
		}
	}

	return false
}

// OnMouseButtonDown is the mouse button down event handler
func (ui *UIManager) OnMouseButtonDown(event d2interface.MouseEvent) bool {
	ui.CursorX, ui.CursorY = event.X(), event.Y()
	if event.Button() == d2enum.MouseButtonLeft {
		// find and press a widget on screen
		ui.pressedWidget = nil
		for _, w := range ui.clickableWidgets {
			if w.Contains(ui.CursorX, ui.CursorY) && w.GetVisible() && w.GetEnabled() {
				w.SetPressed(true)
				ui.pressedWidget = w
				ui.clickSfx.Play()

				break
			}
		}
	}

	if event.Button() == d2enum.MouseButtonRight {
		ui.cursorButtons |= CursorButtonRight
	}

	return false
}

// Render renders all of the UI elements
func (ui *UIManager) Render(target d2interface.Surface) {
	for _, widget := range ui.widgets {
		if widget.GetVisible() {
			widget.Render(target)
		}
	}

	for _, widgetGroup := range ui.widgetsGroups {
		if widgetGroup.GetVisible() {
			widgetGroup.Render(target)
		}
	}
}

// Advance updates all of the UI elements
func (ui *UIManager) Advance(elapsed float64) {
	for _, widget := range ui.widgets {
		if widget.GetVisible() {
			err := widget.Advance(elapsed)
			if err != nil {
				log.Print(err)
			}
		}
	}
}

// CursorButtonPressed determines if the specified button has been pressed
func (ui *UIManager) CursorButtonPressed(button CursorButton) bool {
	return ui.cursorButtons&button > 0
}

// CursorPosition returns the current cursor position
func (ui *UIManager) CursorPosition() (x, y int) {
	return ui.CursorX, ui.CursorY
}

// Renderer returns the renderer for this ui manager
func (ui *UIManager) Renderer() d2interface.Renderer {
	return ui.renderer
}
