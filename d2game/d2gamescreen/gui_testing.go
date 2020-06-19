package d2gamescreen

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2gui"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
)

type GuiTestMain struct{}

func CreateGuiTestMain() *GuiTestMain {
	return &GuiTestMain{}
}

func (g *GuiTestMain) OnLoad() error {
	layout := d2gui.CreateLayout(d2gui.PositionTypeHorizontal)
	//
	layoutLeft := layout.AddLayout(d2gui.PositionTypeVertical)
	layoutLeft.SetHorizontalAlign(d2gui.HorizontalAlignCenter)
	layoutLeft.AddLabel("FontStyle16Units", d2gui.FontStyle16Units)
	layoutLeft.AddSpacerStatic(0, 100)
	layoutLeft.AddLabel("FontStyle30Units", d2gui.FontStyle30Units)
	layoutLeft.AddLabel("FontStyle42Units", d2gui.FontStyle42Units)
	layoutLeft.AddLabel("FontStyleFormal10Static", d2gui.FontStyleFormal10Static)
	layoutLeft.AddLabel("FontStyleFormal11Units", d2gui.FontStyleFormal11Units)
	layoutLeft.AddLabel("FontStyleFormal12Static", d2gui.FontStyleFormal12Static)

	layout.AddSpacerDynamic()

	layoutRight := layout.AddLayout(d2gui.PositionTypeVertical)
	layoutRight.SetHorizontalAlign(d2gui.HorizontalAlignRight)
	layoutRight.AddButton("Medium", d2gui.ButtonStyleMedium)
	layoutRight.AddButton("Narrow", d2gui.ButtonStyleNarrow)
	layoutRight.AddButton("OkCancel", d2gui.ButtonStyleOkCancel)
	layoutRight.AddButton("Short", d2gui.ButtonStyleShort)
	layoutRight.AddButton("Wide", d2gui.ButtonStyleWide)

	layout.SetVerticalAlign(d2gui.VerticalAlignMiddle)
	d2gui.SetLayout(layout)

	return nil
}

func (g *GuiTestMain) Render(screen d2render.Surface) error {
	return nil
}

func (g *GuiTestMain) Advance(tickTime float64) error {
	return nil
}
