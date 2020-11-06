package d2player

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2gui"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
)

// KeyBindingMenu represents the menu to view/edit the
// key bindings
type KeyBindingMenu struct {
	*Box
}

func NewKeyBindingMenu(
	asset *d2asset.AssetManager,
	renderer d2interface.Renderer,
	ui *d2ui.UIManager,
	guiManager *d2gui.GuiManager,
) *KeyBindingMenu {
	return &KeyBindingMenu{
		Box: &Box{
			asset:      asset,
			renderer:   renderer,
			uiManager:  ui,
			guiManager: guiManager,
		},
	}
}
