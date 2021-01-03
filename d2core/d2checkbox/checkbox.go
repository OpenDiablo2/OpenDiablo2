package d2checkbox

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2geom/rectangle"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2label"
)

// set up defaults
const (
	CheckboxDefaultTextOffset = 18
	CheckboxDefaultWidth      = 16
	CheckboxDefaultHeight     = 15
)

type callbackFunc = func(this akara.Component) (preventPropagation bool)

// Checkbox defines a standard wide UI button
type Checkbox struct {
	Layout   CheckboxLayout
	Sprite   d2interface.Sprite
	Label    *d2label.Label
	callback callbackFunc
	pressed  bool
	enabled  bool
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

// GetDefaultLayout returns the default layout of a checkbox.
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
		TextOffset:       CheckboxDefaultTextOffset,
		FixedWidth:       CheckboxDefaultWidth,
		FixedHeight:      CheckboxDefaultHeight,
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

// GetPressed returns the enabled state
func (v *Checkbox) GetPressed() bool {
	return v.pressed
}

// SetPressed sets the enabled state
func (v *Checkbox) SetPressed(pressed bool) {
	v.pressed = pressed
}

// Update updates the checkbox's sprite in accordance with the checkbox's current state.
// This ensures that the checkbox rendered in the UI accurately reflects the state of the checkbox.
func (v *Checkbox) Update() {
	if v.Sprite == nil {
		return
	}

	switch {
	case v.GetEnabled() && v.GetPressed():
		// checked, enabled
		_ = v.Sprite.SetCurrentFrame(1)
	case v.GetEnabled():
		// unchecked, enabled
		_ = v.Sprite.SetCurrentFrame(0)
	case v.GetPressed():
		// checked, disabled
		_ = v.Sprite.SetCurrentFrame(1)
	default:
		// unchecked, disabled
		_ = v.Sprite.SetCurrentFrame(0)
	}
}
