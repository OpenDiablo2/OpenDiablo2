package d2ui

import (
	"image"
	"image/color"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2assetmanager"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
)

// ButtonType defines the type of button
type ButtonType int

const (
	ButtonTypeWide     ButtonType = 1
	ButtonTypeMedium   ButtonType = 2
	ButtonTypeNarrow   ButtonType = 3
	ButtonTypeCancel   ButtonType = 4
	ButtonTypeTall     ButtonType = 5
	ButtonTypeShort    ButtonType = 6
	ButtonTypeOkCancel ButtonType = 7

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
	XSegments        int              //1
	YSegments        int              // 1
	ResourceName     string           // Font Name
	PaletteName      string           // PaletteType
	Toggleable       bool             // false
	BaseFrame        int              // 0
	DisabledFrame    int              // -1
	FontPath         string           // ResourcePaths.FontExocet10
	ClickableRect    *image.Rectangle // nil
	AllowFrameChange bool             // true
	TextOffset       int              // 0
}

// ButtonLayouts define the type of buttons you can have
var ButtonLayouts = map[ButtonType]ButtonLayout{
	ButtonTypeWide:     {2, 1, d2resource.WideButtonBlank, d2resource.PaletteUnits, false, 0, -1, d2resource.FontExocet10, nil, true, 1},
	ButtonTypeShort:    {1, 1, d2resource.ShortButtonBlank, d2resource.PaletteUnits, false, 0, -1, d2resource.FontRediculous, nil, true, -1},
	ButtonTypeMedium:   {1, 1, d2resource.MediumButtonBlank, d2resource.PaletteUnits, false, 0, 0, d2resource.FontExocet10, nil, true, 0},
	ButtonTypeTall:     {1, 1, d2resource.TallButtonBlank, d2resource.PaletteUnits, false, 0, 0, d2resource.FontExocet10, nil, true, 5},
	ButtonTypeOkCancel: {1, 1, d2resource.CancelButton, d2resource.PaletteUnits, false, 0, -1, d2resource.FontRediculous, nil, true, 0},
	/*
		{eButtonType.Wide,  new ButtonLayout { XSegments = 2, ResourceName = ResourcePaths.WideButtonBlank, PaletteName = PaletteDefs.Units } },
		{eButtonType.Narrow, new ButtonLayout { ResourceName = ResourcePaths.NarrowButtonBlank, PaletteName = PaletteDefs.Units } },
		{eButtonType.Cancel, new ButtonLayout { ResourceName = ResourcePaths.CancelButton, PaletteName = PaletteDefs.Units } },
		// Minipanel
		{eButtonType.MinipanelCharacter, new ButtonLayout { ResourceName = ResourcePaths.MinipanelButton, PaletteName = PaletteDefs.Units, BaseFrame = 0 } },
		{eButtonType.MinipanelInventory, new ButtonLayout { ResourceName = ResourcePaths.MinipanelButton, PaletteName = PaletteDefs.Units, BaseFrame = 2 } },
		{eButtonType.MinipanelSkill, new ButtonLayout { ResourceName = ResourcePaths.MinipanelButton, PaletteName = PaletteDefs.Units, BaseFrame = 4 } },
		{eButtonType.MinipanelAutomap, new ButtonLayout { ResourceName = ResourcePaths.MinipanelButton, PaletteName = PaletteDefs.Units, BaseFrame = 8 } },
		{eButtonType.MinipanelMessage, new ButtonLayout { ResourceName = ResourcePaths.MinipanelButton, PaletteName = PaletteDefs.Units, BaseFrame = 10 } },
		{eButtonType.MinipanelQuest, new ButtonLayout { ResourceName = ResourcePaths.MinipanelButton, PaletteName = PaletteDefs.Units, BaseFrame = 12 } },
		{eButtonType.MinipanelMenu, new ButtonLayout { ResourceName = ResourcePaths.MinipanelButton, PaletteName = PaletteDefs.Units, BaseFrame = 14 } },

		{eButtonType.SecondaryInvHand, new ButtonLayout { ResourceName = ResourcePaths.InventoryWeaponsTab, PaletteName = PaletteDefs.Units, ClickableRect = new Rectangle(0, 0, 0, 20), AllowFrameChange = false } },
		{eButtonType.Run, new ButtonLayout { ResourceName = ResourcePaths.RunButton, PaletteName = PaletteDefs.Units, Toggleable = true } },
		{eButtonType.Menu, new ButtonLayout { ResourceName = ResourcePaths.MenuButton, PaletteName = PaletteDefs.Units, Toggleable = true } },
		{eButtonType.GoldCoin, new ButtonLayout { ResourceName = ResourcePaths.GoldCoinButton, PaletteName = PaletteDefs.Units } },
		{eButtonType.Close, new ButtonLayout { ResourceName = ResourcePaths.SquareButton, PaletteName = PaletteDefs.Units, BaseFrame = 10 } },
		{eButtonType.Skill, new ButtonLayout { ResourceName = ResourcePaths.AddSkillButton, PaletteName = PaletteDefs.Units, DisabledFrame = 2
	*/
}

// Button defines a standard wide UI button
type Button struct {
	enabled               bool
	x, y                  int
	width, height         int
	visible               bool
	pressed               bool
	toggled               bool
	normalSurface         d2common.Surface
	pressedSurface        d2common.Surface
	toggledSurface        d2common.Surface
	pressedToggledSurface d2common.Surface
	disabledSurface       d2common.Surface
	buttonLayout          ButtonLayout
	onClick               func()
}

// CreateButton creates an instance of Button
func CreateButton(buttonType ButtonType, text string) Button {
	result := Button{
		width:   0,
		height:  0,
		visible: true,
		enabled: true,
		pressed: false,
	}
	buttonLayout := ButtonLayouts[buttonType]
	result.buttonLayout = buttonLayout
	font := GetFont(buttonLayout.FontPath, d2resource.PaletteUnits)

	animation, _ := d2assetmanager.LoadAnimation(buttonLayout.ResourceName, buttonLayout.PaletteName)
	buttonSprite, _ := d2render.LoadSprite(animation)
	totalButtonTypes := buttonSprite.GetFrameCount() / (buttonLayout.XSegments * buttonLayout.YSegments)
	for i := 0; i < buttonLayout.XSegments; i++ {
		w, _, _ := buttonSprite.GetFrameSize(i)
		result.width += w
	}
	for i := 0; i < buttonLayout.YSegments; i++ {
		_, h, _ := buttonSprite.GetFrameSize(i * buttonLayout.YSegments)
		result.height += h
	}

	_, result.normalSurface = d2render.NewSurface(int(result.width), int(result.height), d2common.FilterNearest)
	_, fontHeight := font.GetTextMetrics(text)
	textY := int((result.height/2)-(int(fontHeight)/2)) + buttonLayout.TextOffset

	buttonSprite.SetPosition(0, 0)
	buttonSprite.SetBlend(true)
	buttonSprite.RenderSegmented(result.normalSurface, buttonLayout.XSegments, buttonLayout.YSegments, buttonLayout.BaseFrame)
	font.Render(0, textY, text, color.RGBA{100, 100, 100, 255}, result.normalSurface)
	if buttonLayout.AllowFrameChange {
		if totalButtonTypes > 1 {
			_, result.pressedSurface = d2render.NewSurface(int(result.width), int(result.height), d2common.FilterNearest)
			buttonSprite.RenderSegmented(result.pressedSurface, buttonLayout.XSegments, buttonLayout.YSegments, buttonLayout.BaseFrame+1)
			font.Render(-2, textY+2, text, color.RGBA{100, 100, 100, 255}, result.pressedSurface)
		}
		if totalButtonTypes > 2 {
			_, result.toggledSurface = d2render.NewSurface(int(result.width), int(result.height), d2common.FilterNearest)
			buttonSprite.RenderSegmented(result.toggledSurface, buttonLayout.XSegments, buttonLayout.YSegments, buttonLayout.BaseFrame+2)
			font.Render(0, textY, text, color.RGBA{100, 100, 100, 255}, result.toggledSurface)
		}
		if totalButtonTypes > 3 {
			_, result.pressedToggledSurface = d2render.NewSurface(int(result.width), int(result.height), d2common.FilterNearest)
			buttonSprite.RenderSegmented(result.pressedToggledSurface, buttonLayout.XSegments, buttonLayout.YSegments, buttonLayout.BaseFrame+3)
			font.Render(0, textY, text, color.RGBA{100, 100, 100, 255}, result.pressedToggledSurface)
		}
		if buttonLayout.DisabledFrame != -1 {
			_, result.disabledSurface = d2render.NewSurface(int(result.width), int(result.height), d2common.FilterNearest)
			buttonSprite.RenderSegmented(result.disabledSurface, buttonLayout.XSegments, buttonLayout.YSegments, buttonLayout.DisabledFrame)
			font.Render(0, textY, text, color.RGBA{100, 100, 100, 255}, result.disabledSurface)
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

// Render renders the button
func (v *Button) Render(target d2common.Surface) {
	target.PushCompositeMode(d2common.CompositeModeSourceAtop)
	target.PushFilter(d2common.FilterNearest)
	target.PushTranslation(v.x, v.y)
	defer target.PopN(3)

	if !v.enabled {
		target.PushColor(color.RGBA{128, 128, 128, 195})
		defer target.Pop()
		target.Render(v.disabledSurface)
	} else if v.toggled && v.pressed {
		target.Render(v.pressedToggledSurface)
	} else if v.pressed {
		target.Render(v.pressedSurface)
	} else if v.toggled {
		target.Render(v.toggledSurface)
	} else {
		target.Render(v.normalSurface)
	}
}

func (v *Button) Advance(elapsed float64) {

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
func (v *Button) GetSize() (int, int) {
	return v.width, v.height
}

// SetPosition moves the button
func (v *Button) SetPosition(x, y int) {
	v.x = x
	v.y = y
}

// GetPosition returns the location of the button
func (v *Button) GetPosition() (x, y int) {
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
