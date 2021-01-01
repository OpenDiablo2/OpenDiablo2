package d2checkbox

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2geom/rectangle"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2label"
	"github.com/gravestench/akara"
	"image/color"
	"math/rand"
)

type callbackFunc = func(this akara.Component) (preventPropagation bool)

// Button defines a standard wide UI button
type Checkbox struct {
	Layout        CheckboxLayout
	Sprite        d2interface.Sprite
	Label         *d2label.Label
	callback      callbackFunc
	width, height int
	pressed       bool
	enabled       bool
}

// CheckboxLayout defines the type of buttons
type CheckboxLayout struct {
	X                float64
	Y                float64
	SpritePath       string
	PalettePath      string
	FontPath         string
	ClickableRect    *rectangle.Rectangle
	XSegments        int
	YSegments        int
	BaseFrame        int
	DisabledFrame    int
	DisabledColor    uint32
	TextOffset       float64
	FixedWidth       int
	FixedHeight      int
	LabelColor       uint32
	Toggleable       bool
	AllowFrameChange bool
	HasImage         bool
	Tooltip          int
	TooltipXOffset   int
	TooltipYOffset   int
}

// New creates an instance of Button
func New() *Checkbox {
	checkbox := &Checkbox{
		Layout: GetDefaultLayout(),
	}

	return checkbox
}

func GetDefaultLayout() CheckboxLayout {
	return CheckboxLayout{
		X:                0,
		Y:                0,
		SpritePath:       d2resource.Checkbox,
		PalettePath:      d2resource.PaletteFechar,
		FontPath:         d2resource.FontExocet10,
		XSegments:        1,
		YSegments:        1,
		BaseFrame:        0,
		DisabledFrame:    -1,
		DisabledColor:    lightGreyAlpha75,
		TextOffset:       18,
		FixedWidth:       16,
		FixedHeight:      15,
		LabelColor:       goldAlpha100,
		Toggleable:       true,
		AllowFrameChange: true,
		HasImage:         true,
		Tooltip:          0,
		TooltipXOffset:   0,
		TooltipYOffset:   0,
	}
}

// OnActivated defines the callback handler for the activate event
func (v *Checkbox) OnActivated(callback callbackFunc) {
	v.callback = callback
}

// Activate calls the on activated callback handler, if any
func (v *Checkbox) Activate(thisComponent akara.Component) bool {
	if v.GetEnabled() {
		v.Toggle()
	}

	if v.callback != nil {
		return v.callback(thisComponent)
	}

	return false
}

// Toggle negates the toggled state of the button
func (v *Checkbox) Toggle() {
	v.SetPressed(!v.GetPressed())
}

// GetEnabled returns the enabled state
func (v *Checkbox) GetEnabled() bool {
	return v.enabled
}

// SetEnabled sets the enabled state
func (v *Checkbox) SetEnabled(enabled bool) {
	v.enabled = enabled
}

// GetChecked returns the enabled state
func (v *Checkbox) GetPressed() bool {
	return v.pressed
}

// SetEnabled sets the enabled state
func (v *Checkbox) SetPressed(pressed bool) {
	v.pressed = pressed
}

func (v *Checkbox) Update() {
	if v.Sprite == nil {
		return
	}

	if v.GetEnabled() && v.GetPressed() {
		// checked, enabled
		_ = v.Sprite.SetCurrentFrame(1)
	} else if v.GetEnabled() {
		// unchecked, enabled
		_ = v.Sprite.SetCurrentFrame(0)
	} else if v.GetPressed() {
		// checked, disabled
		_ = v.Sprite.SetCurrentFrame(1)
		v.Sprite.SetColorMod(color.RGBA{R: uint8(rand.Uint32() % 255), B: uint8(rand.Uint32() % 255), G: uint8(rand.Uint32() % 255), A: 0xff})
		v.Sprite.SetEffect(d2enum.DrawEffectPctTransparency25)
	} else {
		// unchecked, disabled
		_ = v.Sprite.SetCurrentFrame(0)
		v.Sprite.SetColorMod(color.RGBA{R: uint8(rand.Uint32() % 255), B: uint8(rand.Uint32() % 255), G: uint8(rand.Uint32() % 255), A: 0xff})
	}
}
