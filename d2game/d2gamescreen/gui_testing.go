package d2gamescreen

import (
	"fmt"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2gui"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2screen"
)

// GuiTestMain is a playground screen for the gui components
type GuiTestMain struct {
	renderer     d2interface.Renderer
	guiManager   *d2gui.GuiManager
	assetManager *d2asset.AssetManager
}

// CreateGuiTestMain creates a GuiTestMain screen
func CreateGuiTestMain(renderer d2interface.Renderer, guiManager *d2gui.GuiManager, assetManager *d2asset.AssetManager) *GuiTestMain {
	return &GuiTestMain{
		renderer:     renderer,
		guiManager:   guiManager,
		assetManager: assetManager,
	}
}

// OnLoad loads the resources and creates the gui components
func (g *GuiTestMain) OnLoad(loading d2screen.LoadingState) {
	layout := d2gui.CreateLayout(g.renderer, d2gui.PositionTypeHorizontal, g.assetManager)

	loading.Progress(thirtyPercent)
	//
	layoutLeft := layout.AddLayout(d2gui.PositionTypeVertical)
	layoutLeft.SetHorizontalAlign(d2gui.HorizontalAlignCenter)

	if _, err := layoutLeft.AddLabel("FontStyle16Units", d2gui.FontStyle16Units); err != nil {
		fmt.Printf("could not add label: %s to the GuiTestMain screen\n", "FontStyle16Units")
	}

	layoutLeft.AddSpacerStatic(0, 100)

	if _, err := layoutLeft.AddLabel("FontStyle30Units", d2gui.FontStyle30Units); err != nil {
		fmt.Printf("could not add label: %s to the GuiTestMain screen\n", "FontStyle30Units")
	}

	if _, err := layoutLeft.AddLabel("FontStyle42Units", d2gui.FontStyle42Units); err != nil {
		fmt.Printf("could not add label: %s to the GuiTestMain screen\n", "FontStyle42Units")
	}

	if _, err := layoutLeft.AddLabel("FontStyleFormal10Static", d2gui.FontStyleFormal10Static); err != nil {
		fmt.Printf("could not add label: %s to the GuiTestMain screen\n", "FontStyleFormal10Static")
	}

	if _, err := layoutLeft.AddLabel("FontStyleFormal11Units", d2gui.FontStyleFormal11Units); err != nil {
		fmt.Printf("could not add label: %s to the GuiTestMain screen\n", "FontStyleFormal11Units")
	}

	if _, err := layoutLeft.AddLabel("FontStyleFormal12Static", d2gui.FontStyleFormal12Static); err != nil {
		fmt.Printf("could not add label: %s to the GuiTestMain screen\n", "FontStyleFormal12Static")
	}

	loading.Progress(sixtyPercent)

	layout.AddSpacerDynamic()

	layoutRight := layout.AddLayout(d2gui.PositionTypeVertical)
	layoutRight.SetHorizontalAlign(d2gui.HorizontalAlignRight)

	if _, err := layoutRight.AddButton("Medium", d2gui.ButtonStyleMedium); err != nil {
		fmt.Printf("could not add button: %s to the GuiTestMain screen\n", "Medium")
	}

	if _, err := layoutRight.AddButton("Narrow", d2gui.ButtonStyleNarrow); err != nil {
		fmt.Printf("could not add button: %s to the GuiTestMain screen\n", "Narrow")
	}

	if _, err := layoutRight.AddButton("OkCancel", d2gui.ButtonStyleOkCancel); err != nil {
		fmt.Printf("could not add button: %s to the GuiTestMain screen\n", "OkCancel")
	}

	if _, err := layoutRight.AddButton("Short", d2gui.ButtonStyleShort); err != nil {
		fmt.Printf("could not add button: %s to the GuiTestMain screen\n", "Short")
	}

	if _, err := layoutRight.AddButton("Wide", d2gui.ButtonStyleWide); err != nil {
		fmt.Printf("could not add button: %s to the GuiTestMain screen\n", "Wide")
	}

	loading.Progress(ninetyPercent)

	layout.SetVerticalAlign(d2gui.VerticalAlignMiddle)
	g.guiManager.SetLayout(layout)
}

// Render does nothing for the GuiTestMain screen
func (g *GuiTestMain) Render(_ d2interface.Surface) error {
	return nil
}

// Advance does nothing for the GuiTestMain screen
func (g *GuiTestMain) Advance(_ float64) error {
	return nil
}
