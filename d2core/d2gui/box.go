package d2gui

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
)

const (
	boxSpriteHeight = 15 - 5
	boxSpriteWidth  = 14 - 2

	boxBorderSpriteLeftBorderOffset       = 4
	boxBorderSpriteRightBorderOffset      = 7
	boxBorderSpriteTopBorderSectionOffset = 5

	minimumAllowedSectionSize    = 14
	sectionHeightPercentageOfBox = 0.12
	boxBackgroundColor           = 0x000000d0
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

// Box takes a content layout and wraps in
// a box
type Box struct {
	renderer      d2interface.Renderer
	asset         *d2asset.AssetManager
	sprites       []*d2ui.Sprite
	uiManager     *d2ui.UIManager
	layout        *Layout
	contentLayout *Layout
	Options       []*LabelButton
	sfc           d2interface.Surface

	x, y               int
	paddingX, paddingY int
	width, height      int
	disableBorder      bool
	isOpen             bool
	title              string

	*d2util.Logger
}

// NewBox return a new Box instance
func NewBox(
	asset *d2asset.AssetManager,
	renderer d2interface.Renderer,
	ui *d2ui.UIManager,
	contentLayout *Layout,
	width, height int,
	x, y int,
	l d2util.LogLevel,
	title string,
) *Box {
	box := &Box{
		asset:         asset,
		renderer:      renderer,
		uiManager:     ui,
		width:         width,
		height:        height,
		contentLayout: contentLayout,
		sfc:           renderer.NewSurface(width, height),
		title:         title,
		x:             x,
		y:             y,
	}

	box.Logger = d2util.NewLogger()
	box.Logger.SetLevel(l)
	box.Logger.SetPrefix(logPrefix)

	return box
}

// GetLayout returns the box layout
func (box *Box) GetLayout() *Layout {
	return box.layout
}

// Toggle the visibility state of the menu
func (box *Box) Toggle() {
	if box.isOpen {
		box.Close()
	} else {
		box.Open()
	}
}

// SetPadding sets the padding of the box content
func (box *Box) SetPadding(paddingX, paddingY int) {
	box.paddingX = paddingX
	box.paddingY = paddingY
}

// Open will set the isOpen value to true
func (box *Box) Open() {
	box.isOpen = true
}

// Close will hide the help overlay
func (box *Box) Close() {
	box.isOpen = false
}

// IsOpen returns whether or not the box is opened
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
				box.Error(err.Error())
			}

			err = f.SetCurrentFrame(frameIndex)
			if err != nil {
				box.Error(err.Error())
			}

			f.SetPosition(currentX, currentY)
			currentX += boxSpriteWidth

			box.sprites = append(box.sprites, f)

			i++
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
				box.Error(err.Error())
			}

			err = f.SetCurrentFrame(frameIndex)
			if err != nil {
				box.Error(err.Error())
			}

			f.SetPosition(currentX, currentY)
			currentX += boxSpriteWidth

			box.sprites = append(box.sprites, f)

			i++
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

	currentX, currentY := box.x-boxBorderSpriteLeftBorderOffset, box.y+boxSpriteHeight
	maxPieces := box.height / boxSpriteHeight
	i := 0

	for {
		for _, frameIndex := range leftBorderPiece {
			f, err := box.uiManager.NewSprite(d2resource.BoxPieces, d2resource.PaletteSky)
			if err != nil {
				box.Error(err.Error())
			}

			err = f.SetCurrentFrame(frameIndex)
			if err != nil {
				box.Error(err.Error())
			}

			f.SetPosition(currentX, currentY)
			currentY += boxSpriteHeight

			box.sprites = append(box.sprites, f)

			i++
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
	currentX, currentY := box.width+box.x-boxBorderSpriteRightBorderOffset, box.y+boxSpriteHeight
	maxPieces := box.height / boxSpriteHeight

	for {
		for _, frameIndex := range rightBorderPiece {
			f, err := box.uiManager.NewSprite(d2resource.BoxPieces, d2resource.PaletteSky)
			if err != nil {
				box.Error(err.Error())
			}

			err = f.SetCurrentFrame(frameIndex)
			if err != nil {
				box.Error(err.Error())
			}

			f.SetPosition(currentX, currentY)
			currentY += boxSpriteHeight

			box.sprites = append(box.sprites, f)

			i++
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
			box.Error(err.Error())
		}

		err = f.SetCurrentFrame(frameIndex)
		if err != nil {
			box.Error(err.Error())
		}

		switch frameIndex {
		case boxCornerTopLeft:
			f.SetPosition(box.x, box.y+boxSpriteHeight)
		case boxCornerTopRight:
			f.SetPosition(box.x+box.width-boxSpriteWidth, box.y+boxSpriteHeight)
		case boxCornerBottomLeft:
			f.SetPosition(box.x, box.y+box.height)
		case boxCornerBottomRight:
			f.SetPosition(box.x+box.width-boxSpriteWidth, box.y+box.height)
		}

		box.sprites = append(box.sprites, f)
	}
}

// SetOptions sets the box options that will show up at the bottom
func (box *Box) SetOptions(options []*LabelButton) {
	box.Options = options
}

func (box *Box) setupTitle(sectionHeight int) error {
	if !box.disableBorder {
		cornerLeft, err := box.uiManager.NewSprite(d2resource.BoxPieces, d2resource.PaletteSky)
		if err != nil {
			return err
		}

		cornerRight, err := box.uiManager.NewSprite(d2resource.BoxPieces, d2resource.PaletteSky)
		if err != nil {
			return err
		}

		offsetY := box.y + sectionHeight

		if err := cornerLeft.SetCurrentFrame(boxCornerBottomLeft); err != nil {
			return err
		}

		cornerLeft.SetPosition(box.x, offsetY)

		if err := cornerRight.SetCurrentFrame(boxCornerBottomRight); err != nil {
			return err
		}

		cornerRight.SetPosition(box.x+box.width-boxSpriteWidth, offsetY)

		box.sprites = append(box.sprites, cornerLeft, cornerRight)
		box.setupBottomBorder(offsetY)
	}

	contentLayoutW, contentLayoutH := box.contentLayout.GetSize()
	contentLayoutX, contentLayoutY := box.contentLayout.GetPosition()
	box.contentLayout.SetSize(contentLayoutW, contentLayoutH-sectionHeight)
	box.contentLayout.SetPosition(contentLayoutX, contentLayoutY+sectionHeight)

	titleLayout := box.layout.AddLayout(PositionTypeHorizontal)
	titleLayout.SetHorizontalAlign(HorizontalAlignCenter)
	titleLayout.SetVerticalAlign(VerticalAlignMiddle)
	titleLayout.SetPosition(box.x, box.y)
	titleLayout.SetSize(contentLayoutW, sectionHeight)
	titleLayout.AddSpacerDynamic()

	if _, err := titleLayout.AddLabel(box.title, FontStyle30Units); err != nil {
		return err
	}

	titleLayout.AddSpacerDynamic()

	return nil
}

func (box *Box) setupOptions(sectionHeight int) error {
	box.contentLayout.SetSize(box.width, (box.height - sectionHeight))

	if !box.disableBorder {
		cornerLeft, err := box.uiManager.NewSprite(d2resource.BoxPieces, d2resource.PaletteSky)
		if err != nil {
			return err
		}

		cornerRight, err := box.uiManager.NewSprite(d2resource.BoxPieces, d2resource.PaletteSky)
		if err != nil {
			return err
		}

		offsetY := box.y + box.height - sectionHeight + boxSpriteHeight

		if err := cornerLeft.SetCurrentFrame(boxCornerTopLeft); err != nil {
			return err
		}

		cornerLeft.SetPosition(box.x, offsetY)

		if err := cornerRight.SetCurrentFrame(boxCornerTopRight); err != nil {
			return err
		}

		cornerRight.SetPosition(box.x+box.width-boxSpriteWidth, offsetY)
		box.setupTopBorder(box.height - (4 * boxSpriteHeight) + boxSpriteHeight - boxBorderSpriteTopBorderSectionOffset)
		box.sprites = append(box.sprites, cornerLeft, cornerRight)
	}

	buttonsLayoutWrapper := box.layout.AddLayout(PositionTypeAbsolute)
	buttonsLayoutWrapper.SetSize(box.width, sectionHeight)
	buttonsLayoutWrapper.SetPosition(box.x, box.y+box.height-sectionHeight)
	buttonsLayout := buttonsLayoutWrapper.AddLayout(PositionTypeHorizontal)
	buttonsLayout.SetSize(buttonsLayoutWrapper.GetSize())
	buttonsLayout.SetVerticalAlign(VerticalAlignMiddle)
	buttonsLayout.AddSpacerDynamic()

	for _, option := range box.Options {
		option.Load(box.renderer, box.asset)
		buttonsLayout.AddLayoutFromSource(option.GetLayout())
		buttonsLayout.AddSpacerDynamic()
	}

	return nil
}

// Load will setup the layouts and sprites for the box deptending on the parameters
func (box *Box) Load() error {
	box.layout = CreateLayout(box.renderer, PositionTypeAbsolute, box.asset)
	box.layout.SetPosition(box.x, box.y)
	box.layout.SetSize(box.width, box.height)
	box.contentLayout.SetPosition(box.x, box.y)

	if !box.disableBorder {
		box.setupTopBorder(boxSpriteHeight)
		box.setupBottomBorder(box.y + box.height + boxSpriteHeight)
		box.setupLeftBorder()
		box.setupRightBorder()
		box.setupCorners()
	}

	sectionHeight := int(float32(box.height) * sectionHeightPercentageOfBox)

	optionsEnabled := len(box.Options) > 0 && sectionHeight >= minimumAllowedSectionSize
	if optionsEnabled {
		if err := box.setupOptions(sectionHeight); err != nil {
			return err
		}
	} else {
		box.contentLayout.SetSize(box.width, box.height)
	}

	if box.title != "" {
		if err := box.setupTitle(sectionHeight); err != nil {
			return err
		}
	}

	contentLayoutW, contentLayoutH := box.contentLayout.GetSize()
	contentLayoutX, contentLayoutY := box.contentLayout.GetPosition()
	box.contentLayout.SetPosition(contentLayoutX+box.paddingX, contentLayoutY+box.paddingY)
	box.contentLayout.SetSize(contentLayoutW-(2*box.paddingX), contentLayoutH-(2*box.paddingY))

	box.layout.AddLayoutFromSource(box.contentLayout)

	return nil
}

// OnMouseButtonDown will be called whenever a mouse button is triggered
func (box *Box) OnMouseButtonDown(event d2interface.MouseEvent) bool {
	for _, option := range box.Options {
		if option.IsInRect(event.X(), event.Y()) {
			option.callback()
			return true
		}
	}

	return false
}

// Render the box to the given surface
func (box *Box) Render(target d2interface.Surface) error {
	if !box.isOpen {
		return nil
	}

	target.PushTranslation(box.x, box.y)
	target.DrawRect(box.width, box.height, d2util.Color(boxBackgroundColor))
	target.Pop()

	for _, s := range box.sprites {
		s.Render(target)
	}

	return nil
}

// IsInRect checks if the given point is within the box main layout rectangle
func (box *Box) IsInRect(px, py int) bool {
	ww, hh := box.layout.GetSize()
	x, y := box.layout.GetPosition()

	if px >= x && px <= x+ww && py >= y && py <= y+hh {
		return true
	}

	return false
}
