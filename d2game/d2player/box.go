package d2player

import (
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

type BoxOption struct {
	label      string
	callback   func()
	hoverColor color.RGBA
	canHover   bool
}

// Box represents the menu to view/edit the
// key bindings
type Box struct {
	asset            *d2asset.AssetManager
	isOpen           bool
	renderer         d2interface.Renderer
	sprites          []*d2ui.Sprite
	uiManager        *d2ui.UIManager
	layout           *d2gui.Layout
	innerLayout      *d2gui.Layout
	contentLayout    *d2gui.Layout
	scrollbar        *BoxScrollbar
	guiManager       *d2gui.GuiManager
	sfc              d2interface.Surface
	title            string
	width            int
	height           int
	x                int
	y                int
	Options          []*BoxOption
	disableScrollbar bool
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
	box.layout.SetPosition(box.x, box.y-boxSpriteHeight)
	box.layout.SetSize(box.width, box.height)
	box.layout.SetVisible(false)
	box.layout.SetLayer(0)

	box.innerLayout = box.layout.AddLayout(d2gui.PositionTypeAbsolute)
	box.innerLayout.SetPosition(box.x, box.y-boxSpriteHeight-4)

	box.setupTopBorder(0)
	box.setupBottomBorder(box.y + box.height - boxSpriteHeight + 10)
	box.setupLeftBorder()
	box.setupRightBorder()
	box.setupCorners()

	box.Options = append(box.Options, []*BoxOption{
		{label: "Cancel", hoverColor: d2util.Color(0xD03C39FF), canHover: true},
		{label: "Default", hoverColor: d2util.Color(0x5450D1FF), canHover: true},
		{label: "Accept", hoverColor: d2util.Color(0x00D000FF), canHover: true},
	}...)

	sectionHeight := int(float32(box.height) * 0.12)
	optionsEnabled := len(box.Options) > 0 && sectionHeight > 14
	if optionsEnabled {
		box.innerLayout.SetSize(box.width+2, (box.height - sectionHeight))
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

		buttonsLayoutWrapper := box.layout.AddLayout(d2gui.PositionTypeAbsolute)
		buttonsLayoutWrapper.SetSize(box.width+2, sectionHeight)
		buttonsLayoutWrapper.SetPosition(box.x, box.y+box.height-sectionHeight-boxSpriteHeight)
		buttonsLayout := buttonsLayoutWrapper.AddLayout(d2gui.PositionTypeHorizontal)
		buttonsLayout.SetSize(buttonsLayoutWrapper.GetSize())
		buttonsLayout.SetVerticalAlign(d2gui.VerticalAlignMiddle)
		for _, option := range box.Options {
			buttonsLayout.AddSpacerDynamic()
			l, _ := buttonsLayout.AddLabelWithColor(option.label, d2gui.FontStyleFormal11Units, d2util.Color(0xA1925DFF))

			if option.canHover {
				l.SetHoverColor(option.hoverColor)
			}

			l.SetMouseEnterHandler(func(event d2interface.MouseMoveEvent) {
				l.SetIsHovered(true)
			})

			l.SetMouseLeaveHandler(func(event d2interface.MouseMoveEvent) {
				l.SetIsHovered(false)
			})
		}

		buttonsLayout.AddSpacerDynamic()
	} else {
		box.innerLayout.SetSize(box.width+2, box.height)
	}

	if box.title != "" {
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

		innerLayoutW, innerLayoutH := box.innerLayout.GetSize()
		innerLayoutX, innerLayoutY := box.layout.GetPosition()
		box.innerLayout.SetSize(innerLayoutW, innerLayoutH-sectionHeight)
		box.innerLayout.SetPosition(innerLayoutX, innerLayoutY+sectionHeight)

		titleLayout := box.layout.AddLayout(d2gui.PositionTypeHorizontal)
		titleLayout.SetHorizontalAlign(d2gui.HorizontalAlignCenter)
		titleLayout.SetVerticalAlign(d2gui.VerticalAlignMiddle)
		titleLayout.SetPosition(box.x, box.y-boxSpriteHeight)
		titleLayout.SetSize(innerLayoutW, sectionHeight)
		titleLayout.AddSpacerDynamic()
		titleLayout.AddLabel(box.title, d2gui.FontStyle30Units)
		titleLayout.AddSpacerDynamic()

		box.setupBottomBorder(offsetY)
	}

	if !box.disableScrollbar && box.contentLayout != nil {
		box.scrollbar = newBoxScrollBar(
			box.innerLayout,
			box.contentLayout,
		)

		box.innerLayout.AddLayoutFromSource(box.contentLayout)
	}
}

func (box *Box) Update() {
	if box.isOpen && box.scrollbar != nil {
		_, cursorY := box.renderer.GetCursorPos()
		box.scrollbar.update(cursorY)
	}
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

// IsInRect checks if the given point is within the overlay layout rectangle
func (box *Box) IsInRect(px, py int) bool {
	ww, hh := box.layout.GetSize()
	x, y := box.layout.GetPosition()

	if px >= x && px <= x+ww && py >= y && py <= y+hh {
		return true
	}

	return false
}
