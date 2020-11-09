package d2player

import (
	"fmt"
	"image/color"
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2gui"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
)

const (
	boxSpriteHeight = 15 - 3
	boxSpriteWidth  = 14 - 2
)

const (
	boxCornerTopLeft = iota
	boxCornerTopRight
	boxTopHorizontalEdge1
	boxTopHorizontalEdge2
	boxTopHorizontalEdge3
	boxTopHorizontalEdge4
	boxTopHorizontalEdge5
	boxTopHorizontalEdge6
	boxCornerBottomLeft
	boxCornerBottomRight
	boxSideEdge1
	boxSideEdge2
	boxSideEdge3
	boxSideEdge4
	boxSideEdge5
	boxSideEdge6
	boxBottomHorizontalEdge1
	boxBottomHorizontalEdge2
	boxBottomHorizontalEdge3
	boxBottomHorizontalEdge4
	boxBottomHorizontalEdge5
	boxBottomHorizontalEdge6
)

type LabelButton struct {
	label      string
	callback   func()
	hoverColor color.RGBA
	canHover   bool
	isHovered  bool
	layout     *d2gui.Layout
	x, y       int
}

func NewLabelButton(x, y int, text string, col color.RGBA, callback func()) *LabelButton {
	return &LabelButton{
		x:          x,
		y:          y,
		hoverColor: col,
		label:      text,
		callback:   callback,
		canHover:   true,
	}
}

func (lb *LabelButton) Load(renderer d2interface.Renderer, asset *d2asset.AssetManager) {
	mainLayout := d2gui.CreateLayout(renderer, d2gui.PositionTypeAbsolute, asset)
	l, _ := mainLayout.AddLabelWithColor(lb.label, d2gui.FontStyleFormal11Units, d2util.Color(0xA1925DFF))

	if lb.canHover {
		l.SetHoverColor(lb.hoverColor)
	}

	mainLayout.SetMouseClickHandler(func(d d2interface.MouseEvent) {
		if lb.callback != nil {
			lb.callback()
		}
	})

	mainLayout.SetMouseEnterHandler(func(event d2interface.MouseMoveEvent) {
		l.SetIsHovered(true)
	})

	mainLayout.SetMouseLeaveHandler(func(event d2interface.MouseMoveEvent) {
		l.SetIsHovered(false)
	})

	lb.layout = mainLayout
}

func (lb *LabelButton) SetLabel(val string) {
	lb.label = val
}

func (lb *LabelButton) SetHoverColor(col color.RGBA) {
	lb.hoverColor = col
}

func (lb *LabelButton) SetCanHover(val bool) {
	lb.canHover = val
}

func (lb *LabelButton) IsHovered() bool {
	return lb.isHovered
}

func (lb *LabelButton) GetLayout() *d2gui.Layout {
	return lb.layout
}

// Box represents the menu to view/edit the
// key bindings
type Box struct {
	renderer      d2interface.Renderer
	asset         *d2asset.AssetManager
	isOpen        bool
	sprites       []*d2ui.Sprite
	uiManager     *d2ui.UIManager
	layout        *d2gui.Layout
	contentLayout *d2gui.Layout
	guiManager    *d2gui.GuiManager
	sfc           d2interface.Surface
	title         string
	width         int
	height        int
	x             int
	y             int
	Options       []*LabelButton
	disableBorder bool
}

func NewBox(
	asset *d2asset.AssetManager,
	renderer d2interface.Renderer,
	ui *d2ui.UIManager,
	guiManager *d2gui.GuiManager,
	contentLayout *d2gui.Layout,
	width, height int,
	x, y int,
	title string,
) *Box {
	return &Box{
		asset:         asset,
		renderer:      renderer,
		uiManager:     ui,
		guiManager:    guiManager,
		width:         width,
		height:        height,
		contentLayout: contentLayout,
		sfc:           renderer.NewSurface(width, height),
		title:         title,
		x:             x,
		y:             y,
	}
}

// Toggle the visibility state of the menu
func (box *Box) Toggle() {
	if box.isOpen {
		box.Close()
	} else {
		box.open()
	}
}

func (box *Box) open() {
	box.isOpen = true
	box.guiManager.SetLayout(box.layout)
}

// Close will hide the help overlay
func (box *Box) Close() {
	box.isOpen = false
	box.guiManager.SetLayout(nil)
}

// IsOpen returns whether or not the overlay is visible/open
func (box *Box) IsOpen() bool {
	return box.isOpen
}

func (box *Box) setupTopBorder(offsetY int) {
	topEdgePiece := []int{
		boxTopHorizontalEdge1,
		boxTopHorizontalEdge2,
		boxTopHorizontalEdge3,
		boxTopHorizontalEdge4,
		boxTopHorizontalEdge5,
		boxTopHorizontalEdge6,
	}

	i := 0
	maxPieces := box.width / boxSpriteWidth
	currentX, currentY := box.x, box.y+offsetY
	for {
		for _, frameIndex := range topEdgePiece {
			f, err := box.uiManager.NewSprite(d2resource.BoxPieces, d2resource.PaletteSky)
			if err != nil {
				log.Print(err)
			}

			err = f.SetCurrentFrame(frameIndex)
			if err != nil {
				log.Print(err)
			}

			f.SetPosition(currentX, currentY)
			currentX = currentX + boxSpriteWidth
			i++

			box.sprites = append(box.sprites, f)
			if i >= maxPieces {
				break
			}
		}
		if i >= maxPieces {
			break
		}
	}
}
func (box *Box) setupBottomBorder(offsetY int) {
	bottomEdgePiece := []int{
		boxBottomHorizontalEdge1,
		boxBottomHorizontalEdge2,
		boxBottomHorizontalEdge3,
		boxBottomHorizontalEdge4,
		boxBottomHorizontalEdge5,
		boxBottomHorizontalEdge6,
	}

	i := 0
	currentX, currentY := box.x, offsetY
	maxPieces := box.width / boxSpriteWidth
	for {
		for _, frameIndex := range bottomEdgePiece {
			f, err := box.uiManager.NewSprite(d2resource.BoxPieces, d2resource.PaletteSky)
			if err != nil {
				log.Print(err)
			}

			err = f.SetCurrentFrame(frameIndex)
			if err != nil {
				log.Print(err)
			}

			f.SetPosition(currentX, currentY)
			currentX = currentX + boxSpriteWidth
			i++

			box.sprites = append(box.sprites, f)
			if i >= maxPieces {
				break
			}
		}
		if i >= maxPieces {
			break
		}
	}
}

func (box *Box) setupLeftBorder() {
	leftBorderPiece := []int{
		boxSideEdge1,
		boxSideEdge2,
		boxSideEdge3,
	}

	currentX, currentY := box.x-4, box.y
	maxPieces := box.height / boxSpriteHeight
	i := 0
	for {
		for _, frameIndex := range leftBorderPiece {
			f, err := box.uiManager.NewSprite(d2resource.BoxPieces, d2resource.PaletteSky)
			if err != nil {
				log.Print(err)
			}

			err = f.SetCurrentFrame(frameIndex)
			if err != nil {
				log.Print(err)
			}

			f.SetPosition(currentX, currentY)
			currentY = currentY + boxSpriteHeight
			i++

			box.sprites = append(box.sprites, f)
			if i >= maxPieces {
				break
			}
		}
		if i >= maxPieces {
			break
		}
	}
}
func (box *Box) setupRightBorder() {
	rightBorderPiece := []int{
		boxSideEdge4,
		boxSideEdge5,
		boxSideEdge6,
	}

	i := 0
	currentX, currentY := box.width+box.x-7, box.y
	maxPieces := box.height / boxSpriteHeight
	for {
		for _, frameIndex := range rightBorderPiece {
			f, err := box.uiManager.NewSprite(d2resource.BoxPieces, d2resource.PaletteSky)
			if err != nil {
				log.Print(err)
			}

			err = f.SetCurrentFrame(frameIndex)
			if err != nil {
				log.Print(err)
			}

			f.SetPosition(currentX, currentY)
			currentY = currentY + boxSpriteHeight
			i++

			box.sprites = append(box.sprites, f)
			if i >= maxPieces {
				break
			}
		}
		if i >= maxPieces {
			break
		}
	}
}

func (box *Box) setupCorners() {
	cornersFrames := []int{
		boxCornerTopLeft,
		boxCornerTopRight,
		boxCornerBottomLeft,
		boxCornerBottomRight,
	}

	for _, frameIndex := range cornersFrames {
		f, err := box.uiManager.NewSprite(d2resource.BoxPieces, d2resource.PaletteSky)
		if err != nil {
			log.Print(err)
		}

		err = f.SetCurrentFrame(frameIndex)
		if err != nil {
			log.Print(err)
		}

		switch frameIndex {
		case boxCornerTopLeft:
			f.SetPosition(box.x, box.y)
			break
		case boxCornerTopRight:
			f.SetPosition(box.x+box.width-boxSpriteWidth, box.y)
			break
		case boxCornerBottomLeft:
			f.SetPosition(box.x, box.y+box.height-boxSpriteHeight)
			break
		case boxCornerBottomRight:
			f.SetPosition(box.x+box.width-boxSpriteWidth, box.y+box.height-boxSpriteHeight)
			break
		}

		box.sprites = append(box.sprites, f)
	}
}

func (box *Box) Load() {
	box.layout = d2gui.CreateLayout(box.renderer, d2gui.PositionTypeAbsolute, box.asset)
	box.layout.SetPosition(box.x, box.y)
	box.layout.SetSize(box.width, box.height)
	box.contentLayout.SetPosition(box.x, box.y)

	if !box.disableBorder {
		box.setupTopBorder(0)
		box.setupBottomBorder(box.y + box.height - boxSpriteHeight + 10)
		box.setupLeftBorder()
		box.setupRightBorder()
		box.setupCorners()
	}

	box.Options = append(box.Options, []*LabelButton{
		NewLabelButton(0, 0, "Cancel", d2util.Color(0xD03C39FF), func() { fmt.Println("Cancel callback") }),
		NewLabelButton(0, 0, "Default", d2util.Color(0x5450D1FF), func() { fmt.Println("Default callback") }),
		NewLabelButton(0, 0, "Accept", d2util.Color(0x00D000FF), func() { fmt.Println("Accept callback") }),
	}...)

	sectionHeight := int(float32(box.height) * 0.12)
	optionsEnabled := len(box.Options) > 0 && sectionHeight > 14
	if optionsEnabled {
		box.contentLayout.SetSize(box.width+2, (box.height - sectionHeight))
		if !box.disableBorder {
			cornerLeft, err := box.uiManager.NewSprite(d2resource.BoxPieces, d2resource.PaletteSky)
			if err != nil {
				log.Print(err)
			}

			cornerRight, err := box.uiManager.NewSprite(d2resource.BoxPieces, d2resource.PaletteSky)
			if err != nil {
				log.Print(err)
			}

			offsetY := box.y + box.height - sectionHeight
			cornerLeft.SetCurrentFrame(boxCornerTopLeft)
			cornerLeft.SetPosition(box.x, offsetY)
			cornerRight.SetCurrentFrame(boxCornerTopRight)
			cornerRight.SetPosition(box.x+box.width-boxSpriteWidth, offsetY)
			box.setupTopBorder(box.height - (4 * boxSpriteHeight) + 3)
			box.sprites = append(box.sprites, cornerLeft, cornerRight)
		}

		buttonsLayoutWrapper := box.layout.AddLayout(d2gui.PositionTypeAbsolute)
		buttonsLayoutWrapper.SetSize(box.width+2, sectionHeight)
		buttonsLayoutWrapper.SetPosition(box.x, box.y+box.height-sectionHeight)
		buttonsLayout := buttonsLayoutWrapper.AddLayout(d2gui.PositionTypeHorizontal)
		buttonsLayout.SetSize(buttonsLayoutWrapper.GetSize())
		buttonsLayout.SetVerticalAlign(d2gui.VerticalAlignMiddle)
		buttonsLayout.AddSpacerDynamic()
		for _, option := range box.Options {
			option.Load(box.renderer, box.asset)
			buttonsLayout.AddLayoutFromSource(option.GetLayout())
			buttonsLayout.AddSpacerDynamic()
		}
	} else {
		box.contentLayout.SetSize(box.width+2, box.height)
	}

	if box.title != "" {
		if !box.disableBorder {
			cornerLeft, err := box.uiManager.NewSprite(d2resource.BoxPieces, d2resource.PaletteSky)
			if err != nil {
				log.Print(err)
			}

			cornerRight, err := box.uiManager.NewSprite(d2resource.BoxPieces, d2resource.PaletteSky)
			if err != nil {
				log.Print(err)
			}

			offsetY := box.y + sectionHeight

			cornerLeft.SetCurrentFrame(boxCornerBottomLeft)
			cornerLeft.SetPosition(box.x, offsetY-10)

			cornerRight.SetCurrentFrame(boxCornerBottomRight)
			cornerRight.SetPosition(box.x+box.width-boxSpriteWidth, offsetY-10)

			box.sprites = append(box.sprites, cornerLeft, cornerRight)
			box.setupBottomBorder(offsetY)
		}

		contentLayoutW, contentLayoutH := box.contentLayout.GetSize()
		contentLayoutX, contentLayoutY := box.contentLayout.GetPosition()
		box.contentLayout.SetSize(contentLayoutW, contentLayoutH-sectionHeight)
		box.contentLayout.SetPosition(contentLayoutX, contentLayoutY+sectionHeight)

		titleLayout := box.layout.AddLayout(d2gui.PositionTypeHorizontal)
		titleLayout.SetHorizontalAlign(d2gui.HorizontalAlignCenter)
		titleLayout.SetVerticalAlign(d2gui.VerticalAlignMiddle)
		titleLayout.SetPosition(box.x, box.y)
		titleLayout.SetSize(contentLayoutW, sectionHeight)
		titleLayout.AddSpacerDynamic()
		titleLayout.AddLabel(box.title, d2gui.FontStyle30Units)
		titleLayout.AddSpacerDynamic()
	}

	box.layout.AddLayoutFromSource(box.contentLayout)
}

// Render the overlay to the given surface
func (box *Box) Render(target d2interface.Surface) error {
	if !box.isOpen {
		return nil
	}

	target.PushTranslation(box.x, box.y-boxSpriteHeight)
	target.DrawRect(box.width, box.height, d2util.Color(0x000000D0))
	target.Pop()

	for _, s := range box.sprites {
		s.Render(target)
	}

	return nil
}

// func (box *Box) OnMouseButtonDown(event d2interface.MouseEvent) {
//   if box.scrollbar != nil && box.scrollbar.IsInSliderRect(event.X(), event.Y()) {
//     box.scrollbar.SetSliderClicked(true)
//     box.scrollbar.onSliderMouseClick(event)
//   }
// }
//
// func (box *Box) OnMouseMove(event d2interface.MouseMoveEvent) {
//   if box.scrollbar != nil {
//     box.scrollbar.onMouseMove(event)
//   }
// }
//
// func (box *Box) OnMouseButtonUp(event d2interface.MouseEvent) {
//   if box.scrollbar != nil {
//     box.scrollbar.SetSliderClicked(false)
//   }
// }

// IsInRect checks if the given point is within the overlay layout rectangle
func (box *Box) IsInRect(px, py int) bool {
	ww, hh := box.layout.GetSize()
	x, y := box.layout.GetPosition()

	if px >= x && px <= x+ww && py >= y && py <= y+hh {
		return true
	}

	return false
}
