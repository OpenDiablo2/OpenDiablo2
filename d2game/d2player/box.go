package d2player

import (
	"image/color"
	"log"
	"math"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
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

type BoxScrollbar struct {
	x                 int
	y                 int
	maxY              int
	minY              int
	arrowSliderOffset int
	parent            *Box
	mainLayout        *d2gui.Layout
	sliderLayout      *d2gui.Layout
	sprites           []*d2ui.Sprite
	viewportSize      int
	contentSize       int

	sliderClicked      bool
	arrowUpClicked     bool
	arrowDownClicked   bool
	clickedAtY         int
	mouseYOnSlider     int
	lastY              int
	gutterHeight       int
	contentToViewRatio float32
}

const (
	textSliderPartWidth  = 12
	textSliderPartHeight = 13
)

type textSliderPart int

const (
	textSliderPartArrowDownHollow textSliderPart = 8
	textSliderPartArrowUpHollow
	textSliderPartArrowDownFilled
	textSliderPartArrowUpFilled
	textSliderPartSquare
	textSliderPartInnerGutter
	textSliderPartFillingVariation1
)

func newBoxScrollBar(
	box *Box,
	x, y int,
	height int,
	contentSize, viewportSize int,
) *BoxScrollbar {
	ret := &BoxScrollbar{}

	arrowUpLayout := box.layout.AddLayout(d2gui.PositionTypeAbsolute)
	arrowUpLayout.SetSize(textSliderPartWidth, textSliderPartHeight)
	arrowUpLayout.SetPosition(x, y)
	arrowUpLayout.SetMouseClickHandler(ret.onArrowUpClick)

	arrowDownLayout := box.layout.AddLayout(d2gui.PositionTypeAbsolute)
	arrowDownLayout.SetPosition(x, y+height-textSliderPartHeight)
	arrowDownLayout.SetSize(textSliderPartWidth, textSliderPartHeight)
	arrowDownLayout.SetMouseClickHandler(ret.onArrowDownClick)

	viewportPercentage := 1.0 - (float32(contentSize-viewportSize) / float32(contentSize))
	ret.contentToViewRatio = viewportPercentage
	gutterHeight := height - 2*textSliderPartHeight
	ret.contentToViewRatio = float32(contentSize) / float32(gutterHeight)
	ret.gutterHeight = gutterHeight
	sliderHeight := int(float32(gutterHeight) * viewportPercentage)

	sliderLayout := box.layout.AddLayout(d2gui.PositionTypeAbsolute)
	sliderLayout.SetSize(textSliderPartWidth, sliderHeight)
	sliderLayout.SetPosition(x, y+textSliderPartHeight)
	sliderLayout.SetMouseClickHandler(ret.onSliderMouseClick)

	ret.minY = y + textSliderPartHeight
	ret.maxY = (y + height) - (textSliderPartHeight + sliderHeight)

	// ret.mainLayout = layout
	ret.sliderLayout = sliderLayout
	ret.viewportSize = viewportSize
	ret.parent = box
	ret.contentSize = contentSize
	ret.arrowSliderOffset = int(float32(sliderHeight) * 0.02)
	return ret
}

func (boxScroll *BoxScrollbar) onSliderMouseClick(event d2interface.MouseEvent) {
	if event.Button() == d2enum.MouseButtonLeft {
		boxScroll.sliderClicked = !boxScroll.sliderClicked
		if boxScroll.sliderClicked {
			boxScroll.clickedAtY = event.Y()
			boxScroll.lastY = boxScroll.clickedAtY
			boxScroll.mouseYOnSlider = event.Y() - boxScroll.sliderLayout.Sy
		}
	}
}

func (boxScroll *BoxScrollbar) onArrowUpClick(event d2interface.MouseEvent) {
	if event.Button() == d2enum.MouseButtonLeft {
		boxScroll.arrowUpClicked = !boxScroll.arrowUpClicked
	}
}

func (boxScroll *BoxScrollbar) onArrowDownClick(event d2interface.MouseEvent) {
	if event.Button() == d2enum.MouseButtonLeft {
		boxScroll.arrowDownClicked = !boxScroll.arrowDownClicked
	}
}

func (boxScroll *BoxScrollbar) update(cursorPosY int) {
	if boxScroll.arrowDownClicked {
		x, y := boxScroll.sliderLayout.GetPosition()
		newY := y + boxScroll.arrowSliderOffset
		if newY > boxScroll.maxY {
			newY = boxScroll.maxY
		}

		boxScroll.sliderLayout.SetPosition(x, newY)
	}

	if boxScroll.arrowUpClicked {
		x, y := boxScroll.sliderLayout.GetPosition()
		newY := y - boxScroll.arrowSliderOffset
		if newY < boxScroll.minY {
			newY = boxScroll.minY
		}

		boxScroll.sliderLayout.SetPosition(x, newY)
	}

	if boxScroll.sliderClicked {
		offset := cursorPosY - boxScroll.clickedAtY
		x, y := boxScroll.sliderLayout.GetPosition()
		newY := y + offset
		outOfBoundsUp := false
		outOfBoundsDown := false

		if newY > boxScroll.maxY {
			newY = boxScroll.maxY
			outOfBoundsDown = true
		}

		if newY < boxScroll.minY {
			newY = boxScroll.minY
			outOfBoundsUp = true
		}

		boxScroll.sliderLayout.SetPosition(x, newY)
		if !outOfBoundsUp && !outOfBoundsDown {
			boxScroll.clickedAtY = boxScroll.clickedAtY + offset
			if boxScroll.parent.contentLayout != nil {
				contentX, contentY := boxScroll.parent.contentLayout.GetPosition()
				scaledOffset := int(math.Round(float64(float32(offset) * boxScroll.contentToViewRatio)))
				newContentY := contentY - scaledOffset
				boxScroll.parent.contentLayout.SetPosition(contentX, newContentY)
			}
		}

		if outOfBoundsDown && boxScroll.parent.contentLayout != nil {
			newContentY := -boxScroll.contentSize + boxScroll.viewportSize
			boxScroll.parent.contentLayout.SetPosition(0, newContentY)
		}

		if outOfBoundsUp && boxScroll.parent.contentLayout != nil {
			boxScroll.parent.contentLayout.SetPosition(0, 0)
		}

	}
}

type BoxOption struct {
	label      string
	callback   func()
	hoverColor color.RGBA
	canHover   bool
}

// Box represents the menu to view/edit the
// key bindings
type Box struct {
	asset           *d2asset.AssetManager
	isOpen          bool
	renderer        d2interface.Renderer
	sprites         []*d2ui.Sprite
	uiManager       *d2ui.UIManager
	layout          *d2gui.Layout
	innerLayout     *d2gui.Layout
	contentLayout   *d2gui.Layout
	scrollBar       *BoxScrollbar
	guiManager      *d2gui.GuiManager
	sfc             d2interface.Surface
	width           int
	height          int
	x               int
	y               int
	Options         []*BoxOption
	enableScrollbar bool
}

func NewBox(
	asset *d2asset.AssetManager,
	renderer d2interface.Renderer,
	ui *d2ui.UIManager,
	guiManager *d2gui.GuiManager,
	contentLayout *d2gui.Layout,
	width, height int,
	x, y int,
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
	currentX, currentY := box.x, box.y+offsetY
	maxPieces := box.width / boxSpriteWidth
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
func (box *Box) setupBottomBorder() {
	bottomEdgePiece := []int{
		boxBottomHorizontalEdge1,
		boxBottomHorizontalEdge2,
		boxBottomHorizontalEdge3,
		boxBottomHorizontalEdge4,
		boxBottomHorizontalEdge5,
		boxBottomHorizontalEdge6,
	}

	i := 0
	currentX, currentY := box.x, box.y+box.height-boxSpriteHeight+9
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
	box.innerLayout = box.layout.AddLayout(d2gui.PositionTypeAbsolute)

	box.innerLayout.SetPosition(box.x, box.y-boxSpriteHeight-4)
	if box.contentLayout != nil {
		box.innerLayout.AddLayoutFromSource(box.contentLayout)
	}
	box.layout.SetPosition(box.x, box.y-boxSpriteHeight)
	box.layout.SetSize(box.width, box.height)
	box.layout.SetVisible(false)
	box.layout.SetLayer(0)

	box.setupTopBorder(0)
	box.setupBottomBorder()
	box.setupLeftBorder()
	box.setupRightBorder()
	box.setupCorners()

	box.Options = append(box.Options, []*BoxOption{
		{label: "Cancel", hoverColor: d2util.Color(0xD03C39FF), canHover: true},
		{label: "Default", hoverColor: d2util.Color(0x5450D1FF), canHover: true},
		{label: "Accept", hoverColor: d2util.Color(0x00D000FF), canHover: true},
	}...)

	bottomSectionHeight := int(float32(box.height) * 0.12)
	optionsEnabled := len(box.Options) > 0 && bottomSectionHeight > 14
	var innerSectionHeight int
	if optionsEnabled {
		innerSectionHeight = box.height - bottomSectionHeight
		box.innerLayout.SetSize(box.width+2, (box.height - bottomSectionHeight))
		cornerLeft, err := box.uiManager.NewSprite(d2resource.BoxPieces, d2resource.PaletteSky)
		if err != nil {
			log.Print(err)
		}

		cornerRight, err := box.uiManager.NewSprite(d2resource.BoxPieces, d2resource.PaletteSky)
		if err != nil {
			log.Print(err)
		}

		offsetY := box.y + box.height - bottomSectionHeight
		cornerLeft.SetCurrentFrame(boxCornerTopLeft)
		cornerLeft.SetPosition(box.x, offsetY)
		cornerRight.SetCurrentFrame(boxCornerTopRight)
		cornerRight.SetPosition(box.x+box.width-boxSpriteWidth, offsetY)
		box.setupTopBorder(box.height - 4*boxSpriteHeight)
		box.sprites = append(box.sprites, cornerLeft, cornerRight)

		buttonsLayoutWrapper := box.layout.AddLayout(d2gui.PositionTypeAbsolute)
		buttonsLayoutWrapper.SetSize(box.width+2, bottomSectionHeight)
		buttonsLayoutWrapper.SetPosition(box.x, box.y+box.height-bottomSectionHeight-boxSpriteHeight-4)
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

			l.SetLayer(0)
		}

		buttonsLayout.AddSpacerDynamic()
	} else {
		innerSectionHeight = box.height
		box.innerLayout.SetSize(box.width+2, box.height)
	}

	scrollBarX, scrollBarY, scrollBarHeight := box.x+box.width-12, box.y-boxSpriteHeight, box.height-5
	if optionsEnabled {
		scrollBarHeight -= bottomSectionHeight
	}

	if box.contentLayout != nil {
		_, contentH := box.contentLayout.GetSize()
		box.scrollBar = newBoxScrollBar(box, scrollBarX, scrollBarY, scrollBarHeight, contentH, innerSectionHeight)
	}
}

func (box *Box) Update() {
	_, cursorY := box.renderer.GetCursorPos()
	box.scrollBar.update(cursorY)
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
