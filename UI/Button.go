package UI

import (
	"image"
	"image/color"

	"github.com/essial/OpenDiablo2/Common"
	"github.com/essial/OpenDiablo2/Palettes"
	"github.com/essial/OpenDiablo2/ResourcePaths"
	"github.com/hajimehoshi/ebiten"
)

// ButtonType defines the type of button
type ButtonType int

const (
	ButtonTypeWide   ButtonType = 1
	ButtonTypeMedium ButtonType = 2
	ButtonTypeNarrow ButtonType = 3
	ButtonTypeCancel ButtonType = 4
	ButtonTypeTall   ButtonType = 5
	ButtonTypeShort  ButtonType = 6

	// Game UI

	ButtonTypeSkill              ButtonType = 7
	ButtonTypeRun                ButtonType = 8
	ButtonTypeMenu               ButtonType = 9
	ButtonTypeGoldCoin           ButtonType = 10
	ButtonTypeClose              ButtonType = 11
	ButtonTypeSecondaryInvHand   ButtonType = 12
	ButtonTypeMinipanelCharacter ButtonType = 13
	ButtonTypeMinipanelInventory ButtonType = 14
	ButtonTypeMinipanelSkill     ButtonType = 15
	ButtonTypeMinipanelAutomap   ButtonType = 16
	ButtonTypeMinipanelMessage   ButtonType = 17
	ButtonTypeMinipanelQuest     ButtonType = 18
	ButtonTypeMinipanelMen       ButtonType = 19
)

// ButtonLayout defines the type of buttons
type ButtonLayout struct {
	XSegments        int //1
	YSegments        int // 1
	ResourceName     string
	PaletteName      Palettes.Palette
	Toggleable       bool             // false
	BaseFrame        int              // 0
	DisabledFrame    int              // -1
	FontPath         string           // ResourcePaths.FontExocet10
	ClickableRect    *image.Rectangle // nil
	AllowFrameChange bool             // true
}

// ButtonLayouts define the type of buttons you can have
var ButtonLayouts = map[ButtonType]ButtonLayout{
	ButtonTypeWide:  {2, 1, ResourcePaths.WideButtonBlank, Palettes.Units, false, 0, -1, ResourcePaths.FontExocet10, nil, true},
	ButtonTypeShort: {1, 1, ResourcePaths.ShortButtonBlank, Palettes.Units, false, 0, -1, ResourcePaths.FontExocet8, nil, true},
	/*
		{eButtonType.Wide,  new ButtonLayout { XSegments = 2, ResourceName = ResourcePaths.WideButtonBlank, PaletteName = Palettes.Units } },
		{eButtonType.Medium, new ButtonLayout{ ResourceName = ResourcePaths.MediumButtonBlank, PaletteName = Palettes.Units } },
		{eButtonType.Narrow, new ButtonLayout { ResourceName = ResourcePaths.NarrowButtonBlank, PaletteName = Palettes.Units } },
		{eButtonType.Tall, new ButtonLayout { ResourceName = ResourcePaths.TallButtonBlank, PaletteName = Palettes.Units } },

		{eButtonType.Cancel, new ButtonLayout { ResourceName = ResourcePaths.CancelButton, PaletteName = Palettes.Units } },
		// Minipanel
		{eButtonType.MinipanelCharacter, new ButtonLayout { ResourceName = ResourcePaths.MinipanelButton, PaletteName = Palettes.Units, BaseFrame = 0 } },
		{eButtonType.MinipanelInventory, new ButtonLayout { ResourceName = ResourcePaths.MinipanelButton, PaletteName = Palettes.Units, BaseFrame = 2 } },
		{eButtonType.MinipanelSkill, new ButtonLayout { ResourceName = ResourcePaths.MinipanelButton, PaletteName = Palettes.Units, BaseFrame = 4 } },
		{eButtonType.MinipanelAutomap, new ButtonLayout { ResourceName = ResourcePaths.MinipanelButton, PaletteName = Palettes.Units, BaseFrame = 8 } },
		{eButtonType.MinipanelMessage, new ButtonLayout { ResourceName = ResourcePaths.MinipanelButton, PaletteName = Palettes.Units, BaseFrame = 10 } },
		{eButtonType.MinipanelQuest, new ButtonLayout { ResourceName = ResourcePaths.MinipanelButton, PaletteName = Palettes.Units, BaseFrame = 12 } },
		{eButtonType.MinipanelMenu, new ButtonLayout { ResourceName = ResourcePaths.MinipanelButton, PaletteName = Palettes.Units, BaseFrame = 14 } },

		{eButtonType.SecondaryInvHand, new ButtonLayout { ResourceName = ResourcePaths.InventoryWeaponsTab, PaletteName = Palettes.Units, ClickableRect = new Rectangle(0, 0, 0, 20), AllowFrameChange = false } },
		{eButtonType.Run, new ButtonLayout { ResourceName = ResourcePaths.RunButton, PaletteName = Palettes.Units, Toggleable = true } },
		{eButtonType.Menu, new ButtonLayout { ResourceName = ResourcePaths.MenuButton, PaletteName = Palettes.Units, Toggleable = true } },
		{eButtonType.GoldCoin, new ButtonLayout { ResourceName = ResourcePaths.GoldCoinButton, PaletteName = Palettes.Units } },
		{eButtonType.Close, new ButtonLayout { ResourceName = ResourcePaths.SquareButton, PaletteName = Palettes.Units, BaseFrame = 10 } },
		{eButtonType.Skill, new ButtonLayout { ResourceName = ResourcePaths.AddSkillButton, PaletteName = Palettes.Units, DisabledFrame = 2
	*/
}

// Button defines a standard wide UI button
type Button struct {
	enabled             bool
	x, y                int
	width, height       uint32
	visible             bool
	pressed             bool
	toggled             bool
	fileProvider        Common.FileProvider
	normalImage         *ebiten.Image
	pressedImage        *ebiten.Image
	toggledImage        *ebiten.Image
	pressedToggledImage *ebiten.Image
	disabledImage       *ebiten.Image
	onClick             func()
}

// CreateButton creates an instance of Button
func CreateButton(buttonType ButtonType, fileProvider Common.FileProvider, text string) *Button {
	result := &Button{
		fileProvider: fileProvider,
		width:        0,
		height:       0,
		visible:      true,
		enabled:      true,
		pressed:      false,
	}
	buttonLayout := ButtonLayouts[buttonType]
	font := GetFont(buttonLayout.FontPath, Palettes.Units, fileProvider)
	buttonSprite := fileProvider.LoadSprite(buttonLayout.ResourceName, buttonLayout.PaletteName)
	totalButtonTypes := buttonSprite.GetTotalFrames() / (buttonLayout.XSegments * buttonLayout.YSegments)
	for i := 0; i < buttonLayout.XSegments; i++ {
		w, _ := buttonSprite.GetFrameSize(i)
		result.width += w
	}
	for i := 0; i < buttonLayout.YSegments; i++ {
		_, h := buttonSprite.GetFrameSize(i * buttonLayout.YSegments)
		result.height += h
	}

	result.normalImage, _ = ebiten.NewImage(int(result.width), int(result.height), ebiten.FilterNearest)
	result.pressedImage, _ = ebiten.NewImage(int(result.width), int(result.height), ebiten.FilterNearest)
	textWidth, textHeight := font.GetTextMetrics(text)
	textX := (result.width / 2) - (textWidth / 2)
	textY := (result.height / 2) - (textHeight / 2) + 1
	buttonSprite.MoveTo(0, 0)
	buttonSprite.Blend = true
	buttonSprite.DrawSegments(result.normalImage, buttonLayout.XSegments, buttonLayout.YSegments, buttonLayout.BaseFrame)
	font.Draw(int(textX), int(textY), text, color.RGBA{100, 100, 100, 255}, result.normalImage)
	if buttonLayout.AllowFrameChange {
		if totalButtonTypes > 1 {
			buttonSprite.DrawSegments(result.pressedImage, buttonLayout.XSegments, buttonLayout.YSegments, buttonLayout.BaseFrame+1)
			font.Draw(int(textX-2), int(textY+2), text, color.RGBA{100, 100, 100, 255}, result.pressedImage)
		}
		if totalButtonTypes > 2 {
			buttonSprite.DrawSegments(result.toggledImage, buttonLayout.XSegments, buttonLayout.YSegments, buttonLayout.BaseFrame+2)
			font.Draw(int(textX), int(textY), text, color.RGBA{100, 100, 100, 255}, result.toggledImage)
		}
		if totalButtonTypes > 3 {
			buttonSprite.DrawSegments(result.pressedToggledImage, buttonLayout.XSegments, buttonLayout.YSegments, buttonLayout.BaseFrame+3)
			font.Draw(int(textX), int(textY), text, color.RGBA{100, 100, 100, 255}, result.pressedToggledImage)
		}
		if buttonLayout.DisabledFrame != -1 {
			buttonSprite.DrawSegments(result.disabledImage, buttonLayout.XSegments, buttonLayout.YSegments, buttonLayout.DisabledFrame)
			font.Draw(int(textX), int(textY), text, color.RGBA{100, 100, 100, 255}, result.disabledImage)
		}
	}
	return result
}

// OnActivated defines the callback handler for the activate event
func (v *Button) OnActivated(callback func()) {
	v.onClick = callback
}

// Activate calls the on activated callback handler, if any
func (v *Button) Activate() {
	if v.onClick == nil {
		return
	}
	v.onClick()
}

// Draw renders the button
func (v *Button) Draw(target *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{
		CompositeMode: ebiten.CompositeModeSourceAtop,
		Filter:        ebiten.FilterNearest,
	}
	opts.GeoM.Translate(float64(v.x), float64(v.y))

	if !v.enabled {
		target.DrawImage(v.disabledImage, opts)
	} else if v.toggled && v.pressed {
		target.DrawImage(v.pressedToggledImage, opts)
	} else if v.pressed {
		target.DrawImage(v.pressedImage, opts)
	} else if v.toggled {
		target.DrawImage(v.toggledImage, opts)
	} else {
		target.DrawImage(v.normalImage, opts)
	}
}

// GetEnabled returns the enabled state
func (v *Button) GetEnabled() bool {
	return v.enabled
}

// SetEnabled sets the enabled state
func (v *Button) SetEnabled(enabled bool) {
	v.enabled = enabled
}

// GetSize returns the size of the button
func (v *Button) GetSize() (uint32, uint32) {
	return v.width, v.height
}

// MoveTo moves the button
func (v *Button) MoveTo(x, y int) {
	v.x = x
	v.y = y
}

// GetLocation returns the location of the button
func (v *Button) GetLocation() (x, y int) {
	return v.x, v.y
}

// GetVisible returns the visibility of the button
func (v *Button) GetVisible() bool {
	return v.visible
}

// SetVisible sets the visibility of the button
func (v *Button) SetVisible(visible bool) {
	v.visible = visible
}

// GetPressed returns the pressed state of the button
func (v *Button) GetPressed() bool {
	return v.pressed
}

// SetPressed sets the pressed state of the button
func (v *Button) SetPressed(pressed bool) {
	v.pressed = pressed
}
