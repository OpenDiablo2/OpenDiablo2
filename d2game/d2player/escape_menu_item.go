package d2player

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
)

type Item struct {
	label     d2ui.Label
	text      string
	isHovered bool
}

func newItem(text string) *Item {
	return &Item{
		text: text,
	}
}

func (i *Item) Render(target d2render.Surface) {

}

func (i *Item) Load() {
	i.label = d2ui.CreateLabel(d2resource.Font42, d2resource.PaletteSky)
	i.label.SetText(i.text)
	i.label.Alignment = d2ui.LabelAlignCenter
}
