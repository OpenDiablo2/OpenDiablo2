package d2ui

import (
	"fmt"
	"image"
	"image/color"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2gui"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
)

// ButtonType defines the type of button
type ButtonType int

// ButtonType constants
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
	ResourceName     string
	PaletteName      string
	FontPath         string
	XSegments        int
	YSegments        int
	BaseFrame        int
	DisabledFrame    int
	ClickableRect    *image.Rectangle
	TextOffset       int
	Toggleable       bool
	AllowFrameChange bool
}

func getButtonLayouts() map[ButtonType]ButtonLayout {
	return map[ButtonType]ButtonLayout{
		ButtonTypeWide: {
			XSegments: 2, YSegments: 1, ResourceName: d2resource.WideButtonBlank, PaletteName: d2resource.PaletteUnits,
			DisabledFrame: -1, FontPath: d2resource.FontExocet10, AllowFrameChange: true, TextOffset: 1},
		ButtonTypeShort: {
			XSegments: 1, YSegments: 1, ResourceName: d2resource.ShortButtonBlank, PaletteName: d2resource.PaletteUnits,
			DisabledFrame: -1, FontPath: d2resource.FontRediculous, AllowFrameChange: true, TextOffset: -1},
		ButtonTypeMedium: {
			XSegments: 1, YSegments: 1, ResourceName: d2resource.MediumButtonBlank, PaletteName: d2resource.PaletteUnits,
			FontPath: d2resource.FontExocet10, AllowFrameChange: true},
		ButtonTypeTall: {
			XSegments: 1, YSegments: 1, ResourceName: d2resource.TallButtonBlank, PaletteName: d2resource.PaletteUnits,
			FontPath: d2resource.FontExocet10, AllowFrameChange: true, TextOffset: 5},
		ButtonTypeOkCancel: {
			XSegments: 1, YSegments: 1, ResourceName: d2resource.CancelButton, PaletteName: d2resource.PaletteUnits,
			DisabledFrame: -1, FontPath: d2resource.FontRediculous, AllowFrameChange: true},
		ButtonTypeRun: {
			XSegments: 1, YSegments: 1, ResourceName: d2resource.RunButton, PaletteName: d2resource.PaletteSky,
			Toggleable: true, DisabledFrame: -1, FontPath: d2resource.FontRediculous, AllowFrameChange: true},
	}
}

// Button defines a standard wide UI button
type Button struct {
	buttonLayout          ButtonLayout
	normalSurface         d2interface.Surface
	pressedSurface        d2interface.Surface
	toggledSurface        d2interface.Surface
	pressedToggledSurface d2interface.Surface
	disabledSurface       d2interface.Surface
	x                     int
	y                     int
	width                 int
	height                int
	onClick               func()
	enabled               bool
	visible               bool
	pressed               bool
	toggled               bool
}

// CreateButton creates an instance of Button
func CreateButton(renderer d2interface.Renderer, buttonType ButtonType, text string) Button {
	result := Button{
		width:   0,
		height:  0,
		visible: true,
		enabled: true,
		pressed: false,
	}
	buttonLayout := getButtonLayouts()[buttonType]
	result.buttonLayout = buttonLayout
	lbl := CreateLabel(buttonLayout.FontPath, d2resource.PaletteUnits)
	lbl.SetText(text)
	lbl.Color = color.RGBA{R: 100, G: 100, B: 100, A: 255}
	lbl.Alignment = d2gui.HorizontalAlignCenter

	animation, _ := d2asset.LoadAnimation(buttonLayout.ResourceName, buttonLayout.PaletteName)
	buttonSprite, _ := LoadSprite(animation)

	for i := 0; i < buttonLayout.XSegments; i++ {
		w, _, _ := buttonSprite.GetFrameSize(i)
		result.width += w
	}

	for i := 0; i < buttonLayout.YSegments; i++ {
		_, h, _ := buttonSprite.GetFrameSize(i * buttonLayout.YSegments)
		result.height += h
	}

	result.normalSurface, _ = renderer.NewSurface(result.width, result.height, d2enum.FilterNearest)

	buttonSprite.SetPosition(0, 0)
	buttonSprite.SetEffect(d2enum.DrawEffectModulate)

	result.renderFrames(renderer, buttonSprite, &buttonLayout, &lbl)

	return result
}

func (v *Button) renderFrames(renderer d2interface.Renderer, buttonSprite *Sprite, buttonLayout *ButtonLayout, label *Label) {
	totalButtonTypes := buttonSprite.GetFrameCount() / (buttonLayout.XSegments * buttonLayout.YSegments)

	var err error
	err = buttonSprite.RenderSegmented(v.normalSurface, buttonLayout.XSegments, buttonLayout.YSegments, buttonLayout.BaseFrame)

	if err != nil {
		fmt.Printf("failed to render button normalSurface, err: %v\n", err)
	}

	_, labelHeight := label.GetSize()
	textY := v.height/2 - labelHeight/2
	xOffset := v.width / 2

	label.SetPosition(xOffset, textY)
	label.Render(v.normalSurface)

	if buttonLayout.AllowFrameChange {
		if totalButtonTypes > 1 {
			v.pressedSurface, _ = renderer.NewSurface(v.width, v.height, d2enum.FilterNearest)
			err = buttonSprite.RenderSegmented(v.pressedSurface, buttonLayout.XSegments, buttonLayout.YSegments, buttonLayout.BaseFrame+1)

			if err != nil {
				fmt.Printf("failed to render button pressedSurface, err: %v\n", err)
			}

			label.SetPosition(xOffset-2, textY+2)
			label.Render(v.pressedSurface)
		}

		if totalButtonTypes > 2 {
			v.toggledSurface, _ = renderer.NewSurface(v.width, v.height, d2enum.FilterNearest)
			err = buttonSprite.RenderSegmented(v.toggledSurface, buttonLayout.XSegments, buttonLayout.YSegments, buttonLayout.BaseFrame+2)

			if err != nil {
				fmt.Printf("failed to render button toggledSurface, err: %v\n", err)
			}

			label.SetPosition(xOffset, textY)
			label.Render(v.toggledSurface)
		}

		if totalButtonTypes > 3 {
			v.pressedToggledSurface, _ = renderer.NewSurface(v.width, v.height, d2enum.FilterNearest)
			err = buttonSprite.RenderSegmented(v.pressedToggledSurface, buttonLayout.XSegments, buttonLayout.YSegments, buttonLayout.BaseFrame+3)

			if err != nil {
				fmt.Printf("failed to render button pressedToggledSurface, err: %v\n", err)
			}

			label.SetPosition(xOffset, textY)
			label.Render(v.pressedToggledSurface)
		}

		if buttonLayout.DisabledFrame != -1 {
			v.disabledSurface, _ = renderer.NewSurface(v.width, v.height, d2enum.FilterNearest)
			err = buttonSprite.RenderSegmented(v.disabledSurface, buttonLayout.XSegments, buttonLayout.YSegments, buttonLayout.DisabledFrame)

			if err != nil {
				fmt.Printf("failed to render button disabledSurface, err: %v\n", err)
			}

			label.SetPosition(xOffset, textY)
			label.Render(v.disabledSurface)
		}
	}
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
func (v *Button) Render(target d2interface.Surface) error {
	target.PushFilter(d2enum.FilterNearest)
	target.PushTranslation(v.x, v.y)

	defer target.PopN(2)

	var err error

	switch {
	case !v.enabled:
		target.PushColor(color.RGBA{R: 128, G: 128, B: 128, A: 195})
		defer target.Pop()
		err = target.Render(v.disabledSurface)
	case v.toggled && v.pressed:
		err = target.Render(v.pressedToggledSurface)
	case v.pressed:
		err = target.Render(v.pressedSurface)
	case v.toggled:
		err = target.Render(v.toggledSurface)
	default:
		err = target.Render(v.normalSurface)
	}

	if err != nil {
		fmt.Printf("failed to render button surface, err: %v\n", err)
	}

	return nil
}

// Toggle negates the toggled state of the button
func (v *Button) Toggle() {
	v.toggled = !v.toggled
}

// Advance advances the button state
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
func (v *Button) GetSize() (width, height int) {
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
