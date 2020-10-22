package d2player

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapentity"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
)

// SkillSelectMenu is a wrapper for the left + right menu that pop up when a player clicks the left/right skill select.
type SkillSelectMenu struct {
	LeftPanel  *SkillPanel
	RightPanel *SkillPanel
}

// NewSkillSelectMenu creates a skill select menu.
func NewSkillSelectMenu(asset *d2asset.AssetManager, ui *d2ui.UIManager, hero *d2mapentity.Player) *SkillSelectMenu {
	skillSelectMenu := &SkillSelectMenu{
		LeftPanel: NewHeroSkillsPanel(asset, ui, hero, true),
		RightPanel: NewHeroSkillsPanel(asset, ui, hero, false),
	}

	return skillSelectMenu
}

// HandleClick will propagate the click to the panels.
func (sm *SkillSelectMenu) HandleClick(X int, Y int) bool {
	if sm.LeftPanel.HandleClick(X, Y) {
		return true
	}

	if sm.RightPanel.HandleClick(X, Y) {
		return true
	}

	return true
}

// HandleMouseMove will propagate the mouse move event to the panels.
func (sm *SkillSelectMenu) HandleMouseMove(X int, Y int) {
	if sm.LeftPanel.IsOpen() {
		sm.LeftPanel.HandleMouseMove(X, Y)

	} else if sm.RightPanel.IsOpen() {
		sm.RightPanel.HandleMouseMove(X, Y)
	}
}

// RegenerateImageCache will force both panels to re-create the image shown at skill popup menus.
// Somewhat expensive operation, should not be called often.
func (sm *SkillSelectMenu) RegenerateImageCache() {
	sm.LeftPanel.RegenerateImageCache()
	sm.RightPanel.RegenerateImageCache()
}

// Render gets called on every frame
func (sm *SkillSelectMenu) Render(target d2interface.Surface) {
	sm.LeftPanel.Render(target)
	sm.RightPanel.Render(target)
}

// IsOpen returns whether one of the panels(left or right) is open
func (sm *SkillSelectMenu) IsOpen() bool {
	return sm.LeftPanel.IsOpen() || sm.RightPanel.IsOpen()
}

// IsInRect returns whether the coordinates are in one of the panels(left or right)
func (sm *SkillSelectMenu) IsInRect(X int, Y int) bool {
	return sm.LeftPanel.IsInRect(X, Y) || sm.RightPanel.IsInRect(X, Y)
}

// ClosePanels will close both panels
func (sm *SkillSelectMenu) ClosePanels() {
	sm.RightPanel.Close()
	sm.LeftPanel.Close()
}

// OpenLeftPanel will close the right panel and open the left panel.
func (sm *SkillSelectMenu) OpenLeftPanel() {
	sm.RightPanel.Close()
	sm.LeftPanel.Open()
}

// ToggleLeftPanel will close or open the left panel, depending on the current state
func (sm *SkillSelectMenu) ToggleLeftPanel() {
	if sm.LeftPanel.IsOpen() {
		sm.LeftPanel.Close()
	} else {
		sm.OpenLeftPanel()
	}
}

// OpenRightPanel will close the left panel and open the right panel.
func (sm *SkillSelectMenu) OpenRightPanel() {
	sm.LeftPanel.Close()
	sm.RightPanel.Open()
}

// ToggleRightPanel will close or open the right panel, depending on the current state
func (sm *SkillSelectMenu) ToggleRightPanel() {
	if sm.RightPanel.IsOpen() {
		sm.RightPanel.Close()
	} else {
		sm.OpenRightPanel()
	}
}
